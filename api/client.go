package api

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strconv"
)

type Client struct {
	httpClient *http.Client
	apiKey     string
	apiUser    string
	apiPass    string
	restUrl    string
	logger     *slog.Logger
}

type LoadMasterResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
	Status  string `json:"status"`
}

type LoadMasterRequest struct {
	Command string `json:"cmd"`
	ApiUser string `json:"apiuser,omitempty"`
	ApiPass string `json:"apipass,omitempty"`
	ApiKey  string `json:"apikey,omitempty"`
}

type LoadMasterDataResponse struct {
	*LoadMasterResponse
	Data string `json:"data,omitempty"`
}

type LoadMasterError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *LoadMasterError) Error() string {
	return "Code: " + strconv.Itoa(e.Code) + ", Message: " + e.Message
}

type AuthInjectable interface {
	injectAuth(*Client) error
	getCommand(*Client) string
}

func (r *LoadMasterRequest) injectAuth(c *Client) (err error) {
	c.logger.Debug("Injecting authentication credentials in payload")
	if c.apiUser != "" && c.apiPass != "" {
		c.logger.Debug("Using username and password authentication")

		r.ApiUser = c.apiUser
		r.ApiPass = c.apiPass
		return nil
	}

	if c.apiKey != "" {
		c.logger.Debug("Using API key authentication")
		r.ApiKey = c.apiKey
		return nil
	}

	return fmt.Errorf("missing authentication")
}

func (r *LoadMasterRequest) getCommand(c *Client) string {
	return r.Command
}

type HTTPWithResponseCode interface {
	getResponseCode() int
	getResponseMessage() string
}

func (r LoadMasterResponse) getResponseCode() int {
	return r.Code
}

func (r LoadMasterResponse) getResponseMessage() string {
	return r.Message
}

func NewClient(restUrl string, apiKey string, apiUser string, apiPass string) *Client {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	return &Client{
		httpClient: http.DefaultClient,
		apiKey:     apiKey,
		apiUser:    apiUser,
		apiPass:    apiPass,
		restUrl:    restUrl,
		logger:     slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}
}

func NewClientWithUsernamePassword(restUrl string, apiUser string, apiPass string) *Client {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	return &Client{
		httpClient: http.DefaultClient,
		apiUser:    apiUser,
		apiPass:    apiPass,
		restUrl:    restUrl,
		logger:     slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}
}

func NewClientWithApiKey(restUrl string, apiKey string) *Client {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	return &Client{
		httpClient: http.DefaultClient,
		apiKey:     apiKey,
		restUrl:    restUrl,
		logger:     slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}
}

func (c *Client) SetLogger(logger *slog.Logger) {
	c.logger = logger
}

func sendRequest[T HTTPWithResponseCode](c *Client, payload AuthInjectable, response T) (*T, error) {
	c.logger.Info("Initiate communication with LoadMaster API")
	request, err := c.newRequest(payload)
	if err != nil {
		return nil, err
	}

	http_response, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	c.logger.Debug("Unmarshalling response")
	err = json.Unmarshal(http_response, &response)
	if err != nil {
		c.logger.Error("Error unmarshalling response: ", "Error", err)
		return nil, err
	}

	return &response, nil
}

func (c *Client) newRequest(payload AuthInjectable) (*http.Request, error) {
	c.logger.Info("Creating new request for LoadMaster API", "Request", payload)
	err := payload.injectAuth(c)

	if err != nil {
		c.logger.Error("Error injecting authentication credentials in payload", "Error", err)
		return nil, err
	}

	c.logger.Debug("Marshalling payload before sending")
	b, err := json.Marshal(payload)
	if err != nil {
		c.logger.Error("Error marshalling payload to json: ", "Error", err)
		return nil, err
	}
	c.logger.Debug("Payload marshalled successfully", "Payload", string(b))

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/accessv2", c.restUrl), bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	return req, nil

}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	c.logger.Info("Sending request to LoadMaster API", "URL", req.URL.String(), "Method", req.Method)

	res, err := c.httpClient.Do(req)
	if err != nil {
		c.logger.Error("Error sending request:", "Error", err)
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		c.logger.Error("Error reading body of response:", "Error", err)
		return nil, err
	}
	if res.StatusCode < 400 {
		c.logger.Debug("Response", "Status", res.Status, "Headers", res.Header, "Body", string(body))

		return body, nil
	} else {
		c.logger.Error("Error in response:", slog.String("status", res.Status), slog.String("body", string(body)))

		return nil, &LoadMasterError{Code: res.StatusCode, Message: string(body)}
	}
}

func (r LoadMasterRequest) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("cmd", r.Command),
		slog.String("ApiUser", "[redacted]"),
		slog.String("ApiPass", "[redacted]"),
		slog.String("ApiKey", "[redacted]"),
	)
}

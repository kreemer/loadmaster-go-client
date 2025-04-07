package api

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"
)

type Client struct {
	httpClient *http.Client
	apiKey     string
	apiUser    string
	apiPass    string
	restUrl    string
	debug      uint8
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
	slog.Debug("Injecting authentication credentials in payload")
	if c.apiUser != "" && c.apiPass != "" {
		slog.Debug("Using username and password authentication")

		r.ApiUser = c.apiUser
		r.ApiPass = c.apiPass
		return nil
	}

	if c.apiKey != "" {
		slog.Debug("Using API key authentication")
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
		debug:      0,
	}
}

func NewClientWithUsernamePassword(restUrl string, apiUser string, apiPass string) *Client {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	return &Client{
		httpClient: http.DefaultClient,
		apiUser:    apiUser,
		apiPass:    apiPass,
		restUrl:    restUrl,
		debug:      0,
	}
}

func NewClientWithApiKey(restUrl string, apiKey string) *Client {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	return &Client{
		httpClient: http.DefaultClient,
		apiKey:     apiKey,
		restUrl:    restUrl,
		debug:      0,
	}
}

func (c *Client) SetDebugLevel(level uint8) {
	c.debug = level
}

func sendRequest[T HTTPWithResponseCode](c *Client, payload AuthInjectable, response T) (*T, error) {
	slog.Info("Initiate communication with LoadMaster API")
	request, err := c.newRequest(payload)
	if err != nil {
		return nil, err
	}

	http_response, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	slog.Debug("Unmarshalling response")
	err = json.Unmarshal(http_response, &response)
	if err != nil {
		slog.Error("Error unmarshalling response: ", "Error", err)
		return nil, err
	}

	if response.getResponseCode() != http.StatusOK && response.getResponseCode() != http.StatusNoContent {
		return nil, &LoadMasterError{
			Code:    response.getResponseCode(),
			Message: response.getResponseMessage(),
		}
	}

	return &response, nil
}

func (c *Client) newRequest(payload AuthInjectable) (*http.Request, error) {
	slog.Info("Creating new request for LoadMaster API", "Request", payload)
	err := payload.injectAuth(c)

	if err != nil {
		slog.Error("Error injecting authentication credentials in payload", "Error", err)
		return nil, err
	}

	slog.Debug("Marshalling payload before sending")
	b, err := json.Marshal(payload)
	if err != nil {
		slog.Error("Error marshalling payload to json: ", "Error", err)
		return nil, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/accessv2", c.restUrl), bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	return req, nil

}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	slog.Info("Sending request to LoadMaster API", "URL", req.URL.String(), "Method", req.Method)

	res, err := c.httpClient.Do(req)
	if err != nil {
		slog.Error("Error sending request:", "Error", err)
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		slog.Error("Error reading body of response:", "Error", err)
		return nil, err
	}
	if res.StatusCode == http.StatusOK || res.StatusCode == http.StatusNoContent || res.StatusCode == http.StatusUnprocessableEntity {
		slog.Debug("Response", "Status", res.Status, "Headers", res.Header, "Body", string(body))
		return body, err
	} else {
		slog.Error("Error in response:", slog.String("status", res.Status), slog.String("body", string(body)))

		return body, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
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

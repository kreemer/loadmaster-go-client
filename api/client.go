package api

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
)

type Client struct {
	httpClient *http.Client
	apiKey     string
	apiUser    string
	apiPass    string
	restUrl    string
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

type AuthInjectable interface {
	injectAuth(*Client) error
}

func (r *LoadMasterRequest) injectAuth(c *Client) (err error) {
	if c.apiUser != "" && c.apiPass != "" {

		r.ApiUser = c.apiUser
		r.ApiPass = c.apiPass
		return nil
	}

	if c.apiKey != "" {
		r.ApiKey = c.apiKey
		return nil
	}

	return fmt.Errorf("missing authentication")
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
	}
}

func NewClientWithUsernamePassword(restUrl string, apiUser string, apiPass string) *Client {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	return &Client{
		httpClient: http.DefaultClient,
		apiUser:    apiUser,
		apiPass:    apiPass,
		restUrl:    restUrl,
	}
}

func NewClientWithApiKey(restUrl string, apiKey string) *Client {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	return &Client{
		httpClient: http.DefaultClient,
		apiKey:     apiKey,
		restUrl:    restUrl,
	}
}

func sendRequest[T HTTPWithResponseCode](c *Client, payload AuthInjectable, response *T) (*T, error) {

	log.Debug().Msg("Sending request")
	log.Trace().Interface("payload", payload).Msg("Payload")

	http, err := c.newRequest(payload)
	if err != nil {
		log.Error().Err(err).Msg("Error creating request")
		return nil, err
	}

	http_response, err := c.doRequest(http)
	if err != nil {
		log.Error().Err(err).Msg("Error sending request")
		return nil, err
	}

	log.Trace().Interface("response", http_response).Msg("Response")
	err = json.Unmarshal(http_response, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) newRequest(payload AuthInjectable) (*http.Request, error) {
	err := payload.injectAuth(c)

	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/accessv2", c.restUrl), bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	return req, nil

}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode == http.StatusOK || res.StatusCode == http.StatusNoContent {
		return body, err
	} else {
		return body, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}
}

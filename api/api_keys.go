package api

import (
	"encoding/json"
	"log/slog"
)

type DeleteApiKeyRequest struct {
	Key string `json:"key"`
}

type ListApiKeyResponse struct {
	*LoadMasterResponse
	ApiKeys []string `json:"apikeys"`
}

type GenerateApiKeyResponse struct {
	*LoadMasterResponse
	ApiKeys []string `json:"apikeys"`
}

type DeleteApiKeyResponse struct {
	*LoadMasterResponse
	ApiKeys []string `json:"apikeys"`
}

func (c *Client) ListApiKey() (*ListApiKeyResponse, error) {
	slog.Debug("Listing API keys")
	payload := struct {
		*LoadMasterRequest
	}{
		&LoadMasterRequest{
			Command: "listapikeys",
		},
	}

	http, err := c.newRequest(payload)
	if err != nil {
		return nil, err
	}

	http_response, err := c.doRequest(http)
	if err != nil {
		return nil, err
	}

	response := &ListApiKeyResponse{}
	err = json.Unmarshal(http_response, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GenerateApiKey() (*GenerateApiKeyResponse, error) {
	slog.Debug("Generating API key")
	payload := struct {
		*LoadMasterRequest
	}{
		&LoadMasterRequest{
			Command: "addapikey",
		},
	}
	http, err := c.newRequest(payload)
	if err != nil {
		return nil, err
	}

	http_response, err := c.doRequest(http)
	if err != nil {
		return nil, err
	}

	response := &GenerateApiKeyResponse{}
	err = json.Unmarshal(http_response, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) DeleteApiKey(request DeleteApiKeyRequest) (*DeleteApiKeyResponse, error) {
	slog.Debug("Deleting API key")
	payload := struct {
		*LoadMasterRequest
		Key string `json:"key"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "delapikey",
		},
		Key: request.Key,
	}

	http, err := c.newRequest(payload)
	if err != nil {
		return nil, err
	}

	http_response, err := c.doRequest(http)
	if err != nil {
		return nil, err
	}

	response := &DeleteApiKeyResponse{}
	err = json.Unmarshal(http_response, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

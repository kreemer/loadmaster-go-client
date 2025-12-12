package api

import (
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

	response, err := sendRequest(c, payload, ListApiKeyResponse{})
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
	response, err := sendRequest(c, payload, GenerateApiKeyResponse{})
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

	response, err := sendRequest(c, payload, DeleteApiKeyResponse{})
	if err != nil {
		return nil, err
	}

	return response, nil
}

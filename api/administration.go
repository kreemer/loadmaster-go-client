package api

import (
	"encoding/json"
	"log/slog"
)

func (c *Client) Backup() (*LoadMasterDataResponse, error) {
	slog.Debug("Backup")
	payload := struct {
		*LoadMasterRequest
	}{
		&LoadMasterRequest{
			Command: "backup",
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

	response := &LoadMasterDataResponse{}
	err = json.Unmarshal(http_response, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) Restore(data string, restore_type string) (*LoadMasterResponse, error) {
	slog.Debug("Restore base configuration", "type", restore_type)
	payload := struct {
		*LoadMasterRequest
		Data string `json:"data"`
		Type string `json:"type"`
	}{
		&LoadMasterRequest{
			Command: "restore",
		},
		data,
		restore_type,
	}
	http, err := c.newRequest(payload)
	if err != nil {
		return nil, err
	}

	http_response, err := c.doRequest(http)
	if err != nil {
		return nil, err
	}

	response := &LoadMasterResponse{}
	err = json.Unmarshal(http_response, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

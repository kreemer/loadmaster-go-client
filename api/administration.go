package api

import (
	"encoding/json"
	"fmt"
	"log/slog"
)

type BackupResponse struct {
	*LoadMasterResponse
	Data string `json:"data,omitempty"`
}

func (c *Client) Backup() (*BackupResponse, error) {
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

	response := &BackupResponse{}
	err = json.Unmarshal(http_response, response)
	if err != nil {
		return nil, err
	}

	if response.Code >= 400 {
		return nil, fmt.Errorf("error: %s", response.Message)
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

	if response.Code >= 400 {
		return nil, fmt.Errorf("error: %s", response.Message)
	}

	return response, nil
}

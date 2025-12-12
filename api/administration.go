package api

import (
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

	response, err := sendRequest(c, payload, LoadMasterDataResponse{})
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
	response, err := sendRequest(c, payload, LoadMasterResponse{})
	if err != nil {
		return nil, err
	}
	return response, nil
}

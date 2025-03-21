package api

import (
	"encoding/json"
	"fmt"
)

type SubVirtualService struct {
	*VirtualService
	Name    string `json:"Name"`
	Forward string `json:"Forward"`
}

type ShowSubVirtualServiceResponse struct {
	*LoadMasterResponse
	*SubVirtualService
}

func (c *Client) ShowSubVirtualService(identifier int) (*ShowSubVirtualServiceResponse, error) {
	payload := struct {
		*LoadMasterRequest
		VS int `json:"vs"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "showvs",
		},
		VS: identifier,
	}

	http, err := c.newRequest(payload)
	if err != nil {
		return nil, err
	}

	http_response, err := c.doRequest(http)
	if err != nil {
		return nil, err
	}

	response := &ShowSubVirtualServiceResponse{}
	err = json.Unmarshal(http_response, response)
	if err != nil {
		return nil, err
	}

	if response.Code >= 400 {
		return nil, fmt.Errorf("error: %s", response.Message)
	}

	return response, nil
}

func (c *Client) AddSubVirtualService(vs_identifier int, parameters VirtualServiceParameters) (*ShowSubVirtualServiceResponse, error) {
	payload := struct {
		*LoadMasterRequest
		VS          int    `json:"vs"`
		CreateSubVS string `json:"createsubvs"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "modvs",
		},
		VS:          vs_identifier,
		CreateSubVS: "",
	}

	http, err := c.newRequest(payload)
	if err != nil {
		return nil, err
	}

	http_response, err := c.doRequest(http)
	if err != nil {
		return nil, err
	}

	response := &ShowSubVirtualServiceResponse{}
	err = json.Unmarshal(http_response, response)
	if err != nil {
		return nil, err
	}

	if response.Code >= 400 {
		return nil, fmt.Errorf("error: %s", response.Message)
	}

	return response, nil
}

func (c *Client) ModifySubVirtualService(identifier int, parameters VirtualServiceParameters) (*ShowSubVirtualServiceResponse, error) {
	payload := struct {
		*LoadMasterRequest
		VS      int    `json:"vs"`
		Persist string `json:"persist"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "modvs",
		},
		VS:      identifier,
		Persist: "super",
	}

	http, err := c.newRequest(payload)
	if err != nil {
		return nil, err
	}

	http_response, err := c.doRequest(http)
	if err != nil {
		return nil, err
	}

	response := &ShowSubVirtualServiceResponse{}
	err = json.Unmarshal(http_response, response)
	if err != nil {
		return nil, err
	}

	if response.Code >= 400 {
		return nil, fmt.Errorf("error: %s", response.Message)
	}

	return response, nil
}

func (c *Client) DeleteSubVirtualService(identifier int) (*LoadMasterResponse, error) {
	payload := struct {
		*LoadMasterRequest
		VS int `json:"vs"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "delvs",
		},
		VS: identifier,
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

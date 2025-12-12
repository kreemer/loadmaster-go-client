package api

import "log/slog"

type SubVirtualService struct {
	*VirtualService
	Name    string `json:"Name,omitempty"`
	Forward string `json:"Forward,omitempty"`
	VSIndex int32  `json:"VSIndex,omitempty"`
}

type ShowSubVirtualServiceResponse struct {
	*LoadMasterResponse
	*SubVirtualService
}

func (c *Client) ShowSubVirtualService(identifier string) (*ShowSubVirtualServiceResponse, error) {
	slog.Debug("Showing sub virtual service", "identifier", identifier)
	payload := struct {
		*LoadMasterRequest
		VS string `json:"vs"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "showvs",
		},
		VS: identifier,
	}

	response, err := sendRequest(c, payload, ShowSubVirtualServiceResponse{})
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) AddSubVirtualService(vs_identifier string, parameters VirtualServiceParameters) (*ShowSubVirtualServiceResponse, error) {
	slog.Debug("Adding sub virtual service", "vs_identifier", vs_identifier)
	payload := struct {
		*LoadMasterRequest
		*VirtualServiceParameters
		VS          string `json:"vs"`
		CreateSubVS string `json:"createsubvs"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "modvs",
		},
		VS:                       vs_identifier,
		CreateSubVS:              "",
		VirtualServiceParameters: &parameters,
	}

	response, err := sendRequest(c, payload, ShowSubVirtualServiceResponse{})
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ModifySubVirtualService(identifier string, parameters VirtualServiceParameters) (*ShowSubVirtualServiceResponse, error) {
	slog.Debug("Modifying sub virtual service", "identifier", identifier)
	payload := struct {
		*LoadMasterRequest
		*VirtualServiceParameters
		VS string `json:"vs"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "modvs",
		},
		VS:                       identifier,
		VirtualServiceParameters: &parameters,
	}

	response, err := sendRequest(c, payload, ShowSubVirtualServiceResponse{})
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) DeleteSubVirtualService(identifier string) (*LoadMasterResponse, error) {
	slog.Debug("Deleting sub virtual service", "identifier", identifier)

	payload := struct {
		*LoadMasterRequest
		VS string `json:"vs"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "delvs",
		},
		VS: identifier,
	}
	response, err := sendRequest(c, payload, LoadMasterResponse{})
	if err != nil {
		return nil, err
	}

	return response, nil
}

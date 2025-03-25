package api

type SubVirtualService struct {
	*VirtualService
	Name    string `json:"Name,omitempty"`
	Forward string `json:"Forward,omitempty"`
	VSIndex int    `json:"VSIndex,omitempty"`
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

	response, err := sendRequest(c, payload, &ShowSubVirtualServiceResponse{})
	if err != nil {
		return nil, err
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

	response, err := sendRequest(c, payload, &ShowSubVirtualServiceResponse{})
	if err != nil {
		return nil, err
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

	response, err := sendRequest(c, payload, &ShowSubVirtualServiceResponse{})
	if err != nil {
		return nil, err
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
	response, err := sendRequest(c, payload, &LoadMasterResponse{})
	if err != nil {
		return nil, err
	}

	return response, nil
}

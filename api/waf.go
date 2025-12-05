package api

import "log/slog"

func (c *Client) ShowWafRule(filename string) (*LoadMasterDataResponse, error) {
	slog.Debug("Show waf rule", "filename", filename)

	payload := struct {
		*LoadMasterRequest
		Filename string `json:"filename"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "downloadwafcustomrule",
		},
		Filename: filename,
	}

	response, err := sendRequest(c, payload, LoadMasterDataResponse{})
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) AddWafRule(filename string, data string) (*LoadMasterResponse, error) {
	slog.Debug("Add waf rule", "filename", filename)
	payload := struct {
		*LoadMasterRequest
		Filename string `json:"filename"`
		Data     string `json:"data"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "addwafcustomrule",
		},
		Filename: filename,
		Data:     data,
	}

	response, err := sendRequest(c, payload, LoadMasterResponse{})
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) DeleteWafRule(filename string) (*DeleteVirtualServiceResponse, error) {
	slog.Debug("Deleting waf rule", "filename", filename)
	payload := struct {
		*LoadMasterRequest
		Filename string `json:"filename"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "delwafcustomrule",
		},
		Filename: filename,
	}

	response, err := sendRequest(c, payload, DeleteVirtualServiceResponse{})
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ShowWafData(filename string) (*LoadMasterDataResponse, error) {
	slog.Debug("Show waf data", "filename", filename)

	payload := struct {
		*LoadMasterRequest
		Filename string `json:"filename"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "downloadwafcustomdata",
		},
		Filename: filename,
	}

	response, err := sendRequest(c, payload, LoadMasterDataResponse{})
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) AddWafData(filename string, data string) (*LoadMasterResponse, error) {
	slog.Debug("Add waf data", "filename", filename)
	payload := struct {
		*LoadMasterRequest
		Filename string `json:"filename"`
		Data     string `json:"data"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "addwafcustomdata",
		},
		Filename: filename,
		Data:     data,
	}

	response, err := sendRequest(c, payload, LoadMasterResponse{})
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) DeleteWafData(filename string) (*DeleteVirtualServiceResponse, error) {
	slog.Debug("Deleting waf data", "filename", filename)
	payload := struct {
		*LoadMasterRequest
		Filename string `json:"filename"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "delwafcustomdata",
		},
		Filename: filename,
	}

	response, err := sendRequest(c, payload, DeleteVirtualServiceResponse{})
	if err != nil {
		return nil, err
	}

	return response, nil
}

package api

import (
	"log/slog"
)

type OwaspRuleResponse struct {
	*LoadMasterResponse
	Rule OwaspRule `json:"Rule"`
}

type OwaspRule struct {
	Id       string `json:"Id,omitempty"`
	Type     string `json:"Type,omitempty"`
	Name     string `json:"Name,omitempty"`
	Enabled  string `json:"Enabled,omitempty"`
	RunFirst string `json:"Runfirst,omitempty"`
}

func (c *Client) AddOwaspCustomRule(filename string, data string) (*LoadMasterResponse, error) {
	slog.Debug("Adding OWASP custom rule")
	payload := struct {
		*LoadMasterRequest
		Filename string `json:"filename"`
		Data     string `json:"data"`
	}{
		&LoadMasterRequest{
			Command: "addowaspcustomrule",
		},
		filename,
		data,
	}

	response, err := sendRequest(c, payload, LoadMasterResponse{})
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) DeleteOwaspCustomRule(filename string) (*LoadMasterResponse, error) {
	slog.Debug("Delete OWASP custom rule")
	payload := struct {
		*LoadMasterRequest
		Filename string `json:"filename"`
	}{
		&LoadMasterRequest{
			Command: "delowaspcustomrule",
		},
		filename,
	}

	response, err := sendRequest(c, payload, LoadMasterResponse{})
	if err != nil {
		return nil, err
	}


	return response, nil
}

func (c *Client) ShowOwaspCustomRule(filename string) (*LoadMasterDataResponse, error) {
	slog.Debug("Show OWASP custom rule")
	payload := struct {
		*LoadMasterRequest
		Filename string `json:"filename"`
	}{
		&LoadMasterRequest{
			Command: "downloadowaspcustomrule",
		},
		filename,
	}

	response, err := sendRequest(c, payload, LoadMasterDataResponse{})
	if err != nil {
		return nil, err
	}


	return response, nil
}

// AddOwaspCustomData adds an OWASP custom data file to the LoadMaster.
// The filename argument should be with file extension.
func (c *Client) AddOwaspCustomData(filename string, data string) (*LoadMasterResponse, error) {
	slog.Debug("Adding OWASP custom data")
	payload := struct {
		*LoadMasterRequest
		Filename string `json:"filename"`
		Data     string `json:"data"`
	}{
		&LoadMasterRequest{
			Command: "addowaspcustomdata",
		},
		filename,
		data,
	}

	response, err := sendRequest(c, payload, LoadMasterResponse{})
	if err != nil {
		return nil, err
	}


	return response, nil
}

// DeleteOwaspCustomData deletes an OWASP custom data file from the LoadMaster.
// The filename argument should be without file extension.
// For example, if the file is named "test_data.txt", you should pass "test_data" as the filename.
func (c *Client) DeleteOwaspCustomData(filename string) (*LoadMasterResponse, error) {
	slog.Debug("Delete OWASP custom data")
	payload := struct {
		*LoadMasterRequest
		Filename string `json:"filename"`
	}{
		&LoadMasterRequest{
			Command: "delowaspcustomdata",
		},
		filename,
	}

	response, err := sendRequest(c, payload, LoadMasterResponse{})
	if err != nil {
		return nil, err
	}


	return response, nil
}

func (c *Client) ShowOwaspCustomData(filename string) (*LoadMasterDataResponse, error) {
	slog.Debug("Show OWASP custom data")
	payload := struct {
		*LoadMasterRequest
		Filename string `json:"filename"`
	}{
		&LoadMasterRequest{
			Command: "downloadowaspcustomdata",
		},
		filename,
	}

	response, err := sendRequest(c, payload, LoadMasterDataResponse{})
	if err != nil {
		return nil, err
	}


	return response, nil
}

func (c *Client) AddVirtualServiceOwaspCustomRule(vs_identifier string, rule string, run_first bool) (*LoadMasterResponse, error) {
	slog.Debug("Add OWASP custom rule to virtual service")
	run_first_str := "0"
	if run_first {
		run_first_str = "1"
	}
	payload := struct {
		*LoadMasterRequest
		VS       string `json:"vs"`
		Rule     string `json:"rule"`
		Enable   string `json:"enable"`
		RunFirst string `json:"runfirst"`
	}{
		&LoadMasterRequest{
			Command: "owasprules",
		},
		vs_identifier,
		rule,
		"yes",
		run_first_str,
	}

	response, err := sendRequest(c, payload, LoadMasterResponse{})
	if err != nil {
		return nil, err
	}


	return response, nil
}

func (c *Client) DeleteVirtualServiceOwaspCustomRule(vs_identifier string, rule string) (*LoadMasterResponse, error) {
	slog.Debug("Delete OWASP custom rule to virtual service")
	payload := struct {
		*LoadMasterRequest
		VS     string `json:"vs"`
		Rule   string `json:"rule"`
		Enable string `json:"enable"`
	}{
		&LoadMasterRequest{
			Command: "owasprules",
		},
		vs_identifier,
		rule,
		"no",
	}

	response, err := sendRequest(c, payload, LoadMasterResponse{})
	if err != nil {
		return nil, err
	}


	return response, nil
}

func (c *Client) ShowVirtualServiceOwaspRule(vs_identifier string, rule string) (*OwaspRuleResponse, error) {
	slog.Debug("Show OWASP custom rule to virtual service")
	payload := struct {
		*LoadMasterRequest
		VS   string `json:"vs"`
		Rule string `json:"rule"`
	}{
		&LoadMasterRequest{
			Command: "owasprules",
		},
		vs_identifier,
		rule,
	}

	response, err := sendRequest(c, payload, OwaspRuleResponse{})
	if err != nil {
		return nil, err
	}


	return response, nil
}

func (c *Client) AddVirtualServiceOwaspRule(vs_identifier string, rule string) (*LoadMasterResponse, error) {
	slog.Debug("Add OWASP rule to virtual service")
	payload := struct {
		*LoadMasterRequest
		VS     string `json:"vs"`
		Rule   string `json:"rule"`
		Enable string `json:"enable"`
	}{
		&LoadMasterRequest{
			Command: "owasprules",
		},
		vs_identifier,
		rule,
		"yes",
	}

	response, err := sendRequest(c, payload, LoadMasterResponse{})
	if err != nil {
		return nil, err
	}


	return response, nil
}

func (c *Client) DeleteVirtualServiceOwaspRule(vs_identifier string, rule string) (*LoadMasterResponse, error) {
	slog.Debug("Delete OWASP rule to virtual service")
	payload := struct {
		*LoadMasterRequest
		VS     string `json:"vs"`
		Rule   string `json:"rule"`
		Enable string `json:"enable"`
	}{
		&LoadMasterRequest{
			Command: "owasprules",
		},
		vs_identifier,
		rule,
		"no",
	}

	response, err := sendRequest(c, payload, LoadMasterResponse{})
	if err != nil {
		return nil, err
	}


	return response, nil
}

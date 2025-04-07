package api

import "log/slog"

func (c *Client) AddRealServerRule(rs_index string, rule_name string) (*LoadMasterResponse, error) {
	slog.Debug("Adding real server rule", "rs_index", rs_index, "rule_name", rule_name)
	payload := struct {
		*LoadMasterRequest
		Name       string `json:"name"`
		RealServer string `json:"rs"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "addrsrule",
		},
		Name:       rule_name,
		RealServer: rs_index,
	}

	response, err := sendRequest(c, payload, LoadMasterResponse{})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) DeleteRealServerRule(rs_index string, rule_name string) (*LoadMasterResponse, error) {
	slog.Debug("Deleting real server rule", "rs_index", rs_index, "rule_name", rule_name)
	payload := struct {
		*LoadMasterRequest
		Name       string `json:"name"`
		RealServer string `json:"rs"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "delrsrule",
		},
		Name:       rule_name,
		RealServer: rs_index,
	}

	response, err := sendRequest(c, payload, LoadMasterResponse{})

	if err != nil {
		return nil, err
	}

	return response, nil
}

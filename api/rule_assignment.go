package api

func (c *Client) AddRealServerRule(rs_index string, rule_name string) (*LoadMasterResponse, error) {
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

	response, err := sendRequest(c, payload, &LoadMasterResponse{})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) DeleteRealServerRule(rs_index string, rule_name string) (*LoadMasterResponse, error) {
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

	response, err := sendRequest(c, payload, &LoadMasterResponse{})

	if err != nil {
		return nil, err
	}

	return response, nil
}

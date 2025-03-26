package api

type ListRealServerResponse struct {
	*LoadMasterResponse
	Rs []RealServer `json:"Rs"`
}

type RealServer struct {
	VSIndex int `json:"VSIndex,omitempty"`
	RsIndex int `json:"RSIndex,omitempty"`
}

type RealServerParameters struct {
	Address   string `json:"Addr,omitempty"`
	Port      int    `json:"Port,omitempty"`
	DnsName   string `json:"DnsName,omitempty"`
	Forward   string `json:"Forward,omitempty"`
	Weight    int    `json:"Weight,omitempty"`
	Limit     int    `json:"Limit,omitempty"`
	RateLimit int    `json:"RateLimit,omitempty"`
	Follow    int    `json:"Follow,omitempty"`
	Enable    *bool  `json:"Enable,omitempty"`
	Critical  *bool  `json:"Critical,omitempty"`
	Nrules    int    `json:"Nrules,omitempty"`
}

func (c *Client) AddRealServer(vs_identifier string, params RealServerParameters) (*ListRealServerResponse, error) {
	payload := struct {
		*LoadMasterRequest
		*RealServerParameters
		VS string `json:"vs"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "addrs",
		},
		RealServerParameters: &params,
		VS:                   vs_identifier,
	}

	response, err := sendRequest(c, payload, ListRealServerResponse{})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ModifyRealServer(vs_identifier string, params RealServerParameters) (*ListRealServerResponse, error) {
	payload := struct {
		*LoadMasterRequest
		*RealServerParameters
		VS string `json:"vs"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "modrs",
		},
		RealServerParameters: &params,
		VS:                   vs_identifier,
	}

	response, err := sendRequest(c, payload, ListRealServerResponse{})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) DeleteRealServer(vs_identifier string, params RealServerParameters) (*ListRealServerResponse, error) {
	payload := struct {
		*LoadMasterRequest
		*RealServerParameters
		VS string `json:"vs"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "modrs",
		},
		RealServerParameters: &params,
		VS:                   vs_identifier,
	}

	response, err := sendRequest(c, payload, ListRealServerResponse{})

	if err != nil {
		return nil, err
	}

	return response, nil
}

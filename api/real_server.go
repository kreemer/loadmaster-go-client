package api

import "log/slog"

type ListRealServerResponse struct {
	*LoadMasterResponse
	Rs []RealServer `json:"Rs"`
}

type RealServer struct {
	VSIndex    int      `json:"VSIndex,omitempty"`
	RsIndex    int      `json:"RSIndex,omitempty"`
	Address    string   `json:"Addr,omitempty"`
	Port       int      `json:"Port,omitempty"`
	DnsName    string   `json:"DnsName,omitempty"`
	Forward    string   `json:"Forward,omitempty"`
	Weight     int      `json:"Weight,omitempty"`
	Limit      int      `json:"Limit,omitempty"`
	RateLimit  int      `json:"RateLimit,omitempty"`
	Follow     int      `json:"Follow,omitempty"`
	Enable     *bool    `json:"Enable,omitempty"`
	Critical   *bool    `json:"Critical,omitempty"`
	Nrules     int      `json:"Nrules,omitempty"`
	MatchRules []string `json:"MatchRules,omitempty"`
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

func (c *Client) AddRealServer(vs_identifier string, address string, port string, params RealServerParameters) (*ListRealServerResponse, error) {
	slog.Debug("Adding real server", "vs_identifier", vs_identifier, "address", address, "port", port)
	payload := struct {
		*LoadMasterRequest
		*RealServerParameters
		VS        string `json:"vs"`
		RSAddress string `json:"rs"`
		RSPort    string `json:"rsport"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "addrs",
		},
		RealServerParameters: &params,
		VS:                   vs_identifier,
		RSAddress:            address,
		RSPort:               port,
	}

	response, err := sendRequest(c, payload, ListRealServerResponse{})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ShowRealServer(vs_identifier string, rs_identifier string) (*ListRealServerResponse, error) {
	slog.Debug("Showing real server", "vs_identifier", vs_identifier, "rs_identifier", rs_identifier)
	payload := struct {
		*LoadMasterRequest
		*RealServerParameters
		VS string `json:"vs"`
		RS string `json:"rs"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "showrs",
		},
		VS: vs_identifier,
		RS: rs_identifier,
	}

	response, err := sendRequest(c, payload, ListRealServerResponse{})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ModifyRealServer(vs_identifier string, rs_identifier string, params RealServerParameters) (*ListRealServerResponse, error) {
	slog.Debug("Modifying real server", "vs_identifier", vs_identifier, "rs_identifier", rs_identifier)
	payload := struct {
		*LoadMasterRequest
		*RealServerParameters
		VS string `json:"vs"`
		RS string `json:"rs"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "modrs",
		},
		RealServerParameters: &params,
		VS:                   vs_identifier,
		RS:                   rs_identifier,
	}

	response, err := sendRequest(c, payload, ListRealServerResponse{})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) DeleteRealServer(vs_identifier string, rs_identifier string) (*ListRealServerResponse, error) {
	slog.Debug("Deleting real server", "vs_identifier", vs_identifier, "rs_identifier", rs_identifier)
	payload := struct {
		*LoadMasterRequest
		*RealServerParameters
		VS string `json:"vs"`
		RS string `json:"rs"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "delrs",
		},
		VS: vs_identifier,
		RS: rs_identifier,
	}

	response, err := sendRequest(c, payload, ListRealServerResponse{})

	if err != nil {
		return nil, err
	}

	return response, nil
}

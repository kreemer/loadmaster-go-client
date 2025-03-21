package api

import (
	"encoding/json"
	"fmt"
)

type ListRealServerResponse struct {
	*LoadMasterResponse
	Rs []RealServer `json:"Rs"`
}

type RealServer struct {
	VSIndex int `json:"VSIndex"`
	RsIndex int `json:"RSIndex"`
}

type RealServerParameters struct {
	Address   string `json:"Addr"`
	Port      int    `json:"Port"`
	DnsName   string `json:"DnsName"`
	Forward   string `json:"Forward"`
	Weight    int    `json:"Weight"`
	Limit     int    `json:"Limit"`
	RateLimit int    `json:"RateLimit"`
	Follow    int    `json:"Follow"`
	Enable    bool   `json:"Enable"`
	Critical  bool   `json:"Critical"`
	Nrules    int    `json:"Nrules"`
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

	http, err := c.newRequest(payload)
	if err != nil {
		return nil, err
	}

	http_response, err := c.doRequest(http)
	if err != nil {
		return nil, err
	}

	response := &ListRealServerResponse{}
	err = json.Unmarshal(http_response, response)
	if err != nil {
		return nil, err
	}

	if response.Code >= 400 {
		return nil, fmt.Errorf("error: %s", response.Message)
	}

	return response, nil
}

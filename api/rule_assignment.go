package api

import (
	"fmt"
	"log/slog"
	"slices"
)

func (c *Client) AddRealServerRule(vs_identifier string, rs_index string, rule_name string) (*LoadMasterResponse, error) {
	slog.Debug("Adding real server rule", "vs_identifier", vs_identifier, "rs_index", rs_index, "rule_name", rule_name)
	payload := struct {
		*LoadMasterRequest
		Name           string `json:"rule"`
		RealServer     string `json:"rs"`
		VirtualService string `json:"vs"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "addrsrule",
		},
		Name:           rule_name,
		RealServer:     rs_index,
		VirtualService: vs_identifier,
	}

	response, err := sendRequest(c, payload, LoadMasterResponse{})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ShowRealServerRule(vs_identifier string, rs_index string, rule_name string) (*LoadMasterResponse, error) {
	slog.Debug("Showing real server rule", "vs_identifier", vs_identifier, "rs_index", rs_index, "rule_name", rule_name)
	payload := struct {
		*LoadMasterRequest
		RealServer     string `json:"rs"`
		VirtualService string `json:"vs"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "showrs",
		},
		RealServer:     rs_index,
		VirtualService: vs_identifier,
	}

	response, err := sendRequest(c, payload, ListRealServerResponse{})

	if err != nil {
		return nil, err
	}

	rules := response.Rs[len(response.Rs)-1].MatchRules

	if !slices.Contains(rules, rule_name) {
		return nil, fmt.Errorf("rule %s not found in real server %s", rule_name, rs_index)
	}

	return &LoadMasterResponse{response.Code, response.Message, response.Status}, nil
}

func (c *Client) DeleteRealServerRule(vs_identifier string, rs_index string, rule_name string) (*LoadMasterResponse, error) {
	slog.Debug("Deleting real server rule", "vs_identifier", vs_identifier, "rs_index", rs_index, "rule_name", rule_name)
	payload := struct {
		*LoadMasterRequest
		Name           string `json:"rule"`
		RealServer     string `json:"rs"`
		VirtualService string `json:"vs"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "delrsrule",
		},
		Name:           rule_name,
		RealServer:     rs_index,
		VirtualService: vs_identifier,
	}

	response, err := sendRequest(c, payload, LoadMasterResponse{})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) AddSubVirtualServiceRule(vs_identifier string, subvs_identifier string, rule_name string) (*LoadMasterResponse, error) {
	slog.Debug("Adding sub virtual service rule", "vs_identifier", vs_identifier, "subvs_identifier", subvs_identifier, "rule_name", rule_name)
	payload := struct {
		*LoadMasterRequest
		Name              string `json:"rule"`
		SubVirtualService string `json:"rs"`
		VirtualService    string `json:"vs"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "addrsrule",
		},
		Name:              rule_name,
		SubVirtualService: subvs_identifier,
		VirtualService:    vs_identifier,
	}

	response, err := sendRequest(c, payload, LoadMasterResponse{})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ShowSubVirtualServiceRule(vs_identifier string, subvs_identifier string, rule_name string) (*LoadMasterResponse, error) {
	slog.Debug("Showing real server rule", "vs_identifier", vs_identifier, "subvs_identifier", subvs_identifier, "rule_name", rule_name)
	payload := struct {
		*LoadMasterRequest
		SubVirtualService string `json:"rs"`
		VirtualService    string `json:"vs"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "showrs",
		},
		SubVirtualService: subvs_identifier,
		VirtualService:    vs_identifier,
	}

	response, err := sendRequest(c, payload, ShowSubVirtualServiceResponse{})

	if err != nil {
		return nil, err
	}

	rules := response.SubVS[len(response.SubVS)-1].MatchRules

	if !slices.Contains(rules, rule_name) {
		return nil, fmt.Errorf("rule %s not found in sub virtual service %s", rule_name, subvs_identifier)
	}

	return &LoadMasterResponse{response.Code, response.Message, response.Status}, nil
}

func (c *Client) DeleteSubVirtualServiceRule(vs_identifier string, subvs_identifier string, rule_name string) (*LoadMasterResponse, error) {
	slog.Debug("Deleting sub virtual service rule", "vs_identifier", vs_identifier, "subvs_identifier", subvs_identifier, "rule_name", rule_name)
	payload := struct {
		*LoadMasterRequest
		Name              string `json:"rule"`
		SubVirtualService string `json:"rs"`
		VirtualService    string `json:"vs"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "delrsrule",
		},
		Name:              rule_name,
		SubVirtualService: subvs_identifier,
		VirtualService:    vs_identifier,
	}

	response, err := sendRequest(c, payload, LoadMasterResponse{})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) AddVirtualServicePreRule(vs_identifier string, rule_name string) (*LoadMasterResponse, error) {
	slog.Debug("Adding virtual service pre rule", "vs_identifier", vs_identifier, "rule_name", rule_name)
	payload := struct {
		*LoadMasterRequest
		Name           string `json:"rule"`
		VirtualService string `json:"vs"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "addprerule",
		},
		Name:           rule_name,
		VirtualService: vs_identifier,
	}

	response, err := sendRequest(c, payload, LoadMasterResponse{})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ShowVirtualServicePreRule(vs_identifier string, rule_name string) (*LoadMasterResponse, error) {
	slog.Debug("Showing virtual service pre rule", "vs_identifier", vs_identifier, "rule_name", rule_name)
	payload := struct {
		*LoadMasterRequest
		VirtualService string `json:"vs"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "showvs",
		},
		VirtualService: vs_identifier,
	}

	response, err := sendRequest(c, payload, ShowVirtualServiceResponse{})

	if err != nil {
		return nil, err
	}

	rules := response.MatchRules
	if !slices.Contains(rules, rule_name) {
		return nil, fmt.Errorf("rule %s not found in virtual service %s", rule_name, vs_identifier)
	}

	return &LoadMasterResponse{response.Code, response.Message, response.Status}, nil
}

func (c *Client) DeleteVirtualServicePreRule(vs_identifier string, rule_name string) (*LoadMasterResponse, error) {
	slog.Debug("Deleting virtual service pre rule", "vs_identifier", vs_identifier, "rule_name", rule_name)
	payload := struct {
		*LoadMasterRequest
		Name           string `json:"rule"`
		VirtualService string `json:"vs"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "delprerule",
		},
		Name:           rule_name,
		VirtualService: vs_identifier,
	}

	response, err := sendRequest(c, payload, LoadMasterResponse{})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) AddVirtualServiceRequestRule(vs_identifier string, rule_name string) (*LoadMasterResponse, error) {
	slog.Debug("Adding virtual service request rule", "vs_identifier", vs_identifier, "rule_name", rule_name)
	payload := struct {
		*LoadMasterRequest
		Name           string `json:"rule"`
		VirtualService string `json:"vs"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "addrequestrule",
		},
		Name:           rule_name,
		VirtualService: vs_identifier,
	}

	response, err := sendRequest(c, payload, LoadMasterResponse{})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ShowVirtualServiceRequestRule(vs_identifier string, rule_name string) (*LoadMasterResponse, error) {
	slog.Debug("Showing virtual service request rule", "vs_identifier", vs_identifier, "rule_name", rule_name)
	payload := struct {
		*LoadMasterRequest
		VirtualService string `json:"vs"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "showvs",
		},
		VirtualService: vs_identifier,
	}

	response, err := sendRequest(c, payload, ShowVirtualServiceResponse{})

	if err != nil {
		return nil, err
	}

	rules := response.RequestRules
	if !slices.Contains(rules, rule_name) {
		return nil, fmt.Errorf("rule %s not found in virtual service %s", rule_name, vs_identifier)
	}

	return &LoadMasterResponse{response.Code, response.Message, response.Status}, nil
}

func (c *Client) DeleteVirtualServiceRequestRule(vs_identifier string, rule_name string) (*LoadMasterResponse, error) {
	slog.Debug("Deleting virtual service request rule", "vs_identifier", vs_identifier, "rule_name", rule_name)
	payload := struct {
		*LoadMasterRequest
		Name           string `json:"rule"`
		VirtualService string `json:"vs"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "delrequestrule",
		},
		Name:           rule_name,
		VirtualService: vs_identifier,
	}

	response, err := sendRequest(c, payload, LoadMasterResponse{})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) AddVirtualServiceResponseRule(vs_identifier string, rule_name string) (*LoadMasterResponse, error) {
	slog.Debug("Adding virtual service response rule", "vs_identifier", vs_identifier, "rule_name", rule_name)
	payload := struct {
		*LoadMasterRequest
		Name           string `json:"rule"`
		VirtualService string `json:"vs"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "addresponserule",
		},
		Name:           rule_name,
		VirtualService: vs_identifier,
	}

	response, err := sendRequest(c, payload, LoadMasterResponse{})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ShowVirtualServiceResponseRule(vs_identifier string, rule_name string) (*LoadMasterResponse, error) {
	slog.Debug("Showing virtual service response rule", "vs_identifier", vs_identifier, "rule_name", rule_name)
	payload := struct {
		*LoadMasterRequest
		VirtualService string `json:"vs"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "showvs",
		},
		VirtualService: vs_identifier,
	}

	response, err := sendRequest(c, payload, ShowVirtualServiceResponse{})

	if err != nil {
		return nil, err
	}

	rules := response.ResponseRules
	if !slices.Contains(rules, rule_name) {
		return nil, fmt.Errorf("rule %s not found in virtual service %s", rule_name, vs_identifier)
	}

	return &LoadMasterResponse{response.Code, response.Message, response.Status}, nil
}

func (c *Client) DeleteVirtualServiceResponseRule(vs_identifier string, rule_name string) (*LoadMasterResponse, error) {
	slog.Debug("Deleting virtual service response rule", "vs_identifier", vs_identifier, "rule_name", rule_name)
	payload := struct {
		*LoadMasterRequest
		Name           string `json:"rule"`
		VirtualService string `json:"vs"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "delresponserule",
		},
		Name:           rule_name,
		VirtualService: vs_identifier,
	}

	response, err := sendRequest(c, payload, LoadMasterResponse{})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) AddVirtualServiceResponseBodyRule(vs_identifier string, rule_name string) (*LoadMasterResponse, error) {
	slog.Debug("Adding virtual service response body rule", "vs_identifier", vs_identifier, "rule_name", rule_name)
	payload := struct {
		*LoadMasterRequest
		Name           string `json:"rule"`
		VirtualService string `json:"vs"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "addresponsebodyrule",
		},
		Name:           rule_name,
		VirtualService: vs_identifier,
	}

	response, err := sendRequest(c, payload, LoadMasterResponse{})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ShowVirtualServiceResponseBodyRule(vs_identifier string, rule_name string) (*LoadMasterResponse, error) {
	slog.Debug("Showing virtual service response body rule", "vs_identifier", vs_identifier, "rule_name", rule_name)
	payload := struct {
		*LoadMasterRequest
		VirtualService string `json:"vs"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "showvs",
		},
		VirtualService: vs_identifier,
	}

	response, err := sendRequest(c, payload, ShowVirtualServiceResponse{})

	if err != nil {
		return nil, err
	}

	rules := response.MatchBodyRules
	if !slices.Contains(rules, rule_name) {
		return nil, fmt.Errorf("rule %s not found in virtual service %s", rule_name, vs_identifier)
	}

	return &LoadMasterResponse{response.Code, response.Message, response.Status}, nil
}

func (c *Client) DeleteVirtualServiceResponseBodyRule(vs_identifier string, rule_name string) (*LoadMasterResponse, error) {
	slog.Debug("Deleting virtual service response body rule", "vs_identifier", vs_identifier, "rule_name", rule_name)
	payload := struct {
		*LoadMasterRequest
		Name           string `json:"rule"`
		VirtualService string `json:"vs"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "delresponsebodyrule",
		},
		Name:           rule_name,
		VirtualService: vs_identifier,
	}

	response, err := sendRequest(c, payload, LoadMasterResponse{})

	if err != nil {
		return nil, err
	}

	return response, nil
}

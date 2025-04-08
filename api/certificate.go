package api

import (
	"encoding/json"
	"fmt"
	"log/slog"
)

type ListCertResponse struct {
	*LoadMasterResponse
	Cert []CertInfo `json:"cert"`
}

type CertInfo struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Modulus string `json:"modulus"`
}
type ShowCertResponse struct {
	*LoadMasterResponse
	Data string `json:"certificate"`
}

func (c *Client) ListCertificate() (*ListCertResponse, error) {
	slog.Debug("Listing certificates")
	payload := struct {
		*LoadMasterRequest
	}{
		&LoadMasterRequest{
			Command: "listcert",
		},
	}

	http, err := c.newRequest(payload)
	if err != nil {
		return nil, err
	}

	http_response, err := c.doRequest(http)
	if err != nil {
		return nil, err
	}

	response := &ListCertResponse{}
	err = json.Unmarshal(http_response, response)
	if err != nil {
		return nil, err
	}

	if response.Code >= 400 {
		return nil, fmt.Errorf("error: %s", response.Message)
	}

	return response, nil
}

func (c *Client) ListIntermediateCertificate() (*ListCertResponse, error) {
	slog.Debug("Listing intermediate certificates")
	payload := struct {
		*LoadMasterRequest
	}{
		&LoadMasterRequest{
			Command: "listintermediate",
		},
	}
	http, err := c.newRequest(payload)
	if err != nil {
		return nil, err
	}

	http_response, err := c.doRequest(http)
	if err != nil {
		return nil, err
	}

	response := &ListCertResponse{}
	err = json.Unmarshal(http_response, response)
	if err != nil {
		return nil, err
	}

	if response.Code >= 400 {
		return nil, fmt.Errorf("error: %s", response.Message)
	}

	return response, nil
}

func (c *Client) ShowCertificate(name string) (*ShowCertResponse, error) {
	slog.Debug("Show certificates")
	payload := struct {
		*LoadMasterRequest
		Cert string `json:"cert"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "readcert",
		},
		Cert: name,
	}

	http, err := c.newRequest(payload)
	if err != nil {
		return nil, err
	}

	http_response, err := c.doRequest(http)
	if err != nil {
		return nil, err
	}

	response := &ShowCertResponse{}
	err = json.Unmarshal(http_response, response)
	if err != nil {
		return nil, err
	}

	if response.Code >= 400 {
		return nil, fmt.Errorf("error: %s", response.Message)
	}

	return response, nil
}

func (c *Client) ShowIntermediateCertificate(name string) (*ShowCertResponse, error) {
	slog.Debug("Show intermediate certificates")
	payload := struct {
		*LoadMasterRequest
		Cert string `json:"cert"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "readintermediate",
		},
		Cert: name,
	}

	http, err := c.newRequest(payload)
	if err != nil {
		return nil, err
	}

	http_response, err := c.doRequest(http)
	if err != nil {
		return nil, err
	}

	response := &ShowCertResponse{}
	err = json.Unmarshal(http_response, response)
	if err != nil {
		return nil, err
	}

	if response.Code >= 400 {
		return nil, fmt.Errorf("error: %s", response.Message)
	}

	return response, nil
}

func (c *Client) AddCertificate(name string, password *string, data string) (*LoadMasterResponse, error) {
	slog.Debug("Show certificates")
	payload := struct {
		*LoadMasterRequest
		Cert     string  `json:"cert"`
		Data     string  `json:"data"`
		Password *string `json:"password,omitempty"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "addcert",
		},
		Cert:     name,
		Data:     data,
		Password: password,
	}

	http, err := c.newRequest(payload)
	if err != nil {
		return nil, err
	}

	http_response, err := c.doRequest(http)
	if err != nil {
		return nil, err
	}

	response := &LoadMasterResponse{}
	err = json.Unmarshal(http_response, response)
	if err != nil {
		return nil, err
	}

	if response.Code >= 400 {
		return nil, fmt.Errorf("error: %s", response.Message)
	}

	return response, nil
}

func (c *Client) AddIntermediateCertificate(name string, data string) (*LoadMasterResponse, error) {
	slog.Debug("Add intermediate certificates")
	payload := struct {
		*LoadMasterRequest
		Cert string `json:"cert"`
		Data string `json:"data"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "addintermediate",
		},
		Cert: name,
		Data: data,
	}

	http, err := c.newRequest(payload)
	if err != nil {
		return nil, err
	}

	http_response, err := c.doRequest(http)
	if err != nil {
		return nil, err
	}

	response := &LoadMasterResponse{}
	err = json.Unmarshal(http_response, response)
	if err != nil {
		return nil, err
	}

	if response.Code >= 400 {
		return nil, fmt.Errorf("error: %s", response.Message)
	}

	return response, nil
}

func (c *Client) DeleteCertificate(name string) (*LoadMasterResponse, error) {
	slog.Debug("Show certificates")
	payload := struct {
		*LoadMasterRequest
		Cert     string  `json:"cert"`
		Data     string  `json:"data"`
		Password *string `json:"password,omitempty"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "delcert",
		},
		Cert: name,
	}

	http, err := c.newRequest(payload)
	if err != nil {
		return nil, err
	}

	http_response, err := c.doRequest(http)
	if err != nil {
		return nil, err
	}

	response := &LoadMasterResponse{}
	err = json.Unmarshal(http_response, response)
	if err != nil {
		return nil, err
	}

	if response.Code >= 400 {
		return nil, fmt.Errorf("error: %s", response.Message)
	}

	return response, nil
}

func (c *Client) DeleteIntermediateCertificate(name string) (*LoadMasterResponse, error) {
	slog.Debug("Add intermediate certificates")
	payload := struct {
		*LoadMasterRequest
		Cert string `json:"cert"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "delintermediate",
		},
		Cert: name,
	}

	http, err := c.newRequest(payload)
	if err != nil {
		return nil, err
	}

	http_response, err := c.doRequest(http)
	if err != nil {
		return nil, err
	}

	response := &LoadMasterResponse{}
	err = json.Unmarshal(http_response, response)
	if err != nil {
		return nil, err
	}

	if response.Code >= 400 {
		return nil, fmt.Errorf("error: %s", response.Message)
	}

	return response, nil
}

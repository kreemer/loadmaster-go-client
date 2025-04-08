package api

import (
	"encoding/json"
	"fmt"
	"log/slog"
)

type RequestACMECertificateParameters struct {
	Country      string `json:"country,omitempty"`
	State        string `json:"state,omitempty"`
	City         string `json:"city,omitempty"`
	Company      string `json:"company,omitempty"`
	Organization string `json:"organization,omitempty"`
	Email        string `json:"email,omitempty"`
	KeySize      int    `json:"key_size,omitempty"`

	DnsApi       string `json:"dnsapi,omitempty"`
	DnsApiParams string `json:"dnsapiparams,omitempty"`
}

func (c *Client) RegisterLetsEncryptAccount(email *string) (*LoadMasterResponse, error) {
	slog.Debug("Register Lets Encrypt account")
	payload := struct {
		*LoadMasterRequest
		Email *string `json:"email,omitempty"`
		Type  string  `json:"acmetype"`
	}{
		&LoadMasterRequest{
			Command: "registeracmeaccount",
		},
		email,
		"1",
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

func (c *Client) FetchLetsEncryptAccount(password string, data string) (*LoadMasterResponse, error) {
	slog.Debug("Fetching Lets Encrypt account")
	payload := struct {
		*LoadMasterRequest
		Password string `json:"password"`
		Data     string `json:"data"`
	}{
		&LoadMasterRequest{
			Command: "fetchleaccount",
		},
		password,
		data,
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

func (c *Client) SetDigicertKeyId(key_id string) (*LoadMasterResponse, error) {
	slog.Debug("Setting Digicert key ID")
	payload := struct {
		*LoadMasterRequest
		KeyId string `json:"kid"`
		Type  string `json:"acmetype"`
	}{
		&LoadMasterRequest{
			Command: "setacmekid",
		},
		key_id,
		"2",
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

func (c *Client) SetDigicertHMAC(hmac string) (*LoadMasterResponse, error) {
	slog.Debug("Setting Digicert HMAC")
	payload := struct {
		*LoadMasterRequest
		Hmac string `json:"hmac"`
		Type string `json:"acmetype"`
	}{
		&LoadMasterRequest{
			Command: "setacmehmac",
		},
		hmac,
		"2",
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

func (c *Client) RequestACMECertificate(name string, common_name string, vs_identifier string, acme_type string, params *RequestACMECertificateParameters) (*LoadMasterResponse, error) {
	slog.Debug("Request ACME Certificate", "name", name, "common_name", common_name, "vs_identifier", vs_identifier, "acme_type", acme_type)
	payload := struct {
		*LoadMasterRequest
		*RequestACMECertificateParameters
		Name       string `json:"cert"`
		CommonName string `json:"cn"`
		VS         string `json:"vid"`
		AcmeType   string `json:"acmetype"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "addacmecert",
		},
		RequestACMECertificateParameters: params,
		Name:                             name,
		CommonName:                       common_name,
		VS:                               vs_identifier,
		AcmeType:                         acme_type,
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

func (c *Client) DeleteACMECertificate(name string, acme_type string) (*LoadMasterResponse, error) {
	slog.Debug("Delete ACME Certificate", "name", name, "acme_type", acme_type)
	payload := struct {
		*LoadMasterRequest
		Name     string `json:"cert"`
		AcmeType string `json:"acmetype"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "delacmecert",
		},
		Name:     name,
		AcmeType: acme_type,
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

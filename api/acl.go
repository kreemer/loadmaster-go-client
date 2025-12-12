package api

import (
	"encoding/json"
	"log/slog"
)

type ListAclResponse struct {
	*LoadMasterResponse
	List string           `json:"list"`
	IPs  []ListAclAddress `json:"IP"`
}

type ListAclAddress struct {
	Address string `json:"addr"`
	Comment string `json:"comment"`
}

func (c *Client) AddGlobalAclAllow(ip_addr string) (*LoadMasterResponse, error) {
	return c.aclGlobal("allow", "add", ip_addr)
}

func (c *Client) DeleteGlobalAclAllow(ip_addr string) (*LoadMasterResponse, error) {
	return c.aclGlobal("allow", "del", ip_addr)
}

func (c *Client) ListGlobalAclAllow() (*ListAclResponse, error) {
	return c.aclGlobalList("allow")
}

func (c *Client) AddGlobalAclBlock(ip_addr string) (*LoadMasterResponse, error) {
	return c.aclGlobal("block", "add", ip_addr)
}

func (c *Client) DeleteGlobalAclBlock(ip_addr string) (*LoadMasterResponse, error) {
	return c.aclGlobal("block", "del", ip_addr)
}

func (c *Client) ListGlobalAclBlock() (*ListAclResponse, error) {
	return c.aclGlobalList("block")
}

func (c *Client) AddVirtualServiceAclAllow(vs_identifier string, ip_addr string) (*LoadMasterResponse, error) {
	return c.aclVirtualService("allow", "add", vs_identifier, ip_addr)
}

func (c *Client) DeleteVirtualServiceAclAllow(vs_identifier string, ip_addr string) (*LoadMasterResponse, error) {
	return c.aclVirtualService("allow", "del", vs_identifier, ip_addr)
}

func (c *Client) ListVirtualServiceAclAllow(vs_identifier string) (*ListAclResponse, error) {
	return c.aclVirtualServiceList("allow", vs_identifier)
}

func (c *Client) AddVirtualServiceAclBlock(vs_identifier string, ip_addr string) (*LoadMasterResponse, error) {
	return c.aclVirtualService("block", "add", vs_identifier, ip_addr)
}

func (c *Client) DeleteVirtualServiceAclBlock(vs_identifier string, ip_addr string) (*LoadMasterResponse, error) {
	return c.aclVirtualService("block", "del", vs_identifier, ip_addr)
}

func (c *Client) ListVirtualServiceAclBlock(vs_identifier string) (*ListAclResponse, error) {
	return c.aclVirtualServiceList("block", vs_identifier)
}

func (c *Client) aclGlobalList(allow_or_block string) (*ListAclResponse, error) {
	slog.Debug("List global acl allow address", "allow_or_block", allow_or_block)

	payload := struct {
		*LoadMasterRequest
		List    string `json:"list"`
		Address string `json:"addr"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "aclcontrol",
		},
		List: allow_or_block,
	}
	http, err := c.newRequest(payload)
	if err != nil {
		return nil, err
	}

	http_response, err := c.doRequest(http)
	if err != nil {
		return nil, err
	}

	response := &ListAclResponse{}
	err = json.Unmarshal(http_response, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) aclGlobal(allow_or_block string, add_or_delete string, ip_addr string) (*LoadMasterResponse, error) {
	slog.Debug("Modify global acl allow address", "allow_or_block", allow_or_block, "add_or_delete", add_or_delete, "ip_addr", ip_addr)

	var payload AuthInjectable

	if add_or_delete == "add" {
		payload = struct {
			*LoadMasterRequest
			Add     string `json:"add"`
			Address string `json:"addr"`
		}{
			LoadMasterRequest: &LoadMasterRequest{
				Command: "aclcontrol",
			},
			Add:     allow_or_block,
			Address: ip_addr,
		}
	} else {
		payload = struct {
			*LoadMasterRequest
			Del     string `json:"del"`
			Address string `json:"addr"`
		}{
			LoadMasterRequest: &LoadMasterRequest{
				Command: "aclcontrol",
			},
			Del:     allow_or_block,
			Address: ip_addr,
		}
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

	return response, nil
}

func (c *Client) aclVirtualServiceList(allow_or_block string, vs_identifier string) (*ListAclResponse, error) {
	slog.Debug("List global acl allow address", "allow_or_block", allow_or_block, "vs_identifier", vs_identifier)

	payload := struct {
		*LoadMasterRequest
		List string `json:"listvs"`
		VS   string `json:"vs"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "aclcontrol",
		},
		List: allow_or_block,
		VS:   vs_identifier,
	}
	http, err := c.newRequest(payload)
	if err != nil {
		return nil, err
	}

	http_response, err := c.doRequest(http)
	if err != nil {
		return nil, err
	}

	response := &ListAclResponse{}
	err = json.Unmarshal(http_response, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) aclVirtualService(allow_or_block string, add_or_delete string, vs_identifier string, ip_addr string) (*LoadMasterResponse, error) {
	slog.Debug("Modify virtual service acl allow address", "allow_or_block", allow_or_block, "add_or_delete", add_or_delete, "vs_identifier", vs_identifier, "ip_addr", ip_addr)

	var payload AuthInjectable

	if add_or_delete == "add" {
		payload = struct {
			*LoadMasterRequest
			Add     string `json:"addvs"`
			VS      string `json:"vsip"`
			Address string `json:"addr"`
		}{
			LoadMasterRequest: &LoadMasterRequest{
				Command: "aclcontrol",
			},
			Add:     allow_or_block,
			VS:      vs_identifier,
			Address: ip_addr,
		}
	} else {
		payload = struct {
			*LoadMasterRequest
			Del     string `json:"delvs"`
			VS      string `json:"vsip"`
			Address string `json:"addr"`
		}{
			LoadMasterRequest: &LoadMasterRequest{
				Command: "aclcontrol",
			},
			Del:     allow_or_block,
			VS:      vs_identifier,
			Address: ip_addr,
		}
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

	return response, nil
}

package api

import (
	"encoding/json"
	"fmt"
	"net/netip"
)

type ListVirtualServiceResponse struct {
	*LoadMasterResponse
	VS []struct {
		*VirtualServiceIdentifier
		*VirtualServiceParametersProperties
	}
}

type AddVirtualServiceRequest struct {
	*VirtualServiceIdentifier
	*VirtualServiceParametersProperties
}

type AddVirtualServiceResponse struct {
	*LoadMasterResponse
	*VirtualServiceIdentifier
	*VirtualServiceParametersProperties
}

type ModifyVirtualServiceRequest struct {
	*VirtualServiceIdentifier
	*VirtualServiceParametersProperties
}

type ModifyVirtualServiceResponse struct {
	*LoadMasterResponse
	*VirtualServiceIdentifier
	*VirtualServiceParametersProperties
}

type ShowVirtualServiceRequest struct {
}

type ShowVirtualServiceResponse struct {
	*LoadMasterResponse
	*VirtualServiceParametersProperties
}

type DeleteVirtualServiceRequest struct {
}

type DeleteVirtualServiceResponse struct {
	*LoadMasterResponse
}

type DuplicateVirtualServiceRequest struct {
}

type DuplicateVirtualServiceResponse struct {
	*LoadMasterResponse
}

type VirtualServiceIdentifier struct {
	VS         netip.Addr `json:"vs,omitempty"`
	Port       int        `json:"port,omitempty"`
	VSProtocol string     `json:"prot,omitempty"`
}

type VirtualServiceParametersProperties struct {
	VSPort    netip.AddrPort `json:"vsport,omitempty"`
	Protocol  string         `json:"protocol,omitempty"`
	VSAddress netip.Addr     `json:"vsaddress,omitempty"`
}

type VirtualServiceParametersBasicProperties struct {
	Enable   bool   `json:"enable,omitempty"`
	VSType   string `json:"VStype,omitempty"`
	NickName string `json:"NickName,omitempty"`
}

type VirtualServiceParametersStandardOptions struct {
	Cookie             string         `json:"Cookie,omitempty"`
	ForceL7            bool           `json:"ForceL7,omitempty"`
	Idletime           int            `json:"Idletime,omitempty"`
	Persist            string         `json:"Persist,omitempty"`
	SubnetOriginating  bool           `json:"SubnetOriginating,omitempty"`
	PersistTimeout     int            `json:"PersistTimeout,omitempty"`
	Refreshpersist     bool           `json:"Refreshpersist,omitempty"`
	QueryTag           string         `json:"QueryTag,omitempty"`
	Schedule           string         `json:"Schedule,omitempty"`
	Showadaptive       string         `json:"showadaptive,omitempty"`
	AdaptiveInterval   int            `json:"AdaptiveInterval,omitempty"`
	AdaptiveUrl        string         `json:"AdaptiveUrl,omitempty"`
	AdaptivePort       netip.AddrPort `json:"AdaptivePort,omitempty"`
	AdaptiveMinPercent int            `json:"AdaptiveMinPercent,omitempty"`
	ServerInit         int            `json:"ServerInit,omitempty"`
	Transparent        bool           `json:"Transparent,omitempty"`
	UseForSnat         bool           `json:"UseForSnat,omitempty"`
	QoS                string         `json:"QoS,omitempty"`
	StartTLSMode       int            `json:"StartTLSMode,omitempty"`
	ExtraPorts         string         `json:"ExtraPorts,omitempty"`
}

type VirtualServiceParametersSSLProperties struct {
	CertFile              string `json:"CertFile,omitempty"`
	Ciphers               string `json:"Ciphers,omitempty"`
	CipherSet             string `json:"CipherSet,omitempty"`
	Tls13CipherSet        string `json:"Tls13CipherSet,omitempty"`
	ClientCert            int    `json:"ClientCert,omitempty"`
	PassCipher            bool   `json:"PassCipher,omitempty"`
	SSLReencrypt          bool   `json:"SSLReencrypt,omitempty"`
	PassSNI               bool   `json:"PassSNI,omitempty"`
	SSLReverse            bool   `json:"SSLReverse,omitempty"`
	SSLRewrite            string `json:"SSLRewrite,omitempty"`
	ReverseSNIHostname    string `json:"ReverseSNIHostname,omitempty"`
	SecurityHeaderOptions int    `json:"SecurityHeaderOptions,omitempty"`
	SSLAcceleration       bool   `json:"SSLAcceleration,omitempty"`
	OCSPVerify            bool   `json:"OCSPVerify,omitempty"`
	TLSType               int    `json:"TLSType,omitempty"`
	NeedHostName          bool   `json:"NeedHostName,omitempty"`
	IntermediateCerts     string `json:"IntermediateCerts,omitempty"`
}

type VirtualServiceParametersAdvancedProperties struct {
	HTTPReschedule         bool       `json:"HTTPReschedule,omitempty"`
	CopyHdrFrom            string     `json:"CopyHdrFrom,omitempty"`
	CopyHdrTo              string     `json:"CopyHdrTo,omitempty"`
	AddVia                 int        `json:"AddVia,omitempty"`
	AllowHTTP2             bool       `json:"AllowHTTP2,omitempty"`
	Cache                  bool       `json:"Cache,omitempty"`
	Compress               bool       `json:"Compress,omitempty"`
	CachePercent           int        `json:"CachePercent,omitempty"`
	DefaultGW              netip.Addr `json:"DefaultGW,omitempty"`
	ErrorCode              int        `json:"ErrorCode,omitempty"`
	ErrorUrl               string     `json:"ErrorUrl,omitempty"`
	PortFollow             int        `json:"PortFollow,omitempty"`
	FollowVSID             int        `json:"FollowVSID,omitempty"`
	LocalBindAddrs         string     `json:"LocalBindAddrs,omitempty"`
	NRequestRules          int        `json:"NRequestRules,omitempty"`
	NResponseRules         int        `json:"NResponseRules,omitempty"`
	RequestRules           []string   `json:"RequestRules,omitempty"`
	ResponseRules          []string   `json:"ResponseRules,omitempty"`
	StandbyAddr            netip.Addr `json:"StandbyAddr,omitempty"`
	StandbyPort            int        `json:"StandbyPort,omitempty"`
	NonLocalSorryServer    bool       `json:"NonLocalSorryServer,omitempty"`
	Verify                 int        `json:"Verify,omitempty"`
	AltAddress             netip.Addr `json:"AltAddress,omitempty"`
	PreProcPrecedence      string     `json:"PreProcPrecedence,omitempty"`
	PreProcPrecedencePos   int        `json:"PreProcPrecedencePos,omitempty"`
	RequestPrecedence      string     `json:"RequestPrecedence,omitempty"`
	RequestPrecedencePos   int        `json:"RequestPrecedencePos,omitempty"`
	ResponsePrecedence     string     `json:"ResponsePrecedence,omitempty"`
	ResponsePrecedencePos  int        `json:"ResponsePrecedencePos,omitempty"`
	MatchBodyPrecedence    string     `json:"MatchBodyPrecedence,omitempty"`
	MatchBodyPrecedencePos int        `json:"MatchBodyPrecedencePos,omitempty"`
	ResponseStatusRemap    bool       `json:"ResponseStatusRemap,omitempty"`
	ResponseRemapMsgMap    string     `json:"ResponseRemapMsgMap,omitempty"`
	ResponseRemapMsgFormat bool       `json:"ResponseRemapMsgFormat,omitempty"`
	ResponseRemapCodeMap   string     `json:"ResponseRemapCodeMap,omitempty"`
}

func (c *Client) ListVirtualService() (*ListVirtualServiceResponse, error) {
	payload := struct {
		*LoadMasterRequest
	}{
		&LoadMasterRequest{
			Command: "listvs",
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

	response := &ListVirtualServiceResponse{}
	err = json.Unmarshal(http_response, response)
	if err != nil {
		return nil, err
	}

	if response.Code >= 400 {
		return nil, fmt.Errorf("error: %s", response.Message)
	}

	return response, nil
}

func (c *Client) AddVirtualService(request *AddVirtualServiceRequest) (*AddVirtualServiceResponse, error) {
	payload := struct {
		*LoadMasterRequest
		*AddVirtualServiceRequest
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "addvs",
		},
		AddVirtualServiceRequest: request,
	}

	http, err := c.newRequest(payload)
	if err != nil {
		return nil, err
	}

	http_response, err := c.doRequest(http)
	if err != nil {
		return nil, err
	}

	response := &AddVirtualServiceResponse{}
	err = json.Unmarshal(http_response, response)
	if err != nil {
		return nil, err
	}

	if response.Code >= 400 {
		return nil, fmt.Errorf("error: %s", response.Message)
	}

	return response, nil
}

package api

import "log/slog"

type VirtualService struct {
	Index          int32    `json:"Index"`
	Protocol       string   `json:"Protocol"`
	Address        string   `json:"VSAddress"`
	Port           string   `json:"VSPort"`
	MasterVS       *int32   `json:"MasterVS"`
	MasterVSID     int32    `json:"MasterVSID,omitempty"`
	MatchRules     []string `json:"MatchRules,omitempty"`
	MatchBodyRules []string `json:"MatchBodyRules,omitempty"`
	*VirtualServiceParameters
}

type VirtualServiceParameters struct {
	*VirtualServiceParametersBasicProperties
	*VirtualServiceParametersStandardOptions
	*VirtualServiceParametersSSLProperties
	*VirtualServiceParametersAdvancedProperties
	*VirtualServiceParametersWAFSettings
	*VirtualServiceParametersESPOptions
	*VirtualServiceParametersRealServers
	*VirtualServiceParametersMiscellaneous
}

type VirtualServiceParametersBasicProperties struct {
	Enable   *bool  `json:"Enable,omitempty"`
	VSType   string `json:"VStype,omitempty"`
	NickName string `json:"NickName,omitempty"`
}

type VirtualServiceParametersStandardOptions struct {
	Cookie             string `json:"Cookie,omitempty"`
	ForceL7            *bool  `json:"ForceL7,omitempty"`
	Idletime           *int32 `json:"Idletime,omitempty"`
	Persist            string `json:"Persist,omitempty"`
	SubnetOriginating  *bool  `json:"SubnetOriginating,omitempty"`
	PersistTimeout     string `json:"PersistTimeout,omitempty"`
	Refreshpersist     *bool  `json:"Refreshpersist,omitempty"`
	QueryTag           string `json:"QueryTag,omitempty"`
	Schedule           string `json:"Schedule,omitempty"`
	Showadaptive       string `json:"showadaptive,omitempty"`
	AdaptiveInterval   *int32 `json:"AdaptiveInterval,omitempty"`
	AdaptiveUrl        string `json:"AdaptiveUrl,omitempty"`
	AdaptivePort       *int32 `json:"AdaptivePort,omitempty"`
	AdaptiveMinPercent *int32 `json:"AdaptiveMinPercent,omitempty"`
	ServerInit         *int32 `json:"ServerInit,omitempty"`
	Transparent        *bool  `json:"Transparent,omitempty"`
	UseForSnat         *bool  `json:"UseForSnat,omitempty"`
	QoS                *int32 `json:"QoS,omitempty"`
	StartTLSMode       *int32 `json:"StartTLSMode,omitempty"`
	ExtraPorts         string `json:"ExtraPorts,omitempty"`
}

type VirtualServiceParametersSSLProperties struct {
	CertFile              string `json:"CertFile,omitempty"`
	Ciphers               string `json:"Ciphers,omitempty"`
	CipherSet             string `json:"CipherSet,omitempty"`
	Tls13CipherSet        string `json:"Tls13CipherSet,omitempty"`
	ClientCert            *int32 `json:"ClientCert,omitempty"`
	PassCipher            *bool  `json:"PassCipher,omitempty"`
	SSLReencrypt          *bool  `json:"SSLReencrypt,omitempty"`
	PassSNI               *bool  `json:"PassSNI,omitempty"`
	SSLReverse            *bool  `json:"SSLReverse,omitempty"`
	SSLRewrite            string `json:"SSLRewrite,omitempty"`
	ReverseSNIHostname    string `json:"ReverseSNIHostname,omitempty"`
	SecurityHeaderOptions *int32 `json:"SecurityHeaderOptions,omitempty"`
	SSLAcceleration       *bool  `json:"SSLAcceleration,omitempty"`
	OCSPVerify            *bool  `json:"OCSPVerify,omitempty"`
	TLSType               string `json:"TLSType,omitempty"`
	NeedHostName          *bool  `json:"NeedHostName,omitempty"`
	IntermediateCerts     string `json:"IntermediateCerts,omitempty"`
}

type VirtualServiceParametersAdvancedProperties struct {
	HTTPReschedule         *bool    `json:"HTTPReschedule,omitempty"`
	CopyHdrFrom            string   `json:"CopyHdrFrom,omitempty"`
	CopyHdrTo              string   `json:"CopyHdrTo,omitempty"`
	AddVia                 *int32   `json:"AddVia,omitempty"`
	AllowHTTP2             *bool    `json:"AllowHTTP2,omitempty"`
	Cache                  *bool    `json:"Cache,omitempty"`
	Compress               *bool    `json:"Compress,omitempty"`
	CachePercent           *int32   `json:"CachePercent,omitempty"`
	DefaultGW              string   `json:"DefaultGW,omitempty"`
	ErrorCode              string   `json:"ErrorCode,omitempty"`
	ErrorUrl               string   `json:"ErrorUrl,omitempty"`
	PortFollow             *int32   `json:"PortFollow,omitempty"`
	FollowVSID             *int32   `json:"FollowVSID,omitempty"`
	LocalBindAddrs         string   `json:"LocalBindAddrs,omitempty"`
	NRequestRules          *int32   `json:"NRequestRules,omitempty"`
	NResponseRules         *int32   `json:"NResponseRules,omitempty"`
	RequestRules           []string `json:"RequestRules,omitempty"`
	ResponseRules          []string `json:"ResponseRules,omitempty"`
	StandbyAddr            string   `json:"StandbyAddr,omitempty"`
	StandbyPort            *int32   `json:"StandbyPort,omitempty"`
	NonLocalSorryServer    *bool    `json:"NonLocalSorryServer,omitempty"`
	Verify                 *int32   `json:"Verify,omitempty"`
	AltAddress             string   `json:"AltAddress,omitempty"`
	PreProcPrecedence      string   `json:"PreProcPrecedence,omitempty"`
	PreProcPrecedencePos   *int32   `json:"PreProcPrecedencePos,omitempty"`
	RequestPrecedence      string   `json:"RequestPrecedence,omitempty"`
	RequestPrecedencePos   *int32   `json:"RequestPrecedencePos,omitempty"`
	ResponsePrecedence     string   `json:"ResponsePrecedence,omitempty"`
	ResponsePrecedencePos  *int32   `json:"ResponsePrecedencePos,omitempty"`
	MatchBodyPrecedence    string   `json:"MatchBodyPrecedence,omitempty"`
	MatchBodyPrecedencePos *int32   `json:"MatchBodyPrecedencePos,omitempty"`
	ResponseStatusRemap    *bool    `json:"ResponseStatusRemap,omitempty"`
	ResponseRemapMsgMap    string   `json:"ResponseRemapMsgMap,omitempty"`
	ResponseRemapMsgFormat *int32   `json:"ResponseRemapMsgFormat,omitempty"`
	ResponseRemapCodeMap   string   `json:"ResponseRemapCodeMap,omitempty"`
}

type VirtualServiceParametersWAFSettings struct {
	Intercept                      *bool    `json:"Intercept,omitempty"`
	InterceptMode                  *int32   `json:"InterceptMode,omitempty"`
	InterceptOpts                  []string `json:"InterceptOpts,omitempty"`
	InterceptPOSTOtherContentTypes string   `json:"InterceptPOSTOtherContentTypes,omitempty"`
	AlertThreshold                 *int32   `json:"AlertThreshold,omitempty"`
}

type VirtualServiceParametersESPOptions struct {
	AllowedHosts          string `json:"AllowedHosts,omitempty"`
	AllowedDirectories    string `json:"AllowedDirectories,omitempty"`
	Domain                string `json:"Domain,omitempty"`
	Logoff                string `json:"Logoff,omitempty"`
	AddAuthHeader         string `json:"AddAuthHeader,omitempty"`
	DisplayPubPriv        *bool  `json:"DisplayPubPriv,omitempty"`
	DisablePasswordForm   *bool  `json:"DisablePasswordForm,omitempty"`
	Captcha               *bool  `json:"Captcha,omitempty"`
	CaptchaPublicKey      string `json:"CaptchaPublicKey,omitempty"`
	CaptchaPrivateKey     string `json:"CaptchaPrivateKey,omitempty"`
	CaptchaAccessUrl      string `json:"CaptchaAccessUrl,omitempty"`
	CaptchaVerifyUrl      string `json:"CaptchaVerifyUrl,omitempty"`
	ESPLogs               *int32 `json:"ESPLogs,omitempty"`
	SMTPAllowedDomains    string `json:"SMTPAllowedDomains,omitempty"`
	ExcludedDirectories   string `json:"ExcludedDirectories,omitempty"`
	EspEnabled            *bool  `json:"EspEnabled,omitempty"`
	InputAuthMode         *int32 `json:"InputAuthMode,omitempty"`
	OutputAuthMode        *int32 `json:"OutputAuthMode,omitempty"`
	TokenServerFQDN       string `json:"TokenServerFQDN,omitempty"`
	ServerFbaPath         string `json:"ServerFbaPath,omitempty"`
	ServerFBAPost         string `json:"ServerFBAPost,omitempty"`
	ServerFbaUsernameOnly *bool  `json:"ServerFbaUsernameOnly,omitempty"`
	OutConf               string `json:"OutConf,omitempty"`
	SingleSignOnDir       string `json:"SingleSignOnDir,omitempty"`
	SingleSignOnMessage   string `json:"SingleSignOnMessage,omitempty"`
	AllowedGroups         string `json:"AllowedGroups,omitempty"`
	GroupSIDs             string `json:"GroupSIDs,omitempty"`
	IncludeNestedGroups   *bool  `json:"IncludeNestedGroups,omitempty"`
	SteeringGroups        string `json:"SteeringGroups,omitempty"`
	VerifyBearer          *bool  `json:"VerifyBearer,omitempty"`
	BearerCertificateName string `json:"BearerCertificateName,omitempty"`
	BearerText            string `json:"BearerText,omitempty"`
	ExcludedDomains       string `json:"ExcludedDomains,omitempty"`
	AltDomains            string `json:"AltDomains,omitempty"`
	SameSite              *int32 `json:"SameSite,omitempty"`
	UserPwdChangeURL      string `json:"UserPwdChangeURL,omitempty"`
	UserPwdChangeMsg      string `json:"UserPwdChangeMsg,omitempty"`
	UserPwdExpiryWarn     *bool  `json:"UserPwdExpiryWarn,omitempty"`
	UserPwdExpiryWarnDays *int32 `json:"UserPwdExpiryWarnDays,omitempty"`
}

type VirtualServiceParametersRealServers struct {
	CheckType            string              `json:"CheckType,omitempty"`
	LdapEndpoint32       string              `json:"LdapEndpoint,omitempty"`
	CheckHost            string              `json:"CheckHost,omitempty"`
	CheckPattern         string              `json:"CheckPattern,omitempty"`
	CheckUrl             string              `json:"CheckUrl,omitempty"`
	CheckCodes           string              `json:"CheckCodes,omitempty"`
	CheckHeaders         string              `json:"CheckHeaders,omitempty"`
	MatchLen             *int32              `json:"MatchLen,omitempty"`
	CheckUse1_1          *bool               `json:"CheckUse1.1,omitempty"`
	CheckPort            string              `json:"CheckPort,omitempty"`
	NumberOfRSs          *int32              `json:"NumberOfRSs,omitempty"`
	NRules               *int32              `json:"NRules,omitempty"`
	RuleList             string              `json:"RuleList,omitempty"`
	CheckUseGet          *int32              `json:"CheckUseGet,omitempty"`
	ExtraHdrKey          string              `json:"ExtraHdrKey,omitempty"`
	ExtraHdrValue        string              `json:"ExtraHdrValue,omitempty"`
	SubVS                []SubVirtualService `json:"SubVS,omitempty"`
	CheckPostData        string              `json:"CheckPostData,omitempty"`
	RSRulePrecedence     string              `json:"RSRulePrecedence,omitempty"`
	RSRulePrecedencePos  *int32              `json:"RSRulePrecedencePos,omitempty"`
	EnhancedHealthchecks *bool               `json:"EnhancedHealthchecks,omitempty"`
	RsMinimum            *int32              `json:"RsMinimum,omitempty"`
}

type VirtualServiceParametersMiscellaneous struct {
	Adaptive     string `json:"Adaptive,omitempty"`
	MultiConnect *bool  `json:"MultiConnect,omitempty"`
	NonLocal     *bool  `json:"non_local,omitempty"`
}

type ListVirtualServiceResponse struct {
	*LoadMasterResponse
	VS []VirtualService `json:"VS"`
}

type ShowVirtualServiceResponse struct {
	*LoadMasterResponse
	*VirtualService
}

type AddVirtualServiceResponse struct {
	*LoadMasterResponse
	*VirtualService
}

type DeleteVirtualServiceResponse struct {
	*LoadMasterResponse
}

type ModifyVirtualServiceResponse struct {
	*LoadMasterResponse
	*VirtualService
}

func (c *Client) ListVirtualService() (*ListVirtualServiceResponse, error) {
	slog.Debug("Listing virtual services")

	payload := struct {
		*LoadMasterRequest
	}{
		&LoadMasterRequest{
			Command: "listvs",
		},
	}
	response, err := sendRequest(c, payload, ListVirtualServiceResponse{})
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ShowVirtualService(vs_identifier string) (*ShowVirtualServiceResponse, error) {
	slog.Debug("Showing virtual service", "vs_identifier", vs_identifier)

	payload := struct {
		*LoadMasterRequest
		VS string `json:"vs"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "showvs",
		},
		VS: vs_identifier,
	}

	response, err := sendRequest(c, payload, ShowVirtualServiceResponse{})
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) AddVirtualService(address string, port string, protocol string, parameters VirtualServiceParameters) (*AddVirtualServiceResponse, error) {
	slog.Debug("Adding virtual service", "address", address, "port", port, "protocol", protocol)
	payload := struct {
		*LoadMasterRequest
		VS       string `json:"vs"`
		Port     string `json:"port"`
		Protocol string `json:"prot"`
		*VirtualServiceParameters
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "addvs",
		},
		VS:                       address,
		Port:                     port,
		Protocol:                 protocol,
		VirtualServiceParameters: &parameters,
	}

	response, err := sendRequest(c, payload, AddVirtualServiceResponse{})
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) DeleteVirtualService(vs_identifier string) (*DeleteVirtualServiceResponse, error) {
	slog.Debug("Deleting virtual service", "vs_identifier", vs_identifier)
	payload := struct {
		*LoadMasterRequest
		VS string `json:"vs"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "delvs",
		},
		VS: vs_identifier,
	}

	response, err := sendRequest(c, payload, DeleteVirtualServiceResponse{})
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ModifyVirtualService(vs_identifier string, parameters VirtualServiceParameters) (*ModifyVirtualServiceResponse, error) {
	slog.Debug("Modifying virtual service", "vs_identifier", vs_identifier)
	payload := struct {
		*LoadMasterRequest
		VS string `json:"vs"`
		*VirtualServiceParameters
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "modvs",
		},
		VS:                       vs_identifier,
		VirtualServiceParameters: &parameters,
	}

	response, err := sendRequest(c, payload, ModifyVirtualServiceResponse{})
	if err != nil {
		return nil, err
	}

	return response, nil
}

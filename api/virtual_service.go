package api

type VirtualService struct {
	Index    int    `json:"Index"`
	Protocol string `json:"Protocol"`
	Address  string `json:"VSAddress"`
	Port     string `json:"VSPort"`

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
	ForceL7            *bool  `json:"ForceL7"`
	Idletime           int    `json:"Idletime,omitempty"`
	Persist            string `json:"Persist,omitempty"`
	SubnetOriginating  *bool  `json:"SubnetOriginating"`
	PersistTimeout     string `json:"PersistTimeout,omitempty"`
	Refreshpersist     *bool  `json:"Refreshpersist"`
	QueryTag           string `json:"QueryTag,omitempty"`
	Schedule           string `json:"Schedule,omitempty"`
	Showadaptive       string `json:"showadaptive,omitempty"`
	AdaptiveInterval   int    `json:"AdaptiveInterval,omitempty"`
	AdaptiveUrl        string `json:"AdaptiveUrl,omitempty"`
	AdaptivePort       int    `json:"AdaptivePort,omitempty"`
	AdaptiveMinPercent int    `json:"AdaptiveMinPercent,omitempty"`
	ServerInit         int    `json:"ServerInit,omitempty"`
	Transparent        *bool  `json:"Transparent,omitempty"`
	UseForSnat         *bool  `json:"UseForSnat,omitempty"`
	QoS                int    `json:"QoS,omitempty"`
	StartTLSMode       int    `json:"StartTLSMode,omitempty"`
	ExtraPorts         string `json:"ExtraPorts,omitempty"`
}

type VirtualServiceParametersSSLProperties struct {
	CertFile              string `json:"CertFile,omitempty"`
	Ciphers               string `json:"Ciphers,omitempty"`
	CipherSet             string `json:"CipherSet,omitempty"`
	Tls13CipherSet        string `json:"Tls13CipherSet,omitempty"`
	ClientCert            int    `json:"ClientCert,omitempty"`
	PassCipher            *bool  `json:"PassCipher,omitempty"`
	SSLReencrypt          *bool  `json:"SSLReencrypt,omitempty"`
	PassSNI               *bool  `json:"PassSNI,omitempty"`
	SSLReverse            *bool  `json:"SSLReverse,omitempty"`
	SSLRewrite            string `json:"SSLRewrite,omitempty"`
	ReverseSNIHostname    string `json:"ReverseSNIHostname,omitempty"`
	SecurityHeaderOptions int    `json:"SecurityHeaderOptions,omitempty"`
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
	AddVia                 int      `json:"AddVia,omitempty"`
	AllowHTTP2             *bool    `json:"AllowHTTP2,omitempty"`
	Cache                  *bool    `json:"Cache,omitempty"`
	Compress               *bool    `json:"Compress,omitempty"`
	CachePercent           int      `json:"CachePercent,omitempty"`
	DefaultGW              string   `json:"DefaultGW,omitempty"`
	ErrorCode              string   `json:"ErrorCode,omitempty"`
	ErrorUrl               string   `json:"ErrorUrl,omitempty"`
	PortFollow             int      `json:"PortFollow,omitempty"`
	FollowVSID             int      `json:"FollowVSID,omitempty"`
	LocalBindAddrs         string   `json:"LocalBindAddrs,omitempty"`
	NRequestRules          int      `json:"NRequestRules,omitempty"`
	NResponseRules         int      `json:"NResponseRules,omitempty"`
	RequestRules           []string `json:"RequestRules,omitempty"`
	ResponseRules          []string `json:"ResponseRules,omitempty"`
	StandbyAddr            string   `json:"StandbyAddr,omitempty"`
	StandbyPort            int      `json:"StandbyPort,omitempty"`
	NonLocalSorryServer    *bool    `json:"NonLocalSorryServer,omitempty"`
	Verify                 int      `json:"Verify,omitempty"`
	AltAddress             string   `json:"AltAddress,omitempty"`
	PreProcPrecedence      string   `json:"PreProcPrecedence,omitempty"`
	PreProcPrecedencePos   int      `json:"PreProcPrecedencePos,omitempty"`
	RequestPrecedence      string   `json:"RequestPrecedence,omitempty"`
	RequestPrecedencePos   int      `json:"RequestPrecedencePos,omitempty"`
	ResponsePrecedence     string   `json:"ResponsePrecedence,omitempty"`
	ResponsePrecedencePos  int      `json:"ResponsePrecedencePos,omitempty"`
	MatchBodyPrecedence    string   `json:"MatchBodyPrecedence,omitempty"`
	MatchBodyPrecedencePos int      `json:"MatchBodyPrecedencePos,omitempty"`
	ResponseStatusRemap    *bool    `json:"ResponseStatusRemap,omitempty"`
	ResponseRemapMsgMap    string   `json:"ResponseRemapMsgMap,omitempty"`
	ResponseRemapMsgFormat int      `json:"ResponseRemapMsgFormat,omitempty"`
	ResponseRemapCodeMap   string   `json:"ResponseRemapCodeMap,omitempty"`
}

type VirtualServiceParametersWAFSettings struct {
	Intercept                      *bool    `json:"Intercept,omitempty"`
	InterceptMode                  int      `json:"InterceptMode,omitempty"`
	InterceptOpts                  []string `json:"InterceptOpts,omitempty"`
	InterceptPOSTOtherContentTypes string   `json:"InterceptPOSTOtherContentTypes,omitempty"`
	AlertThreshold                 int      `json:"AlertThreshold,omitempty"`
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
	ESPLogs               int    `json:"ESPLogs,omitempty"`
	SMTPAllowedDomains    string `json:"SMTPAllowedDomains,omitempty"`
	ExcludedDirectories   string `json:"ExcludedDirectories,omitempty"`
	EspEnabled            *bool  `json:"EspEnabled,omitempty"`
	InputAuthMode         int    `json:"InputAuthMode,omitempty"`
	OutputAuthMode        int    `json:"OutputAuthMode,omitempty"`
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
	SameSite              int    `json:"SameSite,omitempty"`
	UserPwdChangeURL      string `json:"UserPwdChangeURL,omitempty"`
	UserPwdChangeMsg      string `json:"UserPwdChangeMsg,omitempty"`
	UserPwdExpiryWarn     *bool  `json:"UserPwdExpiryWarn,omitempty"`
	UserPwdExpiryWarnDays int    `json:"UserPwdExpiryWarnDays,omitempty"`
}

type VirtualServiceParametersRealServers struct {
	CheckType            string              `json:"CheckType,omitempty"`
	LdapEndpoint         string              `json:"LdapEndpoint,omitempty"`
	CheckHost            string              `json:"CheckHost,omitempty"`
	CheckPattern         string              `json:"CheckPattern,omitempty"`
	CheckUrl             string              `json:"CheckUrl,omitempty"`
	CheckCodes           string              `json:"CheckCodes,omitempty"`
	CheckHeaders         string              `json:"CheckHeaders,omitempty"`
	MatchLen             int                 `json:"MatchLen,omitempty"`
	CheckUse1_1          *bool               `json:"CheckUse1.1,omitempty"`
	CheckPort            string              `json:"CheckPort,omitempty"`
	NumberOfRSs          int                 `json:"NumberOfRSs,omitempty"`
	NRules               int                 `json:"NRules,omitempty"`
	RuleList             string              `json:"RuleList,omitempty"`
	CheckUseGet          int                 `json:"CheckUseGet,omitempty"`
	ExtraHdrKey          string              `json:"ExtraHdrKey,omitempty"`
	ExtraHdrValue        string              `json:"ExtraHdrValue,omitempty"`
	SubVS                []SubVirtualService `json:"SubVS,omitempty"`
	CheckPostData        string              `json:"CheckPostData,omitempty"`
	RSRulePrecedence     string              `json:"RSRulePrecedence,omitempty"`
	RSRulePrecedencePos  int                 `json:"RSRulePrecedencePos,omitempty"`
	EnhancedHealthchecks *bool               `json:"EnhancedHealthchecks,omitempty"`
	RsMinimum            int                 `json:"RsMinimum,omitempty"`
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
	payload := struct {
		*LoadMasterRequest
	}{
		&LoadMasterRequest{
			Command: "listvs",
		},
	}
	response, err := sendRequest(c, payload, &ListVirtualServiceResponse{})
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ShowVirtualService(vs_identifier int) (*ShowVirtualServiceResponse, error) {
	payload := struct {
		*LoadMasterRequest
		VS int `json:"vs"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "showvs",
		},
		VS: vs_identifier,
	}

	response, err := sendRequest(c, payload, &ShowVirtualServiceResponse{})
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) AddVirtualService(address string, port string, protocol string, parameters VirtualServiceParameters) (*AddVirtualServiceResponse, error) {
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

	response, err := sendRequest(c, payload, &AddVirtualServiceResponse{})
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) DeleteVirtualService(vs_identifier int) (*DeleteVirtualServiceResponse, error) {
	payload := struct {
		*LoadMasterRequest
		VS int `json:"vs"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "delvs",
		},
		VS: vs_identifier,
	}

	response, err := sendRequest(c, payload, &DeleteVirtualServiceResponse{})
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ModifyVirtualService(vs_identifier int, parameters VirtualServiceParameters) (*ModifyVirtualServiceResponse, error) {
	payload := struct {
		*LoadMasterRequest
		VS int `json:"vs"`
		*VirtualServiceParameters
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "modvs",
		},
		VS:                       vs_identifier,
		VirtualServiceParameters: &parameters,
	}

	response, err := sendRequest(c, payload, &ModifyVirtualServiceResponse{})
	if err != nil {
		return nil, err
	}

	return response, nil
}

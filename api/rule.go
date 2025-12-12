package api

import "log/slog"

type RuleResponse struct {
	*LoadMasterResponse
	AddHeaderRules     []AddHeaderRule     `json:"AddHeaderRule,omitempty"`
	DeleteHeaderRules  []DeleteHeaderRule  `json:"DeleteHeaderRule,omitempty"`
	MatchContentRules  []MatchContentRule  `json:"MatchContentRule,omitempty"`
	ModifyURLRules     []ModifyURLRule     `json:"ModifyURLRule,omitempty"`
	ReplaceBodyRules   []ReplaceBodyRule   `json:"ReplaceBodyRule,omitempty"`
	ReplaceHeaderRules []ReplaceHeaderRule `json:"ReplaceHeaderRule,omitempty"`
}

type GeneralRule struct {
	Replacement     *string `json:"replacement,omitempty"`
	Pattern         *string `json:"pattern,omitempty"`
	NoCase          *bool   `json:"nocase,omitempty"`
	CaseIndependent *bool   `json:"caseindependent,omitempty"`
	MatchType       *string `json:"matchtype,omitempty"`
	IncHost         *bool   `json:"inchost,omitempty"`
	Negate          *bool   `json:"negate,omitempty"`
	IncQuery        *bool   `json:"incquery,omitempty"`
	Header          *string `json:"header,omitempty"`
	SetOnMatch      *int32  `json:"setonmatch,omitempty"`
	OnlyOnFlag      *int32  `json:"onlyonflag,omitempty"`
	OnlyOnNoFlag    *int32  `json:"onlyonnoflag,omitempty"`
	MustFail        *bool   `json:"mustfail,omitempty"`
}

type MatchContentRule struct {
	Name            string  `json:"name,omitempty"`
	MatchType       string  `json:"matchtype,omitempty"`
	IncHost         *bool   `json:"addhost,omitempty"`
	CaseIndependent *bool   `json:"CaseIndependent,omitempty"`
	Negate          *bool   `json:"negate,omitempty"`
	IncQuery        *bool   `json:"IncludeQuery,omitempty"`
	Header          *string `json:"header,omitempty"`
	Pattern         string  `json:"pattern,omitempty"`
	SetOnMatch      *int32  `json:"SetFlagOnMatch,omitempty"`
	OnlyOnFlag      *int32  `json:"onlyonflag,omitempty"`
	OnlyOnNoFlag    *int32  `json:"onlyonnoflag,omitempty"`
	MustFail        *bool   `json:"mustfail,omitempty"`
}

type AddHeaderRule struct {
	Name         string  `json:"name,omitempty"`
	Header       *string `json:"header,omitempty"`
	Replacement  string  `json:"HeaderValue,omitempty"`
	OnlyOnFlag   *int32  `json:"onlyonflag,omitempty"`
	OnlyOnNoFlag *int32  `json:"onlyonnoflag,omitempty"`
}

type DeleteHeaderRule struct {
	Name         string `json:"name,omitempty"`
	Pattern      string `json:"pattern,omitempty"`
	OnlyOnFlag   *int32 `json:"onlyonflag,omitempty"`
	OnlyOnNoFlag *int32 `json:"onlyonnoflag,omitempty"`
}

type ReplaceHeaderRule struct {
	Name         string  `json:"name,omitempty"`
	Header       *string `json:"header,omitempty"`
	Replacement  string  `json:"replacement,omitempty"`
	Pattern      string  `json:"pattern,omitempty"`
	OnlyOnFlag   *int32  `json:"onlyonflag,omitempty"`
	OnlyOnNoFlag *int32  `json:"onlyonnoflag,omitempty"`
}

type ModifyURLRule struct {
	Name         string `json:"name,omitempty"`
	Replacement  string `json:"replacement,omitempty"`
	Pattern      string `json:"pattern,omitempty"`
	OnlyOnFlag   *int32 `json:"onlyonflag,omitempty"`
	OnlyOnNoFlag *int32 `json:"onlyonnoflag,omitempty"`
}

type ReplaceBodyRule struct {
	Name            string `json:"name,omitempty"`
	Replacement     string `json:"replacement,omitempty"`
	Pattern         string `json:"pattern,omitempty"`
	OnlyOnFlag      *int32 `json:"onlyonflag,omitempty"`
	OnlyOnNoFlag    *int32 `json:"onlyonnoflag,omitempty"`
	CaseIndependent *bool  `json:"caseindependent,omitempty"`
}

func (c *Client) ListRule() (*RuleResponse, error) {
	slog.Debug("Listing rules")
	payload := struct {
		*LoadMasterRequest
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "showrule",
		},
	}

	response, err := sendRequest(c, payload, RuleResponse{})

	if err != nil {
		return nil, err
	}

	if response.Code >= 400 {
		return nil, &LoadMasterError{
			Code:    response.Code,
			Message: response.Message,
		}
	}

	return response, nil
}

func (c *Client) ShowRule(name string) (*RuleResponse, error) {
	slog.Debug("Showing rule", "name", name)
	payload := struct {
		*LoadMasterRequest
		Name string `json:"name"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "showrule",
		},
		Name: name,
	}

	response, err := sendRequest(c, payload, RuleResponse{})
	if err != nil {
		return nil, err
	}

	if response.Code >= 400 {
		return nil, &LoadMasterError{
			Code:    response.Code,
			Message: response.Message,
		}
	}

	return response, nil
}

func (c *Client) AddRule(rule_type string, name string, params GeneralRule) (*RuleResponse, error) {
	slog.Debug("Adding rule", "name", name, "type", rule_type)
	payload := struct {
		*LoadMasterRequest
		*GeneralRule
		RuleType string `json:"type"`
		Name     string `json:"name"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "addrule",
		},
		Name:        name,
		RuleType:    rule_type,
		GeneralRule: &params,
	}

	response, err := sendRequest(c, payload, RuleResponse{})

	if err != nil {
		return nil, err
	}
	if response.Code >= 400 {
		return nil, &LoadMasterError{
			Code:    response.Code,
			Message: response.Message,
		}
	}

	return response, nil
}

func (c *Client) ModifyRule(name string, params GeneralRule) (*RuleResponse, error) {
	slog.Debug("Modifying rule", "name", name)
	payload := struct {
		*LoadMasterRequest
		*GeneralRule
		Name string `json:"name"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "modrule",
		},
		Name:        name,
		GeneralRule: &params,
	}

	response, err := sendRequest(c, payload, RuleResponse{})
	if err != nil {
		return nil, err
	}
	if response.Code >= 400 {
		return nil, &LoadMasterError{
			Code:    response.Code,
			Message: response.Message,
		}
	}

	return response, nil
}

func (c *Client) DeleteRule(name string) (*LoadMasterResponse, error) {
	slog.Debug("Deleting rule", "name", name)
	payload := struct {
		*LoadMasterRequest
		Name string `json:"name"`
	}{
		LoadMasterRequest: &LoadMasterRequest{
			Command: "delrule",
		},
		Name: name,
	}

	response, err := sendRequest(c, payload, LoadMasterResponse{})

	if err != nil {
		return nil, err
	}
	if response.Code >= 400 {
		return nil, &LoadMasterError{
			Code:    response.Code,
			Message: response.Message,
		}
	}

	return response, nil
}

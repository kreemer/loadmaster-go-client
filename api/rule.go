package api

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
	Replacement     string  `json:"replacement,omitempty"`
	Pattern         string  `json:"pattern,omitempty"`
	Onlyonflag      int     `json:"onlyonflag,omitempty"`
	Onlyonnoflag    int     `json:"onlyonnoflag,omitempty"`
	Caseindependent *bool   `json:"caseindependent,omitempty"`
	Matchtype       string  `json:"matchtype,omitempty"`
	Inchost         *bool   `json:"inchost,omitempty"`
	Nocase          *bool   `json:"nocase,omitempty"`
	Negate          *bool   `json:"negate,omitempty"`
	Incquery        *bool   `json:"incquery,omitempty"`
	Header          *string `json:"header,omitempty"`
	Setonmatch      int     `json:"setonmatch,omitempty"`
	Mustfail        *bool   `json:"mustfail,omitempty"`
}

type MatchContentRule struct {
	Name         string  `json:"name,omitempty"`
	Matchtype    string  `json:"matchtype,omitempty"`
	Inchost      *bool   `json:"inchost,omitempty"`
	Nocase       *bool   `json:"nocase,omitempty"`
	Negate       *bool   `json:"negate,omitempty"`
	Incquery     *bool   `json:"incquery,omitempty"`
	Header       *string `json:"header,omitempty"`
	Pattern      string  `json:"pattern,omitempty"`
	Setonmatch   int     `json:"setonmatch,omitempty"`
	Onlyonflag   int     `json:"onlyonflag,omitempty"`
	Onlyonnoflag int     `json:"onlyonnoflag,omitempty"`
	Mustfail     *bool   `json:"mustfail,omitempty"`
}

type AddHeaderRule struct {
	Name         string  `json:"name,omitempty"`
	Header       *string `json:"header,omitempty"`
	Replacement  string  `json:"replacement,omitempty"`
	Onlyonflag   int     `json:"onlyonflag,omitempty"`
	Onlyonnoflag int     `json:"onlyonnoflag,omitempty"`
	HeaderValue  string  `json:"HeaderValue,omitempty"`
}

type DeleteHeaderRule struct {
	Name         string `json:"name,omitempty"`
	Pattern      string `json:"pattern,omitempty"`
	Onlyonflag   int    `json:"onlyonflag,omitempty"`
	Onlyonnoflag int    `json:"onlyonnoflag,omitempty"`
}

type ReplaceHeaderRule struct {
	Name         string  `json:"name,omitempty"`
	Header       *string `json:"header,omitempty"`
	Replacement  string  `json:"replacement,omitempty"`
	Pattern      string  `json:"pattern,omitempty"`
	Onlyonflag   int     `json:"onlyonflag,omitempty"`
	Onlyonnoflag int     `json:"onlyonnoflag,omitempty"`
}

type ModifyURLRule struct {
	Name         string `json:"name,omitempty"`
	Replacement  string `json:"replacement,omitempty"`
	Pattern      string `json:"pattern,omitempty"`
	Onlyonflag   int    `json:"onlyonflag,omitempty"`
	Onlyonnoflag int    `json:"onlyonnoflag,omitempty"`
}

type ReplaceBodyRule struct {
	Name            string `json:"name,omitempty"`
	Replacement     string `json:"replacement,omitempty"`
	Pattern         string `json:"pattern,omitempty"`
	Onlyonflag      int    `json:"onlyonflag,omitempty"`
	Onlyonnoflag    int    `json:"onlyonnoflag,omitempty"`
	Caseindependent *bool  `json:"caseindependent,omitempty"`
}

func (c *Client) ListRule() (*RuleResponse, error) {
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

	return response, nil
}

func (c *Client) AddRule(rule_type string, name string, params GeneralRule) (*RuleResponse, error) {
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

	return response, nil
}

func (c *Client) DeleteRule(name string) (*LoadMasterResponse, error) {
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

	return response, nil
}

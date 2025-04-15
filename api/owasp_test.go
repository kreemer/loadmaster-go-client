package api

import (
	"encoding/base64"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegration_OwaspRuleVirtualService(t *testing.T) {
	client, function := createClientForIntegration()
	if client == nil || function == nil {
		t.Skip("Skipping test because LOADMASTER_API_KEY or LOADMASTER_IP is not set")
	}
	defer function()

	vs, err := client.AddVirtualService("10.0.0.4", "30500", "tcp", VirtualServiceParameters{VirtualServiceParametersBasicProperties: &VirtualServiceParametersBasicProperties{VSType: "http"}})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	_, err = client.ModifyVirtualService(vs.Index, VirtualServiceParameters{VirtualServiceParametersWAFSettings: &VirtualServiceParametersWAFSettings{InterceptMode: convert2Ptr(2)}})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	t.Run("Assign an owasp custom rule to a virtual service", func(t *testing.T) {
		response, err := client.AddVirtualServiceOwaspRule(strconv.Itoa(vs.Index), "913100")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if response.Code != 200 {
			t.Fatalf("expected 200, got %d", response.Code)
		}

		response_rule, err := client.ShowVirtualServiceOwaspRule(strconv.Itoa(vs.Index), "913100")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		assert.Equal(t, 200, response_rule.Code)
		assert.Equal(t, "913100", response_rule.Rule.Id)
		assert.Equal(t, "CRS", response_rule.Rule.Type)
		assert.Equal(t, "yes", response_rule.Rule.Enabled)
	})

	t.Run("Unassign an owasp custom rule to a virtual service", func(t *testing.T) {
		_, err = client.AddVirtualServiceOwaspRule(strconv.Itoa(vs.Index), "913100")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		response, err := client.ShowVirtualServiceOwaspRule(strconv.Itoa(vs.Index), "913100")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		assert.Equal(t, 200, response.Code)
		assert.Equal(t, "913100", response.Rule.Id)
		assert.Equal(t, "CRS", response.Rule.Type)
		assert.Equal(t, "yes", response.Rule.Enabled)

		_, err = client.DeleteVirtualServiceOwaspRule(strconv.Itoa(vs.Index), "913100")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		response, err = client.ShowVirtualServiceOwaspRule(strconv.Itoa(vs.Index), "913100")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		assert.Equal(t, 200, response.Code)
		assert.Equal(t, "913100", response.Rule.Id)
		assert.Equal(t, "CRS", response.Rule.Type)
		assert.Equal(t, "no", response.Rule.Enabled)

	})

}

func TestIntegration_OwaspCustomRuleVirtualService(t *testing.T) {
	client, function := createClientForIntegration()
	if client == nil || function == nil {
		t.Skip("Skipping test because LOADMASTER_API_KEY or LOADMASTER_IP is not set")
	}
	defer function()

	vs, err := client.AddVirtualService("10.0.0.4", "30500", "tcp", VirtualServiceParameters{VirtualServiceParametersBasicProperties: &VirtualServiceParametersBasicProperties{VSType: "http"}})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	_, err = client.ModifyVirtualService(vs.Index, VirtualServiceParameters{VirtualServiceParametersWAFSettings: &VirtualServiceParametersWAFSettings{InterceptMode: convert2Ptr(2)}})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	custom_rule_content := `
SecMarker BEGIN_ALLOWLIST_login
 
# START allowlisting block for URI /login SecRule REQUEST_URI "!@beginsWith /login" \
    "id:11001,phase:1,pass,t:lowercase,nolog,skipAfter:END_ALLOWLIST_login"
SecRule REQUEST_URI "!@beginsWith /login" \
    "id:11002,phase:2,pass,t:lowercase,nolog,skipAfter:END_ALLOWLIST_login"
 
# Validate HTTP method
SecRule REQUEST_METHOD "!@pm GET HEAD POST OPTIONS" \
    "id:11100,phase:1,deny,status:405,log,tag:'Login Allowlist',\
    msg:'Method %{MATCHED_VAR} not allowed'"
 
# Validate URIs
SecRule REQUEST_FILENAME "@beginsWith /login/static/css" \
    "id:11200,phase:1,pass,nolog,tag:'Login Allowlist',\
    skipAfter:END_ALLOWLIST_URIBLOCK_login"
SecRule REQUEST_FILENAME "@beginsWith /login/static/img" \
    "id:11201,phase:1,pass,nolog,tag:'Login Allowlist',\
    skipAfter:END_ALLOWLIST_URIBLOCK_login"
SecRule REQUEST_FILENAME "@beginsWith /login/static/js" \
    "id:11202,phase:1,pass,nolog,tag:'Login Allowlist',\
    skipAfter:END_ALLOWLIST_URIBLOCK_login"
SecRule REQUEST_FILENAME \
    "@rx ^/login/(displayLogin|login|logout).do$" \
    "id:11250,phase:1,pass,nolog,tag:'Login Allowlist',\
    skipAfter:END_ALLOWLIST_URIBLOCK_login"
 
# If we land here, we are facing an unknown URI, # which is why we will respond using the 404 status code SecAction "id:11299,phase:1,deny,status:404,log,tag:'Login Allowlist',\
    msg:'Unknown URI %{REQUEST_URI}'"
 
SecMarker END_ALLOWLIST_URIBLOCK_login
 
# Validate parameter names
SecRule ARGS_NAMES "!@rx ^(username|password|sectoken)$" \
    "id:11300,phase:2,deny,log,tag:'Login Allowlist',\
    msg:'Unknown parameter: %{MATCHED_VAR_NAME}'"
 
# Validate each parameter's uniqueness
SecRule &ARGS:username  "@gt 1" \
    "id:11400,phase:2,deny,log,tag:'Login Allowlist',\
    msg:'%{MATCHED_VAR_NAME} occurring more than once'"
SecRule &ARGS:password  "@gt 1" \
    "id:11401,phase:2,deny,log,tag:'Login Allowlist',\
    msg:'%{MATCHED_VAR_NAME} occurring more than once'"
SecRule &ARGS:sectoken  "@gt 1" \
    "id:11402,phase:2,deny,log,tag:'Login Allowlist',\
    msg:'%{MATCHED_VAR_NAME} occurring more than once'"
 
# Check individual parameters
SecRule ARGS:username "!@rx ^[a-zA-Z0-9.@_-]{1,64}$" \
    "id:11500,phase:2,deny,log,tag:'Login Allowlist',\
    msg:'Invalid parameter format: %{MATCHED_VAR_NAME} (%{MATCHED_VAR})'"
SecRule ARGS:sectoken "!@rx ^[a-zA-Z0-9]{32}$" \
    "id:11501,phase:2,deny,log,tag:'Login Allowlist',\
    msg:'Invalid parameter format: %{MATCHED_VAR_NAME} (%{MATCHED_VAR})'"
SecRule ARGS:password "@gt 64" \
    "id:11502,phase:2,deny,log,t:length,tag:'Login Allowlist',\
    msg:'Invalid parameter format: %{MATCHED_VAR_NAME} too long (%{MATCHED_VAR} bytes)'"
SecRule ARGS:password "@validateByteRange 33-244" \
    "id:11503,phase:2,deny,log,tag:'Login Allowlist',\
    msg:'Invalid parameter format: %{MATCHED_VAR_NAME} (%{MATCHED_VAR})'"
 
SecMarker END_ALLOWLIST_login
`
	t.Run("Adding a new owasp custom rule", func(t *testing.T) {
		owasp_custom_rule, err := client.AddOwaspCustomRule("test", base64.StdEncoding.EncodeToString([]byte(custom_rule_content)))
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if owasp_custom_rule.Code != 200 {
			t.Fatalf("expected 200, got %d", owasp_custom_rule.Code)
		}
	})
	t.Run("Assign an owasp custom rule to a virtual service", func(t *testing.T) {
		_, err := client.AddOwaspCustomRule("test", base64.StdEncoding.EncodeToString([]byte(custom_rule_content)))
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		response, err := client.AddVirtualServiceOwaspCustomRule(strconv.Itoa(vs.Index), "test", true)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if response.Code != 200 {
			t.Fatalf("expected 200, got %d", response.Code)
		}
	})

	t.Run("Show an owasp custom rule to a virtual service", func(t *testing.T) {
		_, err := client.AddOwaspCustomRule("test", base64.StdEncoding.EncodeToString([]byte(custom_rule_content)))
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		_, err = client.AddVirtualServiceOwaspCustomRule(strconv.Itoa(vs.Index), "test", false)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		response, err := client.ShowVirtualServiceOwaspRule(strconv.Itoa(vs.Index), "test")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		assert.Equal(t, response.Code, 200)
		assert.Equal(t, response.Rule.Name, "test")
		assert.Equal(t, response.Rule.Type, "custom")
		assert.Equal(t, response.Rule.Enabled, "yes")
		assert.Equal(t, response.Rule.RunFirst, "no")
	})

	t.Run("Unassign an owasp custom rule to a virtual service", func(t *testing.T) {
		_, err := client.AddOwaspCustomRule("test", base64.StdEncoding.EncodeToString([]byte(custom_rule_content)))
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		_, err = client.AddVirtualServiceOwaspCustomRule(strconv.Itoa(vs.Index), "test", true)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		response, err := client.ShowVirtualServiceOwaspRule(strconv.Itoa(vs.Index), "test")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		assert.Equal(t, response.Code, 200)
		assert.Equal(t, response.Rule.Name, "test")
		assert.Equal(t, response.Rule.Type, "custom")
		assert.Equal(t, response.Rule.Enabled, "yes")
		assert.Equal(t, response.Rule.RunFirst, "yes")

		_, err = client.DeleteVirtualServiceOwaspCustomRule(strconv.Itoa(vs.Index), "test")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		response, err = client.ShowVirtualServiceOwaspRule(strconv.Itoa(vs.Index), "test")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		assert.Equal(t, response.Code, 200)
		assert.Equal(t, response.Rule.Name, "test")
		assert.Equal(t, response.Rule.Type, "custom")
		assert.Equal(t, response.Rule.Enabled, "no")

	})

	t.Run("Deleting a new owasp custom rule", func(t *testing.T) {
		_, err := client.AddOwaspCustomRule("test", base64.StdEncoding.EncodeToString([]byte(custom_rule_content)))
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		_, err = client.DeleteOwaspCustomRule("test")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
	})
}

func TestIntegration_OwaspCustomData(t *testing.T) {
	client, function := createClientForIntegration()
	if client == nil || function == nil {
		t.Skip("Skipping test because LOADMASTER_API_KEY or LOADMASTER_IP is not set")
	}
	defer function()

	custom_rule_content := `TEST=test`
	t.Run("Adding a new owasp custom data", func(t *testing.T) {
		response, err := client.AddOwaspCustomData("test_data.txt", base64.StdEncoding.EncodeToString([]byte(custom_rule_content)))
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		assert.Equal(t, response.Code, 200)
	})

	t.Run("Deleting a new owasp custom data", func(t *testing.T) {
		response, err := client.AddOwaspCustomData("test_data.txt", base64.StdEncoding.EncodeToString([]byte(custom_rule_content)))
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		assert.Equal(t, response.Code, 200)

		response, err = client.DeleteOwaspCustomData("test_data")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		assert.Equal(t, response.Code, 200)
	})
	t.Run("Download a new owasp custom data", func(t *testing.T) {
		response, err := client.AddOwaspCustomData("test_data.txt", base64.StdEncoding.EncodeToString([]byte(custom_rule_content)))
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		assert.Equal(t, response.Code, 200)

		response_data, err := client.ShowOwaspCustomData("test_data.txt")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		assert.Equal(t, response_data.Code, 200)
		assert.NotEmpty(t, response_data.Data)
	})
}

func TestIntegration_OwaspRuleSubVirtualService(t *testing.T) {
	client, function := createClientForIntegration()
	if client == nil || function == nil {
		t.Skip("Skipping test because LOADMASTER_API_KEY or LOADMASTER_IP is not set")
	}
	defer function()

	vs, err := client.AddVirtualService("10.0.0.4", "30500", "tcp", VirtualServiceParameters{VirtualServiceParametersBasicProperties: &VirtualServiceParametersBasicProperties{VSType: "http"}})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	subvs, err := client.AddSubVirtualService(vs.Index, VirtualServiceParameters{})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	_, err = client.ModifySubVirtualService(subvs.Index, VirtualServiceParameters{VirtualServiceParametersWAFSettings: &VirtualServiceParametersWAFSettings{InterceptMode: convert2Ptr(2)}})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	t.Run("Assign an owasp custom rule to a sub virtual service", func(t *testing.T) {
		response, err := client.AddVirtualServiceOwaspRule(strconv.Itoa(subvs.Index), "913100")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if response.Code != 200 {
			t.Fatalf("expected 200, got %d", response.Code)
		}

		response_rule, err := client.ShowVirtualServiceOwaspRule(strconv.Itoa(subvs.Index), "913100")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		assert.Equal(t, 200, response_rule.Code)
		assert.Equal(t, "913100", response_rule.Rule.Id)
		assert.Equal(t, "CRS", response_rule.Rule.Type)
		assert.Equal(t, "yes", response_rule.Rule.Enabled)
	})

	t.Run("Unassign an owasp custom rule to a sub virtual service", func(t *testing.T) {
		_, err = client.AddVirtualServiceOwaspRule(strconv.Itoa(subvs.Index), "913100")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		response, err := client.ShowVirtualServiceOwaspRule(strconv.Itoa(subvs.Index), "913100")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		assert.Equal(t, 200, response.Code)
		assert.Equal(t, "913100", response.Rule.Id)
		assert.Equal(t, "CRS", response.Rule.Type)
		assert.Equal(t, "yes", response.Rule.Enabled)

		_, err = client.DeleteVirtualServiceOwaspRule(strconv.Itoa(subvs.Index), "913100")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		response, err = client.ShowVirtualServiceOwaspRule(strconv.Itoa(subvs.Index), "913100")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		assert.Equal(t, 200, response.Code)
		assert.Equal(t, "913100", response.Rule.Id)
		assert.Equal(t, "CRS", response.Rule.Type)
		assert.Equal(t, "no", response.Rule.Enabled)

	})

}

func TestIntegration_OwaspCustomRuleSubVirtualService(t *testing.T) {
	client, function := createClientForIntegration()
	if client == nil || function == nil {
		t.Skip("Skipping test because LOADMASTER_API_KEY or LOADMASTER_IP is not set")
	}
	defer function()

	vs, err := client.AddVirtualService("10.0.0.4", "30500", "tcp", VirtualServiceParameters{VirtualServiceParametersBasicProperties: &VirtualServiceParametersBasicProperties{VSType: "http"}})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	subvs, err := client.AddSubVirtualService(vs.Index, VirtualServiceParameters{})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	_, err = client.ModifySubVirtualService(subvs.Index, VirtualServiceParameters{VirtualServiceParametersWAFSettings: &VirtualServiceParametersWAFSettings{InterceptMode: convert2Ptr(2)}})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	custom_rule_content := `
SecMarker BEGIN_ALLOWLIST_login
 
# START allowlisting block for URI /login SecRule REQUEST_URI "!@beginsWith /login" \
    "id:11001,phase:1,pass,t:lowercase,nolog,skipAfter:END_ALLOWLIST_login"
SecRule REQUEST_URI "!@beginsWith /login" \
    "id:11002,phase:2,pass,t:lowercase,nolog,skipAfter:END_ALLOWLIST_login"
 
# Validate HTTP method
SecRule REQUEST_METHOD "!@pm GET HEAD POST OPTIONS" \
    "id:11100,phase:1,deny,status:405,log,tag:'Login Allowlist',\
    msg:'Method %{MATCHED_VAR} not allowed'"
 
# Validate URIs
SecRule REQUEST_FILENAME "@beginsWith /login/static/css" \
    "id:11200,phase:1,pass,nolog,tag:'Login Allowlist',\
    skipAfter:END_ALLOWLIST_URIBLOCK_login"
SecRule REQUEST_FILENAME "@beginsWith /login/static/img" \
    "id:11201,phase:1,pass,nolog,tag:'Login Allowlist',\
    skipAfter:END_ALLOWLIST_URIBLOCK_login"
SecRule REQUEST_FILENAME "@beginsWith /login/static/js" \
    "id:11202,phase:1,pass,nolog,tag:'Login Allowlist',\
    skipAfter:END_ALLOWLIST_URIBLOCK_login"
SecRule REQUEST_FILENAME \
    "@rx ^/login/(displayLogin|login|logout).do$" \
    "id:11250,phase:1,pass,nolog,tag:'Login Allowlist',\
    skipAfter:END_ALLOWLIST_URIBLOCK_login"
 
# If we land here, we are facing an unknown URI, # which is why we will respond using the 404 status code SecAction "id:11299,phase:1,deny,status:404,log,tag:'Login Allowlist',\
    msg:'Unknown URI %{REQUEST_URI}'"
 
SecMarker END_ALLOWLIST_URIBLOCK_login
 
# Validate parameter names
SecRule ARGS_NAMES "!@rx ^(username|password|sectoken)$" \
    "id:11300,phase:2,deny,log,tag:'Login Allowlist',\
    msg:'Unknown parameter: %{MATCHED_VAR_NAME}'"
 
# Validate each parameter's uniqueness
SecRule &ARGS:username  "@gt 1" \
    "id:11400,phase:2,deny,log,tag:'Login Allowlist',\
    msg:'%{MATCHED_VAR_NAME} occurring more than once'"
SecRule &ARGS:password  "@gt 1" \
    "id:11401,phase:2,deny,log,tag:'Login Allowlist',\
    msg:'%{MATCHED_VAR_NAME} occurring more than once'"
SecRule &ARGS:sectoken  "@gt 1" \
    "id:11402,phase:2,deny,log,tag:'Login Allowlist',\
    msg:'%{MATCHED_VAR_NAME} occurring more than once'"
 
# Check individual parameters
SecRule ARGS:username "!@rx ^[a-zA-Z0-9.@_-]{1,64}$" \
    "id:11500,phase:2,deny,log,tag:'Login Allowlist',\
    msg:'Invalid parameter format: %{MATCHED_VAR_NAME} (%{MATCHED_VAR})'"
SecRule ARGS:sectoken "!@rx ^[a-zA-Z0-9]{32}$" \
    "id:11501,phase:2,deny,log,tag:'Login Allowlist',\
    msg:'Invalid parameter format: %{MATCHED_VAR_NAME} (%{MATCHED_VAR})'"
SecRule ARGS:password "@gt 64" \
    "id:11502,phase:2,deny,log,t:length,tag:'Login Allowlist',\
    msg:'Invalid parameter format: %{MATCHED_VAR_NAME} too long (%{MATCHED_VAR} bytes)'"
SecRule ARGS:password "@validateByteRange 33-244" \
    "id:11503,phase:2,deny,log,tag:'Login Allowlist',\
    msg:'Invalid parameter format: %{MATCHED_VAR_NAME} (%{MATCHED_VAR})'"
 
SecMarker END_ALLOWLIST_login
`

	t.Run("Assign an owasp custom rule to a sub virtual service", func(t *testing.T) {
		_, err := client.AddOwaspCustomRule("test", base64.StdEncoding.EncodeToString([]byte(custom_rule_content)))
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		response, err := client.AddVirtualServiceOwaspCustomRule(strconv.Itoa(subvs.Index), "test", true)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if response.Code != 200 {
			t.Fatalf("expected 200, got %d", response.Code)
		}
	})

	t.Run("Show an owasp custom rule to a sub virtual service", func(t *testing.T) {
		_, err := client.AddOwaspCustomRule("test", base64.StdEncoding.EncodeToString([]byte(custom_rule_content)))
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		_, err = client.AddVirtualServiceOwaspCustomRule(strconv.Itoa(subvs.Index), "test", false)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		response, err := client.ShowVirtualServiceOwaspRule(strconv.Itoa(subvs.Index), "test")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		assert.Equal(t, response.Code, 200)
		assert.Equal(t, response.Rule.Name, "test")
		assert.Equal(t, response.Rule.Type, "custom")
		assert.Equal(t, response.Rule.Enabled, "yes")
		assert.Equal(t, response.Rule.RunFirst, "no")
	})

	t.Run("Unassign an owasp custom rule to a sub virtual service", func(t *testing.T) {
		_, err := client.AddOwaspCustomRule("test", base64.StdEncoding.EncodeToString([]byte(custom_rule_content)))
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		_, err = client.AddVirtualServiceOwaspCustomRule(strconv.Itoa(subvs.Index), "test", true)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		response, err := client.ShowVirtualServiceOwaspRule(strconv.Itoa(subvs.Index), "test")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		assert.Equal(t, response.Code, 200)
		assert.Equal(t, response.Rule.Name, "test")
		assert.Equal(t, response.Rule.Type, "custom")
		assert.Equal(t, response.Rule.Enabled, "yes")
		assert.Equal(t, response.Rule.RunFirst, "yes")

		_, err = client.DeleteVirtualServiceOwaspCustomRule(strconv.Itoa(subvs.Index), "test")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		response, err = client.ShowVirtualServiceOwaspRule(strconv.Itoa(subvs.Index), "test")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		assert.Equal(t, response.Code, 200)
		assert.Equal(t, response.Rule.Name, "test")
		assert.Equal(t, response.Rule.Type, "custom")
		assert.Equal(t, response.Rule.Enabled, "no")

	})

}

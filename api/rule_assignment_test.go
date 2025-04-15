package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_AddRealServerRuleAssignment(t *testing.T) {
	testCases := []struct {
		name     string
		response string
		want     *LoadMasterResponse
		wantErr  bool
	}{
		{"success response", `{"code": 200, "message": "OK", "status": "success"}`, &LoadMasterResponse{Code: 200, Message: "OK", Status: "success"}, false},
		{"fail response", `{"code": 400, "message": "NOK", "message": "error"}`, nil, true},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				_, err := rw.Write([]byte(tt.response))
				if err != nil {
					fmt.Printf("Write failed: %v", err)
				}
			}))

			defer server.Close()
			client := createClientForUnit(server, "baz")

			rs, err := client.AddRealServerRule("test", "test", "test")

			if (err != nil) != tt.wantErr {
				t.Errorf("Client.AddRealServerRule() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(rs, tt.want) {
				t.Errorf("Client.AddRealServerRule() = %v, want %v", rs, tt.want)
			}
		})
	}
}

func TestIntegration_RealServerRuleAssignment(t *testing.T) {
	client, function := createClientForIntegration()
	if client == nil || function == nil {
		t.Skip("Skipping test because LOADMASTER_API_KEY or LOADMASTER_IP is not set")
	}
	defer function()

	vs1, err := client.AddVirtualService("10.0.0.4", "30500", "tcp", VirtualServiceParameters{VirtualServiceParametersBasicProperties: &VirtualServiceParametersBasicProperties{VSType: "http"}})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	vs2, err := client.AddVirtualService("10.0.0.4", "30501", "tcp", VirtualServiceParameters{VirtualServiceParametersBasicProperties: &VirtualServiceParametersBasicProperties{VSType: "http"}})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	subvs, err := client.AddSubVirtualService(vs1.Index, VirtualServiceParameters{})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	subvs_id := subvs.SubVS[len(subvs.SubVS)-1].VSIndex
	assert.NotEqual(t, "0", subvs_id)
	assert.NotEqual(t, vs1.Index, subvs_id)

	real_server1, err := client.AddRealServer(strconv.Itoa(subvs_id), "10.0.0.100", "80", RealServerParameters{})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	assert.Equal(t, 200, real_server1.Code)

	real_server2, err := client.AddRealServer(strconv.Itoa(vs2.Index), "10.0.0.101", "80", RealServerParameters{})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	assert.Equal(t, 200, real_server2.Code)

	match_content_rule1, err := client.AddRule("0", "rule1", GeneralRule{
		Pattern: convert2Ptr("old-body"),
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	assert.Equal(t, 200, match_content_rule1.Code)

	match_content_rule2, err := client.AddRule("0", "rule11", GeneralRule{
		Pattern: convert2Ptr("old-body"),
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	assert.Equal(t, 200, match_content_rule2.Code)

	add_header_rule, err := client.AddRule("1", "rule2", GeneralRule{
		Header:      convert2Ptr("X-HEADER"),
		Replacement: convert2Ptr("VALUE"),
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	assert.Equal(t, 200, add_header_rule.Code)

	modify_response_body, err := client.AddRule("5", "rule3", GeneralRule{
		Pattern:     convert2Ptr("X-HEADER"),
		Replacement: convert2Ptr("VALUE"),
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	assert.Equal(t, 200, modify_response_body.Code)

	t.Run("Show a new match content rule assignment for real server", func(t *testing.T) {
		init_rule_assignment, err := client.AddRealServerRule(strconv.Itoa(vs2.Index), "!"+strconv.Itoa(real_server2.Rs[0].RsIndex), match_content_rule1.MatchContentRules[0].Name)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		assert.Equal(t, 200, init_rule_assignment.Code)
		assert.Equal(t, "ok", init_rule_assignment.Status)

		rule_assignment, err := client.ShowRealServerRule(strconv.Itoa(vs2.Index), "!"+strconv.Itoa(real_server2.Rs[0].RsIndex), match_content_rule1.MatchContentRules[0].Name)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		assert.Equal(t, 200, rule_assignment.Code)
		assert.Equal(t, "ok", rule_assignment.Status)
	})

	t.Run("Deleting a new match content rule assignment for real server", func(t *testing.T) {
		init_rule_assignment, err := client.AddRealServerRule(strconv.Itoa(vs2.Index), "!"+strconv.Itoa(real_server2.Rs[0].RsIndex), match_content_rule2.MatchContentRules[0].Name)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		assert.Equal(t, 200, init_rule_assignment.Code)
		assert.Equal(t, "ok", init_rule_assignment.Status)

		rule_assignment, err := client.DeleteRealServerRule(strconv.Itoa(vs2.Index), "!"+strconv.Itoa(real_server2.Rs[0].RsIndex), match_content_rule2.MatchContentRules[0].Name)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		assert.Equal(t, 200, rule_assignment.Code)
		assert.Equal(t, "ok", rule_assignment.Status)

		_, err = client.ShowRealServerRule(strconv.Itoa(vs2.Index), "!"+strconv.Itoa(real_server2.Rs[0].RsIndex), match_content_rule2.MatchContentRules[0].Name)
		if err == nil {
			t.Fatalf("expected error, got no error")
		}
	})

	t.Run("Adding a new add header rule assignment for real server fails", func(t *testing.T) {
		_, err := client.AddRealServerRule(strconv.Itoa(vs2.Index), "!"+strconv.Itoa(real_server2.Rs[0].RsIndex), add_header_rule.AddHeaderRules[0].Name)
		if err == nil {
			t.Fatalf("expected error, got no error")
		}
	})
	t.Run("Adding a new modify response body rule assignment for real server fails", func(t *testing.T) {
		_, err := client.AddRealServerRule(strconv.Itoa(vs2.Index), "!"+strconv.Itoa(real_server2.Rs[0].RsIndex), modify_response_body.ReplaceBodyRules[0].Name)
		if err == nil {
			t.Fatalf("expected error, got no error")
		}
	})
}

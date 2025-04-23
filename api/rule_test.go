package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient_AddRule(t *testing.T) {
	testCases := []struct {
		name     string
		response string
		want     *RuleResponse
		wantErr  bool
	}{
		{"success response", `{"code": 200, "message": "OK", "status": "success"}`, &RuleResponse{LoadMasterResponse: &LoadMasterResponse{Code: 200, Message: "OK", Status: "success"}}, false},
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

			rs, err := client.AddRule("0", "test", GeneralRule{})

			if (err != nil) != tt.wantErr {
				t.Errorf("Client.AddRule() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(rs, tt.want) {
				t.Errorf("Client.AddRule() = %v, want %v", rs, tt.want)
			}
		})
	}
}

func TestIntegration_MatchContentRules(t *testing.T) {
	client, function := createClientForIntegration()
	if client == nil || function == nil {
		t.Skip("Skipping test because LOADMASTER_API_KEY or LOADMASTER_IP is not set")
	}
	defer function()

	t.Run("Adding a new match rule for url", func(t *testing.T) {
		rule, err := client.AddRule("0", "rule1", GeneralRule{
			MatchType:  convert2Ptr("regex"),
			Pattern:    convert2Ptr("^/test"),
			IncHost:    convert2Ptr(true),
			IncQuery:   convert2Ptr(true),
			NoCase:     convert2Ptr(true),
			Negate:     convert2Ptr(false),
			SetOnMatch: convert2Ptr(int32(4)),
			MustFail:   convert2Ptr(false),
		})

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		assert.Equal(t, 200, rule.Code)
		assert.Equal(t, "ok", rule.Status)
		assert.Len(t, rule.MatchContentRules, 1)
		assert.Equal(t, "rule1", rule.MatchContentRules[0].Name)
		assert.Equal(t, "Regex", rule.MatchContentRules[0].MatchType)
		assert.Equal(t, "^/test", rule.MatchContentRules[0].Pattern)

		require.NotNil(t, rule.MatchContentRules[0].IncHost)
		assert.True(t, *rule.MatchContentRules[0].IncHost)

		require.NotNil(t, rule.MatchContentRules[0].IncQuery)
		assert.True(t, *rule.MatchContentRules[0].IncQuery)

		require.NotNil(t, rule.MatchContentRules[0].CaseIndependent)
		assert.True(t, *rule.MatchContentRules[0].CaseIndependent)

		require.NotNil(t, rule.MatchContentRules[0].Negate)
		assert.False(t, *rule.MatchContentRules[0].Negate)

		require.NotNil(t, rule.MatchContentRules[0].SetOnMatch)
		assert.Equal(t, int32(4), *rule.MatchContentRules[0].SetOnMatch)

		require.NotNil(t, rule.MatchContentRules[0].MustFail)
		assert.False(t, *rule.MatchContentRules[0].MustFail)

	})

	t.Run("Adding a new match rule for header", func(t *testing.T) {
		rule, err := client.AddRule("0", "rule2", GeneralRule{
			MatchType: convert2Ptr("prefix"),
			Header:    convert2Ptr("X-Header"),
			Pattern:   convert2Ptr("^/test"),
			IncHost:   convert2Ptr(false),
			IncQuery:  convert2Ptr(false),
			NoCase:    convert2Ptr(true),
			Negate:    convert2Ptr(true),
			MustFail:  convert2Ptr(true),
		})

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		assert.Equal(t, 200, rule.Code)
		assert.Equal(t, "ok", rule.Status)

		assert.Len(t, rule.MatchContentRules, 1)
		assert.Equal(t, "rule2", rule.MatchContentRules[0].Name)
		assert.Equal(t, "prefix", rule.MatchContentRules[0].MatchType)
		require.NotNil(t, rule.MatchContentRules[0].Header)
		assert.Equal(t, "X-Header", *rule.MatchContentRules[0].Header)
		assert.Equal(t, "^/test", rule.MatchContentRules[0].Pattern)

		require.NotNil(t, rule.MatchContentRules[0].IncHost)
		assert.False(t, *rule.MatchContentRules[0].IncHost)

		require.NotNil(t, rule.MatchContentRules[0].IncQuery)
		assert.False(t, *rule.MatchContentRules[0].IncQuery)

		require.NotNil(t, rule.MatchContentRules[0].CaseIndependent)
		assert.True(t, *rule.MatchContentRules[0].CaseIndependent)

		require.NotNil(t, rule.MatchContentRules[0].Negate)
		assert.True(t, *rule.MatchContentRules[0].Negate)

		require.NotNil(t, rule.MatchContentRules[0].MustFail)
		assert.True(t, *rule.MatchContentRules[0].MustFail)
	})

	t.Run("Adding a new match rule for body", func(t *testing.T) {
		rule, err := client.AddRule("0", "rule3", GeneralRule{
			MatchType:    convert2Ptr("postfix"),
			Header:       convert2Ptr("body"),
			Pattern:      convert2Ptr("^/test"),
			IncHost:      convert2Ptr(false),
			IncQuery:     convert2Ptr(false),
			NoCase:       convert2Ptr(true),
			Negate:       convert2Ptr(false),
			MustFail:     convert2Ptr(true),
			SetOnMatch:   convert2Ptr(int32(8)),
			OnlyOnFlag:   convert2Ptr(int32(2)),
			OnlyOnNoFlag: convert2Ptr(int32(3)),
		})

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		assert.Equal(t, 200, rule.Code)
		assert.Equal(t, "ok", rule.Status)

		assert.Len(t, rule.MatchContentRules, 1)
		assert.Equal(t, "rule3", rule.MatchContentRules[0].Name)
		assert.Equal(t, "postfix", rule.MatchContentRules[0].MatchType)
		require.NotNil(t, rule.MatchContentRules[0].Header)
		assert.Equal(t, "body", *rule.MatchContentRules[0].Header)
		assert.Equal(t, "^/test", rule.MatchContentRules[0].Pattern)

		require.NotNil(t, rule.MatchContentRules[0].IncHost)
		assert.False(t, *rule.MatchContentRules[0].IncHost)

		require.NotNil(t, rule.MatchContentRules[0].IncQuery)
		assert.False(t, *rule.MatchContentRules[0].IncQuery)

		require.NotNil(t, rule.MatchContentRules[0].CaseIndependent)
		assert.True(t, *rule.MatchContentRules[0].CaseIndependent)

		require.NotNil(t, rule.MatchContentRules[0].Negate)
		assert.False(t, *rule.MatchContentRules[0].Negate)

		require.NotNil(t, rule.MatchContentRules[0].MustFail)
		assert.True(t, *rule.MatchContentRules[0].MustFail)

		require.NotNil(t, rule.MatchContentRules[0].SetOnMatch)
		assert.Equal(t, int32(8), *rule.MatchContentRules[0].SetOnMatch)

		require.NotNil(t, rule.MatchContentRules[0].OnlyOnFlag)
		assert.Equal(t, int32(2), *rule.MatchContentRules[0].OnlyOnFlag)

		require.NotNil(t, rule.MatchContentRules[0].OnlyOnNoFlag)
		assert.Equal(t, int32(3), *rule.MatchContentRules[0].OnlyOnNoFlag)
	})

}

func TestIntegration_AddHeaderRules(t *testing.T) {
	client, function := createClientForIntegration()
	if client == nil || function == nil {
		t.Skip("Skipping test because LOADMASTER_API_KEY or LOADMASTER_IP is not set")
	}
	defer function()

	t.Run("Adding a new add header rule", func(t *testing.T) {
		rule, err := client.AddRule("1", "rule4", GeneralRule{
			Header:       convert2Ptr("X-HEADER"),
			Replacement:  convert2Ptr("test"),
			OnlyOnFlag:   convert2Ptr(int32(1)),
			OnlyOnNoFlag: convert2Ptr(int32(2)),
		})

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		assert.Equal(t, 200, rule.Code)
		assert.Equal(t, "ok", rule.Status)

		assert.Len(t, rule.AddHeaderRules, 1)
		assert.Equal(t, "rule4", rule.AddHeaderRules[0].Name)
		require.NotNil(t, rule.AddHeaderRules[0].Header)
		assert.Equal(t, "X-HEADER", *rule.AddHeaderRules[0].Header)
		assert.Equal(t, "test", rule.AddHeaderRules[0].Replacement)

		require.NotNil(t, rule.AddHeaderRules[0].OnlyOnFlag)
		assert.Equal(t, int32(1), *rule.AddHeaderRules[0].OnlyOnFlag)

		require.NotNil(t, rule.AddHeaderRules[0].OnlyOnNoFlag)
		assert.Equal(t, int32(2), *rule.AddHeaderRules[0].OnlyOnNoFlag)
	})

	t.Run("Modifying an add header rule", func(t *testing.T) {
		init_rule, err := client.AddRule("1", "rule5", GeneralRule{
			Header:       convert2Ptr("X-HEADER"),
			Replacement:  convert2Ptr("test"),
			OnlyOnFlag:   convert2Ptr(int32(1)),
			OnlyOnNoFlag: convert2Ptr(int32(2)),
		})

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		assert.Equal(t, 200, init_rule.Code)
		assert.Equal(t, "ok", init_rule.Status)

		rule, err := client.ModifyRule("rule5", GeneralRule{
			Header:       convert2Ptr("X-HEADER-MOD"),
			Replacement:  convert2Ptr("test-modified"),
			OnlyOnFlag:   convert2Ptr(int32(3)),
			OnlyOnNoFlag: convert2Ptr(int32(4)),
		})
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		assert.Equal(t, 200, rule.Code)
		assert.Equal(t, "ok", rule.Status)

		assert.Len(t, rule.AddHeaderRules, 1)
		assert.Equal(t, "rule5", rule.AddHeaderRules[0].Name)
		require.NotNil(t, rule.AddHeaderRules[0].Header)
		assert.Equal(t, "X-HEADER-MOD", *rule.AddHeaderRules[0].Header)
		assert.Equal(t, "test-modified", rule.AddHeaderRules[0].Replacement)

		require.NotNil(t, rule.AddHeaderRules[0].OnlyOnFlag)
		assert.Equal(t, int32(3), *rule.AddHeaderRules[0].OnlyOnFlag)

		require.NotNil(t, rule.AddHeaderRules[0].OnlyOnNoFlag)
		assert.Equal(t, int32(4), *rule.AddHeaderRules[0].OnlyOnNoFlag)
	})

	t.Run("Show an add header rule", func(t *testing.T) {
		init_rule, err := client.AddRule("1", "rule6", GeneralRule{
			Header:       convert2Ptr("X-HEADER"),
			Replacement:  convert2Ptr("test"),
			OnlyOnNoFlag: convert2Ptr(int32(5)),
		})

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		assert.Equal(t, 200, init_rule.Code)
		assert.Equal(t, "ok", init_rule.Status)

		rule, err := client.ShowRule("rule6")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		assert.Equal(t, 200, rule.Code)
		assert.Equal(t, "ok", rule.Status)

		assert.Len(t, rule.AddHeaderRules, 1)
		assert.Equal(t, "rule6", rule.AddHeaderRules[0].Name)
		require.NotNil(t, rule.AddHeaderRules[0].Header)
		assert.Equal(t, "X-HEADER", *rule.AddHeaderRules[0].Header)
		assert.Equal(t, "test", rule.AddHeaderRules[0].Replacement)

		require.Nil(t, rule.AddHeaderRules[0].OnlyOnFlag)

		require.NotNil(t, rule.AddHeaderRules[0].OnlyOnNoFlag)
		assert.Equal(t, int32(5), *rule.AddHeaderRules[0].OnlyOnNoFlag)
	})
	t.Run("Delete an add header rule", func(t *testing.T) {
		init_rule, err := client.AddRule("1", "rule7", GeneralRule{
			Header:       convert2Ptr("X-HEADER"),
			Replacement:  convert2Ptr("test"),
			OnlyOnNoFlag: convert2Ptr(int32(5)),
		})

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		assert.Equal(t, 200, init_rule.Code)
		assert.Equal(t, "ok", init_rule.Status)

		rule, err := client.DeleteRule("rule7")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		assert.Equal(t, 200, rule.Code)
		assert.Equal(t, "ok", rule.Status)

		_, err = client.ShowRule("rule7")
		if err == nil {
			t.Fatalf("expected error, got no error")
		}
	})
}

func TestIntegration_DeleteHeaderRules(t *testing.T) {
	client, function := createClientForIntegration()
	if client == nil || function == nil {
		t.Skip("Skipping test because LOADMASTER_API_KEY or LOADMASTER_IP is not set")
	}
	defer function()

	t.Run("Adding a new delete header rule", func(t *testing.T) {
		rule, err := client.AddRule("2", "rule8", GeneralRule{
			Pattern:      convert2Ptr("X-HEADER"),
			OnlyOnFlag:   convert2Ptr(int32(1)),
			OnlyOnNoFlag: convert2Ptr(int32(2)),
		})

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		assert.Equal(t, 200, rule.Code)
		assert.Equal(t, "ok", rule.Status)

		assert.Len(t, rule.DeleteHeaderRules, 1)
		assert.Equal(t, "rule8", rule.DeleteHeaderRules[0].Name)
		require.NotNil(t, rule.DeleteHeaderRules[0].Pattern)
		assert.Equal(t, "X-HEADER", rule.DeleteHeaderRules[0].Pattern)

		require.NotNil(t, rule.DeleteHeaderRules[0].OnlyOnFlag)
		assert.Equal(t, int32(1), *rule.DeleteHeaderRules[0].OnlyOnFlag)

		require.NotNil(t, rule.DeleteHeaderRules[0].OnlyOnNoFlag)
		assert.Equal(t, int32(2), *rule.DeleteHeaderRules[0].OnlyOnNoFlag)
	})

	t.Run("Modifying an delete header rule", func(t *testing.T) {
		init_rule, err := client.AddRule("2", "rule9", GeneralRule{
			Pattern:      convert2Ptr("X-HEADER"),
			OnlyOnFlag:   convert2Ptr(int32(1)),
			OnlyOnNoFlag: convert2Ptr(int32(2)),
		})

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		assert.Equal(t, 200, init_rule.Code)
		assert.Equal(t, "ok", init_rule.Status)

		rule, err := client.ModifyRule("rule9", GeneralRule{
			Pattern:      convert2Ptr("X-HEADER-MOD"),
			OnlyOnFlag:   convert2Ptr(int32(3)),
			OnlyOnNoFlag: convert2Ptr(int32(4)),
		})

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		assert.Equal(t, 200, init_rule.Code)
		assert.Equal(t, "ok", init_rule.Status)

		assert.Len(t, rule.DeleteHeaderRules, 1)
		assert.Equal(t, "rule9", rule.DeleteHeaderRules[0].Name)
		require.NotNil(t, rule.DeleteHeaderRules[0].Pattern)
		assert.Equal(t, "X-HEADER-MOD", rule.DeleteHeaderRules[0].Pattern)

		require.NotNil(t, rule.DeleteHeaderRules[0].OnlyOnFlag)
		assert.Equal(t, int32(3), *rule.DeleteHeaderRules[0].OnlyOnFlag)

		require.NotNil(t, rule.DeleteHeaderRules[0].OnlyOnNoFlag)
		assert.Equal(t, int32(4), *rule.DeleteHeaderRules[0].OnlyOnNoFlag)
	})

	t.Run("Show an delete header rule", func(t *testing.T) {
		init_rule, err := client.AddRule("2", "rule10", GeneralRule{
			Pattern:      convert2Ptr("X-HEADER"),
			OnlyOnNoFlag: convert2Ptr(int32(2)),
		})

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		assert.Equal(t, 200, init_rule.Code)
		assert.Equal(t, "ok", init_rule.Status)

		rule, err := client.ShowRule("rule10")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		assert.Equal(t, 200, rule.Code)
		assert.Equal(t, "ok", rule.Status)

		assert.Len(t, rule.DeleteHeaderRules, 1)
		assert.Equal(t, "rule10", rule.DeleteHeaderRules[0].Name)
		assert.Equal(t, "X-HEADER", rule.DeleteHeaderRules[0].Pattern)

		require.Nil(t, rule.DeleteHeaderRules[0].OnlyOnFlag)

		require.NotNil(t, rule.DeleteHeaderRules[0].OnlyOnNoFlag)
		assert.Equal(t, int32(2), *rule.DeleteHeaderRules[0].OnlyOnNoFlag)
	})

	t.Run("Delete an delete header rule", func(t *testing.T) {
		init_rule, err := client.AddRule("2", "rule11", GeneralRule{
			Pattern:      convert2Ptr("X-HEADER"),
			OnlyOnNoFlag: convert2Ptr(int32(2)),
		})

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		assert.Equal(t, 200, init_rule.Code)
		assert.Equal(t, "ok", init_rule.Status)

		rule, err := client.DeleteRule("rule11")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		assert.Equal(t, 200, rule.Code)
		assert.Equal(t, "ok", rule.Status)

		_, err = client.ShowRule("rule11")
		if err == nil {
			t.Fatalf("expected error, got no error")
		}
	})

}

func TestIntegration_ReplaceHeaderRules(t *testing.T) {
	client, function := createClientForIntegration()
	if client == nil || function == nil {
		t.Skip("Skipping test because LOADMASTER_API_KEY or LOADMASTER_IP is not set")
	}
	defer function()

	t.Run("Adding a new replace header rule", func(t *testing.T) {
		rule, err := client.AddRule("3", "rule12", GeneralRule{
			Header:      convert2Ptr("X-HEADER"),
			Replacement: convert2Ptr("new-value"),
			Pattern:     convert2Ptr("old-value"),
		})

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		assert.Equal(t, 200, rule.Code)
		assert.Equal(t, "ok", rule.Status)

		assert.Len(t, rule.ReplaceHeaderRules, 1)
		assert.Equal(t, "rule12", rule.ReplaceHeaderRules[0].Name)
		require.NotNil(t, rule.ReplaceHeaderRules[0].Header)
		assert.Equal(t, "X-HEADER", *rule.ReplaceHeaderRules[0].Header)
		assert.Equal(t, "new-value", rule.ReplaceHeaderRules[0].Replacement)
		assert.Equal(t, "old-value", rule.ReplaceHeaderRules[0].Pattern)
	})
	t.Run("Modifying a replace header rule", func(t *testing.T) {
		init_rule, err := client.AddRule("3", "rule15", GeneralRule{
			Header:      convert2Ptr("X-HEADER"),
			Replacement: convert2Ptr("initial-value"),
			Pattern:     convert2Ptr("old-value"),
		})

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		assert.Equal(t, 200, init_rule.Code)
		assert.Equal(t, "ok", init_rule.Status)

		rule, err := client.ModifyRule("rule15", GeneralRule{
			Header:      convert2Ptr("X-HEADER-MOD"),
			Replacement: convert2Ptr("modified-value"),
			Pattern:     convert2Ptr("old-value-mod"),
		})

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		assert.Equal(t, 200, rule.Code)
		assert.Equal(t, "ok", rule.Status)

		assert.Len(t, rule.ReplaceHeaderRules, 1)
		assert.Equal(t, "rule15", rule.ReplaceHeaderRules[0].Name)
		require.NotNil(t, rule.ReplaceHeaderRules[0].Header)
		assert.Equal(t, "X-HEADER-MOD", *rule.ReplaceHeaderRules[0].Header)
		assert.Equal(t, "modified-value", rule.ReplaceHeaderRules[0].Replacement)
		assert.Equal(t, "old-value-mod", rule.ReplaceHeaderRules[0].Pattern)
	})

	t.Run("Show a replace header rule", func(t *testing.T) {
		init_rule, err := client.AddRule("3", "rule16", GeneralRule{
			Header:      convert2Ptr("X-HEADER"),
			Replacement: convert2Ptr("value"),
			Pattern:     convert2Ptr("old-value"),
		})

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		assert.Equal(t, 200, init_rule.Code)
		assert.Equal(t, "ok", init_rule.Status)

		rule, err := client.ShowRule("rule16")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		assert.Equal(t, 200, rule.Code)
		assert.Equal(t, "ok", rule.Status)

		assert.Len(t, rule.ReplaceHeaderRules, 1)
		assert.Equal(t, "rule16", rule.ReplaceHeaderRules[0].Name)
		require.NotNil(t, rule.ReplaceHeaderRules[0].Header)
		assert.Equal(t, "X-HEADER", *rule.ReplaceHeaderRules[0].Header)
		assert.Equal(t, "value", rule.ReplaceHeaderRules[0].Replacement)
		assert.Equal(t, "old-value", rule.ReplaceHeaderRules[0].Pattern)
	})

	t.Run("Delete a replace header rule", func(t *testing.T) {
		init_rule, err := client.AddRule("3", "rule17", GeneralRule{
			Header:      convert2Ptr("X-HEADER"),
			Replacement: convert2Ptr("value"),
			Pattern:     convert2Ptr("old-value"),
		})

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		assert.Equal(t, 200, init_rule.Code)
		assert.Equal(t, "ok", init_rule.Status)

		rule, err := client.DeleteRule("rule17")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		assert.Equal(t, 200, rule.Code)
		assert.Equal(t, "ok", rule.Status)

		_, err = client.ShowRule("rule17")
		if err == nil {
			t.Fatalf("expected error, got no error")
		}
	})
}

func TestIntegration_ModifyURLRules(t *testing.T) {
	client, function := createClientForIntegration()
	if client == nil || function == nil {
		t.Skip("Skipping test because LOADMASTER_API_KEY or LOADMASTER_IP is not set")
	}
	defer function()

	t.Run("Adding a new modify URL rule", func(t *testing.T) {
		rule, err := client.AddRule("4", "rule13", GeneralRule{
			Pattern:      convert2Ptr("old-path"),
			Replacement:  convert2Ptr("new-path"),
			OnlyOnFlag:   convert2Ptr(int32(1)),
			OnlyOnNoFlag: convert2Ptr(int32(2)),
		})

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		assert.Equal(t, 200, rule.Code)
		assert.Equal(t, "ok", rule.Status)

		assert.Len(t, rule.ModifyURLRules, 1)
		assert.Equal(t, "rule13", rule.ModifyURLRules[0].Name)
		assert.Equal(t, "old-path", rule.ModifyURLRules[0].Pattern)
		assert.Equal(t, "new-path", rule.ModifyURLRules[0].Replacement)

		require.NotNil(t, rule.ModifyURLRules[0].OnlyOnFlag)
		assert.Equal(t, int32(1), *rule.ModifyURLRules[0].OnlyOnFlag)

		require.NotNil(t, rule.ModifyURLRules[0].OnlyOnNoFlag)
		assert.Equal(t, int32(2), *rule.ModifyURLRules[0].OnlyOnNoFlag)
	})

	t.Run("Modifying a modify URL rule", func(t *testing.T) {
		init_rule, err := client.AddRule("4", "rule18", GeneralRule{
			Pattern:     convert2Ptr("old-path"),
			Replacement: convert2Ptr("new-path"),
		})

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		assert.Equal(t, 200, init_rule.Code)
		assert.Equal(t, "ok", init_rule.Status)

		rule, err := client.ModifyRule("rule18", GeneralRule{
			Pattern:     convert2Ptr("modified-old-path"),
			Replacement: convert2Ptr("modified-new-path"),
		})

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		assert.Equal(t, 200, rule.Code)
		assert.Equal(t, "ok", rule.Status)

		assert.Len(t, rule.ModifyURLRules, 1)
		assert.Equal(t, "rule18", rule.ModifyURLRules[0].Name)
		assert.Equal(t, "modified-old-path", rule.ModifyURLRules[0].Pattern)
		assert.Equal(t, "modified-new-path", rule.ModifyURLRules[0].Replacement)
	})

	t.Run("Show a modify URL rule", func(t *testing.T) {
		init_rule, err := client.AddRule("4", "rule19", GeneralRule{
			Pattern:     convert2Ptr("old-path"),
			Replacement: convert2Ptr("new-path"),
		})

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		assert.Equal(t, 200, init_rule.Code)
		assert.Equal(t, "ok", init_rule.Status)

		rule, err := client.ShowRule("rule19")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		assert.Equal(t, 200, rule.Code)
		assert.Equal(t, "ok", rule.Status)

		assert.Len(t, rule.ModifyURLRules, 1)
		assert.Equal(t, "rule19", rule.ModifyURLRules[0].Name)
		assert.Equal(t, "old-path", rule.ModifyURLRules[0].Pattern)
		assert.Equal(t, "new-path", rule.ModifyURLRules[0].Replacement)
	})

	t.Run("Delete a modify URL rule", func(t *testing.T) {
		init_rule, err := client.AddRule("4", "rule20", GeneralRule{
			Pattern:     convert2Ptr("old-path"),
			Replacement: convert2Ptr("new-path"),
		})

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		assert.Equal(t, 200, init_rule.Code)
		assert.Equal(t, "ok", init_rule.Status)

		rule, err := client.DeleteRule("rule20")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		assert.Equal(t, 200, rule.Code)
		assert.Equal(t, "ok", rule.Status)

		_, err = client.ShowRule("rule20")
		if err == nil {
			t.Fatalf("expected error, got no error")
		}
	})
}

func TestIntegration_ReplaceResponseBodyRules(t *testing.T) {
	client, function := createClientForIntegration()
	if client == nil || function == nil {
		t.Skip("Skipping test because LOADMASTER_API_KEY or LOADMASTER_IP is not set")
	}
	defer function()

	t.Run("Adding a new replace body rule", func(t *testing.T) {
		rule, err := client.AddRule("5", "rule14", GeneralRule{
			Pattern:      convert2Ptr("old-body"),
			Replacement:  convert2Ptr("new-body"),
			OnlyOnFlag:   convert2Ptr(int32(3)),
			OnlyOnNoFlag: convert2Ptr(int32(4)),
		})

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		assert.Equal(t, 200, rule.Code)
		assert.Equal(t, "ok", rule.Status)

		assert.Len(t, rule.ReplaceBodyRules, 1)
		assert.Equal(t, "rule14", rule.ReplaceBodyRules[0].Name)
		assert.Equal(t, "old-body", rule.ReplaceBodyRules[0].Pattern)
		assert.Equal(t, "new-body", rule.ReplaceBodyRules[0].Replacement)

		require.NotNil(t, rule.ReplaceBodyRules[0].OnlyOnFlag)
		assert.Equal(t, int32(3), *rule.ReplaceBodyRules[0].OnlyOnFlag)

		require.NotNil(t, rule.ReplaceBodyRules[0].OnlyOnNoFlag)
		assert.Equal(t, int32(4), *rule.ReplaceBodyRules[0].OnlyOnNoFlag)
	})

	t.Run("Modifying a replace body rule", func(t *testing.T) {
		init_rule, err := client.AddRule("5", "rule21", GeneralRule{
			Pattern:     convert2Ptr("old-body"),
			Replacement: convert2Ptr("new-body"),
		})

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		assert.Equal(t, 200, init_rule.Code)
		assert.Equal(t, "ok", init_rule.Status)

		rule, err := client.ModifyRule("rule21", GeneralRule{
			Pattern:     convert2Ptr("modified-old-body"),
			Replacement: convert2Ptr("modified-new-body"),
		})

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		assert.Equal(t, 200, rule.Code)
		assert.Equal(t, "ok", rule.Status)

		assert.Len(t, rule.ReplaceBodyRules, 1)
		assert.Equal(t, "rule21", rule.ReplaceBodyRules[0].Name)
		assert.Equal(t, "modified-old-body", rule.ReplaceBodyRules[0].Pattern)
		assert.Equal(t, "modified-new-body", rule.ReplaceBodyRules[0].Replacement)
	})

	t.Run("Show a replace body rule", func(t *testing.T) {
		init_rule, err := client.AddRule("5", "rule22", GeneralRule{
			Pattern:     convert2Ptr("old-body"),
			Replacement: convert2Ptr("new-body"),
		})

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		assert.Equal(t, 200, init_rule.Code)
		assert.Equal(t, "ok", init_rule.Status)

		rule, err := client.ShowRule("rule22")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		assert.Equal(t, 200, rule.Code)
		assert.Equal(t, "ok", rule.Status)

		assert.Len(t, rule.ReplaceBodyRules, 1)
		assert.Equal(t, "rule22", rule.ReplaceBodyRules[0].Name)
		assert.Equal(t, "old-body", rule.ReplaceBodyRules[0].Pattern)
		assert.Equal(t, "new-body", rule.ReplaceBodyRules[0].Replacement)
	})

	t.Run("Delete a replace body rule", func(t *testing.T) {
		init_rule, err := client.AddRule("5", "rule23", GeneralRule{
			Pattern:     convert2Ptr("old-body"),
			Replacement: convert2Ptr("new-body"),
		})

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		assert.Equal(t, 200, init_rule.Code)
		assert.Equal(t, "ok", init_rule.Status)

		rule, err := client.DeleteRule("rule23")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		assert.Equal(t, 200, rule.Code)
		assert.Equal(t, "ok", rule.Status)

		_, err = client.ShowRule("rule23")
		if err == nil {
			t.Fatalf("expected error, got no error")
		}
	})

}

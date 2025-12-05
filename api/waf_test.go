package api

import "testing"

func TestIntegration_Waf(t *testing.T) {
	client, function := createClientForIntegration()
	if client == nil || function == nil {
		t.Skip("Skipping test because LOADMASTER_API_KEY or LOADMASTER_IP is not set")
	}

	defer function()

	t.Run("Adding new WAF rule", func(t *testing.T) {
		response, err := client.AddWafRule("test_rule.txt", "This is a test WAF rule.")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if response.Code != 200 || response.Status != "ok" {
			t.Fatalf("expected success response, got code %d and status %s", response.Code, response.Status)
		}
	})

	t.Run("Showing WAF rule", func(t *testing.T) {
		response, err := client.ShowWafRule("test_rule.txt")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if response.Code != 200 || response.Status != "ok" {
			t.Fatalf("expected success response, got code %d and status %s", response.Code, response.Status)
		}

		if response.Data != "This is a test WAF rule." {
			t.Fatalf("expected data to be 'This is a test WAF rule.', got %s", response.Data)
		}
	})
	t.Run("Deleting WAF rule", func(t *testing.T) {
		response, err := client.DeleteWafRule("test_rule.txt")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if response.Code != 200 || response.Status != "ok" {
			t.Fatalf("expected success response, got code %d and status %s", response.Code, response.Status)
		}
	})
}

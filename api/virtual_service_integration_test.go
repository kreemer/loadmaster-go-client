package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	client, function := createClientForIntegration()
	if client == nil || function == nil {
		t.Skip("Skipping test because LOADMASTER_API_KEY or LOADMASTER_IP is not set")
	}

	defer function()

	t.Run("Adding new virtual service", func(t *testing.T) {
		response, err := client.AddVirtualService("10.0.0.4", "20001", "tcp", VirtualServiceParameters{})
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		assert.Equal(t, 200, response.Code)
		assert.Equal(t, "ok", response.Status)
		assert.Equal(t, "10.0.0.4", response.Address)
		assert.Equal(t, "20001", response.Port)
		assert.Equal(t, "tcp", response.Protocol)
		assert.Equal(t, "gen", response.VSType)
	})

	t.Run("Adding new virtual service with defined type", func(t *testing.T) {
		response, err := client.AddVirtualService("10.0.0.4", "20002", "tcp", VirtualServiceParameters{VirtualServiceParametersBasicProperties: &VirtualServiceParametersBasicProperties{VSType: "http"}})
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		assert.Equal(t, 200, response.Code)
		assert.Equal(t, "ok", response.Status)
		assert.Equal(t, "10.0.0.4", response.Address)
		assert.Equal(t, "20002", response.Port)
		assert.Equal(t, "tcp", response.Protocol)
		assert.Equal(t, "http", response.VSType)
	})

	t.Run("Adding new virtual service with defined nickname", func(t *testing.T) {
		response, err := client.AddVirtualService("10.0.0.4", "20003", "tcp", VirtualServiceParameters{VirtualServiceParametersBasicProperties: &VirtualServiceParametersBasicProperties{NickName: "test3"}})
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		assert.Equal(t, 200, response.Code)
		assert.Equal(t, "ok", response.Status)
		assert.Equal(t, "10.0.0.4", response.Address)
		assert.Equal(t, "20003", response.Port)
		assert.Equal(t, "tcp", response.Protocol)
		assert.Equal(t, "test3", response.NickName)
	})

	t.Run("Modify virtual service with name and enabled", func(t *testing.T) {
		init_response, err := client.AddVirtualService("10.0.0.4", "20004", "tcp", VirtualServiceParameters{})
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		assert.Equal(t, 200, init_response.Code)
		assert.Equal(t, "ok", init_response.Status)
		assert.Equal(t, "10.0.0.4", init_response.Address)
		assert.Equal(t, "20004", init_response.Port)
		assert.Equal(t, "tcp", init_response.Protocol)
		assert.Equal(t, "", init_response.NickName)
		assert.True(t, *init_response.Enable)

		response, err := client.ModifyVirtualService(init_response.Index, VirtualServiceParameters{VirtualServiceParametersBasicProperties: &VirtualServiceParametersBasicProperties{Enable: bool2ptr(false), NickName: "test4"}})
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		assert.Equal(t, 200, response.Code)
		assert.Equal(t, "ok", response.Status)
		assert.False(t, *response.Enable)
		assert.Equal(t, "test4", response.NickName)
	})

	t.Run("Delete virtual service", func(t *testing.T) {
		init_response, err := client.AddVirtualService("10.0.0.4", "20005", "tcp", VirtualServiceParameters{})
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		assert.Equal(t, 200, init_response.Code)
		assert.Equal(t, "ok", init_response.Status)
		assert.Equal(t, "10.0.0.4", init_response.Address)
		assert.Equal(t, "20005", init_response.Port)
		assert.Equal(t, "tcp", init_response.Protocol)

		response, err := client.DeleteVirtualService(init_response.Index)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		assert.Equal(t, 200, response.Code)
		assert.Equal(t, "ok", response.Status)

		check, err := client.ListVirtualService()
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		for _, vs := range check.VS {
			assert.NotEqual(t, "20005", vs.Port)
		}
	})

	t.Run("Show virtual service", func(t *testing.T) {
		init_response, err := client.AddVirtualService("10.0.0.4", "20006", "tcp", VirtualServiceParameters{})
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		assert.Equal(t, 200, init_response.Code)
		assert.Equal(t, "ok", init_response.Status)
		assert.Equal(t, "10.0.0.4", init_response.Address)
		assert.Equal(t, "20006", init_response.Port)
		assert.Equal(t, "tcp", init_response.Protocol)

		response, err := client.ShowVirtualService(init_response.Index)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		assert.Equal(t, 200, response.Code)
		assert.Equal(t, "ok", response.Status)
		assert.Equal(t, "10.0.0.4", response.Address)
		assert.Equal(t, "20006", response.Port)
		assert.Equal(t, "tcp", response.Protocol)

	})
}

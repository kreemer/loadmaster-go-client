package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegration_SubVirtualService(t *testing.T) {
	client, function := createClientForIntegration()
	if client == nil || function == nil {
		t.Skip("Skipping test because LOADMASTER_API_KEY or LOADMASTER_IP is not set")
	}
	defer function()

	vs, err := client.AddVirtualService("10.0.0.4", "30000", "tcp", VirtualServiceParameters{})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	t.Run("Adding new sub virtual service", func(t *testing.T) {
		response, err := client.AddSubVirtualService(vs.Index, VirtualServiceParameters{})
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		assert.Equal(t, 200, response.Code)
		assert.Equal(t, "ok", response.Status)
		assert.Equal(t, "10.0.0.4", response.Address)
		assert.Equal(t, "30000", response.Port)
		assert.Equal(t, "tcp", response.Protocol)
		assert.Equal(t, "gen", response.VSType)
	})
	t.Run("Adding new sub virtual service with defined type", func(t *testing.T) {
		init_response, err := client.AddSubVirtualService(vs.Index, VirtualServiceParameters{VirtualServiceParametersBasicProperties: &VirtualServiceParametersBasicProperties{VSType: "http"}})
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		subvs_id := init_response.SubVS[len(init_response.SubVS)-1].VSIndex
		assert.NotEqual(t, "0", subvs_id)
		assert.NotEqual(t, vs.Index, subvs_id)
		response, err := client.ShowSubVirtualService(subvs_id)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		assert.Equal(t, 200, response.Code)
		assert.Equal(t, "ok", response.Status)
		assert.Equal(t, "http", response.VSType)
	})
	t.Run("Modify sub virtual service with defined type", func(t *testing.T) {
		init_response, err := client.AddSubVirtualService(vs.Index, VirtualServiceParameters{VirtualServiceParametersBasicProperties: &VirtualServiceParametersBasicProperties{NickName: "subvs1", VSType: "gen"}})
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		assert.Equal(t, 200, init_response.Code)
		assert.Equal(t, "ok", init_response.Status)

		subvs_id := init_response.SubVS[len(init_response.SubVS)-1].VSIndex
		assert.NotEqual(t, "0", subvs_id)
		assert.NotEqual(t, vs.Index, subvs_id)

		response, err := client.ModifySubVirtualService(subvs_id, VirtualServiceParameters{VirtualServiceParametersBasicProperties: &VirtualServiceParametersBasicProperties{NickName: "subvs2", VSType: "http"}})
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		assert.Equal(t, 200, response.Code)
		assert.Equal(t, "ok", response.Status)

		assert.Equal(t, "tcp", response.Protocol)
		assert.Equal(t, "http", response.VSType)
		assert.Equal(t, "subvs2", response.NickName)
	})
	t.Run("Delete sub virtual service", func(t *testing.T) {
		init_response, err := client.AddSubVirtualService(vs.Index, VirtualServiceParameters{})
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		assert.Equal(t, 200, init_response.Code)
		assert.Equal(t, "ok", init_response.Status)
		assert.Equal(t, "10.0.0.4", init_response.Address)
		assert.Equal(t, "30000", init_response.Port)
		assert.Equal(t, "tcp", init_response.Protocol)

		subvs_id := init_response.SubVS[len(init_response.SubVS)-1].VSIndex
		assert.NotEqual(t, "0", subvs_id)
		assert.NotEqual(t, vs.Index, subvs_id)

		response, err := client.DeleteSubVirtualService(subvs_id)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		assert.Equal(t, 200, response.Code)
	})
	t.Run("Show sub virtual service", func(t *testing.T) {
		init_response, err := client.AddSubVirtualService(vs.Index, VirtualServiceParameters{})
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		assert.Equal(t, 200, init_response.Code)
		assert.Equal(t, "ok", init_response.Status)

		subvs_id := init_response.SubVS[len(init_response.SubVS)-1].VSIndex
		assert.NotEqual(t, "0", subvs_id)
		assert.NotEqual(t, vs.Index, subvs_id)

		response, err := client.ShowSubVirtualService(subvs_id)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		assert.Equal(t, 200, response.Code)
		assert.Equal(t, "ok", response.Status)
	})
}

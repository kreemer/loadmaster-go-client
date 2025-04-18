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

func TestClient_AddSubVirtualService(t *testing.T) {
	testCases := []struct {
		name     string
		response string
		want     *ShowSubVirtualServiceResponse
		wantErr  bool
	}{
		{"success response", `{"code": 200, "message": "OK", "status": "success"}`, &ShowSubVirtualServiceResponse{LoadMasterResponse: &LoadMasterResponse{Code: 200, Message: "OK", Status: "success"}}, false},
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

			rs, err := client.AddSubVirtualService(strconv.Itoa(int(1)), VirtualServiceParameters{})

			if (err != nil) != tt.wantErr {
				t.Errorf("Client.AddSubVirtualService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(rs, tt.want) {
				t.Errorf("Client.AddSubVirtualService() = %v, want %v", rs, tt.want)
			}
		})
	}
}

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
		response, err := client.AddSubVirtualService(strconv.Itoa(int(vs.Index)), VirtualServiceParameters{})
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
		init_response, err := client.AddSubVirtualService(strconv.Itoa(int(vs.Index)), VirtualServiceParameters{VirtualServiceParametersBasicProperties: &VirtualServiceParametersBasicProperties{VSType: "http"}})
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		subvs_id := init_response.SubVS[len(init_response.SubVS)-1].VSIndex
		assert.NotEqual(t, "0", subvs_id)
		assert.NotEqual(t, vs.Index, subvs_id)
		response, err := client.ShowSubVirtualService(strconv.Itoa(int(subvs_id)))
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		assert.Equal(t, 200, response.Code)
		assert.Equal(t, "ok", response.Status)
		assert.Equal(t, "http", response.VSType)
	})
	t.Run("Modify sub virtual service with defined type", func(t *testing.T) {
		init_response, err := client.AddSubVirtualService(strconv.Itoa(int(vs.Index)), VirtualServiceParameters{VirtualServiceParametersBasicProperties: &VirtualServiceParametersBasicProperties{NickName: "subvs1", VSType: "gen"}})
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		assert.Equal(t, 200, init_response.Code)
		assert.Equal(t, "ok", init_response.Status)

		subvs_id := init_response.SubVS[len(init_response.SubVS)-1].VSIndex
		assert.NotEqual(t, "0", subvs_id)
		assert.NotEqual(t, vs.Index, subvs_id)

		response, err := client.ModifySubVirtualService(strconv.Itoa(int(subvs_id)), VirtualServiceParameters{VirtualServiceParametersBasicProperties: &VirtualServiceParametersBasicProperties{NickName: "subvs2", VSType: "http"}})
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
		init_response, err := client.AddSubVirtualService(strconv.Itoa(int(vs.Index)), VirtualServiceParameters{})
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

		response, err := client.DeleteSubVirtualService(strconv.Itoa(int(subvs_id)))
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		assert.Equal(t, 200, response.Code)
	})
	t.Run("Show sub virtual service", func(t *testing.T) {
		init_response, err := client.AddSubVirtualService(strconv.Itoa(int(vs.Index)), VirtualServiceParameters{})
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		assert.Equal(t, 200, init_response.Code)
		assert.Equal(t, "ok", init_response.Status)

		subvs_id := init_response.SubVS[len(init_response.SubVS)-1].VSIndex
		assert.NotEqual(t, "0", subvs_id)
		assert.NotEqual(t, vs.Index, subvs_id)

		response, err := client.ShowSubVirtualService(strconv.Itoa(int(subvs_id)))
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		assert.Equal(t, 200, response.Code)
		assert.Equal(t, "ok", response.Status)
	})
}

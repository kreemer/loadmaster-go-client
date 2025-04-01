package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
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
			client := Client{server.Client(), "bar", "foo", "baz", server.URL}

			rs, err := client.AddSubVirtualService(1, VirtualServiceParameters{})

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

package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestClient_ShowRealServer(t *testing.T) {
	testCases := []struct {
		name     string
		response string
		want     *ListRealServerResponse
		wantErr  bool
	}{
		{"success response without Rs", `{"code": 200, "message": "OK", "status": "success"}`, &ListRealServerResponse{LoadMasterResponse: &LoadMasterResponse{Code: 200, Message: "OK", Status: "success"}}, false},
		{"success response with Rs", `{"code": 200, "message": "OK", "status": "success", "Rs": [ { "RSIndex": 1 } ]}`, &ListRealServerResponse{LoadMasterResponse: &LoadMasterResponse{Code: 200, Message: "OK", Status: "success"}, Rs: []RealServer{{RsIndex: 1}}}, false},
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

			rs, err := client.ShowRealServer("test", "test")

			if (err != nil) != tt.wantErr {
				t.Errorf("Client.ShowRealServer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(rs, tt.want) {
				t.Errorf("Client.ShowRealServer() = %v, want %v", rs, tt.want)
			}
		})
	}
}

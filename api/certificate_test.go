package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestClient_ListCertificate(t *testing.T) {
	testCases := []struct {
		name     string
		response string
		want     *ListCertResponse
		wantErr  bool
	}{
		{"success response", `{"code": 200, "message": "OK", "status": "success", "cert": [ { "name": "Example", "type": "RSA", "modulus": "EXAMPLE" } ]}`, &ListCertResponse{LoadMasterResponse: &LoadMasterResponse{Code: 200, Message: "OK", Status: "success"}, Cert: []CertInfo{{Name: "Example", Type: "RSA", Modulus: "EXAMPLE"}}}, false},
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

			rs, err := client.ListCertificate()

			if (err != nil) != tt.wantErr {
				t.Errorf("Client.ListCertificate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(rs, tt.want) {
				t.Errorf("Client.ListCertificate() = %v, want %v", rs, tt.want)
			}
		})
	}
}

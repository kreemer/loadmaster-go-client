package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestClient_ListApiKey(t *testing.T) {
	testCases := []struct {
		name     string
		response string
		want     *ListApiKeyResponse
		wantErr  bool
	}{
		{"success response", `{"code": 200, "message": "OK", "status": "success", "apikeys": ["foo", "bar"]}`, &ListApiKeyResponse{LoadMasterResponse: &LoadMasterResponse{Code: 200, Message: "OK", Status: "success"}, ApiKeys: []string{"foo", "bar"}}, false},
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
			client := Client{server.Client(), "bar", "foo", "baz", server.URL, 0}

			rs, err := client.ListApiKey()

			if (err != nil) != tt.wantErr {
				t.Errorf("Client.ListApiKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(rs, tt.want) {
				t.Errorf("Client.ListApiKey() = %v, want %v", rs, tt.want)
			}
		})
	}
}

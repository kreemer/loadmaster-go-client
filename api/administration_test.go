package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestClient_Backup(t *testing.T) {
	testCases := []struct {
		name         string
		arguments    []any
		response     string
		responseCode int
		want         *LoadMasterDataResponse
		wantErr      bool
	}{
		{"success response", []any{}, `{"code": 200, "message": "OK", "status": "success", "data": "..."}`, 200, &LoadMasterDataResponse{LoadMasterResponse: &LoadMasterResponse{Code: 200, Message: "OK", Status: "success"}, Data: "..."}, false},
		{"fail response", []any{}, `{"code": 400, "message": "NOK", "message": "error"}`, 400, nil, true},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				rw.WriteHeader(tt.responseCode)
				_, err := rw.Write([]byte(tt.response))
				if err != nil {
					fmt.Printf("Write failed: %v", err)
				}
			}))

			defer server.Close()
			client := createClientForUnit(server, "baz")

			rs, err := client.Backup()

			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Backup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(rs, tt.want) {
				t.Errorf("Client.Backup() = %v, want %v", rs, tt.want)
			}
		})
	}
}

func TestClient_Restore(t *testing.T) {
	testCases := []struct {
		name         string
		arguments    []any
		response     string
		responseCode int
		want         *LoadMasterResponse
		wantErr      bool
	}{
		{"success response", []any{"data", "14"}, `{"code": 200, "message": "OK", "status": "success"}`, 200, &LoadMasterResponse{Code: 200, Message: "OK", Status: "success"}, false},
		{"fail response", []any{"data", "14"}, `{"code": 400, "message": "NOK", "message": "error"}`, 400, nil, true},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				rw.WriteHeader(tt.responseCode)
				_, err := rw.Write([]byte(tt.response))
				if err != nil {
					fmt.Printf("Write failed: %v", err)
				}

			}))

			defer server.Close()
			client := createClientForUnit(server, "baz")

			data, _ := tt.arguments[0].(string)
			restore_type, _ := tt.arguments[1].(string)

			rs, err := client.Restore(data, restore_type)

			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Restore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(rs, tt.want) {
				t.Errorf("Client.Restore() = %v, want %v", rs, tt.want)
			}
		})
	}
}

package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestClient_RegisterLetsEncryptAccount(t *testing.T) {
	testCases := []struct {
		name      string
		arguments []any
		response  string
		want      *LoadMasterResponse
		wantErr   bool
	}{
		{"success response", []any{func(i string) *string { return &i }("mail")}, `{"code": 200, "message": "OK", "status": "success"}`, &LoadMasterResponse{Code: 200, Message: "OK", Status: "success"}, false},
		{"success response", []any{nil}, `{"code": 200, "message": "OK", "status": "success"}`, &LoadMasterResponse{Code: 200, Message: "OK", Status: "success"}, false},
		{"fail response", []any{func(i string) *string { return &i }("mail")}, `{"code": 400, "message": "NOK", "message": "error"}`, nil, true},
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

			email, _ := tt.arguments[0].(*string)

			rs, err := client.RegisterLetsEncryptAccount(email)

			if (err != nil) != tt.wantErr {
				t.Errorf("Client.RegisterLetsEncryptAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(rs, tt.want) {
				t.Errorf("Client.RegisterLetsEncryptAccount() = %v, want %v", rs, tt.want)
			}
		})
	}
}

func TestClient_FetchLetsEncryptAccount(t *testing.T) {
	testCases := []struct {
		name      string
		arguments []any
		response  string
		want      *LoadMasterResponse
		wantErr   bool
	}{
		{"success response", []any{"password", "data"}, `{"code": 200, "message": "OK", "status": "success"}`, &LoadMasterResponse{Code: 200, Message: "OK", Status: "success"}, false},
		{"fail response", []any{"password", "data"}, `{"code": 400, "message": "NOK", "message": "error"}`, nil, true},
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

			password, _ := tt.arguments[0].(string)
			data, _ := tt.arguments[1].(string)

			rs, err := client.FetchLetsEncryptAccount(password, data)

			if (err != nil) != tt.wantErr {
				t.Errorf("Client.FetchLetsEncryptAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(rs, tt.want) {
				t.Errorf("Client.FetchLetsEncryptAccount() = %v, want %v", rs, tt.want)
			}
		})
	}
}

func TestClient_SetDigicertKeyId(t *testing.T) {
	testCases := []struct {
		name      string
		arguments []any
		response  string
		want      *LoadMasterResponse
		wantErr   bool
	}{
		{"success response", []any{"key"}, `{"code": 200, "message": "OK", "status": "success"}`, &LoadMasterResponse{Code: 200, Message: "OK", Status: "success"}, false},
		{"fail response", []any{"key"}, `{"code": 400, "message": "NOK", "message": "error"}`, nil, true},
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

			key, _ := tt.arguments[0].(string)

			rs, err := client.SetDigicertKeyId(key)

			if (err != nil) != tt.wantErr {
				t.Errorf("Client.SetDigicertKeyId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(rs, tt.want) {
				t.Errorf("Client.SetDigicertKeyId() = %v, want %v", rs, tt.want)
			}
		})
	}
}

func TestClient_SetDigicertHMAC(t *testing.T) {
	testCases := []struct {
		name      string
		arguments []any
		response  string
		want      *LoadMasterResponse
		wantErr   bool
	}{
		{"success response", []any{"hmac"}, `{"code": 200, "message": "OK", "status": "success"}`, &LoadMasterResponse{Code: 200, Message: "OK", Status: "success"}, false},
		{"fail response", []any{"hmac"}, `{"code": 400, "message": "NOK", "message": "error"}`, nil, true},
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

			hmac, _ := tt.arguments[0].(string)

			rs, err := client.SetDigicertHMAC(hmac)

			if (err != nil) != tt.wantErr {
				t.Errorf("Client.SetDigicertHMAC() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(rs, tt.want) {
				t.Errorf("Client.SetDigicertHMAC() = %v, want %v", rs, tt.want)
			}
		})
	}
}

func TestClient_RequestACMECertificate(t *testing.T) {
	testCases := []struct {
		name      string
		arguments []any
		response  string
		want      *LoadMasterResponse
		wantErr   bool
	}{
		{"success response", []any{"name", "common", "1", "1", nil}, `{"code": 200, "message": "OK", "status": "success"}`, &LoadMasterResponse{Code: 200, Message: "OK", Status: "success"}, false},
		{"success response with params", []any{"name", "common", "1", "1", &RequestACMECertificateParameters{KeySize: 2048}}, `{"code": 200, "message": "OK", "status": "success"}`, &LoadMasterResponse{Code: 200, Message: "OK", Status: "success"}, false},
		{"fail response", []any{"name", "common", "1", "1", nil}, `{"code": 400, "message": "NOK", "message": "error"}`, nil, true},
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

			name, _ := tt.arguments[0].(string)
			common_name, _ := tt.arguments[1].(string)
			vs_identifier, _ := tt.arguments[2].(string)
			acme_type, _ := tt.arguments[3].(string)
			params, _ := tt.arguments[4].(*RequestACMECertificateParameters)

			rs, err := client.RequestACMECertificate(name, common_name, vs_identifier, acme_type, params)

			if (err != nil) != tt.wantErr {
				t.Errorf("Client.RequestACMECertificate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(rs, tt.want) {
				t.Errorf("Client.RequestACMECertificate() = %v, want %v", rs, tt.want)
			}
		})
	}
}

func TestClient_DeleteACMECertificate(t *testing.T) {
	testCases := []struct {
		name      string
		arguments []any
		response  string
		want      *LoadMasterResponse
		wantErr   bool
	}{
		{"success response", []any{"name", "1"}, `{"code": 200, "message": "OK", "status": "success"}`, &LoadMasterResponse{Code: 200, Message: "OK", Status: "success"}, false},
		{"fail response", []any{"name", "1"}, `{"code": 400, "message": "NOK", "status": "error"}`, nil, true},
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

			name, _ := tt.arguments[0].(string)
			acme_type, _ := tt.arguments[1].(string)

			rs, err := client.DeleteACMECertificate(name, acme_type)

			if (err != nil) != tt.wantErr {
				t.Errorf("Client.RequestACMECertificate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(rs, tt.want) {
				t.Errorf("Client.RequestACMECertificate() = %v, want %v", rs, tt.want)
			}
		})
	}
}

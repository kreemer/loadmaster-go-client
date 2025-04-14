package api

import (
	"log/slog"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type closerFunc func() error

func createClientForUnit(server *httptest.Server, key string) Client {
	logger := slog.New(slog.DiscardHandler)
	client := Client{server.Client(), "bar", "foo", key, server.URL, logger}

	return client
}

func createClientForIntegration() (*Client, closerFunc) {

	api_key := os.Getenv("LOADMASTER_API_KEY")
	ip := os.Getenv("LOADMASTER_IP")

	if api_key == "" || ip == "" {
		return nil, nil
	}
	client := NewClientWithApiKey(ip, api_key)
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}))
	// logger := slog.New(slog.DiscardHandler)

	client.SetLogger(logger)

	data, _ := client.Backup()

	cleanup := func() error {
		_, err := client.Restore(data.Data, "14")

		return err
	}

	return client, cleanup
}

func convert2Ptr[T any](object T) *T {
	return &object
}

func TestLoadMasterRequest_injectAuth(t *testing.T) {
	type fields struct {
		ApiUser string
		ApiPass string
		ApiKey  string
	}
	type args struct {
		c *Client
	}
	tests := []struct {
		name     string
		expected fields
		args     args
		wantErr  bool
	}{
		{"Take api key if defined", fields{ApiUser: "", ApiPass: "", ApiKey: "test"}, args{&Client{apiKey: "test", logger: slog.Default()}}, false},
		{"Take api username / password if defined", fields{ApiUser: "user", ApiPass: "pass", ApiKey: ""}, args{&Client{apiUser: "user", apiPass: "pass", logger: slog.Default()}}, false},
		{"Error if no authentication", fields{}, args{&Client{logger: slog.Default()}}, true},
		{"Error if no username but password", fields{}, args{&Client{apiPass: "pass", logger: slog.Default()}}, true},
		{"Error if no password but username", fields{}, args{&Client{apiUser: "user", logger: slog.Default()}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			r := &LoadMasterRequest{}
			if err := r.injectAuth(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("LoadMasterRequest.injectAuth() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			assert.Equal(t, tt.args.c.apiKey, tt.expected.ApiKey)
			assert.Equal(t, tt.args.c.apiUser, tt.expected.ApiUser)
			assert.Equal(t, tt.args.c.apiPass, tt.expected.ApiPass)
		})
	}
}

package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
		{"Take api key if defined", fields{ApiUser: "", ApiPass: "", ApiKey: "test"}, args{&Client{apiKey: "test"}}, false},
		{"Take api username / password if defined", fields{ApiUser: "user", ApiPass: "pass", ApiKey: ""}, args{&Client{apiUser: "user", apiPass: "pass"}}, false},
		{"Error if no authentication", fields{}, args{&Client{}}, true},
		{"Error if no username but password", fields{}, args{&Client{apiPass: "pass"}}, true},
		{"Error if no password but username", fields{}, args{&Client{apiUser: "user"}}, true},
		// TODO: Add test cases.
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

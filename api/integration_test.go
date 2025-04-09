package api

import (
	"log/slog"
	"os"
)

type closerFunc func() error

func createClientForIntegration() (*Client, closerFunc) {

	api_key := os.Getenv("LOADMASTER_API_KEY")
	ip := os.Getenv("LOADMASTER_IP")

	if api_key == "" || ip == "" {
		return nil, nil
	}
	client := NewClientWithApiKey(ip, api_key)
	client.SetLogger(slog.New(slog.DiscardHandler))

	data, _ := client.Backup()

	cleanup := func() error {
		_, err := client.Restore(data.Data, "14")

		return err
	}

	return client, cleanup
}

func bool2ptr(b bool) *bool {
	return &b
}

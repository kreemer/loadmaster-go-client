package api

import (
	"log/slog"
	"net/http/httptest"
)

func createClientForUnit(server *httptest.Server, key string) Client {
	logger := slog.New(slog.DiscardHandler)
	client := Client{server.Client(), "bar", "foo", key, server.URL, logger}

	return client
}

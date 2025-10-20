package server

import (
	"HttpServer/server/handlers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetRouter(t *testing.T) {
	srv := httptest.NewServer(GetRouter(&handlers.MyHandler{}))
	defer srv.Close()
	tests := []struct {
		name   string
		method string
		url    string
		status int
		body   string
	}{
		{
			name:   "GET last message",
			method: http.MethodGet,
			url:    srv.URL + "/messages/last",
			status: http.StatusOK,
		},
		{
			name:   "POST message",
			method: http.MethodPost,
			url:    srv.URL + "/",
			status: http.StatusBadRequest,
			body:   "asdasd",
		},
		{
			name:   "GET all messages",
			method: http.MethodGet,
			url:    srv.URL + "/messages/all",
			status: http.StatusOK,
		},
		{
			name:   "Non-existent method url",
			method: http.MethodGet,
			url:    srv.URL + "/nonexistent",
			status: http.StatusMethodNotAllowed,
		},
	}
	client := http.Client{}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var reqBody io.Reader
			if tc.body != "" {
				reqBody = strings.NewReader(tc.body)
			}
			req, err := http.NewRequest(tc.method, tc.url, reqBody)
			require.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")
			resp, err := client.Do(req)
			require.NoError(t, err)
			assert.Equal(t, tc.status, resp.StatusCode)
			require.NoError(t, resp.Body.Close())
		})
	}
}

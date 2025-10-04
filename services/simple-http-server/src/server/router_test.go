package server

import (
	"HttpServer/server/handlers"
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
			url:    srv.URL + "/", //ченкуть работает ли на любой
			status: http.StatusCreated,
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
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			switch tc.method {
			case http.MethodGet:
				resp, err := http.Get(tc.url)
				if err != nil || resp.StatusCode != tc.status {
					t.Errorf("GET (%s) failed, StatusCode: %v", tc.url, resp.StatusCode)
				}
			case http.MethodPost:
				resp, err := http.Post(tc.url, "application/json", strings.NewReader(`{"message":"Hello"}`))
				if err != nil || resp.StatusCode != tc.status {
					t.Errorf("POST (%s) failed, StatusCode: %v", tc.url, resp.StatusCode)
				}
			default:
				t.Errorf("GET (%s) method not allowed", tc.method)
			}
		})
	}
}

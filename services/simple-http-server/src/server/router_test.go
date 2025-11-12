package server

import (
	"HttpServer/server/handlers"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRouter(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mh := &handlers.MyHandler{}
	router := GetRouter(mh)

	tests := []struct {
		name       string
		method     string
		path       string
		body       any
		statusCode int
		respBody   any
	}{
		{
			name:       "GET last message empty",
			method:     http.MethodGet,
			path:       "/messages/last",
			statusCode: http.StatusOK,
			respBody:   "",
		},
		{
			name:       "GET all messages empty",
			method:     http.MethodGet,
			path:       "/messages/all",
			statusCode: http.StatusOK,
			respBody:   []string(nil),
		},
		{
			name:       "POST valid message",
			method:     http.MethodPost,
			path:       "/messages/add",
			body:       map[string]string{"message": "Hello"},
			statusCode: http.StatusCreated,
		},
		{
			name:       "POST invalid message",
			method:     http.MethodPost,
			path:       "/messages/add",
			body:       map[string]string{"msg": "oops"},
			statusCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var reqBody []byte
			if tt.body != nil {
				reqBody, _ = json.Marshal(tt.body)
			}

			req := httptest.NewRequest(tt.method, tt.path, bytes.NewReader(reqBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			assert.Equal(t, tt.statusCode, w.Code)

			if tt.respBody != nil {
				respData, _ := io.ReadAll(w.Body)
				switch v := tt.respBody.(type) {
				case string:
					assert.Equal(t, v, string(respData))
				case []string:
					var arr []string
					err := json.Unmarshal(respData, &arr)
					assert.NoError(t, err)
					assert.Equal(t, v, arr)
				}
			}
		})
	}
}

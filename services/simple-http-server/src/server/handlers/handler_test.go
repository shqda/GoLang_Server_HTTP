package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMyHandler_GetLastMessageHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name     string
		messages []string
		result   string
		status   int
	}{
		{"One message", []string{"Hello"}, "Hello", http.StatusOK},
		{"Multiple messages", []string{"Hello", "World"}, "World", http.StatusOK},
		{"Empty message", []string{}, "", http.StatusOK},
		{"Nil message", nil, "", http.StatusOK},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mh := &MyHandler{messages: tc.messages}

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest(http.MethodGet, "/messages/last", nil)

			mh.GetLastMessageHandler(c)

			assert.Equal(t, tc.status, w.Code)
			assert.Equal(t, tc.result, strings.TrimSpace(w.Body.String()))
		})
	}
}

func TestMyHandler_GetAllMessagesHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name     string
		messages []string
		result   []string
		status   int
	}{
		{"One message", []string{"Hello"}, []string{"Hello"}, http.StatusOK},
		{"Multiple messages", []string{"Hello", "World"}, []string{"Hello", "World"}, http.StatusOK},
		{"Empty message", []string{}, []string{}, http.StatusOK},
		{"Nil message", nil, []string{}, http.StatusOK},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mh := &MyHandler{messages: tc.messages}
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest(http.MethodGet, "/messages/all", nil)

			mh.GetAllMessagesHandler(c)

			assert.Equal(t, tc.status, w.Code)

			var resp []string
			err := json.Unmarshal(w.Body.Bytes(), &resp)

			assert.NoError(t, err)
			assert.Equal(t, tc.result, resp)
		})
	}
}

func TestMyHandler_CreateMessageHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name   string
		value  map[string]string
		status int
	}{
		{"Valid JSON", map[string]string{"message": "Hello World"}, http.StatusCreated},
		{"Invalid JSON", map[string]string{"msg": "aboba"}, http.StatusBadRequest},
		{"Empty JSON", map[string]string{}, http.StatusBadRequest},
		{"Nil map", nil, http.StatusBadRequest},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mh := &MyHandler{messages: []string{}}
			var bodyBytes []byte
			if tc.value != nil {
				bodyBytes, _ = json.Marshal(tc.value)
			}
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(bodyBytes))
			c.Request.Header.Set("Content-Type", "application/json")

			mh.CreateMessageHandler(c)
			assert.Equal(t, tc.status, w.Code)
		})
	}
}

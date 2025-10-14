package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMyHandler_GetLastMessageHandler(t *testing.T) {
	tests := []struct {
		name     string
		messages []string
		result   string
		status   int
	}{
		{
			name:     "One message",
			messages: []string{"Hello"},
			result:   "Hello",
			status:   http.StatusOK,
		},
		{
			name:     "Multiple messages",
			messages: []string{"Hello", "World"},
			result:   "World",
			status:   http.StatusOK,
		},
		{
			name:     "Empty message",
			messages: []string{},
			result:   ``,
			status:   http.StatusOK,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mh := &MyHandler{
				messages: tc.messages,
			}
			req := httptest.NewRequest(http.MethodGet, "/messages/last", nil)
			w := httptest.NewRecorder()
			mh.GetLastMessageHandler(w, req)
			assert.Equal(t, tc.status, w.Code)
			assert.Equal(t, w.Body.String(), tc.result)
		})
	}
}

func TestMyHandler_GetAllMessagesHandlerMessageHandler(t *testing.T) {
	tests := []struct {
		name     string
		messages []string
		result   string
		status   int
	}{
		{
			name:     "One message",
			messages: []string{"Hello"},
			result:   "Hello\n",
			status:   http.StatusOK,
		},
		{
			name:     "Multiple messages",
			messages: []string{"Hello", "World"},
			result:   "Hello\nWorld\n",
			status:   http.StatusOK,
		},
		{
			name:     "Empty message",
			messages: []string{},
			result:   ``,
			status:   http.StatusOK,
		},
		{
			name:     "Nil message",
			messages: nil,
			result:   ``,
			status:   http.StatusOK,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mh := &MyHandler{
				messages: tc.messages,
			}
			req := httptest.NewRequest(http.MethodGet, "/messages/all", nil)
			w := httptest.NewRecorder()
			mh.GetAllMessagesHandler(w, req)
			assert.Equal(t, tc.status, w.Code)
			assert.Equal(t, w.Body.String(), tc.result)
		})
	}
}

func TestMyHandler_CreateMessageHandler(t *testing.T) {
	tests := []struct {
		name   string
		value  map[string]string
		status int
	}{
		{
			name:   "Valid JSON",
			value:  map[string]string{"message": "Hello World"},
			status: http.StatusCreated,
		},
		{
			name:   "Invalid JSON",
			value:  map[string]string{"msg": "aboba"},
			status: http.StatusBadRequest,
		},
		{
			name:   "Empty JSON",
			value:  map[string]string{},
			status: http.StatusBadRequest,
		},
		{
			name:   "Nil map",
			value:  nil,
			status: http.StatusBadRequest,
		},
	}
	mh := &MyHandler{
		messages: []string{},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			body, _ := json.Marshal(tc.value)
			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(body))
			w := httptest.NewRecorder()
			mh.CreateMessageHandler(w, req)
			assert.Equal(t, w.Code, tc.status)
		})
	}
}

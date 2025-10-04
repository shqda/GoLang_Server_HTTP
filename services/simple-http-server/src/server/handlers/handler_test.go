package handlers

import (
	"bytes"
	"encoding/json"
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
			if w.Code != tc.status || w.Body.String() != tc.result {
				t.Errorf("handler returned wrong status code: got %v want %v", w.Code, tc.status)
				t.Errorf("handler returned unexpected body: got %v want %v", w.Body, tc.result)
			}
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
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mh := &MyHandler{
				messages: tc.messages,
			}
			req := httptest.NewRequest(http.MethodGet, "/messages/all", nil)
			w := httptest.NewRecorder()
			mh.GetAllMessagesHandler(w, req)
			if w.Code != tc.status || w.Body.String() != tc.result {
				t.Errorf("handler returned wrong status code: got %v want %v", w.Code, tc.status)
				t.Errorf("handler returned unexpected body: got %v want %v", w.Body, tc.result)
			}
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
			if w.Code != tc.status {
				t.Errorf("CreateMessageHandler returned wrong status code: got %v want %v", w.Code, tc.status)
			}
		})
	}
}

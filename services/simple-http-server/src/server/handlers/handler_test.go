package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMyHandler_CreateMessageHandler(t *testing.T) {
	tests := []struct {
		name   string
		value  string
		status int
	}{
		{name: "Valid JSON", value: "{\"message\":\"Hello world!\"}", status: http.StatusCreated},
		{name: "Invalid JSON", value: "{\"msg\":\"Privetik!\"}", status: http.StatusBadRequest},
	}
	_ = tests
	mh := &MyHandler{
		messages: []string{},
	}
	payload := map[string]string{
		"message": "Hello World",
	}
	jsonData, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(jsonData))
	rec := httptest.NewRecorder()
	mh.CreateMessageHandler(rec, req)
	if rec.Code != http.StatusCreated {
		t.Errorf("CreateMessageHandler returned wrong status code: got %v want %v", rec.Code, http.StatusCreated)
	}
	//fmt.Println(rec.Body.String())
}

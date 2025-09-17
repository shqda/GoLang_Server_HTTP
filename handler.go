package main

import (
	"fmt"
	"io"
	"net/http"
)

type MyHandler struct {
	messages []string
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.handleGet(w, r)
	case http.MethodPost:
		h.handlePost(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (h *MyHandler) handleGet(w http.ResponseWriter, r *http.Request) error {
	switch r.URL.Path {
	case "/messages/last":
		w.Write([]byte(h.getLastMessages()))
	default:
		http.NotFound(w, r)
	}
	return nil
}

func (h *MyHandler) getLastMessages() string {
	if len(h.messages) == 0 {
		return ""
	}
	return h.messages[len(h.messages)-1]
}

func (h *MyHandler) handlePost(w http.ResponseWriter, r *http.Request) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Request body reading error: ", http.StatusBadRequest)
		return err
	}
	fmt.Println("Received data:", string(body))
	h.messages = append(h.messages, string(body))
	return nil
}

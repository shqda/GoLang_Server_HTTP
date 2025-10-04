package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type MyHandler struct {
	messages []string
}

func (h *MyHandler) GetLastMessageHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, h.getLastMessage())
	if err != nil {
		log.Println("Error: ", err)
	}
}

func (h *MyHandler) GetAllMessagesHandler(w http.ResponseWriter, r *http.Request) {
	for _, m := range h.messages {
		_, err := fmt.Fprintln(w, m)
		if err != nil {
			log.Println("Error: ", err)
		}
	}
}

func (h *MyHandler) CreateMessageHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error: ", err)
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "Request body reading error", http.StatusBadRequest)
		return
	}
	var m struct {
		Msg string `json:"message"`
	}
	err = json.Unmarshal(body, &m)
	if err != nil || len(body) == 0 || m.Msg == "" {
		log.Println("Invalid json")
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "Request body marshaling error", http.StatusBadRequest)
		return
	} else {
		fmt.Println("Received data:", m.Msg)
		h.messages = append(h.messages, m.Msg)
		w.WriteHeader(http.StatusCreated)
	}
}

func (h *MyHandler) getLastMessage() string {
	if len(h.messages) == 0 {
		return ``
	}
	return h.messages[len(h.messages)-1]
}

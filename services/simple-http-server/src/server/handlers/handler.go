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
	_, err := fmt.Fprint(w, h.getLastMessages())
	if err != nil {
		log.Println("Error: ", err)
	}
}

func (h *MyHandler) GetAllMessagesHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Messages count: %d\n", len(h.messages))
	if err != nil {
		log.Println("Error: ", err)
	}
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
		http.Error(w, "Request body reading error", http.StatusBadRequest)
	}
	type message struct {
		Msg string `json:"message"`
	}
	var m message
	err = json.Unmarshal(body, &m)
	if err != nil || len(body) == 0 {
		log.Println("Invalid json")
		http.Error(w, "Request body marshaling error", http.StatusBadRequest)
	} else {
		fmt.Println("Received data:", m.Msg)
		h.messages = append(h.messages, m.Msg)
	}
}

func (h *MyHandler) getLastMessages() string {
	if len(h.messages) == 0 {
		return ``
	}
	return h.messages[len(h.messages)-1]
}

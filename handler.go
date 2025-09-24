package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type MyHandler struct {
	messages []string
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		err := h.handleGet(w, r)
		if err != nil {
			log.Println("handleGet() error: ", err)
			return
		}
	case http.MethodPost:
		err := h.handlePost(w, r)
		if err != nil {
			log.Println("handlePost() error: ", err)
			return
		}
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (h *MyHandler) handleGet(w http.ResponseWriter, r *http.Request) error {
	switch r.URL.Path {
	case "/messages/last":
		_, err := fmt.Fprint(w, h.getLastMessages())
		if err != nil {
			log.Println("Error: ", err)
			return err
		}
	case "/messages/all":
		_, err := fmt.Fprintf(w, "Messages count: %d\n", len(h.messages))
		if err != nil {
			return err
		}
		for _, m := range h.messages {
			_, err := fmt.Fprintln(w, m)
			if err != nil {
				log.Println("Error: ", err)
				return err
			}
		}
	default:
		http.NotFound(w, r)
	}
	return nil
}

func (h *MyHandler) getLastMessages() string {
	if len(h.messages) == 0 {
		return ``
	}
	return h.messages[len(h.messages)-1]
}

func (h *MyHandler) handlePost(w http.ResponseWriter, r *http.Request) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error: ", err)
		http.Error(w, "Request body reading error", http.StatusBadRequest)
		return err
	}
	body, err = MarshalJSON(body)
	if err != nil || len(body) == 0 {
		log.Println("Invalid json")
		http.Error(w, "Request body marshaling error", http.StatusBadRequest)
		return err
	} else {
		fmt.Println("Received data:", string(body))
		h.messages = append(h.messages, string(body))
		return nil
	}
}

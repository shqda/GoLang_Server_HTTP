package server

import (
	"HttpServer/server/handlers"
	"net/http"
)

func GetRouter(h *handlers.MyHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /messages/last", h.GetLastMessageHandler)
	mux.HandleFunc("GET /messages/all", h.GetAllMessagesHandler)
	mux.HandleFunc("POST /messages/add", h.CreateMessageHandler)
	return mux
}

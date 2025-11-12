package server

import (
	"HttpServer/server/handlers"
	"github.com/gin-gonic/gin"
)

func GetRouter(h *handlers.MyHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/messages/last", h.GetLastMessageHandler)
	r.GET("/messages/all", h.GetAllMessagesHandler)
	r.POST("/messages/add", h.CreateMessageHandler)

	return r
}

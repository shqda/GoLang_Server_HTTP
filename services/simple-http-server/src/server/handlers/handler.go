package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type MyHandler struct {
	messages []string
}

type message struct {
	Msg string `json:"message"`
}

func (h *MyHandler) GetLastMessageHandler(c *gin.Context) {
	last := h.getLastMessage()
	if last == "" {
		c.String(http.StatusOK, "")
		return
	}
	c.String(http.StatusOK, last)
}

func (h *MyHandler) GetAllMessagesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, h.messages)
}

func (h *MyHandler) CreateMessageHandler(c *gin.Context) {
	var m message
	if err := c.BindJSON(&m); err != nil || m.Msg == "" {
		log.Println("Invalid JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	fmt.Println("Received data:", m.Msg)
	h.messages = append(h.messages, m.Msg)
	c.Status(http.StatusCreated)
}

func (h *MyHandler) getLastMessage() string {
	if len(h.messages) == 0 {
		return ""
	}
	return h.messages[len(h.messages)-1]
}

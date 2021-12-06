package handler

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UserIdentity(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		log.Fatal("Empty auth Header")
	}

	headerParts := strings.Split(header, " ")

	if len(headerParts) != 2 {
		log.Fatal("invalid auth Header")
	}

	userId, err := h.services.Authorizaton.ParseToken(headerParts[1])
	if err != nil {
		log.Fatal("invalid ParseToken")
	}
	c.Set("userId", userId)

}

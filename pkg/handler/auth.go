package handler

import (
	"fmt"
	"log"
	"net/http"

	project "github.com/PashaFedyushkin/GoProj"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func (h *Handler) signUp(c *gin.Context) {
	var in project.User

	if err := c.BindJSON(&in); err != nil {
		log.Fatal(err.Error())
	}

	id, err := h.services.Authorizaton.CreateUser(in)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(id)
	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var in signInInput

	if err := c.BindJSON(&in); err != nil {
		log.Fatal(err.Error())
	}

	token, err := h.services.Authorizaton.GenerateToken(in.Username, in.Password)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(token)
	c.JSON(http.StatusOK, map[string]interface{}{"token": token})
}

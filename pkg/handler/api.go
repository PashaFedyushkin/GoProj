package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type gameInput struct {
	Number string `json:"number" binding:"required"`
}

func (h *Handler) letsPlay(c *gin.Context) {
	var in gameInput
	if err := c.BindJSON(&in); err != nil {
		log.Fatal(err.Error())
	}
	a, err := strconv.Atoi(in.Number)
	if err != nil {
		log.Fatal(err)
	}
	res, err := h.services.PlayGame.LetsPlay(int32(a))
	if err != nil {
		log.Fatal(err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{"res": res})

}

func (h *Handler) getStat(c *gin.Context) {

}

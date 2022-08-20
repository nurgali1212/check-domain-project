package handler

import (
	"domain/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) checkDomain(c *gin.Context) {
	var input model.Domain

	if err := c.BindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Ошибка при обработке запроса")
		return
	}

	ch := make(chan bool)

	go h.service.CheckDomainService(input, ch)

	valid := <-ch

	message := "Домен не валидный"
	if valid {
		message = "Домен валидный"
	}

	c.IndentedJSON(http.StatusOK, map[string]interface{}{
		"response": message,
	})
}

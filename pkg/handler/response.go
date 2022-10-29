package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/siruspen/logrus"
)

type error struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	// Блокирует выполнение последующих обработчиков и записывает в ответ статус код в тело сообщения в формате json
	c.AbortWithStatusJSON(statusCode, error{message})
}

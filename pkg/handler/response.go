package handler

import (
	"github.com/NickSarychev/todo-app"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ErrorResponse struct {
	Message string `json:"message"`
}
type StatusResponse struct {
	Status string `json:"status"`
}

type GetAllListsResponse struct {
	Data []todo.TodoList `json:"data"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, ErrorResponse{Message: message})
}

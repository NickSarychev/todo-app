package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Massage string `json:"massage"`
}
type statusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, statusCode int, massage string) {
	logrus.Error(massage)
	c.AbortWithStatusJSON(statusCode, errorResponse{Massage: massage})
}

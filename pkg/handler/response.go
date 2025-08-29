package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponce struct {
	Massage string `json:"massage"`
}
type statusResponce struct {
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, statusCode int, massage string) {
	logrus.Error(massage)
	c.AbortWithStatusJSON(statusCode, errorResponce{Massage: massage})
}

package controllers

import (
	"github.com/aristotelesbr/api-go-gin/models"
	"github.com/gin-gonic/gin"
)

func Person(c *gin.Context) {
	c.JSON(200, models.Students)
}

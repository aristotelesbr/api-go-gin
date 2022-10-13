package routes

import (
	"github.com/aristotelesbr/api-go-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	router := gin.Default()
	router.GET("/", controllers.Person)
	router.Run(":8000")
}

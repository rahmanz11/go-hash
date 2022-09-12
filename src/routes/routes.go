package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rahmanz11/go-hash/src/controllers"
)

func Routes() {
	route := gin.Default()
	route.POST("/hash", controllers.CreateHash)
	route.Run()
}

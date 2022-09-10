package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rahmanz11/go-hash/src/controllers"
)

// Routes function to serve endpoints
func Routes() {
	route := gin.Default()

	route.POST("/hash", controllers.CreateHash)
	// Run route whenever triggered
	route.Run()
}

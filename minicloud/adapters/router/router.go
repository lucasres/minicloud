package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"lucares.github.com/minicloud/minicloud/adapters/router/controllers"
)

func CreateRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/alive", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"alive": "yes, just yes"})
	})

	r.GET("/config/was-configured", controllers.WasConfiguredHandler)

	return r
}

package main

import (
	_ "github.com/devxp-tech/teste-loki/config"
	"github.com/devxp-tech/teste-loki/controllers"
	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.New()
	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  true,
			"message": "Hello world for your app teste-loki is fucking running!",
			"headers": c.Request.Header,
		})
	})
	server.GET("/health-check/liveness", controllers.HealthCheckLiveness)
	server.GET("/health-check/readiness", controllers.HealthCheckReadiness)

	server.Run()
}

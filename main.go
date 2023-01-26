package main

import (
	"log"

	_ "github.com/devxp-tech/teste-loki/config"
	"github.com/devxp-tech/teste-loki/controllers"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)

func main() {

	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://3f0ed90a9758404e89028367d40759a6@o4504350064836608.ingest.sentry.io/4504570239975424",
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	server := gin.New()
	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  true,
			"message": "Hello world for your app teste-loki",
		})
	})
	server.GET("/health-check/liveness", controllers.HealthCheckLiveness)
	server.GET("/health-check/readiness", controllers.HealthCheckReadiness)
	server.Run()
}

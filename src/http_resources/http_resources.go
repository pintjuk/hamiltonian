// Package http_resources
// It is the combination of Interface (http) and application layers in DDD layer architecture
// in this package we handle everything related to the http server, validate input, integrate with external services, handle telemetry and logging, etc
package http_resources

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/routes", getRoutes)
	e.GET("/health", getHealth)
}

// Config configuration values needed to run the service
type Config interface {
	Port() string
}

// I wanted to implement a structured logger for GCloud but did not get around to it
type Logger interface {
	Debugf(msg string, trace string)
	Infof(msg string, trace string)
	Error(err error, trace string)
}

func StartHttpServer(config Config, logger Logger) {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes

	InitRoutes(e) // Initialize API routes

	// Start server
	port := fmt.Sprintf(":%s", config.Port())
	log.Printf("Server started on port %s\n", port)

	log.Fatal(e.Start(port))
}

package main

import (
	"golang-template/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Create a new Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	routes.CreateRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

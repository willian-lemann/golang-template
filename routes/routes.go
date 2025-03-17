package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateAuthRoutes(e *echo.Echo) {
	e.POST("/login", func(c echo.Context) error {
		return c.String(200, "Login")
	})

	e.POST("/register", func(c echo.Context) error {
		return c.String(200, "Register")
	})
}

// RegisterRoutes registers all application routes
func CreateRoutes(e *echo.Echo) {
	// Register all route groups
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/listings/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	CreateAuthRoutes(e)
}

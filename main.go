package main

import (
	"net/http"

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
	e.GET("/", hello)
	e.GET("/users/:id", getUser)
	e.POST("/users", createUser)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func getUser(c echo.Context) error {
	// Get user ID from path parameter
	id := c.Param("id")
	return c.JSON(http.StatusOK, map[string]string{
		"id":   id,
		"name": "John Doe",
	})
}

func createUser(c echo.Context) error {
	// User struct
	type User struct {
		Name  string `json:"name" form:"name"`
		Email string `json:"email" form:"email"`
	}

	u := new(User)
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, u)
}

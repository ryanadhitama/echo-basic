package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func getUsers(c echo.Context) error {
	return c.String(http.StatusOK, "get users")
}

func getUser(c echo.Context) error {
	return c.String(http.StatusOK, c.Param("id"))
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// variant route 
	e.GET("/users", getUsers)
	e.GET("/users/:id", getUser)
	e.POST("/users", func(c echo.Context) error {
		return c.String(http.StatusOK, "create user")
	})
	e.PUT("/users/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "update user by id")
	})
	e.DELETE("/users/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "delete user by id")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
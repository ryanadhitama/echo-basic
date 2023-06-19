package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func getUsers(c echo.Context) error {
	name := c.QueryParam("name")
	return c.String(http.StatusOK, "get users with name : " + name)
}

func getUser(c echo.Context) error {
	return c.String(http.StatusOK, c.Param("id"))
}

// curl -d "name=Joe Smith" -d "email=joe@labstack.com" http://localhost:8080/users
func createUser(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:" + name + ", email:" + email)
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// variant route 
	e.GET("/users", getUsers)
	e.GET("/users/:id", getUser)
	e.POST("/users", createUser)
	e.PUT("/users/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "update user by id")
	})
	e.DELETE("/users/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "delete user by id")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
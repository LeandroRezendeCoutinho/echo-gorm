package main

import (
	"net/http"

	config "main/configs"
	"main/internal/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	config.Connect()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/dogs", handlers.GetDogs)
	e.GET("/dogs/:id", handlers.GetDog)
	e.POST("/dogs", handlers.AddDog)
	e.PUT("/dogs/:id", handlers.UpdateDog)
	e.DELETE("/dogs/:id", handlers.DeleteDog)

	e.Logger.Fatal(e.Start(":5000"))
}

package main

import (
	"net/http"

	config "main/configs"
	"main/internal/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	config.Connect()
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<h1>Hello, World!</h1>")
	})

	e.GET("/dogs", handlers.GetDogs)
	e.GET("/dogs/:id", handlers.GetDog)
	e.POST("/dogs", handlers.AddDog)
	e.PUT("/dogs/:id", handlers.UpdateDog)
	e.DELETE("/dogs/:id", handlers.DeleteDog)

	e.Logger.Fatal(e.Start(":5000"))
}

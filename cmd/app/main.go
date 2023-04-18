package main

import (
	"net/http"

	config "main/configs"
	"main/internal/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	router := echo.New()

	config.Connect()
	router.Use(middleware.Logger())

	router.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<h1>Hello, World!</h1>")
	})

	router.GET("/dogs", handlers.GetDogs)
	router.GET("/dogs/:id", handlers.GetDog)
	router.POST("/dogs", handlers.AddDog)
	router.PUT("/dogs/:id", handlers.UpdateDog)
	router.DELETE("/dogs/:id", handlers.DeleteDog)

	router.Logger.Fatal(router.Start(":5000"))
}

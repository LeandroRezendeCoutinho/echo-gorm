package handlers

import (
	config "main/configs"

	"main/internal/entities"

	"github.com/labstack/echo/v4"
)

func GetDogs(c echo.Context) error {
	var dogs []entities.Dog
	config.Database.Find(&dogs)

	return c.JSON(200, dogs)
}

func GetDog(c echo.Context) error {
	var dog entities.Dog
	config.Database.First(&dog, c.Param("id"))

	return c.JSON(200, dog)
}

func AddDog(c echo.Context) error {
	var dog entities.Dog

	if err := c.Bind(&dog); err != nil {
		return err
	}

	config.Database.Create(&dog)

	return c.JSON(201, dog)
}

func UpdateDog(c echo.Context) error {
	var dog entities.Dog
	id := c.Param("id")

	if err := c.Bind(&dog); err != nil {
		return err
	}

	config.Database.Where("id = ?", id).Updates(&dog)

	return c.JSON(200, dog)
}

func DeleteDog(c echo.Context) error {
	var dog entities.Dog
	id := c.Param("id")

	config.Database.Where("id = ?", id).Delete(&dog)

	return c.JSON(200, dog)
}

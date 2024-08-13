package main

import (

	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/myertek/echotest/internal/models"
)

func home(c echo.Context) error{
	
	drivers, err := models.AllActiveDrivers()
	if err != nil{
		return err
	}
	// Create an instance of a templateData struct holding the snippet data.
	data := templateData{
        Drivers: drivers,
    }


	return c.Render(http.StatusOK,"index", data)

}
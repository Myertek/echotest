package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/myertek/echotest/internal/models"
)

func getDrivers(c echo.Context) error {

	drivers, err := models.AllActiveDrivers()
	if err != nil {
		return err
	}
	// Create an instance of a templateData struct holding the snippet data.
	data := templateData{
		Drivers: drivers,
	}

	return c.Render(http.StatusOK, "index", data)

}

func getDriver(c echo.Context) error {

	driver, err := models.GetDriver(c.Param("id"))
	if err != nil {
		return err
	}
	// Create an instance of a templateData struct holding the snippet data.
	data := templateData{
		Driver: driver,
	}

	return c.Render(http.StatusOK, "driver", data)

}

func getRaces(c echo.Context) error {

	races, err := models.GetRaces()
	if err != nil {
		return err
	}
	// Create an instance of a templateData struct holding the snippet data.
	data := templateData{
		Races: races,
	}

	return c.Render(http.StatusOK, "race", data)

}

func getPredictions(c echo.Context) error {

	predictions, err := models.GetPredictions()
	if err != nil {
		return err
	}
	// Create an instance of a templateData struct holding the snippet data.
	data := templateData{
		Predictions: predictions,
	}

	return c.Render(http.StatusOK, "predictions", data)

}

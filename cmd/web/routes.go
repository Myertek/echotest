package main

import (
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	e.GET("/races", getRaces)
	e.GET("/drivers", getDrivers)
	e.GET("/driver/:id", getDriver)
	e.GET("/predictions", getPredictions)
}

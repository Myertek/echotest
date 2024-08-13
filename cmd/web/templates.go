package main

import "github.com/myertek/echotest/internal/models"

// Define a templateData type to act as the holding structure for
// any dynamic data that we want to pass to our HTML templates.

type templateData struct {
	Driver models.Driver
    Drivers []models.Driver
}
package main

import (

	"github.com/labstack/echo/v4"
	
)

func main() {
	e := echo.New()

	InitDB()

	Routes(e)

	e.Logger.Printf("hello.....")
	e.Logger.Fatal(e.Start(":8000"))
}

func InitDB() {
	panic("unimplemented")
}
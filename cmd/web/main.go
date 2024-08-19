package main

import (
	"io"
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	db "github.com/myertek/echotest/internal/database"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	db.InitDB()

	t := &Template{
		templates: template.Must(template.ParseFiles(
			"ui/html/pages/index.html",
			"ui/html/pages/driver.html",
			"ui/html/pages/race.html",
			"ui/html/pages/predictions.html")),
	}
	e.Renderer = t
	e.Static("/img", "static")
	Routes(e)

	e.Logger.Printf("hello.....")
	e.Logger.Fatal(e.Start(":8000"))
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

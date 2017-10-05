package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/GDGVIT/autodock/echotemplate"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	renderer := echotemplate.New(template.Must(template.ParseGlob("views/*.html")))
	e.Renderer = renderer
	e.Static("/static", "static")
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", map[string]interface{}{})
	})
	log.Fatal(e.Start(":" + os.Getenv("PORT")))
}

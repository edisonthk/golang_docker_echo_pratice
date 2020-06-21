package main

import (
  "net/http"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  "myapp/mypackage"
)

type helloPathViewParameter struct {
  Title string
  Text string
}

func main() {
  // Echo instance
  e := echo.New()

  // Middleware
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())
  e.Static("/", "public")
  e.Renderer = mypackage.GetRenderer()

  mypackage.Print()

  // Routes
  e.GET("/", helloInTemplate)
  e.GET("/string", helloInString)

  // Start server
  e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func helloInString(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, World!")
}

func helloInTemplate(c echo.Context) error {

  param := helloPathViewParameter{
    Title: "views/hello1.html will be rendered as layout name 'hello' is defined on first line.",
    Text: "Below show the result of dot in Golang.",
  }

  return c.Render(http.StatusOK, "hello", param)
}
package main

import (
  "net/http"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  "myapp/mypackage"
  "database/sql"
  "fmt"
  _ "github.com/go-sql-driver/mysql"
)

type helloPathViewParameter struct {
  Title string
  Text string
}

// DB の接続情報
const (
  DRIVER = "mysql" // ドライバ名(mysql固定)
  // user:password@tcp(container-name:port)/dbname 
  DATA_SOURCE = "root:root@tcp(mysql-container:3306)/test"
)

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

  //　DB接続 
  db, err := sql.Open(DRIVER, DATA_SOURCE)
  if err != nil {
    fmt.Println("Openエラー")
    panic(err.Error())
  }
  defer db.Close() // 関数がリターンする直前に呼び出す

  //　DB接続確認
  err = db.Ping()
  if err != nil {
    fmt.Println("接続失敗！")
    panic(err.Error())
  } else {
    fmt.Println("接続成功！")
  }

  // データ取得してみた
  rows, err := db.Query("SELECT * FROM member")
  if err != nil {
    fmt.Println("データ取得失敗！")
    panic(err.Error())
  } 
  for rows.Next() {
    var id string
    var name string
    var age string
    err = rows.Scan(&id, &name, &age)
    fmt.Println(id, name, age)
  }

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
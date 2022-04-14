package main

import (
	"github.com/labstack/echo"
	"net/http"
)

type HelloWorld struct {
	Message string `json:"message"`
}

func main() {
	e := echo.New()
	e.GET("/", Greeting)
	e.Logger.Fatal(e.Start(":3000"))
}

func Greeting(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, HelloWorld{
		Message: "Hello World",
	})
}

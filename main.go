package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

type HelloWorld struct {
	Message string `json:"message"`
}

func main() {
	//initialise Echo framework
	e := echo.New()

	//Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, host=${host}\n",
	}))

	//Handlers
	e.GET("/", Greeting)
	e.GET("/hello/:name", GreetingWithParams)
	e.GET("/hello-queries", GreetingWithQuery)
	e.Logger.Fatal(e.Start(":3000"))
}

func Greeting(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, HelloWorld{
		Message: "Hello World",
	})
}

func GreetingWithParams(ctx echo.Context) error {
	params := ctx.Param("name")
	return ctx.JSON(http.StatusOK, HelloWorld{
		Message: "Hello World, my name is " + params,
	})
}

func GreetingWithQuery(ctx echo.Context) error {
	query := ctx.QueryParam("name")
	return ctx.JSON(http.StatusOK, HelloWorld{
		Message: "Hello World, my name is " + query + " and I'm using queries",
	})
}

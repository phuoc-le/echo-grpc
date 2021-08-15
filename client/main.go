package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	gp "grpc-echo/pkg/grpc"
	"grpc-echo/pkg/router"
)

func main() {

	gp.GreeterService()
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	router.Init(e)
	e.Logger.Fatal(e.Start(":8888"))

}

package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"grpc-echo/pkg/app"
	gp "grpc-echo/pkg/grpc"
	"grpc-echo/pkg/router"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	e := echo.New()
	_, err = app.Get()
	if err != nil {
		e.Logger.Fatal(err)
	}
	gp.GreeterService()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	router.Init(e)
	e.Logger.Fatal(e.Start(":8888"))

}

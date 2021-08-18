package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4/middleware"
	gp "grpc-echo/pkg/grpc"
	"grpc-echo/pkg/router"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	r := router.New()

	gp.GreeterService()
	// Middleware
	r.Use(middleware.Logger())
	//e.Use(middleware.Recover())
	r.Logger.Fatal(r.Start(":8888"))

}

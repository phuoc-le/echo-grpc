package main

import (
	"github.com/joho/godotenv"
	gp "grpc-echo/pkg/grpc"
	"grpc-echo/pkg/router"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	r := router.New()

	gp.GreeterService("Test")

	r.Logger.Fatal(r.Start(":"+os.Getenv("API_PORT")))

}

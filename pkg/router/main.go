package router

import (
	"github.com/labstack/echo/v4"
	"grpc-echo/pkg/handler"
	"net/http"
)

func MainGroup(g *echo.Group) {
	g.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Home!")
	})
	g.GET("/login", handler.Login)
}

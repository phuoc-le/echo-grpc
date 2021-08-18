package router

import (
	"github.com/labstack/echo/v4"
	"grpc-echo/pkg/handler"
	"net/http"
)

func UserGroup(g *echo.Group) {
	g.GET("/users/:id", handler.GetUser)
	g.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, User!")
	})
}

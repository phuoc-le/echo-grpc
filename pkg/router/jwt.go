package router

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func JwtGroup(g *echo.Group) {
	g.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, JWT!")
	})
}

package router

import (
	"github.com/labstack/echo/v4"
	"grpc-echo/pkg/controller"
	"net/http"
)

func Init(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/users/:id", controller.GetUser)
	e.GET("/students", controller.GetStudents)

}

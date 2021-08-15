package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	gp "grpc-echo/pkg/grpc"
	"net/http"
)

func GetUser(c echo.Context) error {
	// User ID from path `users/:id`

	id := c.Param("id")
	log.Info(id)
	gp.TestService(id)
	return c.String(http.StatusOK, id)
}

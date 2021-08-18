package router

import (
	"github.com/labstack/echo/v4"
	"grpc-echo/pkg/handler"
)

func StudentGroup(g *echo.Group) {
	g.GET("/list", handler.GetStudents)
	g.GET("/:id", handler.GetStudent)
	g.POST("/", handler.CreateStudent)
	g.DELETE("/:id", handler.DeleteStudent)
	g.PUT("/:id", handler.UpdateStudent)
}

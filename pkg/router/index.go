package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"grpc-echo/pkg/middlewares"
)

func New() *echo.Echo {
	e := echo.New()
	e.Use(middleware.BodyLimit("2M"))
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// create groups
	adminGroup := e.Group("/admin")
	mainGroup := e.Group("")
	jwtGroup := e.Group("/jwt")
	userGroup := e.Group("/user")
	studentGroup := e.Group("/student")


	// set all middlewares
	//middlewares.SetMainMiddlewares(e)
	middlewares.SetAdminMiddlewares(adminGroup)
	middlewares.SetJwtMiddlewares(jwtGroup)
	middlewares.SetJwtMiddlewares(studentGroup)


	//// set admin routes
	AdminGroup(adminGroup)
	//// set group routes
	JwtGroup(jwtGroup)
	UserGroup(userGroup)
	StudentGroup(studentGroup)
	MainGroup(mainGroup)

	return e
}
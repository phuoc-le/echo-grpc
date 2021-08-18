package router

import (
	"github.com/labstack/echo/v4"
	"grpc-echo/pkg/middlewares"
)

func New() *echo.Echo {
	e := echo.New()

	// create groups
	adminGroup := e.Group("/admin")
	//cookieGroup := e.Group("/cookie")
	mainGroup := e.Group("")
	jwtGroup := e.Group("/jwt")
	userGroup := e.Group("/user")
	studentGroup := e.Group("/student")


	// set all middlewares
	//middlewares.SetMainMiddlewares(e)
	middlewares.SetAdminMiddlewares(adminGroup)
	//middlewares.SetCookieMiddlewares(cookieGroup)
	middlewares.SetJwtMiddlewares(jwtGroup)
	//middlewares.SetMainMiddlewares(e)


	//// set admin routes
	AdminGroup(adminGroup)
	//// set group routes
	//CookieGroup(cookieGroup)
	JwtGroup(jwtGroup)
	UserGroup(userGroup)
	StudentGroup(studentGroup)
	MainGroup(mainGroup)

	return e
}
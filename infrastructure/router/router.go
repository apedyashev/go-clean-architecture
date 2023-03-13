package router

import (
	"posts/adapter/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/posts", func(context echo.Context) error { return c.Post.GetPosts(context) })

	return e
}

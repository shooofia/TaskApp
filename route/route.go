package route

import (
	"TaskApp/controller"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(e *echo.Echo) *echo.Echo {
	e.Use(middleware.Logger())
	e.POST("/login", controller.LoginController)

	taskRoute := e.Group("")
	taskRoute.Use(echojwt.JWT([]byte(os.Getenv("SECRET_KEY"))))
	taskRoute.GET("/Task", controller.GetTaskController)
	taskRoute.POST("/Task", controller.AddTaskController)
	taskRoute.GET("/Task/:id", controller.GetDetailTaskController)
	return e
}

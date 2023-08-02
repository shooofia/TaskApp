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

	taskRoute := e.Group("")
	taskRoute.POST("/login", controller.LoginController)
	taskRoute.Use(echojwt.JWT([]byte(os.Getenv("SECRET_KEY"))))
	taskRoute.GET("/task", controller.GetTaskController)
	taskRoute.POST("/task", controller.AddTaskController)
	taskRoute.GET("/task/:id", controller.GetDetailTaskController)
	return e
}

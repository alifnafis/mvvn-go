package router

import (
	"go-api/controller"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Echo, authController *controller.AuthController) {
	e.POST("/signup", authController.Signup)
	e.POST("/login", authController.Login)
}

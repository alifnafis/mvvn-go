package router

import (
	"go-api/controller"
	"go-api/utils"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo, userController *controller.UserController) {
	protected := e.Group("/users", utils.JWTMiddleware())
	protected.GET("", userController.GetAllUsers)
	protected.PUT("/:id", userController.UpdateUser)
	protected.DELETE("/:id", userController.DeleteUser)
	protected.GET("/openlibrary", userController.GetOpenLibraryData)
}

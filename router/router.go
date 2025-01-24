package router

import (
	"go-api/controller"

	"github.com/labstack/echo/v4"
)

func InitRouter(
	e *echo.Echo,
	authController *controller.AuthController,
	userController *controller.UserController,
	musicController *controller.MusicController,
	playlistController *controller.PlaylistController,
) {
	AuthRoutes(e, authController)
	UserRoutes(e, userController)
	MusicRoutes(e, musicController)
	PlaylistRoutes(e, playlistController)
}

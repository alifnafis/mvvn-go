package router

import (
	"go-api/controller"
	"go-api/utils"

	"github.com/labstack/echo/v4"
)

func MusicRoutes(e *echo.Echo, musicController *controller.MusicController) {
	// Music CRUD routes
	music := e.Group("/music", utils.JWTMiddleware())
	music.POST("", musicController.CreateMusic)
	music.GET("", musicController.GetAllMusic)
	music.GET("/:id", musicController.GetMusicByID)
	music.PUT("/:id", musicController.UpdateMusic)
	music.DELETE("/:id", musicController.DeleteMusic)
}

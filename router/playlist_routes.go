package router

import (
	"go-api/controller"
	"go-api/utils"

	"github.com/labstack/echo/v4"
)

func PlaylistRoutes(e *echo.Echo, playlistController *controller.PlaylistController) {
	// My Playlist CRUD routes
	playlist := e.Group("/playlist", utils.JWTMiddleware())
	playlist.GET("/:user_id", playlistController.GetMyPlaylists)
	playlist.POST("", playlistController.CreatePlaylist)
	playlist.POST("/add-music", playlistController.AddMusicToPlaylist)
	playlist.DELETE("/:id", playlistController.DeletePlaylist)
}

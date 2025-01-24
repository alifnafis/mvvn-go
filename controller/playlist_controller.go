package controller

import (
	"go-api/model"
	"go-api/service"
	"go-api/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PlaylistController struct {
	Service *service.PlaylistService
}

func NewPlaylistController(service *service.PlaylistService) *PlaylistController {
	return &PlaylistController{Service: service}
}

func (c *PlaylistController) GetMyPlaylists(ctx echo.Context) error {
	userIDParam := ctx.Param("user_id")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		return utils.Respond(ctx, http.StatusBadRequest, nil, "Invalid user ID")
	}

	playlists, err := c.Service.GetMyPlaylists(userID)
	if err != nil {
		return utils.Respond(ctx, http.StatusInternalServerError, nil, "Failed to retrieve playlists")
	}
	return utils.Respond(ctx, http.StatusOK, playlists, "Success")
}

func (c *PlaylistController) CreatePlaylist(ctx echo.Context) error {
	var playlist model.Playlist
	if err := ctx.Bind(&playlist); err != nil {
		return utils.Respond(ctx, http.StatusBadRequest, nil, "Invalid playlist data")
	}

	if err := c.Service.CreatePlaylist(playlist); err != nil {
		return utils.Respond(ctx, http.StatusInternalServerError, nil, "Failed to create playlist")
	}
	return utils.Respond(ctx, http.StatusCreated, nil, "Playlist created successfully")
}

func (c *PlaylistController) DeletePlaylist(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return utils.Respond(ctx, http.StatusBadRequest, nil, "Invalid playlist ID")
	}

	if err := c.Service.DeletePlaylist(id); err != nil {
		return utils.Respond(ctx, http.StatusInternalServerError, nil, "Failed to delete playlist")
	}
	return utils.Respond(ctx, http.StatusOK, nil, "Playlist deleted successfully")
}

func (c *PlaylistController) AddMusicToPlaylist(ctx echo.Context) error {
	
	type AddMusicRequest struct {
		PlaylistID int `json:"playlist_id"`
		MusicID    int `json:"music_id"`
	}

	var request AddMusicRequest
	if err := ctx.Bind(&request); err != nil {
		return utils.Respond(ctx, http.StatusBadRequest, nil, "Invalid request body")
	}

	if request.PlaylistID == 0 || request.MusicID == 0 {
		return utils.Respond(ctx, http.StatusBadRequest, nil, "Playlist ID and Music ID are required")
	}

	if err := c.Service.AddMusicToPlaylist(request.PlaylistID, request.MusicID); err != nil {
		return utils.Respond(ctx, http.StatusInternalServerError, nil, "Failed to add music to playlist")
	}

	return utils.Respond(ctx, http.StatusOK, nil, "Music added to playlist successfully")
}


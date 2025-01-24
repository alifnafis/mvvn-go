package controller

import (
	"go-api/model"
	"go-api/service"
	"go-api/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MusicController struct {
	Service *service.MusicService
}

func NewMusicController(service *service.MusicService) *MusicController {
	return &MusicController{Service: service}
}

func (c *MusicController) GetAllMusic(ctx echo.Context) error {
	musicList, err := c.Service.GetAllMusic()
	if err != nil {
		return utils.Respond(ctx, http.StatusInternalServerError, nil, "Failed to retrieve music")
	}
	return utils.Respond(ctx, http.StatusOK, musicList, "Success")
}

func (c *MusicController) CreateMusic(ctx echo.Context) error {
	var music model.Music
	if err := ctx.Bind(&music); err != nil {
		return utils.Respond(ctx, http.StatusBadRequest, nil, "Invalid music data")
	}

	if err := c.Service.CreateMusic(music); err != nil {
		return utils.Respond(ctx, http.StatusInternalServerError, nil, "Failed to create music")
	}
	return utils.Respond(ctx, http.StatusCreated, nil, "Music created successfully")
}

func (c *MusicController) UpdateMusic(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return utils.Respond(ctx, http.StatusBadRequest, nil, "Invalid music ID")
	}

	var music model.Music
	if err := ctx.Bind(&music); err != nil {
		return utils.Respond(ctx, http.StatusBadRequest, nil, "Invalid music data")
	}

	if err := c.Service.UpdateMusic(id, music); err != nil {
		return utils.Respond(ctx, http.StatusInternalServerError, nil, "Failed to update music")
	}
	return utils.Respond(ctx, http.StatusOK, nil, "Music updated successfully")
}

func (c *MusicController) DeleteMusic(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return utils.Respond(ctx, http.StatusBadRequest, nil, "Invalid music ID")
	}

	if err := c.Service.DeleteMusic(id); err != nil {
		return utils.Respond(ctx, http.StatusInternalServerError, nil, "Failed to delete music")
	}
	return utils.Respond(ctx, http.StatusOK, nil, "Music deleted successfully")
}

func (c *MusicController) GetMusicByID(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return utils.Respond(ctx, http.StatusBadRequest, nil, "Invalid music ID")
	}

	music, err := c.Service.GetMusicByID(id)
	if err != nil {
		return utils.Respond(ctx, http.StatusInternalServerError, nil, "Failed to retrieve music")
	}

	return utils.Respond(ctx, http.StatusOK, music, "Music retrieved successfully")
}

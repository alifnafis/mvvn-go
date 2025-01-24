package controller

import (
	"go-api/model"
	"go-api/service"
	"go-api/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	Service *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{Service: service}
}

func (c *UserController) GetAllUsers(ctx echo.Context) error {
	users, err := c.Service.GetAllUsers()
	if err != nil {
		return utils.Respond(ctx, http.StatusInternalServerError, nil, "Failed to retrieve users")
	}
	return utils.Respond(ctx, http.StatusOK, users, "Success")
}

func (c *UserController) UpdateUser(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return utils.Respond(ctx, http.StatusBadRequest, nil, "Invalid user ID")
	}

	var user model.User
	if err := ctx.Bind(&user); err != nil {
		return utils.Respond(ctx, http.StatusBadRequest, nil, err.Error())
	}

	if err := c.Service.UpdateUser(id, user); err != nil {
		return utils.Respond(ctx, http.StatusInternalServerError, nil, err.Error())
	}

	return utils.Respond(ctx, http.StatusOK, nil, "User updated successfully")
}

func (c *UserController) DeleteUser(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		return utils.Respond(ctx, http.StatusBadRequest, nil, "Invalid user ID")
	}

	if err := c.Service.DeleteUser(id); err != nil {
		return utils.Respond(ctx, http.StatusInternalServerError, nil, err.Error())
	}

	return utils.Respond(ctx, http.StatusOK, nil, "User deleted successfully")
}

func (c *UserController) GetOpenLibraryData(ctx echo.Context) error {
	data, err := c.Service.GetOpenLibraryData()
	if err != nil {
		return utils.Respond(ctx, http.StatusInternalServerError, nil, err.Error())
	}

	return utils.Respond(ctx, http.StatusOK, data, "Successfully retrieved data from external API")
}
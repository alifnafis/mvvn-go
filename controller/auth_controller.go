package controller

import (
	"go-api/model"
	"go-api/service"
	"go-api/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	Service *service.AuthService
}

func NewAuthController(service *service.AuthService) *AuthController {
	return &AuthController{Service: service}
}

func (c *AuthController) Signup(ctx echo.Context) error {
	var user model.User
	if err := ctx.Bind(&user); err != nil {
		return utils.Respond(ctx, http.StatusBadRequest, nil, err.Error())
	}
	if err := c.Service.CreateUser(user); err != nil {
		return utils.Respond(ctx, http.StatusInternalServerError, nil, err.Error())
	}
	return utils.Respond(ctx, http.StatusCreated, nil, "User created successfully")
}

func (c *AuthController) Login(ctx echo.Context) error {
	var loginRequest model.User
	if err := ctx.Bind(&loginRequest); err != nil {
		return utils.Respond(ctx, http.StatusBadRequest, nil, err.Error())
	}

	user, err := c.Service.AuthenticateUser(loginRequest.Username, loginRequest.Password)
	if err != nil {
		return utils.Respond(ctx, http.StatusUnauthorized, nil, "Invalid username or password")
	}

	token, err := utils.GenerateToken(user.Username)
	if err != nil {
		return utils.Respond(ctx, http.StatusInternalServerError, nil, "Failed to generate token")
	}

	return utils.Respond(ctx, http.StatusOK, map[string]string{"token": token}, "Login successful")
}
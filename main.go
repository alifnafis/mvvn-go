package main

import (
	"go-api/config"
	"go-api/controller"
	"go-api/repository"
	"go-api/router"
	"go-api/service"

	"github.com/labstack/echo/v4"
)

func main() {
	config.InitDB()
	e := echo.New()
	authRepo := repository.NewAuthRepository(config.DB)
	authService := service.NewAuthService(authRepo)
	authController := controller.NewAuthController(authService)

	userRepo := repository.NewUserRepository(config.DB)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	musicRepo := repository.NewMusicRepository(config.DB)
	musicService := service.NewMusicService(musicRepo)
	musicController := controller.NewMusicController(musicService)

	playlistRepo := repository.NewPlaylistRepository(config.DB)
	playlistService := service.NewPlaylistService(playlistRepo, musicRepo)
	playlistController := controller.NewPlaylistController(playlistService)
	router.InitRouter(e, authController, userController, musicController, playlistController)
	e.Logger.Fatal(e.Start(":8080"))
}
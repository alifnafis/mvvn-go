package service

import (
	"go-api/model"
	"go-api/repository"
)

type PlaylistService struct {
	PlaylistRepo *repository.PlaylistRepository
	MusicRepo    *repository.MusicRepository
}

func NewPlaylistService(playlistRepo *repository.PlaylistRepository, musicRepo *repository.MusicRepository) *PlaylistService {
	return &PlaylistService{PlaylistRepo: playlistRepo, MusicRepo: musicRepo}
}

func (s *PlaylistService) GetMyPlaylists(userID int) ([]model.Playlist, error) {
	return s.PlaylistRepo.GetMyPlaylists(userID)
}

func (s *PlaylistService) CreatePlaylist(playlist model.Playlist) error {
	return s.PlaylistRepo.CreatePlaylist(playlist)
}

func (s *PlaylistService) DeletePlaylist(id int) error {
	return s.PlaylistRepo.DeletePlaylist(id)
}

func (s *PlaylistService) AddMusicToPlaylist(playlistID, musicID int) error {
	_, err := s.MusicRepo.GetMusicByID(musicID)
	if err != nil {
		return err
	}

	return s.PlaylistRepo.AddMusicToPlaylist(playlistID, musicID)
}

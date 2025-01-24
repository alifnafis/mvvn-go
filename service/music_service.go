package service

import (
	"go-api/model"
	"go-api/repository"
)

type MusicService struct {
	MusicRepo *repository.MusicRepository
}

func NewMusicService(musicRepo *repository.MusicRepository) *MusicService {
	return &MusicService{MusicRepo: musicRepo}
}

func (s *MusicService) GetAllMusic() ([]model.Music, error) {
	return s.MusicRepo.GetAllMusic()
}

func (s *MusicService) GetMusicByID(id int) (*model.Music, error) {
	return s.MusicRepo.GetMusicByID(id)
}

func (s *MusicService) CreateMusic(music model.Music) error {
	return s.MusicRepo.CreateMusic(music)
}

func (s *MusicService) UpdateMusic(id int, updatedMusic model.Music) error {
	existingMusic, err := s.MusicRepo.GetMusicByID(id)
	if err != nil {
		return err
	}

	if updatedMusic.Title != "" {
		existingMusic.Title = updatedMusic.Title
	}
	if updatedMusic.Artist != "" {
		existingMusic.Artist = updatedMusic.Artist
	}
	if updatedMusic.Album != "" {
		existingMusic.Album = updatedMusic.Album
	}
	if updatedMusic.Genre != "" {
		existingMusic.Genre = updatedMusic.Genre
	}
	if updatedMusic.Duration > 0 {
		existingMusic.Duration = updatedMusic.Duration
	}

	return s.MusicRepo.UpdateMusic(id, *existingMusic)
}

func (s *MusicService) DeleteMusic(id int) error {
	return s.MusicRepo.DeleteMusic(id)
}

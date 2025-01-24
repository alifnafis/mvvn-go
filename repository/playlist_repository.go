package repository

import (
	"fmt"
	"go-api/model"
	"go-api/utils"

	"github.com/jmoiron/sqlx"
)

type PlaylistRepository struct {
	DB *sqlx.DB
}

func NewPlaylistRepository(db *sqlx.DB) *PlaylistRepository {
	return &PlaylistRepository{DB: db}
}

func (r *PlaylistRepository) GetMyPlaylists(userID int) ([]model.Playlist, error) {
    var playlists []model.Playlist
    err := r.DB.Select(&playlists, utils.GetMyPlaylistsQuery, userID)
    if err != nil {
        return nil, fmt.Errorf("failed to fetch playlists: %w", err)
    }
    return playlists, nil
}

func (r *PlaylistRepository) CreatePlaylist(playlist model.Playlist) error {
	_, err := r.DB.Exec(utils.CreatePlaylistQuery, playlist.Name, playlist.UserID)
	return err
}

func (r *PlaylistRepository) DeletePlaylist(id int) error {
	_, err := r.DB.Exec(utils.DeletePlaylistQuery, id)
	return err
}

func (r *PlaylistRepository) AddMusicToPlaylist(playlistID, musicID int) error {
	_, err := r.DB.Exec(utils.AddMusicToPlaylistQuery, playlistID, musicID)
	return err
}

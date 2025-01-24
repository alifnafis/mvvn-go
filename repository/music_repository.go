package repository

import (
	"go-api/model"
	"go-api/utils"

	"github.com/jmoiron/sqlx"
)

type MusicRepository struct {
	DB *sqlx.DB
}

func NewMusicRepository(db *sqlx.DB) *MusicRepository {
	return &MusicRepository{DB: db}
}

func (r *MusicRepository) GetAllMusic() ([]model.Music, error) {
	var music []model.Music
	err := r.DB.Select(&music, utils.GetAllMusicQuery)
	return music, err
}

func (r *MusicRepository) GetMusicByID(id int) (*model.Music, error) {
	var music model.Music
	err := r.DB.Get(&music, utils.GetMusicByIDQuery, id)
	if err != nil {
		return nil, err
	}
	return &music, nil
}

func (r *MusicRepository) CreateMusic(music model.Music) error {
	_, err := r.DB.Exec(utils.CreateMusicQuery, music.Title, music.Artist, music.Album, music.Genre, music.Duration)
	return err
}

func (r *MusicRepository) UpdateMusic(id int, music model.Music) error {
	_, err := r.DB.Exec(
		utils.UpdateMusicQuery,
		music.Title, music.Artist, music.Album, music.Genre, music.Duration, id,
	)
	return err
}

func (r *MusicRepository) DeleteMusic(id int) error {
	_, err := r.DB.Exec(utils.DeleteMusicQuery, id)
	return err
}

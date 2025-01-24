package repository

import (
	"go-api/model"
	"go-api/utils"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type AuthRepository struct {
    DB *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
    return &AuthRepository{DB: db}
}

func (r *AuthRepository) CreateUser(user model.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	
	if err != nil {
		return err
	}
	_, err = r.DB.Exec(utils.CreateUserQuery, user.Name, user.Username, user.Position, user.Salary, string(hashedPassword))
	return err
}


func (r *AuthRepository) GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	err := r.DB.Get(&user, utils.FindUserByUsernameQuery, username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

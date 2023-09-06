package repository

import (
	"clean-code/model"
	"database/sql"
)

type UserRepository interface {
	FindByUsername(username string) (model.UserCredential, error)
	Save(payload model.UserCredential) error
}

type userRepository struct {
	db *sql.DB
}

// FindByUsername implements UserRepository.
func (u *userRepository) FindByUsername(username string) (model.UserCredential, error) {
	row := u.db.QueryRow("SELECT id, username, password FROM user_credential WHERE username = $1 AND is_active = $2", username, true)

	var userCredential model.UserCredential
	if err := row.Scan(&userCredential.ID, &userCredential.Username, &userCredential.Password); err != nil {
		return model.UserCredential{}, err
	}

	return userCredential, nil
}

// Save implements UserRepository.
func (u *userRepository) Save(payload model.UserCredential) error {
	if _, err := u.db.Exec("INSERT INTO user_credential VALUES($1, $2, $3, $4)", payload.ID, payload.Username, payload.Password, true); err != nil {
		return err
	}

	return nil
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

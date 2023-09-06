package usecase

import (
	"clean-code/model"
	"clean-code/model/dto"
	"clean-code/repository"
	"clean-code/util/common"
	"clean-code/util/security"
	"fmt"
)

type UserUseCase interface {
	GetByUserName(username string) (model.UserCredential, error)
	Register(payload dto.AuthRequest) error
}

type userUseCase struct {
	repo repository.UserRepository
}

// Register implements UserUseCase.
func (u *userUseCase) Register(payload dto.AuthRequest) error {
	hashPassword, err := security.HashPassword(payload.Password)
	if err != nil {
		return err
	}

	userCredential := model.UserCredential{
		ID:       common.GenerateID(),
		Username: payload.Username,
		Password: hashPassword,
	}

	if err := u.repo.Save(userCredential); err != nil {
		return fmt.Errorf("failed save user: %v", err.Error())
	}

	return nil
}

// GetByUserName implements UserUseCase.
func (u *userUseCase) GetByUserName(username string) (model.UserCredential, error) {
	uc, err := u.repo.FindByUsername(username)
	if err != nil {
		return model.UserCredential{}, err
	}

	return uc, nil
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{
		repo: repo,
	}
}

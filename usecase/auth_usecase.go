package usecase

import (
	"clean-code/model/dto"
	"clean-code/repository"
	"clean-code/util/security"
	"fmt"
)

type AuthUseCase interface {
	Login(payload dto.AuthRequest) (dto.AuthResponse, error)
}

type authUseCase struct {
	repo repository.UserRepository
}

// Login implements AuthUseCase.
func (a *authUseCase) Login(payload dto.AuthRequest) (dto.AuthResponse, error) {
	user, err := a.repo.FindByUsername(payload.Username)
	if err != nil {
		return dto.AuthResponse{}, fmt.Errorf("unauthorized: invalid credential")
	}

	if err := security.VerifyPassword(user.Password, payload.Password); err != nil {
		return dto.AuthResponse{}, fmt.Errorf("unauthorized: invalid credential")
	}

	token, err := security.GenerateJwtToken(user)
	if err != nil {
		return dto.AuthResponse{}, err
	}

	return dto.AuthResponse{
		Username: user.Username,
		Token:    token,
	}, nil
}

func NewAuthUseCase(repo repository.UserRepository) AuthUseCase {
	return &authUseCase{
		repo: repo,
	}
}

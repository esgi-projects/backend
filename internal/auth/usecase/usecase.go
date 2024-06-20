package usecase

import (
	"errors"

	"github.com/esgi-challenge/backend/config"
	"github.com/esgi-challenge/backend/internal/auth"
	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/internal/user"
	"github.com/esgi-challenge/backend/pkg/jwt"
	"github.com/esgi-challenge/backend/pkg/logger"
)

type authUseCase struct {
	userRepo user.Repository
	cfg      *config.Config
	logger   logger.Logger
}

func NewAuthUseCase(cfg *config.Config, userRepo user.Repository, logger logger.Logger) auth.UseCase {
	return &authUseCase{cfg: cfg, userRepo: userRepo, logger: logger}
}

func (u *authUseCase) Login(payload *models.AuthLogin) (*models.Auth, error) {
	user, err := u.userRepo.GetByEmail(payload.Email)

	if err != nil {
		return nil, errors.New("Wrong email")
	}

	isPasswordGood := user.CheckPassword(payload.Password)
	if !isPasswordGood {
		return nil, errors.New("Wrong password")
	}

	token, err := jwt.Generate(u.cfg.JwtSecret, user)

	if err != nil {
		return nil, err
	}

	return &models.Auth{
		Token: token,
	}, nil
}

func (u *authUseCase) Register(user *models.User) (*models.Auth, error) {
	user, err := u.userRepo.Create(user)

	if err != nil {
		return nil, err
	}

	token, err := jwt.Generate(u.cfg.JwtSecret, user)

	if err != nil {
		return nil, err
	}

	return &models.Auth{
		Token: token,
	}, nil
}

func (u *authUseCase) ResetPassword(payload *models.AuthResetPassword) error {

	user, err := u.userRepo.GetByResetCode(payload.Code)

	if err != nil {
		return err
	}

	user.Password = payload.Password
	user.PasswordResetCode = ""
	user.HashPassword()

	_, err = u.userRepo.Update(user.ID, user)

	if err != nil {
		return err
	}

	return nil
}

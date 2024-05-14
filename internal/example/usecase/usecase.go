package usecase

import (
	"github.com/esgi-challenge/backend/config"
	"github.com/esgi-challenge/backend/internal/example"
	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/pkg/logger"
)

type exampleUseCase struct {
	exampleRepo example.Repository
	cfg         *config.Config
	logger      logger.Logger
}

func NewExampleUseCase(exampleRepo example.Repository, cfg *config.Config, logger logger.Logger) example.UseCase {
	return &exampleUseCase{exampleRepo: exampleRepo, cfg: cfg, logger: logger}
}

func (u *exampleUseCase) Create(example *models.Example) (*models.Example, error) {
	return u.exampleRepo.Create(example)
}

func (u *exampleUseCase) GetAll() (*[]models.Example, error) {
	return u.exampleRepo.GetAll()
}

func (u *exampleUseCase) GetById(id uint) (*models.Example, error) {
	return u.exampleRepo.GetById(id)
}

func (u *exampleUseCase) Delete(id uint) error {
	_, err := u.GetById(id)
	if err != nil {
		return err
	}

	return u.exampleRepo.Delete(id)
}
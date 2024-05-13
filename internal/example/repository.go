package example

import (
	"github.com/esgi-challenge/backend/internal/models"
)

type Repository interface {
	Create(example *models.Example) (*models.Example, error)
  GetAll() (*[]models.Example, error)
}

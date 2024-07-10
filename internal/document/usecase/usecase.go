package usecase

import (
	"context"
	"net/http"

	"github.com/esgi-challenge/backend/config"
	"github.com/esgi-challenge/backend/internal/document"
	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/pkg/errorHandler"
	"github.com/esgi-challenge/backend/pkg/logger"
	"github.com/esgi-challenge/backend/pkg/storage"
)

type documentUseCase struct {
	documentRepo document.Repository
	cfg          *config.Config
	storage      storage.Storage
	logger       logger.Logger
}

func NewDocumentUseCase(cfg *config.Config, documentRepo document.Repository, logger logger.Logger, storage storage.Storage) document.UseCase {
	return &documentUseCase{
		cfg:          cfg,
		documentRepo: documentRepo,
		logger:       logger,
		storage:      storage,
	}
}

func (u *documentUseCase) Create(user *models.User, document *models.DocumentCreate) (*models.Document, error) {
	filename, err := u.storage.UploadFile(context.Background(), document.Byte)

	if err != nil {
		return nil, err
	}

	return u.documentRepo.Create(&models.Document{
		Path:   filename,
		UserId: user.ID,
	})
}

func (u *documentUseCase) GetById(user *models.User, id uint) (*models.Document, error) {
	document, err := u.documentRepo.GetById(id)

	if err != nil {
		return nil, err
	}

	if document.UserId != user.ID {
		return nil, errorHandler.HttpError{
			HttpStatus: http.StatusForbidden,
			HttpError:  "This document is not yours",
		}
	}

	return document, nil
}

func (u *documentUseCase) Delete(user *models.User, id uint) error {
	// Check not needed but added to handle a not found error because gorm do not return
	// error if delete on a row that does not exist
	_, err := u.GetById(user, id)
	if err != nil {
		return err
	}

	return u.documentRepo.Delete(id)
}

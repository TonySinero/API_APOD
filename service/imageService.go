package service

import (
	"github.com/apod/model"
	"github.com/apod/repository"
)

type ImageService struct {
	repo repository.Repository
}

func NewImageService(repo repository.Repository) *ImageService {
	return &ImageService{repo: repo}
}

// GetAlbumFromDB returns a list of albums in the database.
func (a *ImageService) GetAlbumFromDB() ([]model.Nasa, error) {
	return a.repo.ImageRepo.GetAll()
}

// CreateAlbum when receiving data from api, creates a record in the database.
func (a *ImageService) CreateAlbum(image *model.Nasa) (int, error) {
	return a.repo.ImageRepo.CreateAlbum(image)
}

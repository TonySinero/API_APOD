package service

import (
	"github.com/apod/internal/model"
	"github.com/apod/internal/repository"
)

type ImageService struct {
	repo repository.Repository
}

func NewImageService(repo repository.Repository) *ImageService {
	return &ImageService{repo: repo}
}

func (a *ImageService) GetAlbumFromDB() ([]model.Nasa, error) {
	return a.repo.ImageRepo.GetAll()
}

func (a *ImageService) CreateAlbum(im *model.Nasa) (int, error) {
	return a.repo.ImageRepo.CreateAlbum(im)
}

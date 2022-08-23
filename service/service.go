package service

import (
	"github.com/apod/model"
	"github.com/apod/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/service_mock.go

type Image interface {
	GetAlbumFromDB() ([]model.Nasa, error)
	CreateAlbum(im *model.Nasa) (int, error)
}

type Service struct {
	Image
}

func NewService(rep *repository.Repository) *Service {
	return &Service{
		Image: NewImageService(*rep),
	}
}

package repository

import (
	"database/sql"
	"github.com/apod/model"
)

//go:generate mockgen -source=repository.go -destination=mocks/repository_mock.go

type ImageRepo interface {
	GetAll() ([]model.Nasa, error)
	CreateAlbum(im *model.Nasa) (int, error)
}

type Repository struct {
	ImageRepo
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		ImageRepo: NewImagePostgres(db),
	}
}

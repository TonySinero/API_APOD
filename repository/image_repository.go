package repository

import (
	"database/sql"
	"fmt"
	"github.com/apod/model"
	"github.com/sirupsen/logrus"
)

// ImgPostgres type represents postgres object image structure and behavior.
type ImgPostgres struct {
	db *sql.DB
}

func NewImagePostgres(db *sql.DB) *ImgPostgres {
	return &ImgPostgres{db: db}
}

// GetAll method returns objects model.Nasa from postgres database
func (r *ImgPostgres) GetAll() ([]model.Nasa, error) {
	var images []model.Nasa
	query := "SELECT id, copyright, dates, explanation, hdurl, mediaType, serviceVersion, title, url FROM images"
	rows, err := r.db.Query(query)
	if err != nil {
		logrus.Errorf("GetAllImage: can not executes a query:%s", err)
		return nil, fmt.Errorf("GetAllImage: repository error:%w", err)
	}
	for rows.Next() {
		var image model.Nasa
		if err := rows.Scan(&image.ID, &image.Copyright, &image.Date, &image.Explanation,
			&image.Hdurl, &image.MediaType, &image.ServiceVersion, &image.Title, &image.URL); err != nil {
			logrus.Errorf("GetAllImage: error while scanning for image:%s", err)
			return nil, fmt.Errorf("GetAllImage: repository error:%w", err)
		}
		images = append(images, image)
	}
	return images, nil
}

// CreateAlbum method saves object model.Nasa into postgres database.
func (u *ImgPostgres) CreateAlbum(image *model.Nasa) (int, error) {
	logrus.WithFields(logrus.Fields{
		"Copyright":      image.Copyright,
		"Date":           image.Date,
		"Explanation":    image.Explanation,
		"Hdurl":          image.Hdurl,
		"MediaType":      image.MediaType,
		"ServiceVersion": image.ServiceVersion,
		"Title":          image.Title,
		"URL":            image.URL,
	}).Info("postgres repository: create album")
	var id int
	row := u.db.QueryRow("INSERT INTO images (copyright, dates, explanation, hdurl, mediaType, serviceVersion, title, url) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id",
		image.Copyright, image.Date, image.Explanation, image.Hdurl, image.MediaType, image.ServiceVersion, image.Title, image.URL)
	if err := row.Scan(&id); err != nil {
		logrus.Errorf("CreateAlbum: error while scanning for album:%s", err)
		return 0, fmt.Errorf("CreateAlbum: error while scanning for album:%w", err)
	}
	return id, nil
}

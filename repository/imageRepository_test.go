package repository

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/apod/model"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		logrus.Fatal(err)
	}
	defer db.Close()
	r := NewRepository(db)

	testTable := []struct {
		name           string
		mock           func()
		expectedResult []model.Nasa
		expectedError  bool
	}{
		{
			name: "OK",

			mock: func() {
				rows := sqlmock.NewRows([]string{"1", "copyright", "dates", "explanation", "hdurl", "mediaType", "serviceVersion", "title", "url"}).
					AddRow("1", "test juan", "2022-08-18", "tt", "ttt", "image", "test v1", "test title", "test url").
					AddRow("2", "test juan2", "2022-08-19", "tt2", "ttt2", "image2", "test v2", "test title2", "test url2")
				mock.ExpectQuery("SELECT id, copyright,").WillReturnRows(rows)
			},
			expectedResult: []model.Nasa{
				{
					ID:             1,
					Copyright:      "test juan",
					Date:           "2022-08-18",
					Explanation:    "tt",
					Hdurl:          "ttt",
					MediaType:      "image",
					ServiceVersion: "test v1",
					Title:          "test title",
					URL:            "test url",
				},
				{
					ID:             2,
					Copyright:      "test juan2",
					Date:           "2022-08-19",
					Explanation:    "tt2",
					Hdurl:          "ttt2",
					MediaType:      "image2",
					ServiceVersion: "test v2",
					Title:          "test title2",
					URL:            "test url2",
				},
			},
			expectedError: false,
		},

		{
			name: "Data base error",
			mock: func() {
				mock.ExpectQuery("SELECT id, copyright,").WillReturnError(errors.New("data base error"))
			},
			expectedError: true,
		},
	}
	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			image, err := r.GetAll()
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResult, image)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestCreateAlbum(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		logrus.Fatal(err)
	}
	defer db.Close()
	r := NewRepository(db)

	testTable := []struct {
		name            string
		mock            func(album *model.Nasa)
		InputAlbum      *model.Nasa
		expectedAlbumId int
		expectedError   bool
	}{
		{
			name: "OK",
			mock: func(image *model.Nasa) {
				rows := sqlmock.NewRows([]string{"id"}).
					AddRow(1)
				mock.ExpectQuery("INSERT INTO images").WithArgs(image.Copyright, image.Date, image.Explanation, image.Hdurl, image.MediaType, image.ServiceVersion, image.Title, image.URL).
					WillReturnRows(rows)
			},
			InputAlbum: &model.Nasa{
				ID:             1,
				Copyright:      "test juan",
				Date:           "2022-08-18",
				Explanation:    "tt",
				Hdurl:          "ttt",
				MediaType:      "image",
				ServiceVersion: "test v1",
				Title:          "test title",
				URL:            "test url",
			},
			expectedAlbumId: 1,
			expectedError:   false,
		},
	}
	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.InputAlbum)
			got, err := r.CreateAlbum(tt.InputAlbum)
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedAlbumId, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type PostgresDB struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(database PostgresDB) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		database.Username, database.Password, database.Host, database.Port, database.DBName, database.SSLMode))
	if err != nil {

		return nil, fmt.Errorf("error connecting to database:%s", err)
	}
	err = db.Ping()
	if err != nil {
		logrus.Errorf("DB ping error:%s", err)
		return nil, err
	}
	_, err = db.Exec(ALBUM_SCHEMA)
	if err != nil {
		logrus.Errorf("Error executing initial migration:%s", err)
		return nil, fmt.Errorf("error executing initial migration:%s", err)
	}
	return db, nil
}

const ALBUM_SCHEMA = `
	CREATE TABLE IF NOT EXISTS images (
		id serial not null primary key,
		copyright varchar(225) NOT NULL,
		dates varchar(255) NOT NULL,
	    explanation varchar NOT NULL,
		hdurl varchar(225) NOT NULL,
	    mediaType varchar(225) NOT NULL,
		serviceVersion varchar(225) NOT NULL,
	    title varchar(225) NOT NULL,
		url varchar(225) NOT NULL
	);
`

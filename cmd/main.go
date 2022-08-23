// Nasa daily image API
//
// Documentation for Nasa daily image API
//
// Schemes: http
// Host: localhost:8080
// BasePath: /
// Version: 1.0.0
//
// Consumes:
//  - application/json
//
// Produces:
//  - application/json
//
// swagger:meta

package main

import (
	"context"
	"github.com/apod/handler"
	"github.com/apod/internal/database"
	"github.com/apod/internal/server"
	"github.com/apod/repository"
	"github.com/apod/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Fatalf("Error loading .env file. %s", err.Error())
	}
	db, err := database.NewPostgresDB(database.PostgresDB{
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_DATABASE"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
	})
	if err != nil {
		logrus.Panicf("failed to initialize db:%s", err.Error())
	}

	r := repository.NewRepository(db)
	s := service.NewService(r)
	h := handler.NewHandler(s)

	port := os.Getenv("API_SERVER_PORT")
	serv := new(server.Server)

	go func() {
		err := serv.Run(port, h.InitRoutes())
		if err != nil {
			logrus.Panicf("Error occured while running http server: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	if err := serv.Shutdown(context.Background()); err != nil {
		logrus.Panicf("Error occured while shutting down http server: %s", err.Error())
	}
}

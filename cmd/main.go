package main

import (
	"context"
	"github.com/apod/internal/handler"
	"github.com/apod/internal/repository"
	"github.com/apod/internal/service"
	"github.com/apod/pkg/database"
	"github.com/apod/server"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

// @title APOD Service
// @description NASA daily imaging service

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

	repo := repository.NewRepository(db)
	servic := service.NewService(repo)
	handlers := handler.NewHandler(servic)

	port := os.Getenv("API_SERVER_PORT")
	serv := new(server.Server)

	go func() {
		err := serv.Run(port, handlers.InitRoutes())
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

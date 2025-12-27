package main

import (
	"income-flow/internal/handler"
	"income-flow/internal/repository"
	"income-flow/internal/service"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	incomeDB, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5433",
		Username: "postgres",
		Password: "postgres",
		DBName:   "postgres",
		SSLMode:  "disable",
	})
	if err != nil {
		logrus.Fatalf("error initializing database connection: %s", err.Error())
	}

	repo := repository.NewRepository(incomeDB)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	app := handlers.InitRoutes()

	if err := app.Listen(":8080"); err != nil {
		logrus.Fatal(err)
	}
}

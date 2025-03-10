package main

import (
	"github.com/MaksimPozharskiy/loyalty-service/internal/db"
	"github.com/MaksimPozharskiy/loyalty-service/internal/handlers"
	"github.com/MaksimPozharskiy/loyalty-service/internal/repository"
	"github.com/MaksimPozharskiy/loyalty-service/internal/service"
)

func main() {
	err := parseFlags()
	if err != nil {
		panic(err)
	}

	dbConn, err := db.Connect(flagDBConnectionString)
	if err != nil {
		panic(err)
	} else {
		defer dbConn.Close()
	}

	storageRepository := repository.NewDBStorageRepository(dbConn)
	userService := service.NewUserService(storageRepository)
	handlers := handlers.NewUserHandler(userService)

	// Дальше написать server

}

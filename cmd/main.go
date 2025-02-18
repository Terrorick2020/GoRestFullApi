package main

import (
	"fmt"
	"log"
	"errors"

	"github.com/Terrorick2020/GoRestFullApi/internal"
	"github.com/Terrorick2020/GoRestFullApi/pkg/handler"
	"github.com/Terrorick2020/GoRestFullApi/pkg/repository"
	"github.com/Terrorick2020/GoRestFullApi/pkg/service"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error init configs: %s", err.Error())
	}

	if err := migrateDb(); err != nil {
		log.Fatalf("error migrate to db: %s", err.Error())
	}

	repos := repository.NewRepository()
	service := service.NewService(repos)
	handler := handler.NewHandler(service)

	svr := new(internal.Server)
	if err := svr.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		log.Fatalf("error occured while server run: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func migrateDb() error {
	databaseURL := viper.GetString("db_url")

	if databaseURL == "" {
		log.Fatal("db_url is not set")
	}

	m, err := migrate.New(
		"file://./migrations",
		databaseURL,
	)
	if err != nil {
		return err
	}

    err = m.Up()
    if err != nil {
        if errors.Is(err, migrate.ErrNoChange) {
            fmt.Println("No new migrations to apply.")
            return nil
        }
        return err
    }

	fmt.Println("Migrations applied successfully")

	return nil
}

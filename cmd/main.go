package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/Terrorick2020/GoRestFullApi/internal"
	"github.com/Terrorick2020/GoRestFullApi/pkg/handler"
	"github.com/Terrorick2020/GoRestFullApi/pkg/repository"
	"github.com/Terrorick2020/GoRestFullApi/pkg/service"
	"github.com/sirupsen/logrus"

	"github.com/golang-migrate/migrate/v4"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error init configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error load env: %s", err.Error())
	}

	if err := migrateDb(); err != nil {
		logrus.Fatalf("error migrate to db: %s", err.Error())
	}

	db, err := repository.NewPostgresDb(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   viper.GetString("db.dbname"),
		SslMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handler := handler.NewHandler(service)

	svr := new(internal.Server)
	if err := svr.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while server run: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	fmt.Println("Configs applied successfully")

	return viper.ReadInConfig()
}

func migrateDb() error {
	databaseURL := os.Getenv("DB_URL")

	if databaseURL == "" {
		return errors.New("DB_URL is not set")
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

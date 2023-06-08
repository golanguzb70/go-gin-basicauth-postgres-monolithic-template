package main

import (
	"fmt"

	"github.com/golanguzb70/go-gin-basicauth-postgres-monolithic-template/api"
	"github.com/golanguzb70/go-gin-basicauth-postgres-monolithic-template/config"
	"github.com/golanguzb70/go-gin-basicauth-postgres-monolithic-template/pkg/db"
	"github.com/golanguzb70/go-gin-basicauth-postgres-monolithic-template/pkg/logger"
	"github.com/golanguzb70/go-gin-basicauth-postgres-monolithic-template/storage"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel)

	pgxUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase)

	dbConn, err := db.New(pgxUrl)
	if err != nil {
		panic(err)
	}

	server := api.New(api.Option{
		Conf:     cfg,
		Logger:   log,
		Postgres: storage.New(dbConn, log, &cfg),
	})
	if err := server.Run(":" + cfg.HTTPPort); err != nil {
		log.Fatal("failed to run http server", err)
		panic(err)
	}
}

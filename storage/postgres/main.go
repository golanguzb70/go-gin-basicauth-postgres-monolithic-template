package postgres

import (
	"time"

	"github.com/golanguzb70/go-gin-basicauth-postgres-monolithic-template/config"
	"github.com/golanguzb70/go-gin-basicauth-postgres-monolithic-template/pkg/db"
	"github.com/golanguzb70/go-gin-basicauth-postgres-monolithic-template/pkg/logger"
)

var (
	CreatedAt time.Time
	UpdatedAt time.Time
)

type postgresRepo struct {
	Db  *db.Postgres
	Log *logger.Logger
	Cfg config.Config
}

func New(db *db.Postgres, log *logger.Logger, cfg config.Config) PostgresI {
	return &postgresRepo{
		Db:  db,
		Log: log,
		Cfg: cfg,
	}
}

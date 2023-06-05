package mediarepo

import (
	"context"
	"time"

	"github.com/golanguzb70/go-gin-basicauth-monolithic-template/config"
	"github.com/golanguzb70/go-gin-basicauth-monolithic-template/pkg/db"
	"github.com/golanguzb70/go-gin-basicauth-monolithic-template/pkg/logger"
)

var (
	CreatedAt time.Time
	UpdatedAt time.Time
)

type MediaI interface {
	Create(ctx context.Context, req *CreateReq) (*FullResponse, error)
	FindOne(ctx context.Context, req *FindOneReq) (*FullResponse, error)
	FindList(ctx context.Context, req *FindListReq) ([]*FullResponse, error)
	Update(ctx context.Context, req *UpdateReq) (*FullResponse, error)
	Delete(ctx context.Context, req *DeleteReq) error
}

type MediaRepo struct {
	Db  *db.Postgres
	Log *logger.Logger
	Cfg *config.Config
}

func New(db *db.Postgres, log *logger.Logger, cfg *config.Config) MediaI {
	return &MediaRepo{
		Db:  db,
		Log: log,
		Cfg: cfg,
	}
}

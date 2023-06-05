package templaterepo

import (
	"context"
	"time"

	"github.com/golanguzb70/go-gin-basicauth-monolithic-template/config"
	"github.com/golanguzb70/go-gin-basicauth-monolithic-template/models"
	"github.com/golanguzb70/go-gin-basicauth-monolithic-template/pkg/db"
	"github.com/golanguzb70/go-gin-basicauth-monolithic-template/pkg/logger"
)

var (
	CreatedAt time.Time
	UpdatedAt time.Time
)

type TemplateI interface {
	Create(ctx context.Context, req *models.TemplateCreateReq) (*models.TemplateResponse, error)
	Get(ctx context.Context, req *models.TemplateGetReq) (*models.TemplateResponse, error)
	Find(ctx context.Context, req *models.TemplateFindReq) (*models.TemplateFindResponse, error)
	Update(ctx context.Context, req *models.TemplateUpdateReq) (*models.TemplateResponse, error)
	Delete(ctx context.Context, req *models.TemplateDeleteReq) error
}

type TemplateRepo struct {
	Db  *db.Postgres
	Log *logger.Logger
	Cfg *config.Config
}

func New(db *db.Postgres, log *logger.Logger, cfg *config.Config) TemplateI {
	return &TemplateRepo{
		Db:  db,
		Log: log,
		Cfg: cfg,
	}
}

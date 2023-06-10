package postgres

import (
	"context"

	"github.com/golanguzb70/go-gin-basicauth-postgres-monolithic-template/models"
)

type PostgresI interface {
	TemplateCreate(ctx context.Context, req *models.TemplateCreateReq) (*models.TemplateResponse, error)
	TemplateGet(ctx context.Context, req *models.TemplateGetReq) (*models.TemplateResponse, error)
	TemplateFind(ctx context.Context, req *models.TemplateFindReq) (*models.TemplateFindResponse, error)
	TemplateUpdate(ctx context.Context, req *models.TemplateUpdateReq) (*models.TemplateResponse, error)
	TemplateDelete(ctx context.Context, req *models.TemplateDeleteReq) error
}

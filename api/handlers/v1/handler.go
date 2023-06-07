package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/golanguzb70/go-gin-basicauth-postgres-monolithic-template/config"
	"github.com/golanguzb70/go-gin-basicauth-postgres-monolithic-template/pkg/logger"
	"github.com/golanguzb70/go-gin-basicauth-postgres-monolithic-template/storage"
)

type HandlerV1I interface {
	TemplateCreate(c *gin.Context)
	TemplateGet(c *gin.Context)
	TemplateFind(c *gin.Context)
	TemplateUpdate(c *gin.Context)
	TemplateDelete(c *gin.Context)
}

type handlerV1 struct {
	log     *logger.Logger
	cfg     config.Config
	storage storage.StorageI
}

type HandlerV1Config struct {
	Logger   *logger.Logger
	Cfg      config.Config
	Postgres storage.StorageI
}

// New ...
func New(c *HandlerV1Config) HandlerV1I {
	return &handlerV1{
		log:     c.Logger,
		cfg:     c.Cfg,
		storage: c.Postgres,
	}
}

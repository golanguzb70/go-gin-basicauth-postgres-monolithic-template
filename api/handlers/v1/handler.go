package v1

import (
	"github.com/golanguzb70/go-gin-basicauth-postgres-monolithic-template/config"
	"github.com/golanguzb70/go-gin-basicauth-postgres-monolithic-template/pkg/logger"
	"github.com/golanguzb70/go-gin-basicauth-postgres-monolithic-template/storage"
)

type HandlerV1I interface {
	Template() TemplateI
}

type handlerV1 struct {
	template TemplateI
}

type HandlerV1Config struct {
	Logger   *logger.Logger
	Cfg      config.Config
	Postgres storage.StorageI
}

// New ...
func New(c *HandlerV1Config) HandlerV1I {
	return &handlerV1{
		template: NewTemplateHandler(c),
	}
}

func (h *handlerV1) Template() TemplateI {
	return h.template
}

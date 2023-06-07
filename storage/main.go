package storage

import (
	"github.com/golanguzb70/go-gin-basicauth-postgres-monolithic-template/config"
	"github.com/golanguzb70/go-gin-basicauth-postgres-monolithic-template/pkg/db"
	"github.com/golanguzb70/go-gin-basicauth-postgres-monolithic-template/pkg/logger"
	"github.com/golanguzb70/go-gin-basicauth-postgres-monolithic-template/storage/postgres/templaterepo"
)

type StorageI interface {
	Template() templaterepo.TemplateI
}

type StoragePg struct {
	template templaterepo.TemplateI
}

// NewStoragePg
func New(db *db.Postgres, log *logger.Logger, cfg *config.Config) StorageI {
	return &StoragePg{
		template: templaterepo.New(db, log, cfg),
	}
}

func (s *StoragePg) Template() templaterepo.TemplateI {
	return s.template
}

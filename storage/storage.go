package storage

import (
	"github.com/golanguzb70/go-gin-basicauth-monolithic-template/config"
	"github.com/golanguzb70/go-gin-basicauth-monolithic-template/pkg/db"
	"github.com/golanguzb70/go-gin-basicauth-monolithic-template/pkg/logger"
	"github.com/golanguzb70/go-gin-basicauth-monolithic-template/storage/postgres/aboutrepo"
	"github.com/golanguzb70/go-gin-basicauth-monolithic-template/storage/postgres/mediarepo"
	"github.com/golanguzb70/go-gin-basicauth-monolithic-template/storage/postgres/postsrepo"
)

type StorageI interface {
	About() aboutrepo.AboutI
	Posts() postsrepo.PostsI
	Media() mediarepo.MediaI
}

type StoragePg struct {
	aboutRepo aboutrepo.AboutI
	postsRepo postsrepo.PostsI
	mediaRepo mediarepo.MediaI
}

// NewStoragePg
func New(db *db.Postgres, log *logger.Logger, cfg *config.Config) StorageI {
	return &StoragePg{
		aboutRepo: aboutrepo.New(db, log, cfg),
		postsRepo: postsrepo.New(db, log, cfg),
		mediaRepo: mediarepo.New(db, log, cfg),
	}
}

func (s *StoragePg) About() aboutrepo.AboutI {
	return s.aboutRepo
}

func (s *StoragePg) Posts() postsrepo.PostsI {
	return s.postsRepo
}

func (s *StoragePg) Media() mediarepo.MediaI {
	return s.mediaRepo
}

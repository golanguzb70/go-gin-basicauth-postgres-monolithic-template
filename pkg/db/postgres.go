// Package postgres implements postgres connection.
package db

import (
	"fmt"
	"log"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/golanguzb70/go-gin-basicauth-postgres-monolithic-template/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	_defaultConnAttempts = 10
	_defaultConnTimeout  = time.Second
)

// Postgres -.
type Postgres struct {
	connAttempts int
	connTimeout  time.Duration

	Builder squirrel.StatementBuilderType
	Db      *sqlx.DB
}

// New -.
func New(cfg config.Config, opts ...Option) (*Postgres, error) {
	pg := &Postgres{
		connAttempts: _defaultConnAttempts,
		connTimeout:  _defaultConnTimeout,
	}

	// Custom options
	for _, opt := range opts {
		opt(pg)
	}

	pgxUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	)

	pg.Builder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	var err error
	for pg.connAttempts > 0 {
		pg.Db, err = sqlx.Connect("postgres", pgxUrl)
		if err == nil {
			break
		}

		log.Printf("Postgres is trying to connect, attempts left: %d", pg.connAttempts)

		time.Sleep(pg.connTimeout)

		pg.connAttempts--
	}

	if err != nil {
		return nil, fmt.Errorf("postgres - NewPostgres - connAttempts == 0: %w", err)
	}

	return pg, nil
}

// Close -.
func (p *Postgres) Close() {
	if p.Db != nil {
		p.Db.Close()
	}
}

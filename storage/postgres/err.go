package postgres

import (
	"database/sql"

	"github.com/golanguzb70/go-gin-basicauth-postgres-monolithic-template/pkg/logger"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func HandleDatabaseError(err error, log *logger.Logger, message string) error {
	if err == nil {
		return nil
	}
	log.Error(message + ": " + err.Error())
	switch err {
	case sql.ErrNoRows:
		return status.Error(codes.NotFound, "This information is not exists.")
	case sql.ErrConnDone:
		return err
	case sql.ErrTxDone:
		return err
	}

	switch e := err.(type) {
	case *pq.Error:
		// Handle Postgres-specific errors
		switch e.Code.Name() {
		case "unique_violation":
			return status.Error(codes.AlreadyExists, "Already exists")
		case "foreign_key_violation":
			return status.Error(codes.InvalidArgument, "Oops something went wrong")
		default:
			return err
		}
	default:
		// Handle all other errors
		return err
	}
}

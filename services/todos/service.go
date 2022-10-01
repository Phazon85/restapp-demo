package todos

import (
	"database/sql"

	"go.uber.org/zap"
)

type Service struct {
	Logger *zap.Logger
	DB     *sql.DB
}

func New(logger *zap.Logger, db *sql.DB) *Service {
	return &Service{
		Logger: logger,
		DB:     db,
	}
}

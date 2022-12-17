package sqlite

import (
	"database/sql"

	"github.com/ilfey/go-back/internal/pkg/store"
	"github.com/sirupsen/logrus"
)

func New(db *sql.DB, logger *logrus.Logger) *store.Store {
	return &store.Store{
		User: &userRepository{
			db: db,
			logger: logger.WithFields(logrus.Fields{
				"repository": "user",
			}),
		},
	}
}

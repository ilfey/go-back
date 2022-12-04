package pgsql

import (
	"github.com/ilfey/go-back/internal/pkg/store"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

func New(db *pgx.Conn, logger *logrus.Logger) *store.Store {
	return &store.Store{
		User: &userRepository{
			db: db,
			logger: logger.WithFields(logrus.Fields{
				"repository": "user",
			}),
		},
	}
}

// func User(db *pgx.Conn, logger *logrus.Logger) store.UserRepository {

// 	return &userRepository{
// 		db: db,
// 		logger: logger.WithFields(logrus.Fields{
// 			"repository": "user",
// 		}),
// 	}
// }

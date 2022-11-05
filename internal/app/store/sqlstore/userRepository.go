package sqlstore

import (
	"context"
	"fmt"

	"github.com/ilfey/go-back/internal/app/store/models"
	"github.com/jackc/pgx/v5/pgconn"
)

type userRepository struct {
	store *Store
}

func (r *userRepository) Create(ctx context.Context, u *models.User) error {
	q := "INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id"

	r.store.logger.Tracef("SQL Query: %s", q)

	if err := r.store.db.QueryRow(ctx, q, u.Username, u.Email, u.Password).Scan(&u.Id); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			newErr := fmt.Errorf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			r.store.logger.Error(newErr)
			return newErr
		}

		return err
	}

	return nil
}

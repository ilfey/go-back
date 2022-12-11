package pgsql

import (
	"context"
	"fmt"

	"github.com/ilfey/go-back/internal/pkg/store/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/sirupsen/logrus"
)

type userRepository struct {
	db     *pgx.Conn
	logger *logrus.Entry
}

func (r *userRepository) Create(ctx context.Context, u *models.User) error {
	q := "INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id"

	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	r.logger.Tracef("SQL Query: %s", q)

	if err := r.db.QueryRow(ctx, q, u.Username, u.Email, u.Password).Scan(&u.Id); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			newErr := fmt.Errorf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			r.logger.Error(newErr)
			return newErr
		}

		return err
	}

	return nil
}

func (r *userRepository) FindById(ctx context.Context, id int) (*models.User, error) {

	q := "SELECT id, username, email, password, is_deleted FROM users WHERE id = $1"

	r.logger.Tracef("SQL Query: %s", q)

	var u models.User

	if err := r.db.QueryRow(ctx, q, id).Scan(&u.Id, &u.Username, &u.Email, &u.Password, &u.IsDeleted); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			newErr := fmt.Errorf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			r.logger.Error(newErr)
			return nil, newErr
		}

		return nil, err
	}

	return &u, nil
}

func (r *userRepository) FindByUsername(ctx context.Context, username string) (*models.User, error) {

	q := "SELECT id, username, email, password, is_deleted FROM users WHERE username = $1"

	r.logger.Tracef("SQL Query: %s", q)

	var u models.User

	if err := r.db.QueryRow(ctx, q, username).Scan(&u.Id, &u.Username, &u.Email, &u.Password, &u.IsDeleted); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			newErr := fmt.Errorf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			r.logger.Error(newErr)
			return nil, newErr
		}

		return nil, err
	}

	return &u, nil
}

func (r *userRepository) FindByUsernameWithPassword(ctx context.Context, username, password string) (*models.User, error) {

	q := "SELECT id, username, email, password, is_deleted FROM users WHERE username = $1 AND password = $2"

	r.logger.Tracef("SQL Query: %s", q)

	var u models.User

	if err := r.db.QueryRow(ctx, q, username, password).Scan(&u.Id, &u.Username, &u.Email, &u.Password, &u.IsDeleted); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			newErr := fmt.Errorf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			r.logger.Error(newErr)
			return nil, newErr
		}

		return nil, err
	}

	return &u, nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	q := "SELECT id, username, email, password, is_deleted FROM users WHERE email = $1"

	r.logger.Tracef("SQL Query: %s", q)

	var u models.User

	if err := r.db.QueryRow(ctx, q, email).Scan(&u.Id, &u.Username, &u.Email, &u.Password, &u.IsDeleted); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			newErr := fmt.Errorf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			r.logger.Error(newErr)
			return nil, newErr
		}

		return nil, err
	}

	return &u, nil
}

func (r *userRepository) FindByEmailWithPassword(ctx context.Context, email, password string) (*models.User, error) {
	q := "SELECT id, username, email, password, is_deleted FROM users WHERE email = $1 AND password = $2"

	r.logger.Tracef("SQL Query: %s", q)

	var u models.User

	if err := r.db.QueryRow(ctx, q, email, password).Scan(&u.Id, &u.Username, &u.Email, &u.Password, &u.IsDeleted); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			newErr := fmt.Errorf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			r.logger.Error(newErr)
			return nil, newErr
		}

		return nil, err
	}

	return &u, nil
}

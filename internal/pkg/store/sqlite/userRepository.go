package sqlite

import (
	"context"
	"database/sql"

	"github.com/ilfey/go-back/internal/pkg/store/models"
	"github.com/sirupsen/logrus"
)

type userRepository struct {
	db     *sql.DB
	logger *logrus.Entry
}

func (r *userRepository) Create(ctx context.Context, u *models.User) error {
	q := "INSERT INTO users (username, email, password) VALUES (?, ?, ?) RETURNING id"

	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	r.logger.Tracef("SQL Query: %s", q)

	if err := r.db.QueryRowContext(ctx, q, u.Username, u.Email, u.Password).Scan(&u.Id); err != nil {
		r.logger.Error(err.Error())
		return err
	}

	return nil
}

func (r *userRepository) FindById(ctx context.Context, id int) (*models.User, error) {

	q := "SELECT id, username, email, password, is_deleted FROM users WHERE id = ?"

	r.logger.Tracef("SQL Query: %s", q)

	var u models.User

	if err := r.db.QueryRowContext(ctx, q, id).Scan(&u.Id, &u.Username, &u.Email, &u.Password, &u.IsDeleted); err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return &u, nil
}

func (r *userRepository) FindByUsername(ctx context.Context, username string) (*models.User, error) {

	q := "SELECT id, username, email, password, is_deleted FROM users WHERE username = ?"

	r.logger.Tracef("SQL Query: %s", q)

	var u models.User

	if err := r.db.QueryRowContext(ctx, q, username).Scan(&u.Id, &u.Username, &u.Email, &u.Password, &u.IsDeleted); err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return &u, nil
}

func (r *userRepository) FindByUsernameWithPassword(ctx context.Context, username, password string) (*models.User, error) {

	q := "SELECT id, username, email, password, is_deleted FROM users WHERE username = ? AND password = ?"

	r.logger.Tracef("SQL Query: %s", q)

	var u models.User

	if err := r.db.QueryRowContext(ctx, q, username, password).Scan(&u.Id, &u.Username, &u.Email, &u.Password, &u.IsDeleted); err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return &u, nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	q := "SELECT id, username, email, password, is_deleted FROM users WHERE email = ?"

	r.logger.Tracef("SQL Query: %s", q)

	var u models.User

	if err := r.db.QueryRowContext(ctx, q, email).Scan(&u.Id, &u.Username, &u.Email, &u.Password, &u.IsDeleted); err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return &u, nil
}

func (r *userRepository) FindByEmailWithPassword(ctx context.Context, email, password string) (*models.User, error) {
	q := "SELECT id, username, email, password, is_deleted FROM users WHERE email = ? AND password = ?"

	r.logger.Tracef("SQL Query: %s", q)

	var u models.User

	if err := r.db.QueryRowContext(ctx, q, email, password).Scan(&u.Id, &u.Username, &u.Email, &u.Password, &u.IsDeleted); err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return &u, nil
}

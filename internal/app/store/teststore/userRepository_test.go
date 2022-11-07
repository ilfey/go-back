package teststore_test

import (
	"context"
	"testing"

	"github.com/ilfey/go-back/internal/app/store/models"
	"github.com/ilfey/go-back/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()
	u := models.TestUser(t)
	assert.NoError(t, s.User().Create(context.Background(), u))
	assert.NotNil(t, u.Id)
}

func TestUserRepository_FindById(t *testing.T) {
	s := teststore.New()
	u1 := models.TestUser(t)
	s.User().Create(context.Background(), u1)
	u2, err := s.User().FindById(context.Background(), u1.Id)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}

func TestUserRepository_FindByUsername(t *testing.T) {
	s := teststore.New()
	u1 := models.TestUser(t)
	s.User().Create(context.Background(), u1)
	u2, err := s.User().FindByUsername(context.Background(), u1.Username)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}

func TestUserRepository_FindByUsernameWithPassword(t *testing.T) {
	s := teststore.New()
	u1 := models.TestUser(t)
	s.User().Create(context.Background(), u1)
	u2, err := s.User().FindByUsernameWithPassword(context.Background(), u1.Username, u1.Password)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s := teststore.New()
	u1 := models.TestUser(t)
	_, err := s.User().FindByEmail(context.Background(), u1.Email)
	assert.EqualError(t, err, teststore.ErrNotFound.Error())

	s.User().Create(context.Background(), u1)
	u2, err := s.User().FindByEmailWithPassword(context.Background(), u1.Email, u1.Password)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}

func TestUserRepository_FindByEmailWithPassword(t *testing.T) {
	s := teststore.New()
	u1 := models.TestUser(t)
	_, err := s.User().FindByEmailWithPassword(context.Background(), u1.Email, u1.Password)
	assert.EqualError(t, err, teststore.ErrNotFound.Error())

	s.User().Create(context.Background(), u1)
	u2, err := s.User().FindByEmailWithPassword(context.Background(), u1.Email, u1.Password)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}

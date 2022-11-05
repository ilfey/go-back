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

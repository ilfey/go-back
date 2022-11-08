package models

import "testing"

func TestUser(t *testing.T) *User {
	t.Helper()

	return &User{
		Username: "ilfey",
		Email:    "ilfey@example.org",
		Password: "password",
	}
}

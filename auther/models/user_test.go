package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	user := User{
		Password: "secret",
	}
	_, err := user.HashUserPassword()
	assert.NoError(t, err)
}

func TestCreateUserRecord(t *testing.T) {
	userPassword := "mySecretPass"

	user := User{
		Name:     "Name",
		Email:    "me@example.com",
		Password: userPassword,
	}
	userHashedPassword, err := user.HashUserPassword()
	assert.NoError(t, err)

	assert.Equal(t, ComparePassword(userPassword, user.Password), true)
	assert.Equal(t, ComparePassword("totally not user's password", userHashedPassword), false)
}

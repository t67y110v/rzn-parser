package store_test

import (
	"restApi/internal/app/model"
	"restApi/internal/app/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")
	u, err := s.User().Create(&model.User{
		Email:             "userTest@test.org",
		EncryptedPassword: "qwerty",
	})
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")
	email := "userTest@test.org"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)
	s.User().Create(&model.User{
		Email:             "userTest@test.org",
		EncryptedPassword: "qwerty",
	})
	u, err := s.User().FindByEmail(email)

	assert.NotNil(t, u)
	assert.NoError(t, err)
	assert.Equal(t, u.Email, "userTest@test.org")
}

package sqlstore_test

import (
	"restApi/internal/app/model"
	"restApi/internal/app/store"
	"restApi/internal/app/store/sqlstore"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")
	s := sqlstore.New(db)
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")
	s := sqlstore.New(db)
	email := "userTest1@test.org"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = email
	s.User().Create(u)
	u, err = s.User().FindByEmail(email)

	assert.NotNil(t, u)
	assert.NoError(t, err)
	assert.Equal(t, u.Email, "userTest1@test.org")
}

func TestUserRepository_UpdateRoleAdmin(t *testing.T) {
	var err error
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")
	s := sqlstore.New(db)
	u := model.TestUser(t)
	u.Isadmin = false
	s.User().Create(u)
	u, err = s.User().UpdateRoleAdmin(u.Email)
	assert.Equal(t, u.Isadmin, true)
	assert.NotNil(t, u)
	assert.NoError(t, err)
}

func TestUserRepository_UpdateRoleManager(t *testing.T) {
	var err error
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")
	s := sqlstore.New(db)
	u := model.TestUser(t)
	u.Isadmin = true
	s.User().Create(u)
	u, err = s.User().UpdateRoleManager(u.Email)
	assert.Equal(t, u.Isadmin, false)
	assert.NotNil(t, u)
	assert.NoError(t, err)
}

func TestUserRepository_ChangePassword(t *testing.T) {
	var err error
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")
	s := sqlstore.New(db)
	u := model.TestUser(t)
	u.Password = "123"
	s.User().Create(u)
	s.User().ChangePassword(u)
	//u, err = s.User().FindByEmail(u.Email)
	assert.NoError(t, err)
	//assert.Equal(t, u.ComparePassword(u.Password), true)
	assert.NotNil(t, u)
}

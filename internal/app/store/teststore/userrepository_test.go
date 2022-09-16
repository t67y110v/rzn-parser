package teststore_test

import (
	"restApi/internal/app/model"
	"restApi/internal/app/store"
	"restApi/internal/app/store/teststore"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s := teststore.New()
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
	//email := "user@test.com"
	s := teststore.New()
	u := model.TestUser(t)
	u.Isadmin = false
	var err error
	u, err = s.User().UpdateRoleAdmin(u.Email)
	assert.NoError(t, err)
	assert.Equal(t, u.Isadmin, true)

}

func TestUserRepository_UpdateRoleManager(t *testing.T) {
	//email := "user@test.com"
	s := teststore.New()
	u := model.TestUser(t)
	u.Isadmin = true
	var err error
	u, err = s.User().UpdateRoleManager(u.Email)
	assert.NoError(t, err)
	assert.Equal(t, u.Isadmin, false)

}

func TestUserRepository_ChangePassword(t *testing.T) {
	var err error
	s := teststore.New()
	u := model.TestUser(t)
	u.Password = "123"
	s.User().Create(u)
	s.User().ChangePassword(u)
	//u, err = s.User().FindByEmail(u.Email)
	assert.NoError(t, err)
	//assert.Equal(t, u.ComparePassword(u.Password), true)
	assert.NotNil(t, u)

}

func TestUserRepository_DepartmentCondition(t *testing.T) {
	var err error
	s := teststore.New()
	u := model.TestUser(t)
	u.Password = "123"
	s.User().Create(u)
	u, err = s.User().DepartmentCondition(u.Email)
	//u, err = s.User().FindByEmail(u.Email)
	assert.NoError(t, err)
	//assert.Equal(t, u.ComparePassword(u.Password), true)
	assert.NotNil(t, u)
}

func TestUserReposytory_DepartmentUpdate(t *testing.T) {
	var err error
	s := teststore.New()
	u := model.TestUser(t)
	s.User().Create(u)
	u.Department.EducationDepartment = false
	u.Department.DbDepartment = true
	u, err = s.User().DepartmentUpdate(u.Email, u.Name, u.SeccondName, true, true, true, false, false, false, false, false)
	assert.NoError(t, err)
	assert.Equal(t, u.Department.EducationDepartment, true)
	assert.Equal(t, u.Department.DbDepartment, false)
}

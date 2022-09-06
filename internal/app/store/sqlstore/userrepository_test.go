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

func TestUserReposetiryDelete(t *testing.T) {
	var err error
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown("users")
	s := sqlstore.New(db)
	u := model.TestUser(t)
	u.Email = "test21@mail.com"
	s.User().Create(u)
	err = s.User().Delete(u.Email)
	assert.NoError(t, err)

}

func TestUserRepositoryDepartmentCondition(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown("users")
	s := sqlstore.New(db)
	u := model.TestUser(t)
	u.Email = "test21@mail.com"
	u.Name = "oleg"
	u.SeccondName = "test"
	u.EducationDepartment = true
	u.SourceTrackingDepartment = false
	u.PeriodicReportingDepartment = false
	u.InternationalDepartment = false
	u.DocumentationDepartment = false
	u.NrDepartment = false
	u.DbDepartment = false
	u.Isadmin = true
	s.User().Create(u)
	s.User().DepartmentCondition(u.Email)
	assert.Equal(t, u.Name, "oleg")
	assert.Equal(t, u.SeccondName, "test")
	assert.Equal(t, u.EducationDepartment, true)
	assert.Equal(t, u.SourceTrackingDepartment, false)
	assert.NotEqual(t, u.PeriodicReportingDepartment, true)
	assert.Equal(t, u.InternationalDepartment, false)
	assert.Equal(t, u.DocumentationDepartment, false)
	assert.Equal(t, u.NrDepartment, false)
	assert.Equal(t, u.NrDepartment, false)

}

func TestUserRepositoryDepartmentUpdate(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown("users")
	s := sqlstore.New(db)
	u := model.TestUser(t)
	u.Email = "test21@mail.com"
	u.Name = "oleg"
	u.SeccondName = "test"
	u.EducationDepartment = true
	u.SourceTrackingDepartment = false
	u.PeriodicReportingDepartment = false
	u.InternationalDepartment = false
	u.DocumentationDepartment = false
	u.NrDepartment = false
	u.DbDepartment = false
	u.Isadmin = true
	s.User().Create(u)
	_, err := s.User().DepartmentUpdate(u.Email, false, false, false, false, false, false, false, false)
	assert.NoError(t, err)
}

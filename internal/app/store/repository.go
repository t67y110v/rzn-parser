package store

import (
	"restApi/internal/app/model"
	dp "restApi/internal/app/model/departments"
)

type UserRepository interface {
	Create(*model.User) error
	Delete(string) error
	FindByEmail(string) (*model.User, error)
	ChangePassword(*model.User) error
	DepartmentCondition(string) (*model.User, error)
	DepartmentUpdate(string, string, string, bool, bool, bool, bool, bool, bool, bool, bool, string, bool, int) (*model.User, error)
}

type DepartmentRepository interface {
	NrDepartmentAddNewPosition(*dp.NrDepartment) error
}

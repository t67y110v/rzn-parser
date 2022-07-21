package store

import "restApi/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
	UpdateRoleAdmin(string) (*model.User, error)
	UpdateRoleManager(string) (*model.User, error)
	ChangePassword(*model.User) error
}

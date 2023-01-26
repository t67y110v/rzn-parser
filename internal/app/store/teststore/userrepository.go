package teststore

import (
	"restApi/internal/app/model"
	"restApi/internal/app/store"
)

type UserRepository struct {
	store *Store
	users map[string]*model.User
}

func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}
	if err := u.BeforeCreate(); err != nil {
		return err
	}

	r.users[u.Email] = u

	u.ID = len(r.users)
	return nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u, ok := r.users[email]
	if !ok {
		return nil, store.ErrRecordNotFound
	}
	return u, nil
}

func (r *UserRepository) ChangePassword(u *model.User) error {
	if err := u.BeforeCreate(); err != nil {
		return err
	}
	r.users[u.Email] = u
	return nil
}

func (r *UserRepository) DepartmentCondition(string) (*model.User, error) {
	u := &model.User{}

	r.users[u.Email] = u
	return u, nil
}

func (r *UserRepository) DepartmentUpdate(string, string, string, bool, bool, bool, bool, bool, bool, bool, bool, string, bool, int) (*model.User, error) {
	u := &model.User{}

	if err := u.BeforeCreate(); err != nil {
		return u, err
	}
	r.users[u.Email] = u
	return u, nil
}
func (r *UserRepository) Delete(email string) error {
	u := &model.User{}

	if err := u.BeforeCreate(); err != nil {
		return err
	}
	r.users[u.Email] = u
	return nil
}

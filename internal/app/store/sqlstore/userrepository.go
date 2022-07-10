package sqlstore

import (
	"database/sql"
	"restApi/internal/app/model"
	"restApi/internal/app/store"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}
	if err := u.BeforeCreate(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID)
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password, isadmin FROM users WHERE email = $1",
		email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
		&u.Isadmin,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}
	return u, nil

}
func (r *UserRepository) UpdateRoleAdmin(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"UPDATE users SET isadmin = true WHERE email = $1 RETURNING id,isadmin",
		email,
	).Scan(
		&u.ID,
		&u.Isadmin,
	); err != nil {
		return nil, err
	}
	u.Email = email
	return u, nil
}

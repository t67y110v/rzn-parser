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
func (r *UserRepository) UpdateRoleManager(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"UPDATE users SET isadmin = false WHERE email = $1 RETURNING id,isadmin",
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

func (r *UserRepository) ChangePassword(u *model.User) error {
	if err := u.BeforeCreate(); err != nil {
		return err
	}
	return r.store.db.QueryRow(
		"UPDATE users SET encrypted_password = $1 WHERE email = $2 RETURNING ID",
		u.EncryptedPassword,
		u.Email,
	).Scan(&u.ID)
}

func (r *UserRepository) UpdateEducationDepartment(u *model.User, educationDepartment bool) error {
	return r.store.db.QueryRow(
		"UPDATE users SET educationDepartment = $1 WHERE email = $2 RETURNING ID",
		educationDepartment,
		u.Email,
	).Scan(&u.ID)
}

func (r *UserRepository) DepartmentCondition(email string) (*model.User, error) {

	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT educationDepartment, sourceTrackingDepartment,periodicReportingDepartment, internationalDepartment ,documentationDepartment, nrDepartment, dbDepartment FROM users WHERE email = $1",
		email,
	).Scan(
		&u.EducationDepartment,
		&u.SourceTrackingDepartment,
		&u.PeriodicReportingDepartment,
		&u.InternationalDepartment,
		&u.DocumentationDepartment,
		&u.NrDepartment,
		&u.DbDepartment,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}
	return u, nil

}

func (r *UserRepository) DepartmentUpdate(email string, educationDepartment bool, sourceTrackingDepartment bool, periodicReportingDepartment bool, internationalDepartment bool, documentationDepartment bool, nrDepartment bool, dbDepartment bool) (*model.User, error) {

	u := &model.User{}
	if err := r.store.db.QueryRow(
		"UPDATE users SET educationDepartment = $1, sourceTrackingDepartment = $2, periodicReportingDepartment = $3, internationalDepartment = $4, documentationDepartment = $5, nrDepartment = $6, dbDepartment = $7 WHERE email = $8 RETURNING  educationDepartment,sourceTrackingDepartment,periodicReportingDepartment,internationalDepartment,documentationDepartment,nrDepartment,dbDepartment ",
		educationDepartment,
		sourceTrackingDepartment,
		periodicReportingDepartment,
		internationalDepartment,
		documentationDepartment,
		nrDepartment,
		dbDepartment,
		email,
	).Scan(
		&u.EducationDepartment,
		&u.SourceTrackingDepartment,
		&u.PeriodicReportingDepartment,
		&u.InternationalDepartment,
		&u.DocumentationDepartment,
		&u.NrDepartment,
		&u.DbDepartment,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}
	return u, nil

}

//"$2a$04$Wfw9nXhI.4cM40JXvhw7CePo6BbrXaF8dTwRCWtDiHGYfbfIipDEa"
/*

 educationDepartment boolean DEFAULT false,
  sourceTrackingDepartment boolean DEFAULT false,
  periodicReportingDepartment boolean DEFAULT false,
  internationalDepartment boolean DEFAULT false,
  documentationDepartment boolean DEFAULT false,
  nrDepartment boolean DEFAULT false,
  dbDepartment boolean DEFAULT false
*/

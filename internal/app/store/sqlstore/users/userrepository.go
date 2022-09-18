package sqlstore

import (
	"database/sql"
	"restApi/internal/app/model"
	"restApi/internal/app/store"
)

type UserRepository struct {
	store *Store
}

//Созданик пользователя и занесение его в БД

func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}
	if err := u.BeforeCreate(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"INSERT INTO users (email, encrypted_password, userName, seccondName) VALUES ($1, $2,$3,$4) RETURNING id",
		u.Email,
		u.EncryptedPassword,
		u.Name,
		u.SeccondName,
	).Scan(&u.ID)
}

//Поиск пользователя в базе по почте

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password, isadmin, username, seccondname, educationDepartment, sourceTrackingDepartment,periodicReportingDepartment, internationalDepartment ,documentationDepartment, nrDepartment, dbDepartment FROM users WHERE email = $1",
		email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
		&u.Isadmin,
		&u.Name,
		&u.SeccondName,
		&u.Department.EducationDepartment,
		&u.Department.SourceTrackingDepartment,
		&u.Department.PeriodicReportingDepartment,
		&u.Department.InternationalDepartment,
		&u.Department.DocumentationDepartment,
		&u.Department.NrDepartment,
		&u.Department.DbDepartment,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}
	return u, nil

}

//обновление роли пользователя по почте

func (r *UserRepository) UpdateRoleAdmin(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"UPDATE users SET isadmin = true, educationDepartment = true, sourceTrackingDepartment = true, periodicReportingDepartment = true, internationalDepartment = true, documentationDepartment = true, nrDepartment = true, dbDepartment = true WHERE email = $1 RETURNING id,isadmin",
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
		"UPDATE users SET isadmin = false, educationDepartment = false, sourceTrackingDepartment = false, periodicReportingDepartment = false, internationalDepartment = false, documentationDepartment = false, nrDepartment = false, dbDepartment = false WHERE email = $1 RETURNING id,isadmin",
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

//Смена пароля пользователя

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

//Обновление конкретного отедла у пользвоаотеля

func (r *UserRepository) UpdateEducationDepartment(u *model.User, educationDepartment bool) error {
	return r.store.db.QueryRow(
		"UPDATE users SET educationDepartment = $1 WHERE email = $2 RETURNING ID",
		educationDepartment,
		u.Email,
	).Scan(&u.ID)
}

//Выдает информацию по пренадлежности пользователя к отделам булевые переменные

func (r *UserRepository) DepartmentCondition(email string) (*model.User, error) {

	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT educationDepartment, sourceTrackingDepartment,periodicReportingDepartment, internationalDepartment ,documentationDepartment, nrDepartment, dbDepartment FROM users WHERE email = $1",
		email,
	).Scan(
		&u.Department.EducationDepartment,
		&u.Department.SourceTrackingDepartment,
		&u.Department.PeriodicReportingDepartment,
		&u.Department.InternationalDepartment,
		&u.Department.DocumentationDepartment,
		&u.Department.NrDepartment,
		&u.Department.DbDepartment,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	/*if !(u.EducationDepartment && u.SourceTrackingDepartment && u.PeriodicReportingDepartment && u.InternationalDepartment && u.DocumentationDepartment && u.NrDepartment && u.DbDepartment) {
		return nil, store.ErrEmptyValue
	} */
	return u, nil

}

//Обновляет  данные об отелах пользователя чья почта передана первым аргументом

func (r *UserRepository) DepartmentUpdate(email string, name string, seccondname string, educationDepartment bool, sourceTrackingDepartment bool, periodicReportingDepartment bool, internationalDepartment bool, documentationDepartment bool, nrDepartment bool, dbDepartment bool, isadmin bool) (*model.User, error) {

	u := &model.User{}
	if err := r.store.db.QueryRow(
		"UPDATE users SET educationDepartment = $1, sourceTrackingDepartment = $2, periodicReportingDepartment = $3, internationalDepartment = $4, documentationDepartment = $5, nrDepartment = $6, dbDepartment = $7, isadmin = $9, username = $10 , seccondname = $11  WHERE email = $8 RETURNING  isadmin,educationDepartment,sourceTrackingDepartment,periodicReportingDepartment,internationalDepartment,documentationDepartment,nrDepartment,dbDepartment ",
		educationDepartment,
		sourceTrackingDepartment,
		periodicReportingDepartment,
		internationalDepartment,
		documentationDepartment,
		nrDepartment,
		dbDepartment,
		email,
		isadmin,
		name,
		seccondname,
	).Scan(
		&u.Department.EducationDepartment,
		&u.Department.SourceTrackingDepartment,
		&u.Department.PeriodicReportingDepartment,
		&u.Department.InternationalDepartment,
		&u.Department.DocumentationDepartment,
		&u.Department.NrDepartment,
		&u.Department.DbDepartment,
		&u.Isadmin,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}
	return u, nil

}

//Удалеие пользователя по почте

func (r *UserRepository) Delete(email string) error {

	return r.store.db.QueryRow(
		"DELETE FROM users  WHERE email = $1",
		email,
	).Err()
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

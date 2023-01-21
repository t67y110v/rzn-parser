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
		"INSERT INTO users (email, encrypted_password, user_name, seccond_name) VALUES ($1, $2,$3,$4) RETURNING id",
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
		"SELECT id, email, encrypted_password, user_role, user_name, seccond_name,client_department,  education_department, source_tracking_department,periodic_reporting_department, international_department ,documentation_department, nr_department, db_department, monitoring_specialist,monitoring_responsible FROM users WHERE email = $1",
		email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
		&u.Role,
		&u.Name,
		&u.SeccondName,
		&u.Department.ClientDepartment,
		&u.Department.EducationDepartment,
		&u.Department.SourceTrackingDepartment,
		&u.Department.PeriodicReportingDepartment,
		&u.Department.InternationalDepartment,
		&u.Department.DocumentationDepartment,
		&u.Department.NrDepartment,
		&u.Department.DbDepartment,
		&u.MonitoringSpecialist,
		&u.MonitoringResponsible,
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
		"UPDATE users SET user_role = 'admin', education_department = true, source_tracking_department = true, periodic_reporting_department = true, international_department = true, documentation_department = true, nr_department = true, db_department = true WHERE email = $1 RETURNING id,user_role",
		email,
	).Scan(
		&u.ID,
		&u.Role,
	); err != nil {
		return nil, err
	}

	u.Email = email
	return u, nil
}
func (r *UserRepository) UpdateRoleManager(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"UPDATE users SET user_role = 'manager', education_department = false, source_tracking_department = false, periodic_reporting_department = false, international_department = false, documentation_department = false, nr_department = false, db_department = false WHERE email = $1 RETURNING id,role",
		email,
	).Scan(
		&u.ID,
		&u.Role,
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
		"UPDATE users SET education_department = $1 WHERE email = $2 RETURNING ID",
		educationDepartment,
		u.Email,
	).Scan(&u.ID)
}

//Выдает информацию по пренадлежности пользователя к отделам булевые переменные

func (r *UserRepository) DepartmentCondition(email string) (*model.User, error) {

	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT education_department, source_tracking_department,periodic_reporting_department, international_department ,documentation_department, nr_department, db_department, monitoring_specialist,monitoring_responsible FROM users WHERE email = $1",
		email,
	).Scan(
		&u.Department.EducationDepartment,
		&u.Department.SourceTrackingDepartment,
		&u.Department.PeriodicReportingDepartment,
		&u.Department.InternationalDepartment,
		&u.Department.DocumentationDepartment,
		&u.Department.NrDepartment,
		&u.Department.DbDepartment,
		&u.MonitoringSpecialist,
		&u.MonitoringResponsible,
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

func (r *UserRepository) DepartmentUpdate(
	email string,
	name string,
	seccondname string,
	clientDepartment bool,
	educationDepartment bool,
	sourceTrackingDepartment bool,
	periodicReportingDepartment bool,
	internationalDepartment bool,
	documentationDepartment bool,
	nrDepartment bool,
	dbDepartment bool,
	role string,
	monitoringSpecialist bool,
	monitoringResponsible int,
) (*model.User, error) {

	u := &model.User{}
	if err := r.store.db.QueryRow(
		"UPDATE users SET  client_department = $1,  education_department = $2, source_tracking_department = $3, periodic_reporting_department = $4, international_department = $5, documentation_department = $6, nr_department = $7, db_department = $8, user_role = $9, user_name = $10 , seccond_name = $11,monitoring_specialist = $12,monitoring_responsible = $13  WHERE email = $14 RETURNING  user_role,education_department,source_tracking_department,periodic_reporting_department,international_department,documentation_department,nr_department,db_department,monitoring_specialist,monitoring_responsible, client_department",
		clientDepartment,
		educationDepartment,
		sourceTrackingDepartment,
		periodicReportingDepartment,
		internationalDepartment,
		documentationDepartment,
		nrDepartment,
		dbDepartment,
		role,
		name,
		seccondname,
		monitoringSpecialist,
		monitoringResponsible,
		email,
	).Scan(
		&u.Role,
		&u.Department.EducationDepartment,
		&u.Department.SourceTrackingDepartment,
		&u.Department.PeriodicReportingDepartment,
		&u.Department.InternationalDepartment,
		&u.Department.DocumentationDepartment,
		&u.Department.NrDepartment,
		&u.Department.DbDepartment,
		&u.MonitoringSpecialist,
		&u.MonitoringResponsible,
		&u.Department.ClientDepartment,
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

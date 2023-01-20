package model

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	//"github.com/go-ozzo/ozzo-validation/v4/is"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                    int    `json:"id"`
	Email                 string `json:"email"`
	Name                  string `json:"name"`
	SeccondName           string `json:"seccond_name"`
	Password              string `json:"password,omitempty"`
	Role                  string `json:"role"`
	EncryptedPassword     string `json:"-"`
	MonitoringSpecialist  bool   `json:"monitoring_specialist"`
	MonitoringResponsible int    `json:"monitoring_responsible"`
	Department            struct {
		EducationDepartment         bool `json:"education_department"`
		SourceTrackingDepartment    bool `json:"source_tracking_department"`
		PeriodicReportingDepartment bool `json:"periodic_reporting_department"`
		InternationalDepartment     bool `json:"international_department"`
		DocumentationDepartment     bool `json:"documentation_department"`
		NrDepartment                bool `json:"nr_department"`
		DbDepartment                bool `json:"db_department"`
	}
}

/*
sourceTrackingDepartment boolean DEFAULT false,

	periodicReportingDepartment boolean DEFAULT false,
	internationalDepartment boolean DEFAULT false,
	documentationDepartment boolean DEFAULT false,
	nrDepartment boolean DEFAULT false,
	dbDepartment boolean DEFAULT false
*/
func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		//validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.By(requiredIf(u.EncryptedPassword == "")), validation.Length(6, 100)),
	)
}

func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password)
		if err != nil {
			return err
		}

		u.EncryptedPassword = enc
	}

	return nil
}

func (u *User) Sanitize() {
	u.Password = ""
}

func (u *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password)) == nil
}

func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

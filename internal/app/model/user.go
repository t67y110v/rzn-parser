package model

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                          int    `json:"id"`
	Email                       string `json:"email"`
	Password                    string `json:"password,omitempty"`
	Isadmin                     bool   `json:"isadmin"`
	EncryptedPassword           string `json:"-"`
	EducationDepartment         bool   `json:"educationDepartment"`
	SourceTrackingDepartment    bool   `json:"sourceTrackingDepartment"`
	PeriodicReportingDepartment bool   `json:"periodicReportingDepartment"`
	InternationalDepartment     bool   `json:"internationalDepartment"`
	DocumentationDepartment     bool   `json:"documentationDepartment"`
	NrDepartment                bool   `json:"nrDepartment"`
	DbDepartment                bool   `json:"dbDepartment"`
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
		validation.Field(&u.Email, validation.Required, is.Email),
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

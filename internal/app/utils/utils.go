package utils

import "restApi/internal/app/model"

func CheckThatUserIsAdmin(u *model.User) bool {
	return u.Department.ClientDepartment && u.Department.EducationDepartment && u.Department.SourceTrackingDepartment && u.Department.PeriodicReportingDepartment && u.Department.InternationalDepartment && u.Department.DocumentationDepartment && u.Department.NrDepartment && u.Department.DbDepartment
}

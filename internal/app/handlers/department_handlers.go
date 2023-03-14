package handlers

import (
	"restApi/internal/app/handlers/responses"

	"github.com/gofiber/fiber/v2"
)

// @Summary Department condition
// @Description getting current department condition
// @Tags         Department
//
//	@Accept       json
//
// @Produce json
// @Param        email   path      string  true  "Email"
// @Success 200 {object} responses.DepartmentRes
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /department/condition/{email} [get]
func (h *Handlers) HandleDepartmentCondition() fiber.Handler {

	return func(c *fiber.Ctx) error {
		email := c.Params("email")
		u, err := h.store.User().DepartmentCondition(email)
		if err != nil {
			return c.JSON(fiber.Map{
				"message": err,
			})
		}

		res := &responses.DepartmentRes{}
		res.Departments.ClientDepartment = u.Department.ClientDepartment
		res.Departments.EducationDepartment = u.Department.EducationDepartment
		res.Departments.SourceTrackingDepartment = u.Department.SourceTrackingDepartment
		res.Departments.PeriodicReportingDepartment = u.Department.PeriodicReportingDepartment
		res.Departments.InternationalDepartment = u.Department.InternationalDepartment
		res.Departments.DocumentationDepartment = u.Department.DocumentationDepartment
		res.Departments.NrDepartment = u.Department.NrDepartment
		res.Departments.DbDepartment = u.Department.DbDepartment
		res.MonitoringSpecialist = u.MonitoringSpecialist
		res.MonitoringResponsible = u.MonitoringResponsible
		return c.JSON(res)
	}
}

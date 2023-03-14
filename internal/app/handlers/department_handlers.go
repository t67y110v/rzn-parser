package handlers

import (
	"bytes"
	"encoding/json"
	"restApi/internal/app/model"

	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) HandleDepartmentCondition() fiber.Handler {
	type request struct {
		Email string `json:"email"`
	}
	return func(c *fiber.Ctx) error {
		req := &request{}
		reader := bytes.NewReader(c.Body())

		if err := json.NewDecoder(reader).Decode(req); err != nil {
			h.logger.Warningf("handle register, status :%d, error :%e", fiber.StatusBadRequest, err)
		}
		u, err := h.store.User().DepartmentCondition(req.Email)
		if err != nil {
			return c.JSON(fiber.Map{
				"message": err,
			})
		}
		type resp struct {
			Departments           model.Department
			MonitoringSpecialist  bool `json:"monitoring_specialist"`
			MonitoringResponsible int  `json:"monitoring_responsible"`
		}
		res := &resp{}
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

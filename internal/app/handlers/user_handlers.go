package handlers

import (
	"bytes"
	"encoding/json"
	"restApi/internal/app/model"

	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) HandleUsersCreate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		type request struct {
			Email                 string           `json:"email"`
			Role                  string           `json:"user_role"`
			Password              string           `json:"password"`
			Name                  string           `json:"name"`
			SeccondName           string           `json:"seccond_name"`
			Departments           model.Department `json:"departments"`
			MonitoringSpecialist  bool             `json:"monitoring_specialist"`
			MonitoringResponsible int              `json:"monitoring_responsible"`
		}
		req := &request{}
		reader := bytes.NewReader(c.Body())

		if err := json.NewDecoder(reader).Decode(req); err != nil {
			h.logger.Warningf("handle register, status :%d, error :%e", fiber.StatusBadRequest, err)
		}
		u := &model.User{
			Email:       req.Email,
			Password:    req.Password,
			Name:        req.Name,
			SeccondName: req.SeccondName,
		}
		if err := h.store.User().Create(u); err != nil {
			return c.JSON(fiber.Map{
				"message": err,
			})
		}
		u.Sanitize()
		_, err := h.store.User().DepartmentUpdate(
			req.Email,
			req.Name,
			req.SeccondName,
			req.Departments.ClientDepartment,
			req.Departments.EducationDepartment,
			req.Departments.SourceTrackingDepartment,
			req.Departments.PeriodicReportingDepartment,
			req.Departments.InternationalDepartment,
			req.Departments.DocumentationDepartment,
			req.Departments.NrDepartment,
			req.Departments.DbDepartment,
			req.Role,
			req.MonitoringSpecialist,
			req.MonitoringResponsible,
		)
		if err != nil {
			//fmt.Println("тут")

			return c.JSON(fiber.Map{
				"message": err,
			})
		}
		type resp struct {
			ID          int    `json:"id"`
			Email       string `json:"email"`
			Name        string `json:"name"`
			SeccondName string `json:"seccond_name"`
		}
		res := &resp{}
		res.ID = u.ID
		res.Email = u.Email
		res.Name = u.Name
		res.SeccondName = u.SeccondName
		return c.JSON(res)
	}
}

func (h *Handlers) HandleSessionsCreate() fiber.Handler {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	return func(c *fiber.Ctx) error {
		req := &request{}
		reader := bytes.NewReader(c.Body())

		if err := json.NewDecoder(reader).Decode(req); err != nil {
			h.logger.Warningf("handle register, status :%d, error :%e", fiber.StatusBadRequest, err)
		}
		u, err := h.store.User().FindByEmail(req.Email)
		if err != nil || !u.ComparePassword(req.Password) {
			return c.JSON(fiber.Map{
				"message": err,
			})
		}

		return c.JSON(u)
		//h.logger.Infof("handle /sessions, status :%d", http.StatusOK)
	}
}

func (h *Handlers) HandlePasswordChange() fiber.Handler {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	return func(c *fiber.Ctx) error {
		req := &request{}
		reader := bytes.NewReader(c.Body())

		if err := json.NewDecoder(reader).Decode(req); err != nil {
			h.logger.Warningf("handle register, status :%d, error :%e", fiber.StatusBadRequest, err)
		}
		u := &model.User{
			Email:    req.Email,
			Password: req.Password,
		}
		if err := h.store.User().ChangePassword(u); err != nil {
			return c.JSON(fiber.Map{
				"message": err,
			})
		}
		u.Sanitize()
		return c.JSON(u)
	}
}

func (h *Handlers) HandleUserUpdate() fiber.Handler {

	type request struct {
		Email                 string           `json:"email"`
		Name                  string           `json:"name"`
		SeccondName           string           `json:"seccond_name"`
		Role                  string           `json:"user_role"`
		Departments           model.Department `json:"departments"`
		MonitoringSpecialist  bool             `json:"monitoring_specialist"`
		MonitoringResponsible int              `json:"monitoring_responsible"`
	}
	return func(c *fiber.Ctx) error {

		req := &request{}

		reader := bytes.NewReader(c.Body())

		if err := json.NewDecoder(reader).Decode(req); err != nil {
			h.logger.Warningf("handle register, status :%d, error :%e", fiber.StatusBadRequest, err)
		}
		u, err := h.store.User().DepartmentUpdate(req.Email,
			req.Name,
			req.SeccondName,
			req.Departments.ClientDepartment,
			req.Departments.EducationDepartment,
			req.Departments.SourceTrackingDepartment,
			req.Departments.PeriodicReportingDepartment,
			req.Departments.InternationalDepartment,
			req.Departments.DocumentationDepartment,
			req.Departments.NrDepartment,
			req.Departments.DbDepartment,
			req.Role,
			req.MonitoringSpecialist,
			req.MonitoringResponsible)
		if err != nil {
			return c.JSON(fiber.Map{
				"message": err,
			})
		}
		type resp struct {
			Role                  string           `json:"user_role"`
			Departments           model.Department `json:"departments"`
			MonitoringSpecialist  bool             `json:"monitoring_specialist"`
			MonitoringResponsible int              `json:"monitoring_responsible"`
		}
		res := &resp{}
		res.Role = u.Role
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

func (h *Handlers) HandleUserDelete() fiber.Handler {
	type request struct {
		Email string `json:"email"`
	}
	return func(c *fiber.Ctx) error {
		req := &request{}

		reader := bytes.NewReader(c.Body())

		if err := json.NewDecoder(reader).Decode(req); err != nil {
			h.logger.Warningf("handle register, status :%d, error :%e", fiber.StatusBadRequest, err)
		}
		err := h.store.User().Delete(req.Email)
		if err != nil {
			return c.JSON(fiber.Map{
				"message": err,
			})
		}
		type resp struct {
			Result bool `json:"result"`
		}
		res := &resp{}
		if err == nil {
			res.Result = true
		} else {
			res.Result = false
		}
		return c.JSON(res)
		//h.logger.Infof("handle /userDelete, status :%d", http.StatusOK)
	}
}

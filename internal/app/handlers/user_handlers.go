package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"restApi/internal/app/handlers/requests"
	"restApi/internal/app/handlers/responses"
	"restApi/internal/app/model"

	"github.com/gofiber/fiber/v2"
)

// @Summary User Create
// @Description creation of user
// @Tags         User
//
//	@Accept       json
//
// @Produce json
// @Param  data body requests.CreateUser  true "create new user"
// @Success 200 {object} responses.CreateUser
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /user/create [post]
func (h *Handlers) HandleUsersCreate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := &requests.CreateUser{}
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

		res := &responses.CreateUser{}
		res.ID = u.ID
		res.Email = u.Email
		res.Name = u.Name
		res.SeccondName = u.SeccondName
		return c.JSON(res)
	}
}

// @Summary Session Create
// @Description creation new session
// @Tags         User
//
//	@Accept       json
//
// @Produce json
// @Param  data body requests.EmailPassword  true "create new session"
// @Success 200 {object} responses.CreateUser
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /user/session [post]
func (h *Handlers) HandleSessionsCreate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := &requests.EmailPassword{}
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

// @Summary Password Change
// @Description change users passwrod
// @Tags         User
//
//	@Accept       json
//
// @Produce json
// @Param  data body requests.EmailPassword  true "change users password"
// @Success 200 {object} responses.CreateUser
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /user/change/password [put]
func (h *Handlers) HandlePasswordChange() fiber.Handler {

	return func(c *fiber.Ctx) error {
		req := &requests.EmailPassword{}
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

// @Summary User Update
// @Description update users info
// @Tags         User
//
//	@Accept       json
//
// @Produce json
// @Param  data body requests.UpdateUser  true "update users information"
// @Success 200 {object} responses.UserUpdate
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /user/update [put]
func (h *Handlers) HandleUserUpdate() fiber.Handler {

	return func(c *fiber.Ctx) error {

		req := &requests.UpdateUser{}

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
				"message": fmt.Sprintf("%e", err),
			})
		}

		res := &responses.UserUpdate{}
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

// @Summary User Delete
// @Description delete user from system
// @Tags         User
//
//	@Accept       json
//
// @Produce json
// @Param  data body requests.Email true "delete user from system"
// @Success 200 {object} responses.Result
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /user/delete [delete]
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

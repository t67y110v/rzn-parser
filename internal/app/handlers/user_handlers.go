package handlers

import (
	"encoding/json"
	"net/http"
	"restApi/internal/app/model"
)

func (h *Handlers) HandleUsersCreate() http.HandlerFunc {
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
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			h.error(w, r, http.StatusBadRequest, err)
			h.logger.Warningf("handle /userCreate, status :%d, error :%e", http.StatusBadRequest, err)
			return
		}
		u := &model.User{
			Email:       req.Email,
			Password:    req.Password,
			Name:        req.Name,
			SeccondName: req.SeccondName,
		}
		if err := h.store.User().Create(u); err != nil {
			h.error(w, r, http.StatusUnprocessableEntity, err)
			h.logger.Warningf("handle /userCreate, status :%d, error :%e", http.StatusUnprocessableEntity, err)
			return
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
			h.error(w, r, http.StatusBadRequest, err)
			h.logger.Warningf("handle /userCreate, status :%d, error :%e", http.StatusBadRequest, err)
			return
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
		h.respond(w, r, http.StatusCreated, res)
		h.logger.Infof("handle /userCreate, status :%d", http.StatusCreated)
	}
}

func (h *Handlers) HandleSessionsCreate() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			h.error(w, r, http.StatusBadRequest, err)
			h.logger.Warningf("handle /sessions, status :%d, error :%e", http.StatusBadRequest, err)
			return
		}
		u, err := h.store.User().FindByEmail(req.Email)
		if err != nil || !u.ComparePassword(req.Password) {
			h.error(w, r, http.StatusUnauthorized, errorIncorrectEmailOrPassword)
			h.logger.Warningf("handle /sessions, status :%d, error :%e", http.StatusUnauthorized, err)
			return
		}

		h.respond(w, r, http.StatusOK, u)
		h.logger.Infof("handle /sessions, status :%d", http.StatusOK)
	}
}

func (h *Handlers) HandlePasswordChange() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			h.error(w, r, http.StatusBadRequest, err)
			h.logger.Warningf("handle /changePassword, status :%d, error :%e", http.StatusBadRequest, err)
			return
		}
		u := &model.User{
			Email:    req.Email,
			Password: req.Password,
		}
		if err := h.store.User().ChangePassword(u); err != nil {
			h.error(w, r, http.StatusUnprocessableEntity, err)
			h.logger.Warningf("handle /changePassword, status :%d, error :%e", http.StatusUnprocessableEntity, err)
			return
		}
		u.Sanitize()
		h.respond(w, r, http.StatusOK, u)
		h.logger.Infof("handle /changePassword, status :%d", http.StatusOK)
	}
}

func (h *Handlers) HandleUserUpdate() http.HandlerFunc {

	type request struct {
		Email                 string           `json:"email"`
		Name                  string           `json:"name"`
		SeccondName           string           `json:"seccond_name"`
		Role                  string           `json:"user_role"`
		Departments           model.Department `json:"departments"`
		MonitoringSpecialist  bool             `json:"monitoring_specialist"`
		MonitoringResponsible int              `json:"monitoring_responsible"`
	}
	return func(w http.ResponseWriter, r *http.Request) {

		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			h.error(w, r, http.StatusBadRequest, err)
			h.logger.Warningf("handle /userUpdate, status :%d, error :%e", http.StatusBadRequest, err)
			return
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
			h.error(w, r, http.StatusBadRequest, err)
			h.logger.Warningf("handle /userUpdate, status :%d, error :%e", http.StatusBadRequest, err)
			return
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

		h.respond(w, r, http.StatusOK, res)
		h.logger.Infof("handle /userUpdate, status :%d", http.StatusOK)
	}

}

func (h *Handlers) HandleUserDelete() http.HandlerFunc {
	type request struct {
		Email string `json:"email"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			h.error(w, r, http.StatusBadRequest, err)
			h.logger.Warningf("handle /userDelete, status :%d, error :%e", http.StatusBadRequest, err)
			return
		}
		err := h.store.User().Delete(req.Email)
		if err != nil {
			h.error(w, r, http.StatusUnauthorized, errorIncorrectEmailOrPassword)
			h.logger.Warningf("handle /userDelete, status :%d, error :%e", http.StatusBadRequest, err)
			return
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
		h.respond(w, r, http.StatusOK, res)
		h.logger.Infof("handle /userDelete, status :%d", http.StatusOK)
	}
}

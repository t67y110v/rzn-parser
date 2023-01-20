package apiserver

import (
	"encoding/json"
	"net/http"
	"restApi/internal/app/model"
)

func (s *Server) handleUsersCreate() http.HandlerFunc {
	type request struct {
		Email                 string           `json:"email"`
		Role                  string           `json:"role"`
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
			s.error(w, r, http.StatusBadRequest, err)
			s.logger.Warningf("handle /userCreate, status :%d, error :%e", http.StatusBadRequest, err)
			return
		}
		u := &model.User{
			Email:       req.Email,
			Password:    req.Password,
			Name:        req.Name,
			SeccondName: req.SeccondName,
		}
		if err := s.store.User().Create(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			s.logger.Warningf("handle /userCreate, status :%d, error :%e", http.StatusUnprocessableEntity, err)
			return
		}
		u.Sanitize()
		_, err := s.store.User().DepartmentUpdate(
			req.Email,
			req.Name, req.SeccondName,
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
			s.error(w, r, http.StatusBadRequest, err)
			s.logger.Warningf("handle /userCreate, status :%d, error :%e", http.StatusBadRequest, err)
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
		s.respond(w, r, http.StatusCreated, res)
		s.logger.Infof("handle /userCreate, status :%d", http.StatusCreated)
	}
}

func (s *Server) handleSessionsCreate() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			s.logger.Warningf("handle /sessions, status :%d, error :%e", http.StatusBadRequest, err)
			return
		}
		u, err := s.store.User().FindByEmail(req.Email)
		if err != nil || !u.ComparePassword(req.Password) {
			s.error(w, r, http.StatusUnauthorized, errorIncorrectEmailOrPassword)
			s.logger.Warningf("handle /sessions, status :%d, error :%e", http.StatusUnauthorized, err)
			return
		}

		s.respond(w, r, http.StatusOK, u)
		s.logger.Infof("handle /sessions, status :%d", http.StatusOK)
	}
}

func (s *Server) handleAdminUpdate() http.HandlerFunc {
	type request struct {
		Email string `json:"email"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			s.logger.Warningf("handle /makeAdmin, status :%d, error :%e", http.StatusBadRequest, err)
			return
		}
		u, err := s.store.User().UpdateRoleAdmin(req.Email)
		if err != nil {
			s.error(w, r, http.StatusUnauthorized, errorIncorrectEmailOrPassword)
			s.logger.Warningf("handle /makeAdmin, status :%d, error :%e", http.StatusUnauthorized, errorIncorrectEmailOrPassword)
			return
		}
		type resp struct {
			Role string `json:"role"`
		}
		res := &resp{}
		if u.Role == "manager" {
			res.Role = "manager"
		} else if u.Role == "admin" {
			res.Role = "admin"
		} else if u.Role == "reviewer" {
			res.Role = "reviewer"
		} else if u.Role == "writer" {
			res.Role = "writer"
		}
		s.respond(w, r, http.StatusOK, res)
		s.logger.Infof("handle /makeAdmin, status :%d", http.StatusOK)
	}
}

func (s *Server) handleManagerUpdate() http.HandlerFunc {
	type request struct {
		Email string `json:"email"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			s.logger.Warningf("handle /makeManager, status :%d, error :%e", http.StatusBadRequest, err)
			return
		}
		u, err := s.store.User().UpdateRoleManager(req.Email)
		if err != nil {
			s.error(w, r, http.StatusUnauthorized, errorIncorrectEmailOrPassword)
			s.logger.Warningf("handle /makeManager, status :%d, error :%e", http.StatusUnauthorized, errorIncorrectEmailOrPassword)
			return
		}
		type resp struct {
			Role string `json:"role"`
		}
		res := &resp{}
		if u.Role == "admin" {
			res.Role = "admin"
		} else if u.Role == "manager" {
			res.Role = "manager"
		} else if u.Role == "reviewer" {
			res.Role = "reviewer"
		} else if u.Role == "writer" {
			res.Role = "writer"
		}
		s.respond(w, r, http.StatusOK, res)
		s.logger.Infof("handle /makeManager, status :%d", http.StatusOK)
	}
}

func (s *Server) handlePasswordChange() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			s.logger.Warningf("handle /changePassword, status :%d, error :%e", http.StatusBadRequest, err)
			return
		}
		u := &model.User{
			Email:    req.Email,
			Password: req.Password,
		}
		if err := s.store.User().ChangePassword(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			s.logger.Warningf("handle /changePassword, status :%d, error :%e", http.StatusUnprocessableEntity, err)
			return
		}
		u.Sanitize()
		s.respond(w, r, http.StatusOK, u)
		s.logger.Infof("handle /changePassword, status :%d", http.StatusOK)
	}
}

func (s *Server) handleUserUpdate() http.HandlerFunc {

	type request struct {
		Email                 string           `json:"email"`
		Name                  string           `json:"name"`
		SeccondName           string           `json:"seccond_name"`
		Role                  string           `json:"role"`
		Departments           model.Department `json:"departments"`
		MonitoringSpecialist  bool             `json:"monitoring_specialist"`
		MonitoringResponsible int              `json:"monitoring_responsible"`
	}
	return func(w http.ResponseWriter, r *http.Request) {

		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			s.logger.Warningf("handle /userUpdate, status :%d, error :%e", http.StatusBadRequest, err)
			return
		}
		u, err := s.store.User().DepartmentUpdate(req.Email,
			req.Name, req.SeccondName,
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
			s.error(w, r, http.StatusBadRequest, err)
			s.logger.Warningf("handle /userUpdate, status :%d, error :%e", http.StatusBadRequest, err)
			return
		}
		type resp struct {
			Role                  string           `json:"role"`
			Departments           model.Department `json:"departments"`
			MonitoringSpecialist  bool             `json:"monitoring_specialist"`
			MonitoringResponsible int              `json:"monitoring_responsible"`
		}
		res := &resp{}
		res.Role = u.Role
		res.Departments.EducationDepartment = u.Department.EducationDepartment
		res.Departments.SourceTrackingDepartment = u.Department.SourceTrackingDepartment
		res.Departments.PeriodicReportingDepartment = u.Department.PeriodicReportingDepartment
		res.Departments.InternationalDepartment = u.Department.InternationalDepartment
		res.Departments.DocumentationDepartment = u.Department.DocumentationDepartment
		res.Departments.NrDepartment = u.Department.NrDepartment
		res.Departments.DbDepartment = u.Department.DbDepartment
		res.MonitoringSpecialist = u.MonitoringSpecialist
		res.MonitoringResponsible = u.MonitoringResponsible

		s.respond(w, r, http.StatusOK, res)
		s.logger.Infof("handle /userUpdate, status :%d", http.StatusOK)
	}

}

func (s *Server) handleUserDelete() http.HandlerFunc {
	type request struct {
		Email string `json:"email"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			s.logger.Warningf("handle /userDelete, status :%d, error :%e", http.StatusBadRequest, err)
			return
		}
		err := s.store.User().Delete(req.Email)
		if err != nil {
			s.error(w, r, http.StatusUnauthorized, errorIncorrectEmailOrPassword)
			s.logger.Warningf("handle /userDelete, status :%d, error :%e", http.StatusBadRequest, err)
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
		s.respond(w, r, http.StatusOK, res)
		s.logger.Infof("handle /userDelete, status :%d", http.StatusOK)
	}
}

package handlers

import (
	"encoding/json"
	"net/http"
	"restApi/internal/app/model"
)

func (h *Handlers) HandleDepartmentCondition() http.HandlerFunc {
	type request struct {
		Email string `json:"email"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			h.error(w, r, http.StatusBadRequest, err)
			h.logger.Warningf("handle /departmentCondition, status :%d, error :%e", http.StatusBadRequest, err)
			return
		}
		u, err := h.store.User().DepartmentCondition(req.Email)
		if err != nil {
			h.logger.Warningf("handle /departmentCondition, status :%d, error :%e", http.StatusBadRequest, err)
			return
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
		h.respond(w, r, http.StatusOK, res)
		h.logger.Infof("handle /departmentCondition, status :%d", http.StatusOK)
	}
}

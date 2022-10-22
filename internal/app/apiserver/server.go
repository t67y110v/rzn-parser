package apiserver

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"restApi/internal/app/logging"
	mail "restApi/internal/app/mailservice"
	"restApi/internal/app/model"
	"restApi/internal/app/store"

	"github.com/gorilla/mux"
	//"github.com/sirupsen/logrus"
)

var (
	errorIncorrectEmailOrPassword = errors.New("incorrect email or password")
)

type server struct {
	router *mux.Router
	logger logging.Logger
	store  store.UserStore
	config *Config
}

func newServer(store store.UserStore, config *Config) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logging.GetLogger(),
		store:  store,
		config: config,
	}

	s.configureRouter()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	//s.router.Host("{subdomain:[a-z]+}.example.com")
	s.router.HandleFunc("/userCreate", s.handleUsersCreate()).Methods("POST")                  //почта + пароль + имя + фамилия + булевые значения для каждого отдела -> статус:201 json {"id":27,"email":"test3@gmail.com","isadmin":false}
	s.router.HandleFunc("/userUpdate", s.handleUserUpdate()).Methods("PUT")                    //почта + булевые значения для каждого отдела  ->  статус:200 json {"isadmin":false,"educationDepartment":true,"sourceTrackingDepartment":true,"periodicReportingDepartment":false,"internationalDepartment":false,"documentationDepartment":false,"nrDepartment":false,"dbDepartment":true}
	s.router.HandleFunc("/userDelete", s.handleUserDelete()).Methods("DELETE")                 //почта  -> статус:200 json  {result : true}
	s.router.HandleFunc("/sessions", s.handleSessionsCreate()).Methods("POST")                 //почта + пароль -> статус:200 json {"isAdmin":"false"}
	s.router.HandleFunc("/makeAdmin", s.handleAdminUpdate()).Methods("PUT")                    //почта  -> статус:200 json {isAdmin:true}
	s.router.HandleFunc("/makeManager", s.handleManagerUpdate()).Methods("PUT")                //почта  -> статус:200 json {isAdmin:false}
	s.router.HandleFunc("/changePassword", s.handlePasswordChange()).Methods("PUT")            //почта + новый пароль -> статус:200 json {Модель пользователя с очищенным полем пароля}
	s.router.HandleFunc("/departmentCondition", s.handleDepartmentCondition()).Methods("POST") //почта  -> статус:200 json {"isadmin":false,"educationDepartment":true,"sourceTrackingDepartment":true,"periodicReportingDepartment":false,"internationalDepartment":false,"documentationDepartment":false,"nrDepartment":false,"dbDepartment":true}
	s.router.HandleFunc("/sendEmail", s.handleSendEmail()).Methods("POST")
}

func (s *server) handleUsersCreate() http.HandlerFunc {
	type request struct {
		Email                 string           `json:"email"`
		IsAdmin               bool             `json:"isadmin"`
		Password              string           `json:"password"`
		Name                  string           `json:"name"`
		SeccondName           string           `json:"seccondName"`
		Departments           model.Department `json:"Departments"`
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
			req.IsAdmin,
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
			SeccondName string `json:"seccondName"`
		}
		res := &resp{}
		res.ID = u.ID
		res.Email = u.Email
		res.Name = u.Name
		res.SeccondName = u.SeccondName
		log.Printf("")
		s.respond(w, r, http.StatusCreated, res)
		s.logger.Infof("handle /userCreate, status :%d", http.StatusCreated)
	}
}

func (s *server) handleSessionsCreate() http.HandlerFunc {
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
			s.logger.Warningf("handle /users, status :%d, error :%e", http.StatusUnauthorized, errorIncorrectEmailOrPassword)
			return
		}
		/*type resp struct {
			IsAdmin bool `json:"isAdmin"`
		}
		res := &resp{}
		if u.Isadmin {
			res.IsAdmin = true
		} else {
			res.IsAdmin = false
		} */
		s.respond(w, r, http.StatusOK, u)
		s.logger.Infof("handle /sessions, status :%d", http.StatusOK)
	}
}

func (s *server) handleAdminUpdate() http.HandlerFunc {
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
			IsAdmin bool `json:"isAdmin"`
		}
		res := &resp{}
		if u.Isadmin {
			res.IsAdmin = true
		} else {
			res.IsAdmin = false
		}
		s.respond(w, r, http.StatusOK, res)
		s.logger.Infof("handle /makeAdmin, status :%d", http.StatusOK)
	}
}

func (s *server) handleManagerUpdate() http.HandlerFunc {
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
			IsAdmin bool `json:"isAdmin"`
		}
		res := &resp{}
		if u.Isadmin {
			res.IsAdmin = true
		} else {
			res.IsAdmin = false
		}
		s.respond(w, r, http.StatusOK, res)
		s.logger.Infof("handle /makeManager, status :%d", http.StatusOK)
	}
}

func (s *server) handlePasswordChange() http.HandlerFunc {
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

func (s *server) handleDepartmentCondition() http.HandlerFunc {
	type request struct {
		Email string `json:"email"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			s.logger.Warningf("handle /departmentCondition, status :%d, error :%e", http.StatusBadRequest, err)
			return
		}
		u, err := s.store.User().DepartmentCondition(req.Email)
		if err != nil {
			s.logger.Warningf("handle /departmentCondition, status :%d, error :%e", http.StatusBadRequest, err)
			return
		}
		type resp struct {
			Departments           model.Department
			MonitoringSpecialist  bool `json:"monitoring_specialist"`
			MonitoringResponsible int  `json:"monitoring_responsible"`
		}
		res := &resp{}
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
		s.logger.Infof("handle /departmentCondition, status :%d", http.StatusOK)
	}
}

func (s *server) handleUserUpdate() http.HandlerFunc {

	type request struct {
		Email                 string           `json:"email"`
		Name                  string           `json:"name"`
		SeccondName           string           `json:"seccondName"`
		IsAdmin               bool             `json:"isadmin"`
		Departments           model.Department `json:"Departments"`
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
			req.IsAdmin,
			req.MonitoringSpecialist,
			req.MonitoringResponsible)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			s.logger.Warningf("handle /userUpdate, status :%d, error :%e", http.StatusBadRequest, err)
			return
		}
		type resp struct {
			IsAdmin               bool             `json:"isadmin"`
			Departments           model.Department `json:"Departments"`
			MonitoringSpecialist  bool             `json:"monitoring_specialist"`
			MonitoringResponsible int              `json:"monitoring_responsible"`
		}
		res := &resp{}
		res.IsAdmin = u.Isadmin
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

func (s *server) handleUserDelete() http.HandlerFunc {
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

func (s *server) handleSendEmail() http.HandlerFunc {
	type request struct {
		RecipientMail string `json:"recipient_mail"`
		Subject       string `json:"subject"`
		Body          string `json:"body"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			s.logger.Warningf("handle /sendEmail, status :%d, error :%e", http.StatusBadRequest, err)
			return
		}
		err := mail.SendEmailMessage(s.config.EmailSender, s.config.PasswordSender, s.config.SmtpEmail, req.RecipientMail, req.Subject, req.Body, s.logger)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, errorIncorrectEmailOrPassword)
			s.logger.Warningf("handle /sendEmail, status :%d, error :%e", http.StatusBadRequest, err)
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
		s.logger.Infof("handle /sendEmail, status :%d", http.StatusOK)
	}
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}
func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

package apiserver

import (
	"encoding/json"
	"errors"
	"net/http"
	"restApi/internal/app/model"
	"restApi/internal/app/store"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var (
	errorIncorrectEmailOrPassword = errors.New("incorrect email or password")
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
	store  store.UserStore
}

func newServer(store store.UserStore) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:  store,
	}

	s.configureRouter()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	//s.router.Host("{subdomain:[a-z]+}.example.com")
	s.router.HandleFunc("/users", s.handleUsersCreate()).Methods("POST")            //почта + пароль + имя + фамилия + булевые значения для каждого отдела -> статус:201 json {"id":27,"email":"test3@gmail.com","isadmin":false}
	s.router.HandleFunc("/sessions", s.handleSessionsCreate()).Methods("POST")      //почта + пароль -> статус:200 json {"isAdmin":"false"}
	s.router.HandleFunc("/makeAdmin", s.handleAdminUpdate()).Methods("PUT")         //почта  -> статус:200 json {isAdmin:true}
	s.router.HandleFunc("/makeManager", s.handleManagerUpdate()).Methods("PUT")     //почта  -> статус:200 json {isAdmin:false}
	s.router.HandleFunc("/changePassword", s.handlePasswordChange()).Methods("PUT") //почта + новый пароль -> статус:200 json {Модель пользователя с очищенным полем пароля}
	s.router.HandleFunc("/departmentCondition", s.handleDepartmentCondition()).Methods("POST")
	s.router.HandleFunc("/departmentUpdate", s.handleDepartmentUpdate()).Methods("PUT")
	s.router.HandleFunc("/deleteUser", s.handleUserDelete()).Methods("DELETE")
}

func (s *server) handleUsersCreate() http.HandlerFunc {
	type request struct {
		Email                       string `json:"email"`
		IsAdmin                     bool   `json:"isadmin"`
		Password                    string `json:"password"`
		Name                        string `json:"name"`
		SeccondName                 string `json:"seccondName"`
		EducationDepartment         bool   `json:"educationDepartment"`
		SourceTrackingDepartment    bool   `json:"sourceTrackingDepartment"`
		PeriodicReportingDepartment bool   `json:"periodicReportingDepartment"`
		InternationalDepartment     bool   `json:"internationalDepartment"`
		DocumentationDepartment     bool   `json:"documentationDepartment"`
		NrDepartment                bool   `json:"nrDepartment"`
		DbDepartment                bool   `json:"dbDepartment"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
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
			return
		}
		u.Sanitize()
		_, err := s.store.User().DepartmentUpdate(req.Email, req.EducationDepartment, req.SourceTrackingDepartment, req.PeriodicReportingDepartment, req.InternationalDepartment, req.DocumentationDepartment, req.NrDepartment, req.DbDepartment, req.IsAdmin)
		if err != nil {
			//fmt.Println("тут")
			s.error(w, r, http.StatusBadRequest, err)
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
		s.respond(w, r, http.StatusCreated, res)
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
			return
		}
		u, err := s.store.User().FindByEmail(req.Email)
		if err != nil || !u.ComparePassword(req.Password) {
			s.error(w, r, http.StatusUnauthorized, errorIncorrectEmailOrPassword)
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
			return
		}
		u, err := s.store.User().UpdateRoleAdmin(req.Email)
		if err != nil {
			s.error(w, r, http.StatusUnauthorized, errorIncorrectEmailOrPassword)
			return
		}
		type resp struct {
			IsAdmin string `json:"isAdmin"`
		}
		res := &resp{}
		if u.Isadmin {
			res.IsAdmin = "true"
		} else {
			res.IsAdmin = "false"
		}
		s.respond(w, r, http.StatusOK, res)
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
			return
		}
		u, err := s.store.User().UpdateRoleManager(req.Email)
		if err != nil {
			s.error(w, r, http.StatusUnauthorized, errorIncorrectEmailOrPassword)
			return
		}
		type resp struct {
			IsAdmin string `json:"isAdmin"`
		}
		res := &resp{}
		if u.Isadmin {
			res.IsAdmin = "true"
		} else {
			res.IsAdmin = "false"
		}
		s.respond(w, r, http.StatusOK, res)
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
			return
		}
		u := &model.User{
			Email:    req.Email,
			Password: req.Password,
		}
		if err := s.store.User().ChangePassword(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		u.Sanitize()
		s.respond(w, r, http.StatusOK, u)
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
			return
		}
		u, err := s.store.User().DepartmentCondition(req.Email)
		if err != nil {
			return
		}
		type resp struct {
			EducationDepartment         bool `json:"educationDepartment"`
			SourceTrackingDepartment    bool `json:"sourceTrackingDepartment"`
			PeriodicReportingDepartment bool `json:"periodicReportingDepartment"`
			InternationalDepartment     bool `json:"internationalDepartment"`
			DocumentationDepartment     bool `json:"documentationDepartment"`
			NrDepartment                bool `json:"nrDepartment"`
			DbDepartment                bool `json:"dbDepartment"`
		}
		res := &resp{}
		res.EducationDepartment = u.EducationDepartment
		res.SourceTrackingDepartment = u.SourceTrackingDepartment
		res.PeriodicReportingDepartment = u.PeriodicReportingDepartment
		res.InternationalDepartment = u.InternationalDepartment
		res.DocumentationDepartment = u.DocumentationDepartment
		res.NrDepartment = u.NrDepartment
		res.DbDepartment = u.DbDepartment
		s.respond(w, r, http.StatusOK, res)
	}
}

func (s *server) handleDepartmentUpdate() http.HandlerFunc {

	type request struct {
		Email                       string `json:"email"`
		IsAdmin                     bool   `json:"isadmin"`
		EducationDepartment         bool   `json:"educationDepartment"`
		SourceTrackingDepartment    bool   `json:"sourceTrackingDepartment"`
		PeriodicReportingDepartment bool   `json:"periodicReportingDepartment"`
		InternationalDepartment     bool   `json:"internationalDepartment"`
		DocumentationDepartment     bool   `json:"documentationDepartment"`
		NrDepartment                bool   `json:"nrDepartment"`
		DbDepartment                bool   `json:"dbDepartment"`
	}
	return func(w http.ResponseWriter, r *http.Request) {

		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		u, err := s.store.User().DepartmentUpdate(req.Email, req.EducationDepartment, req.SourceTrackingDepartment, req.PeriodicReportingDepartment, req.InternationalDepartment, req.DocumentationDepartment, req.NrDepartment, req.DbDepartment, req.IsAdmin)
		if err != nil {
			//fmt.Println("тут")
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		type resp struct {
			IsAdmin                     bool `json:"isadmin"`
			EducationDepartment         bool `json:"educationDepartment"`
			SourceTrackingDepartment    bool `json:"sourceTrackingDepartment"`
			PeriodicReportingDepartment bool `json:"periodicReportingDepartment"`
			InternationalDepartment     bool `json:"internationalDepartment"`
			DocumentationDepartment     bool `json:"documentationDepartment"`
			NrDepartment                bool `json:"nrDepartment"`
			DbDepartment                bool `json:"dbDepartment"`
		}
		res := &resp{}
		res.EducationDepartment = u.EducationDepartment
		res.SourceTrackingDepartment = u.SourceTrackingDepartment
		res.PeriodicReportingDepartment = u.PeriodicReportingDepartment
		res.InternationalDepartment = u.InternationalDepartment
		res.DocumentationDepartment = u.DocumentationDepartment
		res.NrDepartment = u.NrDepartment
		res.DbDepartment = u.DbDepartment
		s.respond(w, r, http.StatusOK, res)
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
			return
		}
		err := s.store.User().Delete(req.Email)
		if err != nil {
			s.error(w, r, http.StatusUnauthorized, errorIncorrectEmailOrPassword)
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

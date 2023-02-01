package apiserver

import (
	"encoding/json"
	"errors"
	"net/http"
	"restApi/internal/app/logging"

	"restApi/internal/app/store"

	"github.com/gorilla/mux"
	//"github.com/sirupsen/logrus"
)

var (
	errorIncorrectEmailOrPassword = errors.New("incorrect email or password")
	errorThisUserIsNotAdmin       = errors.New("this user is not an admin")
)

type Server struct {
	router *mux.Router
	logger logging.Logger
	store  store.UserStore
	config *Config
}

func newServer(store store.UserStore, config *Config) *Server {
	s := &Server{
		router: mux.NewRouter(),
		logger: logging.GetLogger(),
		store:  store,
		config: config,
	}

	s.configureRouter()
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) configureRouter() {
	//s.router.Host("{subdomain:[a-z]+}.example.com")
	s.router.HandleFunc("/userCreate", s.handleUsersCreate()).Methods("POST")                  //почта + пароль + имя + фамилия + булевые значения для каждого отдела -> статус:201 json {"id":27,"email":"test3@gmail.com","isadmin":false}
	s.router.HandleFunc("/userUpdate", s.handleUserUpdate()).Methods("PUT")                    //почта + булевые значения для каждого отдела  ->  статус:200 json {"isadmin":false,"educationDepartment":true,"sourceTrackingDepartment":true,"periodicReportingDepartment":false,"internationalDepartment":false,"documentationDepartment":false,"nrDepartment":false,"dbDepartment":true}
	s.router.HandleFunc("/userDelete", s.handleUserDelete()).Methods("DELETE")                 //почта  -> статус:200 json  {result : true}
	s.router.HandleFunc("/sessions", s.handleSessionsCreate()).Methods("POST")                 //почта + пароль -> статус:200 json {"isAdmin":"false"}
	s.router.HandleFunc("/changePassword", s.handlePasswordChange()).Methods("PUT")            //почта + новый пароль -> статус:200 json {Модель пользователя с очищенным полем пароля}
	s.router.HandleFunc("/departmentCondition", s.handleDepartmentCondition()).Methods("POST") //почта  -> статус:200 json {"isadmin":false,"educationDepartment":true,"sourceTrackingDepartment":true,"periodicReportingDepartment":false,"internationalDepartment":false,"documentationDepartment":false,"nrDepartment":false,"dbDepartment":true}
	s.router.HandleFunc("/sendEmail", s.handleSendEmail()).Methods("POST")
	s.router.HandleFunc("/adminAccess", s.handleAdminAccess()).Methods("POST")
	s.router.HandleFunc("/parse", s.handleParser()).Methods("POST")
}

func (s *Server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}
func (s *Server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

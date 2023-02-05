package apiserver

import (
	"encoding/json"

	"net/http"
	"restApi/internal/app/handlers"
	"restApi/internal/app/logging"

	"restApi/internal/app/store"

	"github.com/gorilla/mux"
	//"github.com/sirupsen/logrus"
)

type Server struct {
	router   *mux.Router
	logger   logging.Logger
	handlers handlers.Handlers
	config   *Config
}

func newServer(store store.UserStore, config *Config, logger logging.Logger) *Server {

	s := &Server{
		router:   mux.NewRouter(),
		handlers: *handlers.NewHandlers(store, logger),
		logger:   logging.GetLogger(),
		config:   config,
	}

	s.configureRouter()
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) configureRouter() {
	//s.router.Host("{subdomain:[a-z]+}.example.com")
	s.router.HandleFunc("/userCreate", s.handlers.HandleUsersCreate()).Methods("POST")                  //почта + пароль + имя + фамилия + булевые значения для каждого отдела -> статус:201 json {"id":27,"email":"test3@gmail.com","isadmin":false}
	s.router.HandleFunc("/userUpdate", s.handlers.HandleUserUpdate()).Methods("PUT")                    //почта + булевые значения для каждого отдела  ->  статус:200 json {"isadmin":false,"educationDepartment":true,"sourceTrackingDepartment":true,"periodicReportingDepartment":false,"internationalDepartment":false,"documentationDepartment":false,"nrDepartment":false,"dbDepartment":true}
	s.router.HandleFunc("/userDelete", s.handlers.HandleUserDelete()).Methods("DELETE")                 //почта  -> статус:200 json  {result : true}
	s.router.HandleFunc("/sessions", s.handlers.HandleSessionsCreate()).Methods("POST")                 //почта + пароль -> статус:200 json {"isAdmin":"false"}
	s.router.HandleFunc("/changePassword", s.handlers.HandlePasswordChange()).Methods("PUT")            //почта + новый пароль -> статус:200 json {Модель пользователя с очищенным полем пароля}
	s.router.HandleFunc("/departmentCondition", s.handlers.HandleDepartmentCondition()).Methods("POST") //почта  -> статус:200 json {"isadmin":false,"educationDepartment":true,"sourceTrackingDepartment":true,"periodicReportingDepartment":false,"internationalDepartment":false,"documentationDepartment":false,"nrDepartment":false,"dbDepartment":true}
	s.router.HandleFunc("/sendEmail", s.handlers.HandleSendEmail(s.config.EmailSender, s.config.PasswordSender, s.config.SmtpEmail)).Methods("POST")
	s.router.HandleFunc("/adminAccess", s.handlers.HandleAdminAccess()).Methods("POST")
	s.router.HandleFunc("/parse", s.handlers.HandleParser()).Methods("POST")
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

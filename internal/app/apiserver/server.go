package apiserver

import (
	"restApi/internal/app/handlers"
	"restApi/internal/app/logging"

	"restApi/internal/app/store"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	//"github.com/sirupsen/logrus"
)

type Server struct {
	router   *fiber.App
	logger   logging.Logger
	handlers handlers.Handlers
	config   *Config
}

func newServer(store store.UserStore, config *Config, logger logging.Logger) *Server {

	s := &Server{
		router:   fiber.New(fiber.Config{ServerHeader: "PVsystem API", AppName: "API v1.0.0"}),
		handlers: *handlers.NewHandlers(store, logger),
		logger:   logging.GetLogger(),
		config:   config,
	}

	s.configureRouter()
	return s
}

func (s *Server) configureRouter() {

	// s.router.Use(cors.New(cors.Config{
	// 	AllowHeaders:     "Origin, Content-Type, Accept",
	// 	AllowCredentials: true,
	// }))

	s.router.Get("/swagger/", swagger.HandlerDefault)

	user := s.router.Group("/user")
	user.Use(logger.New())
	user.Post("/create", s.handlers.HandleUsersCreate())
	user.Delete("/delete", s.handlers.HandleUserDelete())
	user.Post("/session", s.handlers.HandleSessionsCreate())
	user.Put("/update", s.handlers.HandleUserUpdate())
	user.Put("/changePassword", s.handlers.HandlePasswordChange())

	email := s.router.Group("/email")
	email.Post("/send", s.handlers.HandleSendEmail(s.config.EmailSender, s.config.PasswordSender, s.config.SmtpEmail))

	parser := s.router.Group("/parser")
	parser.Post("/parse", s.handlers.HandleParser())

	department := s.router.Group("/department")
	department.Post("/condition", s.handlers.HandleDepartmentCondition())

	admin := s.router.Group("admin")
	admin.Post("/access", s.handlers.HandleAdminAccess())
	//s.router.Host("{subdomain:[a-z]+}.example.com")
	//	s.router.HandleFunc("/userCreate", s.handlers.HandleUsersCreate()).Methods("POST")                  //почта + пароль + имя + фамилия + булевые значения для каждого отдела -> статус:201 json {"id":27,"email":"test3@gmail.com","isadmin":false}
	//	s.router.HandleFunc("/userUpdate", s.handlers.HandleUserUpdate()).Methods("PUT")                    //почта + булевые значения для каждого отдела  ->  статус:200 json {"isadmin":false,"educationDepartment":true,"sourceTrackingDepartment":true,"periodicReportingDepartment":false,"internationalDepartment":false,"documentationDepartment":false,"nrDepartment":false,"dbDepartment":true}
	//	s.router.HandleFunc("/userDelete", s.handlers.HandleUserDelete()).Methods("DELETE")                 //почта  -> статус:200 json  {result : true}
	//	s.router.HandleFunc("/sessions", s.handlers.HandleSessionsCreate()).Methods("POST")                 //почта + пароль -> статус:200 json {"isAdmin":"false"}
	//	s.router.HandleFunc("/changePassword", s.handlers.HandlePasswordChange()).Methods("PUT")            //почта + новый пароль -> статус:200 json {Модель пользователя с очищенным полем пароля}
	//	s.router.HandleFunc("/departmentCondition", s.handlers.HandleDepartmentCondition()).Methods("POST") //почта  -> статус:200 json {"isadmin":false,"educationDepartment":true,"sourceTrackingDepartment":true,"periodicReportingDepartment":false,"internationalDepartment":false,"documentationDepartment":false,"nrDepartment":false,"dbDepartment":true}
	//s.router.HandleFunc("/sendEmail", s.handlers.HandleSendEmail(s.config.EmailSender, s.config.PasswordSender, s.config.SmtpEmail)).Methods("POST")
	//s.router.HandleFunc("/adminAccess", s.handlers.HandleAdminAccess()).Methods("POST")
	//s.router.HandleFunc("/parse", s.handlers.HandleParser()).Methods("POST")
}

package apiserver

import (
	"restApi/internal/app/handlers"
	"restApi/internal/app/logging"

	_ "restApi/docs"
	"restApi/internal/app/store"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	s.router.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))

	s.router.Get("/swagger/*", swagger.HandlerDefault)

	user := s.router.Group("/user")
	user.Use(logger.New())
	user.Post("/create", s.handlers.HandleUsersCreate())
	user.Delete("/delete", s.handlers.HandleUserDelete())
	user.Post("/session", s.handlers.HandleSessionsCreate())
	user.Put("/update", s.handlers.HandleUserUpdate())
	user.Put("/change/password", s.handlers.HandlePasswordChange())

	email := s.router.Group("/email")
	email.Post("/send", s.handlers.HandleSendEmail(s.config.EmailSender, s.config.PasswordSender, s.config.SmtpEmail))

	parser := s.router.Group("/parser")
	parser.Post("/parse", s.handlers.HandleParser())

	department := s.router.Group("/department")
	department.Get("/condition/:email", s.handlers.HandleDepartmentCondition())

	admin := s.router.Group("admin")
	admin.Post("/access", s.handlers.HandleAdminAccess())
}

package apiserver

import (
	"restApi/internal/app/handlers"

	_ "restApi/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/sirupsen/logrus"
)

type Server struct {
	router   *fiber.App
	logger   *logrus.Logger
	handlers handlers.Handlers
	config   *Config
}

func newServer(config *Config, logger *logrus.Logger) *Server {

	s := &Server{
		router:   fiber.New(fiber.Config{ServerHeader: "PVsystem parser API", AppName: "API v2.0.0"}),
		handlers: *handlers.NewHandlers(logger),
		logger:   logger,
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

	parser := s.router.Group("/parser")
	parser.Post("/parse", s.handlers.HandleParser())

}

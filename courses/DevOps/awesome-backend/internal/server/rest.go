package server

import (
	"context"
	"fmt"

	"awesome-backend/internal/service"

	jwt "git.itmo.su/go-modules/jwt-module/v4"
	"git.itmo.su/go-modules/log/v4"
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.uber.org/zap"
)

type Server struct {
	port string

	auth jwt.Service
	svc  service.Service
	app  *fiber.App
}

func New(
	port string,
	auth jwt.Service,
	service service.Service) *Server {
	srv := Server{
		port: port,
		auth: auth,
		svc:  service,
	}

	cfg := fiber.Config{
		DisableStartupMessage: true,
		CaseSensitive:         true,
		StrictRouting:         false,
	}

	app := fiber.New(cfg)
	app.Use(
		recover.New(recover.Config{
			EnableStackTrace: true,
			StackTraceHandler: func(_ *fiber.Ctx, e interface{}) {
				log.Error("panic handled", zap.Any("error", e), zap.Stack("stack"))
			},
		}),
		fiberzap.New(fiberzap.Config{Logger: log.Sugared().Desugar()}),
		srv.auth.FiberMiddleware(),
	)

	srv.app = app

	srv.initRoutes()
	return &srv
}

func (s *Server) Start(_ context.Context) error {
	log.Infof("starting server with port: %s", s.port)
	if err := s.app.Listen(":" + s.port); err != nil {
		return fmt.Errorf("HTTP server stopped execution: %w", err)
	}
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	return s.app.ShutdownWithContext(ctx)
}

func (s *Server) Name() string {
	return "rest-server"
}

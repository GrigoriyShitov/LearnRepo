package server

import (
	"net/http"

	"git.itmo.su/go-modules/swaga/v4"
	"github.com/gofiber/fiber/v2"
)

func (s *Server) initRoutes() {
	// healthcheck endpoint
	app := s.app.Group("/awesome-backend")
	app.Get("/healthcheck", s.HealthCheck)

	v1 := app.Group("/api/v1")
	swaga.Setup(v1, swaga.Config{
		URL:       "/awesome-backend/api/v1",
		MountPath: "/docs",
		Title:     "awesome-backend API",
	}, swaga.WithSpecFile("./docs/swagger.json"),
	)
	s.RmqRoutes(v1)
	s.RedisRoutes(v1)
}

func (s *Server) HealthCheck(c *fiber.Ctx) error {
	if err := s.svc.Ping(c.Context()); err != nil {
		return c.SendStatus(http.StatusServiceUnavailable)
	}
	return c.SendStatus(http.StatusOK)
}

package server

import (
	"net/http"

	"git.itmo.su/go-modules/swaga/v4"
	"github.com/gofiber/fiber/v2"
)

func (s *Server) initRoutes() {
	// healthcheck endpoint
	app := s.app.Group("/awesome-backend")
	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusOK)
	})

	v1 := app.Group("/api/v1")
	swaga.Setup(v1, swaga.Config{
		URL:       "/awesome-backend/api/v1",
		MountPath: "/docs",
		Title:     "awesome-backend API",
	}, swaga.WithSpecFile("./docs/swagger.json"),
	)
	//rabbitMQ routes
	s.RmqRoutes(v1)
}

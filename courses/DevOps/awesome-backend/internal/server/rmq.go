package server

import (
	pkgerr "awesome-backend/pkg/errors"

	jwt "git.itmo.su/go-modules/jwt-module/v4"
	"github.com/gofiber/fiber/v2"
)

// RmqRoutes rabbitMQ routes
func (s *Server) RmqRoutes(router fiber.Router) {
	rmq := router.Group("/rmq")

	rmq.Get("/get", jwt.NoAuth(s.getRmqData))
	rmq.Post("/send", jwt.NoAuth(s.sendRmqData))
}

// getRmqData get all data from RabbitMq. Queue name from path parameter.
func (s *Server) getRmqData(c *fiber.Ctx) error {
	ctx := c.Context()
	data, err := s.svc.GetRMQData(ctx, c.Query("queue_name", ""))
	return s.response(c, data, err)
}

// sendRmqData send data to RabbitMq. Queue name from path parameter.
func (s *Server) sendRmqData(c *fiber.Ctx) error {
	ctx := c.Context()
	var input struct {
		Message string `json:"message"`
	}
	if err := c.BodyParser(&input); err != nil {
		return s.response(c, nil, pkgerr.ErrValidationFailed.WithCause(err))
	}
	err := s.svc.SendRmqData(ctx, c.Query("queue_name", ""), input.Message)
	return s.response(c, nil, err)
}

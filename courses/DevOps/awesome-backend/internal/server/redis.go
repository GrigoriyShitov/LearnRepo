package server

import (
	pkgerr "awesome-backend/pkg/errors"

	jwt "git.itmo.su/go-modules/jwt-module/v4"
	"github.com/gofiber/fiber/v2"
)

func (s *Server) RedisRoutes(router fiber.Router) {
	r := router.Group("/redis")

	r.Get("/get", jwt.NoAuth(s.getRedisData))
	r.Post("/send", jwt.NoAuth(s.sendRedisData))
}

// getRedisData get all data from Redis. Key name from path parameter.
func (s *Server) getRedisData(c *fiber.Ctx) error {
	ctx := c.Context()
	data, err := s.svc.GetRedisData(ctx, c.Query("key", ""))
	return s.response(c, data, err)
}

// sendRedisData send data to Redis. Queue name from path parameter.
func (s *Server) sendRedisData(c *fiber.Ctx) error {
	ctx := c.Context()
	var input struct {
		Message string `json:"message"`
	}
	if err := c.BodyParser(&input); err != nil {
		return s.response(c, nil, pkgerr.ErrValidationFailed.WithCause(err))
	}
	err := s.svc.SendRedisData(ctx, c.Query("key", ""), input.Message)
	return s.response(c, nil, err)
}

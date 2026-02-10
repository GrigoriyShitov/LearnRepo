package service

import (
	"awesome-backend/internal/store"
	"awesome-backend/pkg/rmq"

	"github.com/gofiber/fiber/v2"
)

type (
	Service interface {
		GetRMQData(c *fiber.Ctx, queueName string) (interface{}, error)
		SendRmqData(c *fiber.Ctx, query string, body string) error
	}

	serviceImpl struct {
		s          store.Store
		rmqHandler rmq.RabbitHandler
	}
)

func New(
	s store.Store,
	rmqHandler rmq.RabbitHandler,
) Service {
	return &serviceImpl{
		s:          s,
		rmqHandler: rmqHandler,
	}
}

func (s *serviceImpl) SendRmqData(c *fiber.Ctx, query string, body string) error {
	return s.rmqHandler.SendData(c, query, body)
}

func (s *serviceImpl) GetRMQData(c *fiber.Ctx, queueName string) (interface{}, error) {
	return s.rmqHandler.GetItems(c, queueName)
}

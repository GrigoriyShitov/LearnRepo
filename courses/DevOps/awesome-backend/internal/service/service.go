package service

import (
	"awesome-backend/internal/store"
	"awesome-backend/pkg/cache"
	"awesome-backend/pkg/rmq"
	"context"
	"errors"

	"git.itmo.su/go-modules/log/v4"
)

type (
	Service interface {
		Ping(ctx context.Context) error
		GetRMQData(ctx context.Context, queueName string) (interface{}, error)
		SendRmqData(ctx context.Context, query string, body string) error
		SendRedisData(ctx context.Context, key string, message string) error
		GetRedisData(ctx context.Context, key string) (interface{}, error)
	}

	serviceImpl struct {
		s          store.Store
		rmqHandler rmq.Handler
		cache      cache.Handler
	}
)

func (s *serviceImpl) Ping(ctx context.Context) error {
	returnValue := errors.New("service unavailable")
	var cacheFlag, brokerFlag bool
	if err := s.cache.Ping(ctx); err != nil {
		log.Warnf("cache ping failed: %v", err)
		cacheFlag = true
	}
	if err := s.rmqHandler.Ping(ctx); err != nil {
		log.Warnf("rmq handler unavailable")
		brokerFlag = true
	}
	if cacheFlag && brokerFlag {
		return returnValue
	}
	return nil
}

func New(
	s store.Store,
	rmqHandler rmq.Handler,
	cache cache.Handler,
) Service {
	return &serviceImpl{
		s:          s,
		rmqHandler: rmqHandler,
		cache:      cache,
	}
}

func (s *serviceImpl) SendRmqData(ctx context.Context, query string, body string) error {
	if err := s.rmqHandler.Ping(ctx); err != nil {
		log.Errorf("Error pinging rabbitmq handler: %v", err)
		log.Warnf("Sending message to redis")
		return s.cache.SendRedisData(ctx, query, body)
	}
	return s.rmqHandler.SendData(ctx, query, body)

}

func (s *serviceImpl) GetRMQData(ctx context.Context, queueName string) (interface{}, error) {
	if err := s.rmqHandler.Ping(ctx); err != nil {
		log.Errorf("Error pinging rabbitmq handler: %v", err)
		log.Warnf("Getting message from redis")
		return s.cache.GetRedisData(ctx, queueName)
	}
	return s.rmqHandler.GetItems(ctx, queueName)
}

func (s *serviceImpl) SendRedisData(ctx context.Context, key string, message string) error {
	if err := s.cache.Ping(ctx); err != nil {
		log.Errorf("Error pinging cache: %v", err)
		log.Warnf("Sending message to RabbitMq...")
		return s.rmqHandler.SendData(ctx, key, message)
	}

	return s.cache.SendRedisData(ctx, key, message)
}

func (s *serviceImpl) GetRedisData(ctx context.Context, key string) (interface{}, error) {
	if err := s.cache.Ping(ctx); err != nil {
		log.Errorf("Error pinging cache: %v\nReading", err)
		log.Warnf("Getting message from RabbitMq...")
		return s.GetRMQData(ctx, key)
	}
	return s.cache.GetRedisData(ctx, key)
}

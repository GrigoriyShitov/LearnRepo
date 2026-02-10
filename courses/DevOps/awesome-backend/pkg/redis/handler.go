package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Handler interface {
	Stop(ctx context.Context) error
	Start(ctx context.Context) error
}

func InitHandler(host, port, password string, db int) *redis.Options {
	return &redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       db,
	}

}

func NewRabbitMqHandler(options *redis.Options) Handler {
	if options == nil {
		panic("nil redis handler")
	}
	return &redisHandler{options: options}
}

type redisHandler struct {
	options *redis.Options
}

func (r redisHandler) Stop(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (r redisHandler) Start(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

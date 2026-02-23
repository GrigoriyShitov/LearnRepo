package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"git.itmo.su/go-modules/log/v4"
	"github.com/redis/go-redis/v9"
)

type Handler interface {
	Stop(ctx context.Context) error
	Start(ctx context.Context) error
	SendRedisData(ctx context.Context, key string, message string) error
	GetRedisData(ctx context.Context, key string) (interface{}, error)
	Ping(ctx context.Context) error
}

func InitHandler(host, port, password string, db int) *redis.Options {
	return &redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       db,
	}

}

func NewRedisHandler(options *redis.Options, ttl time.Duration) Handler {
	if options == nil {
		panic("nil redis handler")
	}
	return &redisHandler{
		options: options,
		ttl:     ttl,
	}
}

type redisHandler struct {
	options *redis.Options
	ttl     time.Duration
	client  *redis.Client
}

func (r *redisHandler) Ping(ctx context.Context) error {
	// Защита от долгого зависания
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	if r.client == nil {
		return fmt.Errorf("redis client not initialized")
	}

	if err := r.client.Ping(ctx).Err(); err != nil {
		return fmt.Errorf("redis ping failed: %w", err)
	}

	return nil
}

func (r *redisHandler) SendRedisData(ctx context.Context, key string, message string) error {
	data, _ := json.Marshal(message)
	if err := r.client.Set(ctx, key, data, r.ttl).Err(); err != nil {
		log.Errorf("redis send data error: %v", err)
		return err
	}
	return nil
}

func (r *redisHandler) GetRedisData(ctx context.Context, key string) (interface{}, error) {
	val, err := r.client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return val, nil
}

func (r *redisHandler) Stop(ctx context.Context) error {
	return r.client.Close()
}

func (r *redisHandler) Start(ctx context.Context) error {
	r.client = redis.NewClient(r.options)

	// Проверка соединения
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := r.client.Ping(ctx).Err(); err != nil {
		return errors.New("не удалось подключиться к Redis: " + err.Error())
	}
	log.WithField("Redis", "ok").Info("redis ping ok")
	return nil
}

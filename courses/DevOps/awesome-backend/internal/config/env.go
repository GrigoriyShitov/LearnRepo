//nolint:all
package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

var (
	App      appConfig
	Jwt      jwtConfig
	Postgres postgresConfig
	Client   clientConfig
	RabbitMq rabbitmqConfig
)

type (
	appConfig struct {
		Env                 string `env:"ENV"                    envDefault:"dev"`
		Port                string `env:"HTTP_PORT"              envDefault:"8080"`
		StartErrorCodesFrom int    `env:"START_ERROR_CODES_FROM" envDefault:"100"`
	}

	postgresConfig struct {
		Host     string `env:"POSTGRES_HOST"     envDefault:"localhost"`
		DbName   string `env:"POSTGRES_DB"       envDefault:"devopsdb"`
		User     string `env:"POSTGRES_USER"     envDefault:"awesome-backend"`
		Password string `env:"POSTGRES_PASSWORD" envDefault:"postgres"`
		Port     string `env:"POSTGRES_PORT"     envDefault:"5432"`
		SslMode  string `env:"POSTGRES_SSL_MODE" envDefault:"disable"`
	}
	clientConfig struct {
		ClientID     string `env:"CLIENT_ID"`
		ClientSecret string `env:"CLIENT_SECRET"`
		TokenURL     string `env:"TOKEN_SOURCE"`
		Scopes       string `env:"SCOPES"`
	}

	jwtConfig struct {
		Url            string `env:"PUBLIC_KEY_URL"  envDefault:"https://id.itmo.ru/auth/realms/itmo/protocol/openid-connect/certs"`
		CanImpersonate bool   `env:"CAN_IMPERSONATE" envDefault:"true"`
	}
	rabbitmqConfig struct {
		Host     string `env:"RABBITMQ_HOST"     envDefault:"localhost"`
		Port     string `env:"RABBITMQ_PORT"     envDefault:"5672"`
		User     string `env:"RABBITMQ_USER"     envDefault:"guest"`
		Password string `env:"RABBITMQ_PASSWORD" envDefault:"guest"`
	}
)

func InitApp() {
	if err := env.Parse(&App); err != nil {
		panic(fmt.Errorf("failed to parse app config: %w", err))
	}
}
func InitPostgres() {
	if err := env.Parse(&Postgres); err != nil {
		panic(fmt.Errorf("failed to parse postgres config: %w", err))
	}
}

func (p postgresConfig) DSN() string {
	return fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s sslmode=%s TimeZone=Europe/Moscow",
		p.Host, p.User, p.DbName, p.Password, p.Port, p.SslMode)
}
func InitClient() {
	if err := env.Parse(&Client); err != nil {
		panic(fmt.Errorf("failed to parse client config: %w", err))
	}
}

func InitJwt() {
	if err := env.Parse(&Jwt); err != nil {
		panic(fmt.Errorf("failed to parse jwt config: %w", err))
	}
}

func InitRabbitMQ() {
	if err := env.Parse(&RabbitMq); err != nil {
		panic(fmt.Errorf("failed to parse rabbitmq config: %w", err))
	}
}

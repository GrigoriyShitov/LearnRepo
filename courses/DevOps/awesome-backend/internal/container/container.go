//nolint:mnd
package container

import (
	"awesome-backend/pkg/rmq"
	"context"
	"net/http"
	"sync"
	"time"

	"awesome-backend/internal/app"
	"awesome-backend/internal/config"
	"awesome-backend/internal/server"
	"awesome-backend/internal/service"
	"awesome-backend/internal/store"

	"git.itmo.su/go-modules/databases/v4/migrator"
	"git.itmo.su/go-modules/databases/v4/pgx"
	pkghttp "git.itmo.su/go-modules/http-client/v4"

	jwt "git.itmo.su/go-modules/jwt-module/v4"
	"git.itmo.su/go-modules/log/v4"
	"github.com/jackc/pgx/v5/tracelog"
)

type (
	Container interface {
		Starters() []app.StarterFunc
		Stoppers() []app.StopperFunc
		Jwt() jwt.Service
		Server() *server.Server
		DBClient() *pgx.Pgx
		Migrator() *migrator.Migrator
	}

	Starter interface {
		Start(ctx context.Context) error
	}

	Stopper interface {
		Stop(ctx context.Context) error
	}

	containerImpl struct {
		appName string

		starters   []app.StarterFunc
		starterMux sync.Mutex

		stoppers   []app.StopperFunc
		stopperMux sync.Mutex

		httpClient     *http.Client
		httpClientOnce sync.Once

		jwt        jwt.Service
		jwtOnce    sync.Once
		server     *server.Server
		serverOnce sync.Once

		loader     jwt.Loader
		loaderOnce sync.Once

		rmqHandler     rmq.RabbitHandler
		rmqHandlerSync sync.Once

		service     service.Service
		serviceOnce sync.Once

		store        store.Store
		storeOnce    sync.Once
		dbClient     *pgx.Pgx
		dbClientOnce sync.Once
		migrator     *migrator.Migrator
		migratorOnce sync.Once
	}
)

func New(appName string) Container {
	c := &containerImpl{
		appName: appName,
	}

	// in order to start (db, jwt, server)
	c.registerStarter("rabbitMQ", c.RabbitHandler())
	//c.registerStarter("migrator", c.Migrator())
	//c.registerStarter("db", c.DBClient())
	//c.registerStarter("auth", c.Jwt())

	//c.registerStopper("db", c.DBClient())
	//c.registerStopper("auth", c.Jwt())
	c.registerStopper("rabbitMQ", c.RabbitHandler())
	return c
}

func (c *containerImpl) Starters() []app.StarterFunc {
	c.starterMux.Lock()
	defer c.starterMux.Unlock()
	return c.starters
}

func (c *containerImpl) Stoppers() []app.StopperFunc {
	c.stopperMux.Lock()
	defer c.stopperMux.Unlock()
	return c.stoppers
}
func (c *containerImpl) HTTPClient() *http.Client {
	if c.httpClient == nil {
		c.httpClientOnce.Do(func() {
			c.httpClient = pkghttp.New(time.Second * 5)
		})
	}
	return c.httpClient
}

func (c *containerImpl) Jwt() jwt.Service {
	if c.jwt == nil {
		c.jwtOnce.Do(func() {
			config.InitJwt()
			c.jwt = jwt.NewJwtService(c.HTTPClient(), config.Jwt.Url, jwt.Settings{ServiceImpersonation: config.Jwt.CanImpersonate})
		})
	}
	return c.jwt
}

func (c *containerImpl) Loader() jwt.Loader {
	if c.loader == nil {
		c.loaderOnce.Do(func() {
			config.InitClient()
			c.loader = jwt.NewLoader(c.HTTPClient(), jwt.LoaderSettings{
				TokenSource:  config.Client.TokenURL,
				ClientId:     config.Client.ClientID,
				ClientSecret: config.Client.ClientSecret,
				Scopes:       config.Client.Scopes,
			})
		})
	}
	return c.loader
}

func (c *containerImpl) Server() *server.Server {
	if c.server == nil {
		c.serverOnce.Do(func() {
			config.InitApp()
			c.server = server.New(config.App.Port, c.Jwt(), c.Service())
		})
	}
	return c.server
}

func (c *containerImpl) Service() service.Service {
	if c.service == nil {
		c.serviceOnce.Do(func() {
			c.service = service.New(c.Store(), c.RabbitHandler())
		})
	}
	return c.service
}

func (c *containerImpl) Store() store.Store {
	if c.store == nil {
		c.storeOnce.Do(func() {
			c.store = store.New(
				c.DBClient(),
			)
		})
	}
	return c.store
}
func (c *containerImpl) DBClient() *pgx.Pgx {
	if c.dbClient == nil {
		c.dbClientOnce.Do(func() {
			config.InitApp()
			config.InitPostgres()
			traceLogLevel := tracelog.LogLevelDebug
			if config.App.Env == "prod" {
				traceLogLevel = tracelog.LogLevelInfo
			}
			c.dbClient = pgx.New(config.Postgres.DSN(),
				pgx.WithTraceLogLevel(traceLogLevel),
				pgx.WithMinConns(5),
				pgx.WithMaxConns(50),
			)
		})
	}
	return c.dbClient
}
func (c *containerImpl) Migrator() *migrator.Migrator {
	if c.migrator == nil {
		c.migratorOnce.Do(func() {
			config.InitPostgres()
			c.migrator = migrator.New(config.Postgres.DSN(), migrator.DialectPostgres,
				migrator.WithLogger(log.Sugared()),
				migrator.WithTable("goose"),
			)
		})
	}
	return c.migrator
}

func (c *containerImpl) RabbitHandler() rmq.RabbitHandler {
	if c.rmqHandler == nil {
		c.rmqHandlerSync.Do(func() {
			config.InitRabbitMQ()
			c.rmqHandler = rmq.NewRabbitMqHandler(
				rmq.InitHandler(
					config.RabbitMq.User,
					config.RabbitMq.Password,
					config.RabbitMq.Host,
					config.RabbitMq.Port,
				),
			)
		})
	}
	return c.rmqHandler
}

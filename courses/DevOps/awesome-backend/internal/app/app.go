package app

import (
	"context"
	"net"
	"os/signal"
	"syscall"
	"time"

	"awesome-backend/internal/config"
	"awesome-backend/internal/server"

	"git.itmo.su/go-modules/log/v4"
	"golang.org/x/sync/errgroup"
)

type App struct {
	srv *server.Server

	onStart []StarterFunc
	onStop  []StopperFunc
}

const (
	startTimeout = time.Second * 10
	stopTimeout  = time.Second * 5
)

func New(srv *server.Server, onStart []StarterFunc, onStop []StopperFunc) *App {
	return &App{
		srv:     srv,
		onStart: onStart,
		onStop:  onStop,
	}
}

func (a *App) Init() error {
	startTime := time.Now()
	log.Info("starting app...")

	gracefulCtx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	startCtx, cancel := context.WithTimeout(gracefulCtx, startTimeout)
	defer cancel()

	for i := range a.onStart {
		err := a.onStart[i](startCtx)
		if err != nil {
			return err
		}
	}

	runG, runCtx := errgroup.WithContext(gracefulCtx)
	runG.Go(func() error {
		return a.srv.Start(runCtx)
	})

	runG.Go(a.GracefulShutdown(runCtx, stopTimeout))

	waitForPort(config.App.Port)
	log.WithField("name", "rest-server").Info("started")

	took := time.Since(startTime)
	log.WithField(log.LogFieldDurationMs, took.Milliseconds()).Info("application started")

	return runG.Wait()
}

func (a *App) GracefulShutdown(ctx context.Context, timeout time.Duration) func() error {
	return func() error {
		<-ctx.Done()
		log.Info("processing graceful shutdown...")

		shutdownCtx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()

		err := a.srv.Stop(shutdownCtx)
		if err != nil {
			return err
		}
		log.WithField("name", "rest-server").Info("stopped")

		for i := len(a.onStop) - 1; i >= 0; i-- {
			err = a.onStop[i](shutdownCtx)
			if err != nil {
				return err
			}
		}

		log.Info("gracefully stopped app")
		return nil
	}
}

func waitForPort(port string) {
	d := &net.Dialer{}
	backoff := 50 * time.Millisecond //nolint:mnd
	const (
		attemptTimeout = 300 * time.Millisecond
		maxBackoff     = 500 * time.Millisecond
	)

	for {
		perTry := attemptTimeout
		ctx, cancel := context.WithTimeout(context.Background(), perTry)
		conn, err := d.DialContext(ctx, "tcp", ":"+port)
		cancel()

		if err == nil {
			_ = conn.Close()
			return
		}

		sleep := backoff
		time.Sleep(sleep)

		if backoff < maxBackoff {
			backoff *= 2
			if backoff > maxBackoff {
				backoff = maxBackoff
			}
		}
	}
}

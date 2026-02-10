package container

import (
	"context"

	"github.com/pkg/errors"

	"git.itmo.su/go-modules/log/v4"
)

func (c *containerImpl) registerStarter(name string, starter Starter) {
	c.starterMux.Lock()
	defer c.starterMux.Unlock()
	c.starters = append(c.starters, func(ctx context.Context) error {
		log.WithField("name", name).Info("starting...")
		err := starter.Start(ctx)
		if err != nil {
			return errors.Wrap(err, name)
		}
		log.WithField("name", name).Info("started")
		return nil
	})
}

func (c *containerImpl) registerStopper(name string, stopper Stopper) {
	c.stopperMux.Lock()
	defer c.stopperMux.Unlock()
	c.stoppers = append(c.stoppers, func(ctx context.Context) error {
		log.WithField("name", name).Info("stopping...")
		err := stopper.Stop(ctx)
		if err != nil {
			return errors.Wrap(err, name)
		}
		log.WithField("name", name).Info("stopped")
		return nil
	})
}

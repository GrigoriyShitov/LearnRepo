package storage

import (
	"context"
	"errors"
)

func AdminValid(ctx context.Context, whoChange uint) error {
	for _, user := range Users {
		if user.ID == whoChange && user.Admin {
			return nil
		}
	}
	err := errors.New("not admin")
	return err
}

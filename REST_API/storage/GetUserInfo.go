package storage

import (
	"context"
	"errors"
)

func GetUserInfo(ctx context.Context, id uint) (*User, error) { // to do (*storage.User, error)

	for _, user := range Users {
		if user.ID == id {
			return &user, nil // to do &user
		}
	}
	err := errors.New("user not found")
	return nil, err //nil
}

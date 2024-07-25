package methods

import (
	"RestApi/storage"
	"context"
	"errors"
)

func GetUserInfo(ctx context.Context, id uint32) (*storage.User, error) { // to do (*storage.User, error)

	for _, user := range storage.Users {
		if user.ID == uint32(id) {
			return &user, nil // to do &user
		}
	}
	err := errors.New("user not found")
	return nil, err //nil
}

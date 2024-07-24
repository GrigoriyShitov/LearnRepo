package methods

import (
	"RestApi/storage"
	"context"
)

func GetUserInfo(ctx context.Context, id uint32) storage.User { // to do (*storage.User, error)

	for _, user := range storage.Users {
		if user.ID == uint32(id) {
			return user // to do &user
		}
	}
	return storage.User{} //nil
}

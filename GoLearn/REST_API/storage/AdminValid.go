package storage

import (
	"context"
	"errors"
)

func AdminValid(ctx context.Context, whoChange uint) error {

	var user User
	result := db.Find(&user, "ID = ?", whoChange)
	if result.Error != nil {
		return result.Error
	}
	if user.Role == "admin" {
		return nil
	}
	err := errors.New("no access")
	return err
}

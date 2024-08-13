package storage

import "context"

func CreateNewUser(ctx context.Context, id uint, name string) error {
	var user User
	user.ID = id
	user.Username = name
	user.Role = "user"
	err := db.Create(&user)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

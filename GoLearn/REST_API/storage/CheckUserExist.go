package storage

import "context"

func CheckUserExist(ctx context.Context, id uint) error {
	var user User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return err
	}
	return nil
}

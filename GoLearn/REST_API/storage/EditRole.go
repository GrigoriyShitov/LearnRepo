package storage

import (
	"context"
	"log"
)

func EditRole(ctx context.Context, idUser uint, role string) {
	//Users[idUser-1].Role = role
	result := db.Model(&User{}).Where("id = ?", idUser).Update("role", role)
	if result.Error != nil {
		return
	} else {
		log.Printf("Role %d changed to %s", idUser, role)
	}
}

package storage

import (
	"context"
)

func EditRole(ctx context.Context, idUser uint, role string) {
	Users[idUser-1].Role = role
}

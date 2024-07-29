package storage

import (
	"context"
)

func EditRole(ctx context.Context, idUser uint, role bool) {
	Users[idUser-1].Admin = role
}

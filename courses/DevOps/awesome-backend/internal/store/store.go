package store

import (
	"git.itmo.su/go-modules/databases/v4/pgx"
)

type (
	Store interface{}

	storeImpl struct{}
)

func New(
	client *pgx.Pgx,
) Store {
	return &storeImpl{}
}

package storage

import (
	"context"
)

func GetOperationOnWallet(ctx context.Context, WalletId uint) ([]Operation, error) {
	retOps := make([]Operation, 0)
	for _, operations := range Operations {
		if operations.WalletId == WalletId {
			retOps = append(retOps, operations)
		}
	}
	return retOps, nil
}

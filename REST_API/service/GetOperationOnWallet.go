package service

import (
	"RestApi/storage"
	"context"
)

func GetOperationOnWallet(ctx context.Context, WalletId uint) ([]storage.Operation, error) {
	retOps := []storage.Operation{}
	for _, operations := range storage.Operations {
		if operations.WalletId == WalletId {
			retOps = append(retOps, operations)
		}
	}
	return retOps, nil
}

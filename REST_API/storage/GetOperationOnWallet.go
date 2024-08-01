package storage

import (
	"context"
)

func GetOperationOnWallet(ctx context.Context, WalletId uint) ([]Operation, error) {
	retOps := make([]Operation, 0)
	result := db.Find(&retOps, "wallet_id = ?", WalletId)
	if result.Error != nil {
		return nil, result.Error
	}
	return retOps, nil
}

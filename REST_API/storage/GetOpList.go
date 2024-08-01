package storage

import "context"

func GetOpList(ctx context.Context, walletId uint) ([]Operation, error) {
	operationList := make([]Operation, 0)
	result := db.Find(&operationList, "wallet_id = ?", walletId)
	if result.Error != nil {
		return nil, result.Error
	}

	return operationList, nil
}

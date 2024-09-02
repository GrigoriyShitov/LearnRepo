package storage

import "context"

func GetOpList(ctx context.Context, id uint, WalletNum uint) ([]Operation, error) {

	userWallets, err := WalletValid(ctx, id, WalletNum)
	if err != nil {
		return nil, err
	}

	walletId := userWallets[WalletNum-1].Id
	operationList := make([]Operation, 0)
	result := db.Find(&operationList, "wallet_id = ?", walletId)
	if result.Error != nil {
		return nil, result.Error
	}

	return operationList, nil
}

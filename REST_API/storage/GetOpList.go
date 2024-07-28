package storage

import "context"

func GetOpList(ctx context.Context, walletId uint) ([]Operation, error) {
	operationList := make([]Operation, 0)
	//err := db.Select(&operationList, `SELECT * FROM operation WHERE wallet_id = $1`, walletId)
	for _, operation := range Operations {
		if operation.WalletId == walletId {
			operationList = append(operationList, operation)
		}
	}
	return operationList, nil
}

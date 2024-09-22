package storage

import "errors"

func AddNewOperation(userID uint, WalletNum uint, Type string, Amount float64, Category string) error {
	var operationToAdd Operation
	Wallets := []Wallet{}
	result := db.Find(&Wallets, "user_id = ?", userID)
	if result.Error != nil {
		return result.Error
	}

	if len(Wallets) == 0 {
		return errors.New("no wallets")
	}
	WalletId := Wallets[WalletNum-1].Id

	operationToAdd.WalletId = WalletId
	operationToAdd.Amount = Amount
	operationToAdd.OperationType = Type
	operationToAdd.OperationCategory = Category
	result = db.Create(&operationToAdd)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

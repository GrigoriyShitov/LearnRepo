package storage

func AddNewOperation(IdUser uint, IdWallet uint, Type string, Amount float64, Category string) error {
	var operationToAdd Operation

	operationToAdd.WalletId = IdWallet
	operationToAdd.Amount = Amount
	operationToAdd.OperationType = Type
	operationToAdd.OperationCategory = Category
	// Operations = append(Operations, operationToAdd)
	result := db.Create(&operationToAdd)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

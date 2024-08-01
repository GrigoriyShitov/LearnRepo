package storage

func AddNewOperation(AddOperation *NewOperation) error {
	var operationToAdd Operation

	operationToAdd.WalletId = AddOperation.IdWallet
	operationToAdd.Amount = AddOperation.Amount
	operationToAdd.OperationType = AddOperation.Type
	operationToAdd.OperationCategory = AddOperation.Category
	// Operations = append(Operations, operationToAdd)
	result := db.Create(&operationToAdd)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

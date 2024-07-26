package storage

func AddNewOperation(AddOperation *NewOperation) error {
	var operationToAdd Operation
	lastId := Operations[len(Operations)-1].OperationId
	operationToAdd.OperationId = lastId + 1
	operationToAdd.WalletId = AddOperation.WalletId
	operationToAdd.Amount = AddOperation.Amount
	operationToAdd.OperationType = AddOperation.Type
	Operations = append(Operations, operationToAdd)
	return nil
}

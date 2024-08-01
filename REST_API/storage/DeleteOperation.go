package storage

func DeleteOperation(idWallet uint, idOperation uint) error {

	// for i, operation := range Operations {
	// 	if operation.OperationId == idOperation && operation.WalletId == idWallet {
	// 		Operations = append(Operations[:i], Operations[i+1:]...)
	// 		return nil
	// 	}
	// }
	// err := errors.New("operation not found")

	var OpToDelete Operation
	result := db.First(&OpToDelete, "operation_id = ? AND wallet_id = ?", idOperation, idWallet)
	if result.Error != nil {
		return result.Error
	}
	db.Delete(&OpToDelete)
	return nil

}

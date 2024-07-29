package storage

import "errors"

func DeleteOperation(idWallet uint, idOperation uint) error {

	for i, operation := range Operations {
		if operation.OperationId == idOperation && operation.WalletId == idWallet {
			Operations = append(Operations[:i], Operations[i+1:]...)
			return nil
		}
	}
	err := errors.New("operation not found")
	return err

}

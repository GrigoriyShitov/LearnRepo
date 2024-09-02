package storage

import (
	"context"
	"errors"
)

func OpCheckAccess(ctx context.Context, userID uint, WalletNum uint, idOperation uint) error {
	var user User
	if err := db.Find(&user, "id = ?", userID).Error; err != nil {
		return err
	}
	if user.Role == "admin" {
		return nil
	}
	Wallets := []Wallet{}
	result := db.Find(&Wallets, "user_id = ?", userID)
	if result.Error != nil {
		return result.Error
	}
	if len(Wallets) == 0 {
		return errors.New("no wallets")
	}

	for _, wallet := range Wallets {
		if wallet.Id == WalletNum {
			var operations []Operation
			err := db.Find(&operations, "wallet_id = ?", WalletNum).Error
			if err != nil {
				return err
			}
			for _, operation := range operations {
				if operation.ID == idOperation {
					return nil
				}
			}
		}
	}
	err := errors.New("no access")
	return err

}

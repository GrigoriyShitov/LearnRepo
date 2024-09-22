package service

import (
	"context"
)

func WalletBalanceInfo(ctx context.Context, WalletId uint) (float64, error) {

	balance, err := CalculateBalance(ctx, WalletId)
	return balance, err

}

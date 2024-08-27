package service

import (
	"RestApi/storage"
	"context"
	"encoding/json"
)

type OperationToOut struct {
	OperationType     string
	Amount            float64
	OperationCategory string
}

func OperationList(ctx context.Context, id uint, idWallet uint) ([]byte, error) {
	err := storage.WalletValid(ctx, id, idWallet)
	if err != nil {
		return nil, err
	}
	ListOfOperations, _ := storage.GetOpList(ctx, idWallet)

	data := ListToOut(ctx, ListOfOperations)
	Output, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return nil, err
	}
	return Output, nil
}

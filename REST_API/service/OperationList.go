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

func OperationList(ctx context.Context, id uint, WalletNum uint) ([]byte, error) {

	ListOfOperations, _ := storage.GetOpList(ctx, id, WalletNum)

	data := ListToOut(ctx, ListOfOperations)
	Output, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return nil, err
	}
	return Output, nil
}

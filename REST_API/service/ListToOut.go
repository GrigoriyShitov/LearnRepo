package service

import (
	"RestApi/storage"
	"context"
)

func ListToOut(ctx context.Context, list []storage.Operation) []OperationToOut {
	var out []OperationToOut
	for _, operation := range list {
		out = append(out, OperationToOut{
			OperationType:     operation.OperationType,
			Amount:            operation.Amount,
			OperationCategory: operation.OperationCategory,
		})
	}
	return out
}

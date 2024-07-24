package storage

type Operation struct {
	OperationId   uint32 `json:"operation_id"`
	WalletId      uint32 `json:"wallet_id"`
	Amount        int32  `json:"amount"`
	OperationType string `json:"operation_type"`
	OperationInfo string `json:"operation_info"`
}

var Operations = []Operation{

	{
		OperationId:   1,
		WalletId:      1,
		Amount:        100,
		OperationType: "deposit",
		OperationInfo: "some info",
	},

	{
		OperationId:   2,
		WalletId:      1,
		Amount:        100,
		OperationType: "withdraw",
		OperationInfo: "some info",
	},
}

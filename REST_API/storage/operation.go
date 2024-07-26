package storage

type Operation struct {
	OperationId   uint    `json:"operation_id"`
	WalletId      uint    `json:"wallet_id"`
	Amount        float64 `json:"amount"`
	OperationType string  `json:"operation_type"`
	OperationInfo string  `json:"operation_info"`
}

type NewOperation struct {
	UserId   uint
	WalletId uint
	Type     string
	Amount   float64
}

var Operations = []Operation{

	{
		OperationId:   1,
		WalletId:      1,
		Amount:        500,
		OperationType: "deposit",
		OperationInfo: "some info",
	},

	{
		OperationId:   2,
		WalletId:      1,
		Amount:        250,
		OperationType: "withdraw",
		OperationInfo: "some info",
	},
}

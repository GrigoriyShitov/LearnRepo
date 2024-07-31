package storage

type Operation struct {
	OperationId       uint    `json:"operation_id" gorm:"primaryKey"`
	WalletId          uint    `json:"wallet_id"`
	Amount            float64 `json:"amount"`
	OperationType     string  `json:"operation_type"`
	OperationCategory string  `json:"operation_category"`
}

type NewOperation struct {
	UserId   uint
	WalletId uint
	Type     string
	Amount   float64
	Category string
}

var Operations = []Operation{

	{
		OperationId:       1,
		WalletId:          1,
		Amount:            500,
		OperationType:     "deposit",
		OperationCategory: "some info",
	},

	{
		OperationId:       2,
		WalletId:          1,
		Amount:            250,
		OperationType:     "withdraw",
		OperationCategory: "some info",
	},
}

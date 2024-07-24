package storage

type Wallet struct {
	WalletId uint32 `json:"wallet_id"`
	UserId   uint32 `json:"user_id"`
	Balance  uint32 `json:"balance"`
}

var Wallets = []Wallet{
	{
		WalletId: 1,
		UserId:   1,
		Balance:  0,
	},

	{
		WalletId: 2,
		UserId:   1,
		Balance:  0,
	},
}

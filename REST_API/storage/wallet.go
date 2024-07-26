package storage

type Wallet struct {
	WalletId uint `json:"wallet_id"`
	UserId   uint `json:"user_id"`
}

var Wallets = []Wallet{
	{
		WalletId: 1,
		UserId:   1,
	},

	{
		WalletId: 2,
		UserId:   1,
	},
	{
		WalletId: 3,
		UserId:   2,
	},

	{
		WalletId: 4,
		UserId:   2,
	},
}

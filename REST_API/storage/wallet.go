package storage

type Wallet struct {
	Id     uint `json:"wallet_id" gorm:"primaryKey"`
	UserId uint `json:"user_id"`
}

var Wallets = []Wallet{
	{
		Id:     1,
		UserId: 1,
	},

	{
		Id:     2,
		UserId: 1,
	},
	{
		Id:     3,
		UserId: 2,
	},

	{
		Id:     4,
		UserId: 2,
	},
}

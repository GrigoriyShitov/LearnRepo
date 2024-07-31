package storage

type User struct {
	ID       uint     `json:"id" gorm:"primaryKey"`
	Username string   `json:"username"`
	Role     string   `json:"role"`
	Wallets  []Wallet `gorm:"foreignKey:UserId;references:ID"`
}

type WalletsInfo struct {
	WalletInfo Wallet  `json:"wallet"`
	Balance    float64 `json:"balance"`
}
type UserInfo struct {
	User        *User
	WalletsList []Wallet
}

var Users = []User{
	{ID: 1, Username: "admin", Role: "admin"},
	{ID: 2, Username: "user", Role: "user"},
	{ID: 3, Username: "guest", Role: "user"},
}

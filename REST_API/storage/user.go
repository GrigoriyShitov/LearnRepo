package storage

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Admin    bool   `json:"admin"`
}

type WalletsInfo struct {
	WalletInfo Wallet  `json:"wallet"`
	Balance    float64 `json:"balance"`
}
type UserInfo struct {
	User        *User
	WalletsList []WalletsInfo
}

var Users = []User{
	{ID: 1, Username: "admin", Admin: true},
	{ID: 2, Username: "user", Admin: false},
	{ID: 3, Username: "guest", Admin: false},
}

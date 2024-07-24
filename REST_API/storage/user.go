package storage

type User struct {
	ID       uint32 `json:"id"`
	Username string `json:"username"`
	Admin    bool   `json:"admin"`
}

var Users = []User{
	{ID: 1, Username: "admin", Admin: true},
	{ID: 2, Username: "user", Admin: false},
	{ID: 3, Username: "guest", Admin: false},
}

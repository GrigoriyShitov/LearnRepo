package service

import (
	"context"
	"encoding/json"
)

type UserInfo struct {
	Username    string
	Role        string
	WalletsList []WalletsInfo
}

func UserInfoToOut(ctx context.Context, Username string, Role string, WalletsList []WalletsInfo) ([]byte, error) {
	data := make([]UserInfo, 0, 1)

	data = append(data, UserInfo{
		Username:    Username,
		Role:        Role,
		WalletsList: WalletsList,
	})
	return json.MarshalIndent(data, "", "    ")
}

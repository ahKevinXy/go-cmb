package notice

import (
	"encoding/json"
)

// Account
//  @Description:   解析account 账户 通知
//  @param s
//  @return *models.AccountNotice
//  @return error
//  @Author  ahKevinXy
//  @Date  2023-04-14 17:39:48
func Account(s string) (*models.AccountNotice, error) {
	var account *models.AccountNotice
	if err := json.Unmarshal([]byte(s), &account); err != nil {

		return nil, err
	}

	return account, nil

}

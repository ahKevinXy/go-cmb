package notice

import (
	"encoding/json"
	"github.com/ahKevinXy/go-cmb/models"
)

// Pay
//  @Description:  支付通知 序列化
//  @param s
//  @return *models.PayResultNotice
//  @return error
//  @Author  ahKevinXy
//  @Date  2023-05-21 20:31:50
func Pay(s string) (*models.PayResultNotice, error) {
	var pay *models.PayResultNotice

	if err := json.Unmarshal([]byte(s), &pay); err != nil {
		return nil, err
	}

	return pay, nil
}

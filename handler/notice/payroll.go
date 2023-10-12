package notice

import (
	"encoding/json"
)

// Payroll
//  @Description:  代发信息通知解析
//  @param s
//  @return *models.PayResultNotice
//  @return error
//  @Author  ahKevinXy
//  @Date  2023-04-14 17:41:51
func Payroll(s string) (*models.PayrollResultNotice, error) {
	var pay *models.PayrollResultNotice

	if err := json.Unmarshal([]byte(s), &pay); err != nil {
		return nil, err
	}

	return pay, nil
}

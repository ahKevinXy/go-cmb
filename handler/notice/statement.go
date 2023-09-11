package notice

import (
	"encoding/json"
	"github.com/ahKevinXy/go-cmb/models"
)

// Statement
//
//	@Description:   回单提醒解析
//	@param s
//	@return *models.StatementNotice
//	@return error
//	@Author  ahKevinXy
//	@Date  2023-04-14 18:16:47
func Statement(s string) (*models.StatementNotice, error) {
	var statement *models.StatementNotice

	if err := json.Unmarshal([]byte(s), &statement); err != nil {
		return nil, err
	}

	return statement, nil
}

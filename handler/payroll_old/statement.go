package payroll_old

import (
	"encoding/json"
	"github.com/ahKevinXy/go-cmb/v1/cmb_errors"
	"github.com/ahKevinXy/go-cmb/v1/constants"
	"github.com/ahKevinXy/go-cmb/v1/help"
	"github.com/ahKevinXy/go-cmb/v1/models"
	"strconv"
	"time"
)

// QueryPayrollStatementDownloadUrl
//
//	@Description:   获取回单地址
//	@param userId
//	@param sm4PrivateKey
//	@param userPrivateKey
//	@param taskid 查询ID
//	@return *models.QueryBatchTransListResponse
//	@return error
//	@Author  ahKevinXy
//	@Date  2023-04-14 17:33:31
func QueryPayrollStatementDownloadUrl(userId, sm4PrivateKey, userPrivateKey, taskid string) (*models.QueryBatchTransListResponse, error) {
	reqData := new(models.QueryPayrollStatementDownloadUrlRequest)
	reqData.Request.Head.Reqid = time.Now().Format("20060102150405000") + strconv.Itoa(time.Now().Nanosecond())
	reqData.Request.Head.Funcode = constants.CmbPayrollStatementDownloadUrl
	reqData.Request.Head.Userid = userId
	reqData.Signature.Sigtim = time.Now().Format("20060102150405")
	reqData.Signature.Sigdat = "__signature_sigdat__"
	reqData.Request.Body.Taskid = taskid
	req, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	res, err := help.CmbSignRequest(string(req), constants.CmbPayrollStatementDownloadUrl, userId, userPrivateKey, sm4PrivateKey)
	if err != nil {
		return nil, err
	}
	if res == "" {
		return nil, cmb_errors.SystemError
	}

	var resp models.QueryBatchTransListResponse

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, err
	}

	return &resp, err
}

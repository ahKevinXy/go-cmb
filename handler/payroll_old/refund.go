package payroll_old

import (
	"encoding/json"
	"github.com/ahKevinXy/go-cmb/constants"
	"github.com/ahKevinXy/go-cmb/help"
	"github.com/ahKevinXy/go-cmb/models"
	"strconv"
	"time"
)

// Refund
//  @Description:   获取退票
//  @param userId
//  @param sm4PrivateKey
//  @param userPrivateKey
//  @param taskId
//  @return *models.GetPayrollPdfResponse
//  @return error
//  @Author  ahKevinXy
//  @Date  2023-06-05 16:43:16
func Refund(
	userId, sm4PrivateKey, userPrivateKey, taskId string) (*models.GetPayrollPdfResponse, error) {
	reqData := new(models.PayrollOldPdfFileRequest)
	reqData.Request.Head.Reqid = time.Now().Format("20060102150405000") + strconv.Itoa(time.Now().Nanosecond())
	reqData.Request.Head.Funcode = constants.CmbPayrollOldQueryTransRefund
	reqData.Request.Head.Userid = userId
	reqData.Signature.Sigtim = time.Now().Format("20060102150405")
	reqData.Signature.Sigdat = "__signature_sigdat__"
	reqData.Request.Body.Taskid = taskId

	req, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	res, err := help.CmbSignRequest(string(req), constants.CmbPayrollOldQueryTransRefund, userId, userPrivateKey, sm4PrivateKey)
	if err != nil {
		return nil, err
	}
	if res == "" {
		return nil, err
	}

	var resp models.GetPayrollPdfResponse

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		print(err)
	}

	return &resp, err

}

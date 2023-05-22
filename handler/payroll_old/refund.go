package payroll_old

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/ahKevinXy/go-cmb/constants"
	"github.com/ahKevinXy/go-cmb/help"
	"github.com/ahKevinXy/go-cmb/models"
)

func Refund(
	userId, sm4Key, userPrivateKey, taskId string) (*models.GetPayrollPdfResponse, error) {
	reqData := new(models.PayrollPdfFileRequest)
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

	//  todo
	res := help.CmbSignRequest(string(req), constants.CmbPayrollOldQueryTransRefund, userId, userPrivateKey, sm4Key)

	if res == "" {
		return nil, err
	}

	var resp models.GetPayrollPdfResponse

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		print(err)
	}

	return &resp, err

}

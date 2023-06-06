package payroll

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/ahKevinXy/go-cmb/cmb_errors"
	"github.com/ahKevinXy/go-cmb/constants"
	"github.com/ahKevinXy/go-cmb/help"
	"github.com/ahKevinXy/go-cmb/models"
)

// UnitPayrollPayment
//
//	@Description:   代发
//	@param userId
//	@param sm4PrivateKey
//	@param userPrivateKey
//	@param payMod
//	@param totalInfo
//	@param payList
//	@Author  ahKevinXy
//	@Date  2023-04-13 19:23:11
func UnitPayrollPayment(userId, sm4PrivateKey, userPrivateKey string, payMod []*models.Bb6Busmod, totalInfo []*models.Bb6cdcbhx1, payList []*models.Bb6cdcdlx1) (*models.UnitPayrollPaymentResponse, error) {
	reqData := new(models.UnitPayrollPaymentRequest)
	reqData.Request.Head.Reqid = time.Now().Format("20060102150405000") + strconv.Itoa(time.Now().Nanosecond())
	reqData.Request.Head.Funcode = constants.CmbPayroll
	reqData.Request.Head.Userid = userId
	reqData.Signature.Sigtim = time.Now().Format("20060102150405")
	reqData.Signature.Sigdat = "__signature_sigdat__"
	reqData.Request.Body.Bb6Busmod = payMod
	reqData.Request.Body.Bb6cdcbhx1 = totalInfo
	reqData.Request.Body.Bb6cdcdlx1 = payList
	req, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	res, err := help.CmbSignRequest(string(req), constants.CmbPayroll, userId, userPrivateKey, sm4PrivateKey)
	if err != nil {
		return nil, err
	}
	if res == "" {
		return nil, cmb_errors.SystemError
	}

	var resp models.UnitPayrollPaymentResponse

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, err
	}

	return &resp, err
}

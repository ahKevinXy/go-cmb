package payroll

import (
	"encoding/json"
	"github.com/ahKevinXy/go-cmb/constants"
	"github.com/ahKevinXy/go-cmb/help"
	"github.com/ahKevinXy/go-cmb/models"
	"strconv"
	"time"
)

// UnitPayrollPayment
//  @Description:   代发
//  @param userId
//  @param asePrivateKey
//  @param userPrivateKey
//  @param payMod
//  @param totalInfo
//  @param payList
//  @Author  ahKevinXy
//  @Date  2023-04-13 19:23:11
func UnitPayrollPayment(userId, asePrivateKey, userPrivateKey string, payMod []*models.Bb6Busmod, totalInfo []*models.Bb6Aclakx1, payList []*models.Bb6Aclaky1) (*models.UnitPayrollPaymentResponse, error) {
	reqData := new(models.UnitPayrollPaymentRequest)
	reqData.Request.Head.Reqid = time.Now().Format("20060102150405000") + strconv.Itoa(time.Now().Nanosecond())
	reqData.Request.Head.Funcode = constants.CmbPayroll
	reqData.Request.Head.Userid = userId
	reqData.Signature.Sigtim = time.Now().Format("20060102150405")
	reqData.Signature.Sigdat = "__signature_sigdat__"
	reqData.Request.Body.Bb6Busmod = payMod
	reqData.Request.Body.Bb6Aclakx1 = totalInfo
	reqData.Request.Body.Bb6Aclaky1 = payList
	req, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	//  todo
	res := help.CmbSignRequest(string(req), constants.CmbPayroll, userId, userPrivateKey, asePrivateKey)

	if res == "" {
		return nil, err
	}

	var resp models.UnitPayrollPaymentResponse

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, err
	}

	return &resp, err
}

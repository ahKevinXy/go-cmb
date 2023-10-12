package unit_manager

import (
	"encoding/json"
	"github.com/ahKevinXy/go-cmb/v1/cmb_errors"
	"github.com/ahKevinXy/go-cmb/v1/constants"
	"github.com/ahKevinXy/go-cmb/v1/help"
	"github.com/ahKevinXy/go-cmb/v1/models"
	"strconv"
	"time"
)

// UnitAccountPayIn
//
//	@Description:  子账户内部调账
//	@param userId
//	@param sm4PrivateKey
//	@param userPrivateKey
//	@param accnbr
//	@param dmadbt
//	@param dmacrt
//	@param trxamt
//	@param trxtxt
//	@return *models.UnitAccountPayInResponse
//	@return error
//	@Author  ahKevinXy
//	@Date2023-04-13 17:23:59
func UnitAccountPayIn(userId,
	sm4PrivateKey,
	userPrivateKey,
	accnbr,
	dmadbt,
	dmacrt,
	trxamt,
	trxtxt string) (*models.UnitAccountPayInResponse, error) {
	reqData := new(models.AccountUnitTransInRequest)
	reqData.Request.Head.Reqid = time.Now().Format("20060102150405000") + strconv.Itoa(time.Now().Nanosecond())
	reqData.Request.Head.Funcode = constants.CmbUnitManageAccountPayIn
	reqData.Request.Head.Userid = userId
	reqData.Signature.Sigtim = time.Now().Format("20060102150405")
	reqData.Signature.Sigdat = "__signature_sigdat__"
	reqData.Request.Body.Ntdmatrxx1 = append(reqData.Request.Body.Ntdmatrxx1, &models.Ntdmatrxx1{
		Accnbr: accnbr,
		Dmadbt: dmadbt,
		Dmacrt: dmacrt,
		Trxamt: trxamt,
		Trxtxt: trxtxt,
	})

	req, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	//  todo 判断错误信息
	res, err := help.CmbSignRequest(string(req), constants.CmbUnitManageAccountPayIn, userId, userPrivateKey, sm4PrivateKey)
	if err != nil {
		return nil, err
	}
	if res == "" {
		return nil, cmb_errors.SystemError
	}

	var resp models.UnitAccountPayInResponse

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

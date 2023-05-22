package unit_manager

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/ahKevinXy/go-cmb/cmb_errors"
	"github.com/ahKevinXy/go-cmb/constants"
	"github.com/ahKevinXy/go-cmb/help"
	"github.com/ahKevinXy/go-cmb/models"
)

// UnitAccountPayIn
//
//	@Description:  子账户内部调账
//	@param userId
//	@param sm4Key
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
	sm4Key,
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
	res := help.CmbSignRequest(string(req), constants.CmbUnitManageAccountPayIn, userId, userPrivateKey, sm4Key)

	if res == "" {
		return nil, cmb_errors.SystemError
	}

	var resp models.UnitAccountPayInResponse

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

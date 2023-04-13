package unit_manager

import (
	"encoding/json"
	"github.com/ahKevinXy/go-cmb/constants"
	"github.com/ahKevinXy/go-cmb/help"
	"github.com/ahKevinXy/go-cmb/models"
	"strconv"
	"time"
)

// UnitAccountPayIn
//  @Description:  子账户内部调账
//  @param userId
//  @param asePrivateKey
//  @param userPrivateKey
//  @param accnbr
//  @param dmadbt
//  @param dmacrt
//  @param trxamt
//  @param trxtxt
//  @return *models.UnitAccountPayInResponse
//  @return error
//  @Author  ahKevinXy
//  @Date2023-04-13 17:23:59
func UnitAccountPayIn(userId,
	asePrivateKey,
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
	res := help.CmbSignRequest(string(req), constants.CmbUnitManageAccountPayIn, userId, userPrivateKey, asePrivateKey)

	if res == "" {
		return nil, nil
	}

	var resp models.UnitAccountPayInResponse

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

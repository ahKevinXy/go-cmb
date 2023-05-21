package account

import (
	"encoding/json"
	"github.com/ahKevinXy/go-cmb/cmb_errors"
	"github.com/ahKevinXy/go-cmb/constants"
	"github.com/ahKevinXy/go-cmb/help"
	"github.com/ahKevinXy/go-cmb/models"
	"strconv"
	"time"
)

// MainAccountPaySingle
//  @Description:   单笔对公支付
//  @param userId
//  @param asePrivateKey
//  @param userPrivateKey
//  @param dbAcc
//  @param buscode
//  @param busMode 交易模式
//  @param dmaNbr
//  @param crtAcc
//  @param crtNam
//  @param crtBnk
//  @param crtAdr
//  @param bnkFlg
//  @param trsAmt
//  @param brdNbr
//  @param nusAge
//  @param //
//  @param trsTyp
//  @param busNar
//  @param category
//  @return *models.MainAccountSinglePayResponse
//  @return error
//  @Author  ahKevinXy
//  @Date 2023-04-10 13:56:28
func MainAccountPaySingle(userId,
	asePrivateKey, userPrivateKey,
	dbAcc,
	buscode,
	busMode,
	dmaNbr,
	crtAcc,
	crtNam,
	crtBnk,
	crtAdr,
	bnkFlg,
	trsAmt,
	brdNbr,
	nusAge, // 用途
	yurRef,
	trsTyp,
	busNar string) (*models.MainAccountSinglePayResponse, error) {
	reqData := new(models.MainAccountSinglePayRequest)
	reqData.Request.Head.Reqid = time.Now().Format("20060102150405000") + strconv.Itoa(time.Now().Nanosecond())
	reqData.Request.Head.Funcode = constants.CmbAccountSinglePay
	reqData.Request.Head.Userid = userId
	reqData.Signature.Sigtim = time.Now().Format("20060102150405")
	reqData.Signature.Sigdat = "__signature_sigdat__"
	reqData.Request.Body.Bb1Paybmx1 = append(reqData.Request.Body.Bb1Paybmx1, &models.Bb1Paybmx1{
		BusCod: buscode,
		BusMod: busMode,
	})

	reqData.Request.Body.Bb1Payopx1 = append(reqData.Request.Body.Bb1Payopx1, &models.Bb1Payopx1{
		DbtAcc: dbAcc, //转出账户
		DmaNbr: dmaNbr,
		CrtAcc: crtAcc,
		CrtNam: crtNam,
		CrtBnk: crtBnk,
		CrtAdr: crtAdr,
		TrsAmt: trsAmt,
		BrdNbr: brdNbr,
		CcyNbr: "10",
		NusAge: nusAge,
		YurRef: yurRef,
		TrsTyp: trsTyp,
		BnkFlg: bnkFlg,
		RcvChk: "1",
		BusNar: busNar,
	})

	req, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	res := help.CmbSignRequest(string(req), constants.CmbAccountSinglePay, userId, userPrivateKey, asePrivateKey)

	if res == "" {
		return nil, cmb_errors.SystemError
	}

	var resp models.MainAccountSinglePayResponse

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// MainAccountBatchPay
//  @Description:  批量支付
//  @param userId
//  @param asePrivateKey
//  @param userPrivateKey
//  @param busCode
//  @param busMode
//  @param dtlNbr
//  @param ctnFlg
//  @param ctnSts
//  @param bthNbr
//  @param payList
//  @param category
//  @return *models.MainAccountSinglePayResponse
//  @return error
//  @Author  ahKevinXy
//  @Date 2023-04-10 13:57:21
func MainAccountBatchPay(userId,
	asePrivateKey, userPrivateKey,
	busCode, //业务代码
	busMode, // 支付模式
	dtlNbr,
	ctnFlg,
	ctnSts,
	bthNbr string,

	payList []*models.Bb1paybhx1,
) (*models.MainAccountSinglePayResponse, error) {
	reqData := new(models.MainAccountBatchPayRequest)
	reqData.Request.Head.Reqid = time.Now().Format("20060102150405000") + strconv.Itoa(time.Now().Nanosecond())
	reqData.Request.Head.Funcode = constants.CmbAccountBatchPayQuery
	reqData.Request.Head.Userid = userId
	reqData.Signature.Sigtim = time.Now().Format("20060102150405")
	reqData.Signature.Sigdat = "__signature_sigdat__"
	reqData.Request.Body.Bb1bmdbhx1 = append(reqData.Request.Body.Bb1bmdbhx1, &models.Bb1bmdbhx1{
		BusCod: busCode,
		BusMod: busMode,
		DtlNbr: dtlNbr,
		BthNbr: bthNbr,
		CtnFlg: ctnFlg,
		CtnSts: ctnSts,
	})

	reqData.Request.Body.Bb1paybhx1 = payList

	req, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}
	
	res := help.CmbSignRequest(string(req), constants.CmbAccountBatchPayQuery, userId, userPrivateKey, asePrivateKey)

	if res == "" {
		return nil, cmb_errors.SystemError
	}

	var resp models.MainAccountSinglePayResponse

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

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

// QueryAccountPaymentTransInfo
//  @Description:  企银批量支付批次查询
//  @param userId
//  @param asePrivateKey
//  @param userPrivateKey
//  @param begDat
//  @param endDat
//  @param autStr
//  @param rtnStr
//  @param
//  @return *models.MainAccountSinglePayResponse
//  @return error
//  @Author  ahKevinXy
//  @Date 2023-04-13 16:47:10
func QueryAccountPaymentTransInfo(userId,
	asePrivateKey, userPrivateKey,
	begDat,
	endDat,
	autStr,
	rtnStr string,

) (*models.MainAccountSinglePayResponse, error) {
	reqData := new(models.QueryAccountPaymentTransInfoRequest)
	reqData.Request.Head.Reqid = time.Now().Format("20060102150405000") + strconv.Itoa(time.Now().Nanosecond())
	reqData.Request.Head.Funcode = constants.CmbAccountBatchPayQueryInfo
	reqData.Request.Head.Userid = userId
	reqData.Signature.Sigtim = time.Now().Format("20060102150405")
	reqData.Signature.Sigdat = "__signature_sigdat__"
	reqData.Request.Body.Bb1qrybtx1 = append(reqData.Request.Body.Bb1qrybtx1, &models.Bb1qrybtx1{
		BegDat: begDat,
		EndDat: endDat,
		AutStr: autStr,
		RtnStr: rtnStr,
	})

	req, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	//  todo
	res := help.CmbSignRequest(string(req), constants.CmbAccountBatchPayQueryInfo, userId, userPrivateKey, asePrivateKey)

	if res == "" {
		return nil, cmb_errors.SystemError
	}

	var resp models.MainAccountSinglePayResponse

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// QueryAccountPaymentDetail
//  @Description:   获取交易明细
//  @param userId
//  @param asePrivateKey
//  @param userPrivateKey
//  @param bthNbr
//  @param autStr
//  @param rtnStr
//  @param ctnKey
//  @param
//  @return *models.MainAccountSinglePayResponse
//  @return error
//  @Author  ahKevinXy
//  @Date 2023-04-13 16:51:27
func QueryAccountPaymentDetail(
	userId,
	asePrivateKey, userPrivateKey,
	bthNbr,
	autStr,
	rtnStr, ctnKey string,

) (*models.QueryAccountPaymentDetailResponse, error) {
	reqData := new(models.QueryAccountPaymentDetailRequest)
	reqData.Request.Head.Reqid = time.Now().Format("20060102150405000") + strconv.Itoa(time.Now().Nanosecond())
	reqData.Request.Head.Funcode = constants.CmbAccountBatchPayQueryDetail
	reqData.Request.Head.Userid = userId
	reqData.Signature.Sigtim = time.Now().Format("20060102150405")
	reqData.Signature.Sigdat = "__signature_sigdat__"
	reqData.Request.Body.Bb1qrybdy1 = append(reqData.Request.Body.Bb1qrybdy1, &models.Bb1qrybdy1{
		BthNbr: bthNbr,
		AutStr: autStr,
		CtnKey: ctnKey,
		RtnStr: rtnStr,
	})

	req, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	//  todo
	res := help.CmbSignRequest(string(req), constants.CmbAccountBatchPayQueryDetail, userId, userPrivateKey, asePrivateKey)

	if res == "" {
		return nil, err
	}

	var resp models.QueryAccountPaymentDetailResponse

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

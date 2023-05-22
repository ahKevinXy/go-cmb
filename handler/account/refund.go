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

// QueryAccountPaymentRefund
//  @Description:  跨行退票查询
//  @param userId
//  @param sm2PrivateKey
//  @param userPrivateKey
//  @param bbkNbr 开户行
//  @param bgnDat 开始时间
//  @param endDat 结束时间
//  @param reqNbr 请求批次
//  @param ctnKey 续传key
//  @param rsv50z 保留字
//  @param
//  @return *models.QueryAccountPaymentRefundResponse
//  @return error
//  @Author  ahKevinXy
//  @Date 2023-04-13 16:59:34
func QueryAccountPaymentRefund(
	userId,
	sm2PrivateKey, userPrivateKey,
	bbkNbr,
	bgnDat,
	endDat,
	reqNbr,
	ctnKey,

	rsv50z string,

) (*models.QueryAccountPaymentRefundResponse, error) {
	reqData := new(models.QueryAccountPaymentRefundRequest)
	reqData.Request.Head.Reqid = time.Now().Format("20060102150405000") + strconv.Itoa(time.Now().Nanosecond())
	reqData.Request.Head.Funcode = constants.CmbAccountBatchPayRefund
	reqData.Request.Head.Userid = userId
	reqData.Signature.Sigtim = time.Now().Format("20060102150405")
	reqData.Signature.Sigdat = "__signature_sigdat__"
	reqData.Request.Body.Bb1payqby1 = append(reqData.Request.Body.Bb1payqby1, &models.Bb1payqby1{
		BbkNbr: bbkNbr,
		BgnDat: bgnDat,
		EndDat: endDat,
		ReqNbr: reqNbr,
		CtnKey: ctnKey,
		Rsv50z: rsv50z,
	})

	req, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	//  todo
	res := help.CmbSignRequest(string(req), constants.CmbAccountBatchPayRefund, userId, userPrivateKey, sm2PrivateKey)

	if res == "" {
		return nil, cmb_errors.SystemError
	}

	var resp models.QueryAccountPaymentRefundResponse

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

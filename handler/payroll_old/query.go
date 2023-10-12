package payroll_old

import (
	"encoding/json"
	"github.com/ahKevinXy/go-cmb/v1/constants"
	"github.com/ahKevinXy/go-cmb/v1/help"
	"github.com/ahKevinXy/go-cmb/v1/models"

	"strconv"
	"time"
)

// QueryOldPayRollOrder
//
//	@Description:   获取代发 概要信息
//	@param userId
//	@param sm4PrivateKey
//	@param userPrivateKey
//	@param buscod 业务类型
//	@param busmod 业务模式
//	@param bgndat 起始时间
//	@param enddat 结束时间
//	@param datflg 日期类型 默认 经办  A 经办日期 B  期望日期
//	@param ctnkey 续传key
//	@return *models.QueryOldPayRollOrderResponse
//	@return error
//	@Author  ahKevinXy
//	@Date  2023-06-05 17:35:27
func QueryOldPayRollOrder(userId, sm4PrivateKey, userPrivateKey, buscod, busmod, bgndat, enddat, datflg, ctnkey string) (*models.QueryOldPayRollOrderResponse, error) {
	reqData := new(models.QueryOldPayRollOrderRequest)
	reqData.Request.Head.Reqid = time.Now().Format("20060102150405000") + strconv.Itoa(time.Now().Nanosecond())
	reqData.Request.Head.Funcode = constants.CmbPayrollOldQueryTrans
	reqData.Request.Head.Userid = userId
	reqData.Signature.Sigtim = time.Now().Format("20060102150405")
	reqData.Signature.Sigdat = "__signature_sigdat__"
	reqData.Request.Body.Ntagcinny1 = append(reqData.Request.Body.Ntagcinny1, &models.Ntagcinny1{
		Buscod: buscod,
		Busmod: busmod,
		Bgndat: bgndat,
		Enddat: enddat,
		Datflg: datflg,
		Ctnkey: ctnkey,
	})

	req, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	res, err := help.CmbSignRequest(string(req), constants.CmbPayrollOldQueryTrans, userId, userPrivateKey, sm4PrivateKey)
	if err != nil {
		return nil, err
	}
	if res == "" {
		return nil, err
	}

	var resp models.QueryOldPayRollOrderResponse

	if err := json.Unmarshal([]byte(res), &resp); err != nil {

		return nil, err
	}

	return &resp, err

}

func QueryOldPayRollOrderDetail(userId, sm4PrivateKey, userPrivateKey, reqnbr, bthnbr, trxseq, histag string) (*models.QueryOldPayRollOrderDetailResponse, error) {
	reqData := new(models.QueryOldPayRollOrderDetailRequest)
	reqData.Request.Head.Reqid = time.Now().Format("20060102150405000") + strconv.Itoa(time.Now().Nanosecond())
	reqData.Request.Head.Funcode = constants.CmbPayrollOldQueryTrans
	reqData.Request.Head.Userid = userId
	reqData.Signature.Sigtim = time.Now().Format("20060102150405")
	reqData.Signature.Sigdat = "__signature_sigdat__"
	reqData.Request.Body.Ntagdinfy1 = append(reqData.Request.Body.Ntagdinfy1, &models.Ntagdinfy1{
		Reqnbr: reqnbr,
		Bthnbr: bthnbr,
		Trxseq: trxseq,
		Histag: histag,
	})

	req, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	//  todo
	res, err := help.CmbSignRequest(string(req), constants.CmbPayrollOldQueryTrans, userId, userPrivateKey, sm4PrivateKey)
	if err != nil {
		return nil, err
	}
	if res == "" {
		return nil, err
	}

	var resp models.QueryOldPayRollOrderDetailResponse

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		print(err)
	}

	return &resp, err

}

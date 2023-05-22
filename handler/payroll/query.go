package payroll

import (
	"encoding/json"
	"github.com/ahKevinXy/go-cmb/cmb_errors"
	"github.com/ahKevinXy/go-cmb/constants"
	"github.com/ahKevinXy/go-cmb/help"
	"github.com/ahKevinXy/go-cmb/models"
	"strconv"
	"time"
)

// QueryBatchTransInfo
//  @Description: 批次与明细查询
//  @param userId
//  @param sm2PrivateKey
//  @param userPrivateKey
//  @param buscode 业务编码
//  @param yurref 业务参考号
//  @param bthnbr 续传批次号
//  @param trxseq 续传明细序号
//  @param histga 续传历史查询标志
//  @return *models.UnitPayrollPaymentResponse
//  @return error
//  @Author  ahKevinXy
//  @Date  2023-04-14 14:47:18
func QueryBatchTransInfo(userId, sm2PrivateKey, userPrivateKey, buscode, yurref, bthnbr, trxseq, histga string) (*models.QueryBatchTransInfoResponse, error) {
	reqData := new(models.QueryBatchTransInfoRequest)
	reqData.Request.Head.Reqid = time.Now().Format("20060102150405000") + strconv.Itoa(time.Now().Nanosecond())
	reqData.Request.Head.Funcode = constants.CmbPayrollQuery
	reqData.Request.Head.Userid = userId
	reqData.Signature.Sigtim = time.Now().Format("20060102150405")
	reqData.Signature.Sigdat = "__signature_sigdat__"
	reqData.Request.Body.Bb6bpdqyy1 = append(reqData.Request.Body.Bb6bpdqyy1, &models.Bb6bpdqyy1{
		Buscod: buscode,
		Yurref: yurref,
		Bthnbr: bthnbr,
		Trxseq: trxseq,
		Histag: histga,
		Rsv50z: "",
	})
	req, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	//  todo
	res := help.CmbSignRequest(string(req), constants.CmbPayrollQuery, userId, userPrivateKey, sm2PrivateKey)

	if res == "" {
		return nil, cmb_errors.SystemError
	}

	var resp models.QueryBatchTransInfoResponse

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, err
	}

	return &resp, err
}

// QueryBatchTransList
//  @Description: 代发批次查询
//  @param userId
//  @param sm2PrivateKey
//  @param userPrivateKey
//  @param buscode 业务编码
//  @param yurref 业务参考号
//  @param bgndat 起始日期
//  @param enddat 结束日期
//  @param cntkey 续传键值
//  @param dattyp 日期类型
//  @return *models.QueryBatchTransListResponse
//  @return error
//  @Author  ahKevinXy
//  @Date  2023-04-14 14:57:28
func QueryBatchTransList(userId, sm2PrivateKey, userPrivateKey, buscode, yurref, bgndat, enddat, cntkey, dattyp string) (*models.QueryBatchTransListResponse, error) {
	reqData := new(models.QueryBatchTransListRequest)
	reqData.Request.Head.Reqid = time.Now().Format("20060102150405000") + strconv.Itoa(time.Now().Nanosecond())
	reqData.Request.Head.Funcode = constants.CmbPayrollQueryBatchList
	reqData.Request.Head.Userid = userId
	reqData.Signature.Sigtim = time.Now().Format("20060102150405")
	reqData.Signature.Sigdat = "__signature_sigdat__"
	reqData.Request.Body.Bb6bthqyy1 = append(reqData.Request.Body.Bb6bthqyy1, &models.Bb6bthqyy1{
		Buscod: buscode,
		Yurref: yurref,
		Bgndat: bgndat,
		Enddat: enddat,
		Cntkey: cntkey,
		Dattyp: dattyp,
	})
	req, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	//  todo
	res := help.CmbSignRequest(string(req), constants.CmbPayrollQueryBatchList, userId, userPrivateKey, sm2PrivateKey)

	if res == "" {
		return nil, cmb_errors.SystemError
	}

	var resp models.QueryBatchTransListResponse

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, err
	}

	return &resp, err
}

// QueryPayrollTransDetail
//  @Description: 代发明细查询
//  @param userId
//  @param sm2PrivateKey
//  @param userPrivateKey
//  @param reqnbr
//  @param bthnbr
//  @param trxseq
//  @param histag
//  @return *models.QueryBatchTransListResponse
//  @return error
//  @Author  ahKevinXy
//  @Date  2023-04-14 15:02:42
func QueryPayrollTransDetail(userId, sm2PrivateKey, userPrivateKey, reqnbr, bthnbr, trxseq, histag string) (*models.QueryPayrollTransDetailResponse, error) {
	reqData := new(models.QueryPayrollTransDetailRequest)
	reqData.Request.Head.Reqid = time.Now().Format("20060102150405000") + strconv.Itoa(time.Now().Nanosecond())
	reqData.Request.Head.Funcode = constants.CmbPayrollQueryDetail
	reqData.Request.Head.Userid = userId
	reqData.Signature.Sigtim = time.Now().Format("20060102150405")
	reqData.Signature.Sigdat = "__signature_sigdat__"
	reqData.Request.Body.Bb6dtlqyy1 = append(reqData.Request.Body.Bb6dtlqyy1, &models.Bb6dtlqyy1{
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
	res := help.CmbSignRequest(string(req), constants.CmbPayrollQueryDetail, userId, userPrivateKey, sm2PrivateKey)

	if res == "" {
		return nil, cmb_errors.SystemError
	}

	var resp models.QueryPayrollTransDetailResponse

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, err
	}

	return &resp, err
}

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

// QueryPayrollStatement
//
//	@Description: 代发明细对账单查询请求
//	@param userId
//	@param sm4Key
//	@param userPrivateKey
//	@param payeac 付款账户
//	@param begdat 开始日期
//	@param enddat 结束日期
//	@param buscod 业务类型
//	@param busmod 业务模式
//	@param eacnam 收方户名
//	@param ptyref 业务参考号
//	@param trxsrl 流水号
//	@param minamt 最小金额
//	@param maxamt 最大金额
//	@param prtmod 打印模式
//	@param begidx 查询标记
//	@return *models.QueryPayrollStatementResponse
//	@return error
//	@Author  ahKevinXy
//	@Date  2023-04-14 17:28:32
func QueryPayrollStatement(userId, sm4Key, userPrivateKey, payeac, begdat, enddat, buscod, busmod, eacnam, ptyref, trxsrl, minamt, maxamt, prtmod, begidx string) (*models.QueryPayrollStatementResponse, error) {
	reqData := new(models.QueryPayrollStatementRequest)
	reqData.Request.Head.Reqid = time.Now().Format("20060102150405000") + strconv.Itoa(time.Now().Nanosecond())
	reqData.Request.Head.Funcode = constants.CmbPayrollStatement
	reqData.Request.Head.Userid = userId
	reqData.Signature.Sigtim = time.Now().Format("20060102150405")
	reqData.Signature.Sigdat = "__signature_sigdat__"
	reqData.Request.Body.Payeac = payeac
	reqData.Request.Body.Begdat = begdat
	reqData.Request.Body.Enddat = enddat
	reqData.Request.Body.Buscod = buscod
	reqData.Request.Body.Busmod = busmod
	reqData.Request.Body.Eacnam = eacnam
	reqData.Request.Body.Ptyref = ptyref
	reqData.Request.Body.Trxsrl = trxsrl
	reqData.Request.Body.Minamt = minamt
	reqData.Request.Body.Maxamt = maxamt
	reqData.Request.Body.Prtmod = prtmod
	reqData.Request.Body.Begidx = begidx

	req, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	//  todo
	res := help.CmbSignRequest(string(req), constants.CmbPayrollStatement, userId, userPrivateKey, sm4Key)

	if res == "" {
		return nil, cmb_errors.SystemError
	}

	var resp models.QueryPayrollStatementResponse

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, err
	}

	return &resp, err
}

// QueryPayrollStatementDownloadUrl
//
//	@Description:   获取回单地址
//	@param userId
//	@param sm4Key
//	@param userPrivateKey
//	@param taskid 查询ID
//	@return *models.QueryBatchTransListResponse
//	@return error
//	@Author  ahKevinXy
//	@Date  2023-04-14 17:33:31
func QueryPayrollStatementDownloadUrl(userId, sm4Key, userPrivateKey, taskid string) (*models.QueryBatchTransListResponse, error) {
	reqData := new(models.QueryPayrollStatementDownloadUrlRequest)
	reqData.Request.Head.Reqid = time.Now().Format("20060102150405000") + strconv.Itoa(time.Now().Nanosecond())
	reqData.Request.Head.Funcode = constants.CmbPayrollStatementDownloadUrl
	reqData.Request.Head.Userid = userId
	reqData.Signature.Sigtim = time.Now().Format("20060102150405")
	reqData.Signature.Sigdat = "__signature_sigdat__"
	reqData.Request.Body.Taskid = taskid
	req, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	//  todo
	res := help.CmbSignRequest(string(req), constants.CmbPayrollStatementDownloadUrl, userId, userPrivateKey, sm4Key)

	if res == "" {
		return nil, cmb_errors.SystemError
	}

	var resp models.QueryBatchTransListResponse

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, err
	}

	return &resp, err
}

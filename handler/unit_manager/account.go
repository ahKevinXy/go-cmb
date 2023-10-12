package unit_manager

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/ahKevinXy/go-cmb/v1/cmb_errors"
	"github.com/ahKevinXy/go-cmb/v1/constants"
	"github.com/ahKevinXy/go-cmb/v1/help"
	"github.com/ahKevinXy/go-cmb/v1/models"
)

// 记账单元 管理

// CloseUnitAccount
//
//	@Description:  关闭记账单元
//	@param userId
//	@param sm4PrivateKey
//	@param userPrivateKey
//	@param accnbr 账户
//	@param dmanbr 子单元账号
//	@param busMod 业务模式
//	@param
//	@return *models.CloseUnitAccountResponse
//	@return error
//	@Author  ahKevinXy
//	@Date 2023-04-13 17:46:41
func CloseUnitAccount(
	userId,
	sm4PrivateKey, userPrivateKey,
	accnbr,
	bbknbr,
	dmanbr,
	busMod,
	yurref string,

) (*models.CloseUnitAccountResponse, error) {
	reqData := new(models.AccountCloseUnitRequest)
	reqData.Request.Head.Reqid = time.Now().Format("20060102150405000") + strconv.Itoa(time.Now().Nanosecond())
	reqData.Request.Head.Funcode = constants.CmbUnitManageCloseAccountV2
	reqData.Request.Head.Userid = userId
	reqData.Signature.Sigtim = time.Now().Format("20060102150405")
	reqData.Signature.Sigdat = "__signature_sigdat__"

	reqData.Request.Body.Ntbusmody = append(reqData.Request.Body.Ntbusmody, &models.Ntbusmody{Busmod: busMod})
	reqData.Request.Body.Ntdumdltx1 = append(reqData.Request.Body.Ntdumdltx1, &models.Ntdumdltx1{
		Inbacc: accnbr,
		Bbknbr: bbknbr,
	})
	reqData.Request.Body.Ntdumdltx2 = append(reqData.Request.Body.Ntdumdltx2, &models.Ntdumdltx2{
		Dyanbr: dmanbr,
		Yurref: yurref,
	})

	req, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	//  todo
	res, err := help.CmbSignRequest(string(req), constants.CmbUnitManageCloseAccountV2, userId, userPrivateKey, sm4PrivateKey)
	if err != nil {
		return nil, err
	}
	if res == "" {
		return nil, cmb_errors.SystemError
	}

	var resp models.CloseUnitAccountResponse

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		print(err)
		return nil, err
	}
	return &resp, err
}

// QueryUnitAccountInfo
//
//	@Description:  获取账户信息 可以获取实时余额
//	@param userId
//	@param sm4PrivateKey
//	@param userPrivateKey
//	@param accnbr 账号
//	@param dmanbr 子单元
//	@return *models.AccountUnitInfoResponse
//	@return error
//	@Author  ahKevinXy
//	@Date  2023-04-13 17:51:01
func QueryUnitAccountInfo(userId, sm4PrivateKey, userPrivateKey, accnbr, dmanbr string) (*models.AccountUnitInfoResponse, error) {
	reqData := new(models.AccountUnitInfoRequest)
	reqData.Request.Head.Reqid = time.Now().Format("20060102150405000") + strconv.Itoa(time.Now().Nanosecond())
	reqData.Request.Head.Funcode = constants.CmbUnitManageAccountQueryV1
	reqData.Request.Head.Userid = userId
	reqData.Signature.Sigtim = time.Now().Format("20060102150405")
	reqData.Signature.Sigdat = "__signature_sigdat__"
	reqData.Request.Body.Ntdmabalx = append(reqData.Request.Body.Ntdmabalx, &models.Ntdmabalx{
		Accnbr: accnbr,
		Dmanbr: dmanbr,
	})

	req, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	//  todo
	res, err := help.CmbSignRequest(string(req), constants.CmbUnitManageAccountQueryV1, userId, userPrivateKey, sm4PrivateKey)
	if err != nil {
		return nil, err
	}
	if res == "" {
		return nil, cmb_errors.SystemError
	}

	var resp models.AccountUnitInfoResponse

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, err
	}
	//fmt.Println(res)
	return &resp, err
}

// QueryUnitAccountBalanceHistory
//
//	@Description:   获取子账户历史余额
//	@param userId
//	@param sm4PrivateKey
//	@param userPrivateKey
//	@param accnbr //账户
//	@param bbknbr  分行号
//	@param qrydat 时间
//	@param dmanbr 续传子单元
//	@return *models.QueryUnitAccountBalanceHistoryResponse
//	@return error
//	@Author  ahKevinXy
//	@Date  2023-05-18 16:56:54
func QueryUnitAccountBalanceHistory(userId, sm4PrivateKey, userPrivateKey, accnbr, bbknbr, qrydat, dmanbr string) (*models.QueryUnitAccountBalanceHistoryResponse, error) {
	reqData := new(models.QueryUnitAccountBalanceHistoryRequest)
	reqData.Request.Head.Reqid = time.Now().Format("20060102150405000") + strconv.Itoa(time.Now().Nanosecond())
	reqData.Request.Head.Funcode = constants.CmbUnitAllHistoryBalanceV2
	reqData.Request.Head.Userid = userId
	reqData.Signature.Sigtim = time.Now().Format("20060102150405")
	reqData.Signature.Sigdat = "__signature_sigdat__"
	reqData.Request.Body.Ntdmahbdx1 = append(reqData.Request.Body.Ntdmahbdx1, &models.Ntdmahbdx1{
		Accnbr: accnbr,
		Qrydat: qrydat,
		Bbknbr: bbknbr,
		Dmanbr: dmanbr,
	})

	req, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	//  todo
	res, err := help.CmbSignRequest(string(req), constants.CmbUnitAllHistoryBalanceV2, userId, userPrivateKey, sm4PrivateKey)
	if err != nil {
		return nil, err
	}
	if res == "" {
		return nil, cmb_errors.SystemError
	}

	var resp models.QueryUnitAccountBalanceHistoryResponse

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, err
	}
	//fmt.Println(res)
	return &resp, err
}

// QueryUnitAccountSingleBalanceHistory
//
//	@Description: 获取子单元 余额列表
//	@param userId
//	@param sm4PrivateKey
//	@param userPrivateKey
//	@param bbknbr 分行号
//	@param accnbr 账号
//	@param dmanbr 子单元
//	@param begdat 开始时间
//	@param enddat 结束时间
//	@return *models.QueryUnitAccountSingleBalanceHistoryResponse
//	@return error
//	@Author  ahKevinXy
//	@Date  2023-05-19 18:02:13
func QueryUnitAccountSingleBalanceHistory(userId, sm4PrivateKey, userPrivateKey, accnbr, bbknbr, dmanbr, begdat, enddat string) (*models.QueryUnitAccountSingleBalanceHistoryResponse, error) {
	reqData := new(models.QueryUnitAccountSingleBalanceHistoryRequest)
	reqData.Request.Head.Reqid = time.Now().Format("20060102150405000") + strconv.Itoa(time.Now().Nanosecond())
	reqData.Request.Head.Funcode = constants.CmbUnitHistoryBalanceV2
	reqData.Request.Head.Userid = userId
	reqData.Signature.Sigtim = time.Now().Format("20060102150405")
	reqData.Signature.Sigdat = "__signature_sigdat__"
	reqData.Request.Body.Ntdmahadx1 = append(reqData.Request.Body.Ntdmahadx1, &models.Ntdmahadx1{
		Accnbr: accnbr,

		Bbknbr: bbknbr,
		Dmanbr: dmanbr,
		Begdat: begdat,
		Enddat: enddat,
	})

	req, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	//  todo
	res, err := help.CmbSignRequest(string(req), constants.CmbUnitHistoryBalanceV2, userId, userPrivateKey, sm4PrivateKey)
	if err != nil {
		return nil, err
	}
	if res == "" {
		return nil, cmb_errors.SystemError
	}

	var resp models.QueryUnitAccountSingleBalanceHistoryResponse

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, err
	}

	return &resp, err
}

// AccountSetWhitePay
// @Description:   设置白名单
// @param userId
// @param sm4PrivateKey
// @param userPrivateKey
// @param busmod 业务模式
// @param accnbr 账户
// @param bbknbr 开户行
// @param dumnbr 记账子单元编号
// @param rltnam 关联户名
// @param rltacc 关联账户
// @param chktyp 白名单校验类型 1:帐号 2:户名
// @param rcvtyp 入账方式 N：非关联收款入账默认子单元 R：非关联收款拒绝入账
// @param yurref 批次号
// @return *models.AccountSetWhitePayResponse
// @return error
// @Author  colornote
// @Date  2023-06-27 10:00:00
func AccountSetWhitePay(userId,
	sm4PrivateKey,
	userPrivateKey,
	accnbr,
	busmod,
	bbknbr,
	dumnbr,
	rltnam,
	rltacc,
	chktyp,
	rcvtyp,
	yurref string) (*models.AccountSetWhitePayResponse, error) {
	reqData := new(models.AccountSetWhitePayRequest)
	reqData.Request.Head.Reqid = time.Now().Format("20060102150405000") + strconv.Itoa(time.Now().Nanosecond())
	reqData.Request.Head.Funcode = constants.CmbUnitManageAccountSetWhitePay
	reqData.Request.Head.Userid = userId
	reqData.Signature.Sigtim = time.Now().Format("20060102150405")
	reqData.Signature.Sigdat = "__signature_sigdat__"
	reqData.Request.Body.Ntbusmody = append(reqData.Request.Body.Ntbusmody, &models.Ntbusmody{
		Busmod: busmod,
	})
	reqData.Request.Body.Ntdumrlax1 = append(reqData.Request.Body.Ntdumrlax1, &models.Ntdumrlax1{
		Bbknbr: bbknbr,
		Accnbr: accnbr,
		Dumnbr: dumnbr,
		Rltnam: rltnam,
		Rltacc: rltacc,
		Chktyp: chktyp,
		Rcvtyp: rcvtyp,
		Yurref: yurref,
	})

	req, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	//  todo
	res, err := help.CmbSignRequest(string(req), constants.CmbUnitManageAccountSetWhitePay, userId, userPrivateKey, sm4PrivateKey)
	if err != nil {
		return nil, err
	}
	if res == "" {
		return nil, cmb_errors.SystemError
	}

	var resp models.AccountSetWhitePayResponse

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, err
	}

	return &resp, err

}

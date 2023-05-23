package unit_manager

import (
	"encoding/json"
	"github.com/ahKevinXy/go-cmb/cmb_errors"
	"github.com/ahKevinXy/go-cmb/constants"
	"github.com/ahKevinXy/go-cmb/help"
	"github.com/ahKevinXy/go-cmb/models"
	"strconv"
	"time"
)

// 记账单元 管理

// AddUnitAccountV1
//  @Description:  新增记账单元
//  @param userId
//  @param sm4PrivateKey
//  @param userPrivateKey
//  @param accnbr
//  @param dmanam
//  @param dmanbr
//  @param
//  @return *models.AddUnitAccountV1Response
//  @return error
//  @Author  ahKevinXy
//  @Date 2023-04-13 17:28:18
func AddUnitAccountV1(
	userId, sm4PrivateKey, userPrivateKey, accnbr, dmanam, dmanbr string,

) (*models.AddUnitAccountV1Response, error) {
	reqData := new(models.AccountAddUnitV1Request)
	reqData.Request.Head.Reqid = time.Now().Format("20060102150405000") + strconv.Itoa(time.Now().Nanosecond())
	reqData.Request.Head.Funcode = constants.CmbUnitManageAddAccountV1
	reqData.Request.Head.Userid = userId
	reqData.Signature.Sigtim = time.Now().Format("20060102150405")
	reqData.Signature.Sigdat = "__signature_sigdat__"
	reqData.Request.Body.Ntdmaaddx = append(reqData.Request.Body.Ntdmaaddx, &models.Ntdmaaddx{
		Accnbr: accnbr,
		Dmanam: dmanam,
		Dmanbr: dmanbr,
	})

	req, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	//  todo
	res := help.CmbSignRequest(string(req), constants.CmbUnitManageAddAccountV1, userId, userPrivateKey, sm4PrivateKey)

	if res == "" {
		return nil, cmb_errors.SystemError
	}

	var resp models.AddUnitAccountV1Response

	if err := json.Unmarshal([]byte(res), &resp); err != nil {

		return nil, err
	}

	return &resp, nil
}

// CloseUnitAccount
//  @Description:  关闭记账单元
//  @param userId
//  @param sm4PrivateKey
//  @param userPrivateKey
//  @param accnbr 账户
//  @param dmanbr 子单元账号
//  @param busMod 业务模式
//  @param
//  @return *models.CloseUnitAccountResponse
//  @return error
//  @Author  ahKevinXy
//  @Date 2023-04-13 17:46:41
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
	reqData.Request.Head.Funcode = constants.CmbUnitManageCloseAccount
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
	res := help.CmbSignRequest(string(req), constants.CmbUnitManageCloseAccount, userId, userPrivateKey, sm4PrivateKey)

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
//  @Description:  获取账户信息 可以获取实时余额
//  @param userId
//  @param sm4PrivateKey
//  @param userPrivateKey
//  @param accnbr 账号
//  @param dmanbr 子单元
//  @return *models.AccountUnitInfoResponse
//  @return error
//  @Author  ahKevinXy
//  @Date  2023-04-13 17:51:01
func QueryUnitAccountInfo(userId, sm4PrivateKey, userPrivateKey, accnbr, dmanbr string) (*models.AccountUnitInfoResponse, error) {
	reqData := new(models.AccountUnitInfoRequest)
	reqData.Request.Head.Reqid = time.Now().Format("20060102150405000") + strconv.Itoa(time.Now().Nanosecond())
	reqData.Request.Head.Funcode = constants.CmbUnitManageAccountQueryV2
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
	res := help.CmbSignRequest(string(req), constants.CmbUnitManageAccountQueryV2, userId, userPrivateKey, sm4PrivateKey)

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
//  @Description:   获取子账户历史余额
//  @param userId
//  @param sm4PrivateKey
//  @param userPrivateKey
//  @param accnbr //账户
//  @param bbknbr  分行号
//  @param qrydat 时间
//  @param dmanbr 续传子单元
//  @return *models.QueryUnitAccountBalanceHistoryResponse
//  @return error
//  @Author  ahKevinXy
//  @Date  2023-05-18 16:56:54
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
	res := help.CmbSignRequest(string(req), constants.CmbUnitAllHistoryBalanceV2, userId, userPrivateKey, sm4PrivateKey)

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
//  @Description: 获取子单元 余额列表
//  @param userId
//  @param sm4PrivateKey
//  @param userPrivateKey
//  @param bbknbr 分行号
//  @param accnbr 账号
//  @param dmanbr 子单元
//  @param begdat 开始时间
//  @param enddat 结束时间
//  @return *models.QueryUnitAccountSingleBalanceHistoryResponse
//  @return error
//  @Author  ahKevinXy
//  @Date  2023-05-19 18:02:13
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
	res := help.CmbSignRequest(string(req), constants.CmbUnitHistoryBalanceV2, userId, userPrivateKey, sm4PrivateKey)

	if res == "" {
		return nil, cmb_errors.SystemError
	}

	var resp models.QueryUnitAccountSingleBalanceHistoryResponse

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, err
	}

	return &resp, err
}

// UpdateUnitAccountV1
//  @Description:   更新记账单元(记账)
//  @param userId
//  @param sm4PrivateKey
//  @param userPrivateKey
//  @param busmod 业务模式
//  @param accnbr 账户
//  @param bbknbr 开户行
//  @param dmanbr 子单元
//  @param dmanam 子单元名称
//  @param ovrctl 额度控制标志
//  @param ballmt 余额上限
//  @param yurref 批次号
//  @param bcktyp 退票处理
//  @param clstyp 余额为零是否关闭
//  @param lmtflg 额度标志
//  @return *models.QueryUnitAccountSingleBalanceHistoryResponse
//  @return error
//  @Author  ahKevinXy
//  @Date  2023-05-23 18:35:29
func UpdateUnitAccountV1(userId,
	sm4PrivateKey,
	userPrivateKey,
	accnbr,
	busmod,

	bbknbr,
	dmanbr,
	dmanam,
	ovrctl,
	ballmt,
	yurref,
	lmtflg,
	bcktyp,
	clstyp string) (*models.UpdateUnitAccountV1Response, error) {
	reqData := new(models.UpdateUnitAccountV1Request)
	reqData.Request.Head.Reqid = time.Now().Format("20060102150405000") + strconv.Itoa(time.Now().Nanosecond())
	reqData.Request.Head.Funcode = constants.CmbUnitManageUpdateAccountV1
	reqData.Request.Head.Userid = userId
	reqData.Signature.Sigtim = time.Now().Format("20060102150405")
	reqData.Signature.Sigdat = "__signature_sigdat__"
	reqData.Request.Body.Ntbusmody = append(reqData.Request.Body.Ntbusmody, &models.Ntbusmody{
		Busmod: busmod,
	})
	reqData.Request.Body.Ntdmamntx1 = append(reqData.Request.Body.Ntdmamntx1, &models.Ntdmamntx1{
		Bbknbr: bbknbr,
		Accnbr: accnbr,
		Dmanbr: dmanbr,
		Dmanam: dmanam,
		Ovrctl: ovrctl,
		Clstyp: clstyp,
		Yurref: yurref,
		Ballmt: ballmt,
		Lmtflg: lmtflg,
		Bcktyp: bcktyp,
	})
	req, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	//  todo
	res := help.CmbSignRequest(string(req), constants.CmbUnitManageUpdateAccountV1, userId, userPrivateKey, sm4PrivateKey)

	if res == "" {
		return nil, cmb_errors.SystemError
	}

	var resp models.UpdateUnitAccountV1Response

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, err
	}

	return &resp, err
}

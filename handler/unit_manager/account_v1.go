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

// 区分交易管家 管理和记账

// 交易管家 管理

// UpdateUnitAccountV1
//
//	@Description:   更新记账单元(记账)
//	@param userId
//	@param sm4PrivateKey
//	@param userPrivateKey
//	@param busmod 业务模式
//	@param accnbr 账户
//	@param bbknbr 开户行
//	@param dmanbr 子单元
//	@param dmanam 子单元名称
//	@param ovrctl 额度控制标志
//	@param ballmt 余额上限
//	@param yurref 批次号
//	@param bcktyp 退票处理
//	@param clstyp 余额为零是否关闭
//	@param lmtflg 额度标志
//	@return *models.QueryUnitAccountSingleBalanceHistoryResponse
//	@return error
//	@Author  ahKevinXy
//	@Date  2023-05-23 18:35:29
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
	res, err := help.CmbSignRequest(string(req), constants.CmbUnitManageUpdateAccountV1, userId, userPrivateKey, sm4PrivateKey)
	if err != nil {
		return nil, err
	}
	if res == "" {
		return nil, cmb_errors.SystemError
	}

	var resp models.UpdateUnitAccountV1Response

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, err
	}

	return &resp, err
}

// AddUnitAccountV1
//
//	@Description:  新增记账单元
//	@param userId
//	@param sm4PrivateKey
//	@param userPrivateKey
//	@param accnbr
//	@param dmanam
//	@param dmanbr
//	@param
//	@return *models.AddUnitAccountV1Response
//	@return error
//	@Author  ahKevinXy
//	@Date 2023-04-13 17:28:18
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
	res, err := help.CmbSignRequest(string(req), constants.CmbUnitManageAddAccountV1, userId, userPrivateKey, sm4PrivateKey)
	if err != nil {
		return nil, err
	}
	if res == "" {
		return nil, cmb_errors.SystemError
	}

	var resp models.AddUnitAccountV1Response

	if err := json.Unmarshal([]byte(res), &resp); err != nil {

		return nil, err
	}

	return &resp, nil
}

func CloseUnitAccountV1(userId, sm4PrivateKey, userPrivateKey, accnbr, dmanbr string) (*models.CloseUnitAccountV1Response, error) {
	reqData := new(models.AccountCloseUnitV1Request)
	reqData.Request.Head.Reqid = time.Now().Format("20060102150405000") + strconv.Itoa(time.Now().Nanosecond())
	reqData.Request.Head.Funcode = constants.CmbUnitManageAccountQueryV2
	reqData.Request.Head.Userid = userId
	reqData.Signature.Sigtim = time.Now().Format("20060102150405")
	reqData.Signature.Sigdat = "__signature_sigdat__"
	reqData.Request.Body.Ntdmadltx2 = append(reqData.Request.Body.Ntdmadltx2, &models.Ntdmadltx2{
		Dmanbr: dmanbr,
	})
	reqData.Request.Body.Ntdmadltx1 = append(reqData.Request.Body.Ntdmadltx1, &models.Ntdmadltx1{Accnbr: accnbr})

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

	var resp models.CloseUnitAccountV1Response

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, err
	}

	return &resp, err
}

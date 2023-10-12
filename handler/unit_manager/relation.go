package unit_manager

import (
	"encoding/json"

	"github.com/ahKevinXy/go-cmb/cmb_errors"
	"github.com/ahKevinXy/go-cmb/constants"
	"github.com/ahKevinXy/go-cmb/help"
	"github.com/ahKevinXy/go-cmb/models"
)

// SetUnitAccountRelation
//
//	@Description: 设置记账子单元关联付款方信息
//	@param userId
//	@param sm4PrivateKey
//	@param userPrivateKey
//	@param busmod 业务模式
//	@param bbknbr 支行号
//	@param accnbr 账号
//	@param dmanbr 记账子单元编号
//	@param tlyopr Y：非关联账号入对应记账子单元账 N：非关联账号入默认记账子单元账 R：非关联账号拒绝入账
//	@param dbtacc 付款方账号
//	@param yurref 己方提供的唯一编号
//
// @return *models.SetUnitAccountRelationResponse
// @return error
// @Author  darwintree
// @Date  2023-05-24 16:00:00
func SetUnitAccountRelation(
	userId, sm4PrivateKey, userPrivateKey, busmod, bbknbr, accnbr, dmanbr, tlyopr, dbtacc, yurref string,
) (*models.SetUnitAccountRelationResponse, error) {
	funcode := constants.CmbUnitManageSetRelationV2
	type reqType = models.SetUnitAccountRelationRequest
	type resType = models.SetUnitAccountRelationResponse

	reqData := new(reqType)
	reqData.Request.Head.Reqid = help.GetReqidString()
	reqData.Request.Head.Funcode = funcode
	reqData.Request.Head.Userid = userId
	reqData.Signature.Sigtim = help.GetSigtimString()
	reqData.Signature.Sigdat = "__signature_sigdat__"
	reqData.Request.Body.Ntbusmody = append(reqData.Request.Body.Ntbusmody, &models.Ntbusmody{Busmod: busmod})
	reqData.Request.Body.Ntdmarltx1 = append(reqData.Request.Body.Ntdmarltx1, &models.Ntdmarltx1{
		Bbknbr: bbknbr,
		Accnbr: accnbr,
		Dmanbr: dmanbr,
		Tlyopr: tlyopr,
		Dbtacc: dbtacc,
		Yurref: yurref,
	})

	req, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	res, err := help.CmbSignRequest(string(req), funcode, userId, userPrivateKey, sm4PrivateKey)
	if err != nil {
		return nil, err
	}
	if res == "" {
		return nil, cmb_errors.SystemError
	}

	var resp resType

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateUnitAccountRelation
//
//	@Description:  更新管理记账单元模式
//	@param userId
//	@param sm4PrivateKey
//	@param userPrivateKey
//	@param bbknbr 银联号
//	@param accnbr 账号
//	@param dmanbr 记账子单元
//	@param tlyopr 模式
//	@param
//	@return *models.SetUnitAccountRelationResponse
//	@return error
//	@Author  ahKevinXy
//	@Date  2023-10-12 12:45:10
func UpdateUnitAccountRelation(
	userId,
	sm4PrivateKey,
	userPrivateKey,
	bbknbr,
	accnbr,
	dmanbr,
	tlyopr string,
) (*models.UpdateUnitAccountRelationResponse, error) {
	funCode := constants.CmbUnitManageModRelationV2
	type reqType = models.UpdateUnitAccountRelationRequest
	type resType = models.UpdateUnitAccountRelationResponse

	reqData := new(reqType)
	reqData.Request.Head.Reqid = help.GetReqidString()
	reqData.Request.Head.Funcode = funCode
	reqData.Request.Head.Userid = userId
	reqData.Signature.Sigtim = help.GetSigtimString()
	reqData.Signature.Sigdat = "__signature_sigdat__"

	reqData.Request.Body.Ntdmatmnx1 = append(reqData.Request.Body.Ntdmatmnx1, &models.Ntdmatmnx1{
		Bbknbr: bbknbr,
		Accnbr: accnbr,
		Dmanbr: dmanbr,
		Tlyopr: tlyopr,
	})

	req, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	res, err := help.CmbSignRequest(string(req), funCode, userId, userPrivateKey, sm4PrivateKey)
	if err != nil {
		return nil, err
	}
	if res == "" {
		return nil, cmb_errors.SystemError
	}

	var resp resType

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// DelUnitAccountRelation
//
//	@Description:  删除关联支付账户
//	@param userId
//	@param sm4PrivateKey
//	@param userPrivateKey
//	@param bbknbr 银联号
//	@param accnbr 账号
//	@param dmanbr 记账单元
//	@param
//	@return *models.DelUnitAccountRelationResponse
//	@return error
//	@Author  ahKevinXy
//	@Date  2023-10-12 12:31:28
func DelUnitAccountRelation(
	userId,
	sm4PrivateKey,
	userPrivateKey,
	bbknbr, accnbr, dmanbr, dbtacc string,
) (*models.DelUnitAccountRelationResponse, error) {
	funCode := constants.CmbUnitManageDelRelationV2
	type reqType = models.DelUnitAccountRelationRequest
	type resType = models.DelUnitAccountRelationResponse

	reqData := new(reqType)
	reqData.Request.Head.Reqid = help.GetReqidString()
	reqData.Request.Head.Funcode = funCode
	reqData.Request.Head.Userid = userId
	reqData.Signature.Sigtim = help.GetSigtimString()
	reqData.Signature.Sigdat = "__signature_sigdat__"
	reqData.Request.Body.Ntdmarldx1 = append(reqData.Request.Body.Ntdmarldx1, &models.Ntdmarldx1{
		Bbknbr: bbknbr,
		Accnbr: accnbr,
		Dmanbr: dmanbr,
		Dbtacc: dbtacc,
	})

	req, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	res, err := help.CmbSignRequest(string(req), funCode, userId, userPrivateKey, sm4PrivateKey)
	if err != nil {
		return nil, err
	}
	if res == "" {
		return nil, cmb_errors.SystemError
	}

	var resp resType

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// QueryUnitAccountRelation
//
//	@Description:  查询关联账户
//	@param userId
//	@param sm4PrivateKey
//	@param userPrivateKey
//	@param bbknbr
//	@param accnbr
//	@param dmanbr
//	@param cntkey
//	@param
//	@return *models.QueryUnitAccountRelationResponse
//	@return error
//	@Author  ahKevinXy
//	@Date  2023-10-12 12:59:38
func QueryUnitAccountRelation(
	userId,
	sm4PrivateKey,
	userPrivateKey,
	bbknbr, accnbr, dmanbr, cntkey string,
) (*models.QueryUnitAccountRelationResponse, error) {
	funCode := constants.CmbUnitManageQueryRelationV2
	type reqType = models.QueryUnitAccountRelationRequest
	type resType = models.QueryUnitAccountRelationResponse

	reqData := new(reqType)
	reqData.Request.Head.Reqid = help.GetReqidString()
	reqData.Request.Head.Funcode = funCode
	reqData.Request.Head.Userid = userId
	reqData.Signature.Sigtim = help.GetSigtimString()
	reqData.Signature.Sigdat = "__signature_sigdat__"
	reqData.Request.Body.Ntdmarlqy1 = append(reqData.Request.Body.Ntdmarlqy1, &models.Ntdmarlqy1{
		Bbknbr: bbknbr,
		Accnbr: accnbr,
		Dmanbr: dmanbr,
		Cntkey: cntkey,
	})

	req, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	res, err := help.CmbSignRequest(string(req), funCode, userId, userPrivateKey, sm4PrivateKey)
	if err != nil {
		return nil, err
	}
	if res == "" {
		return nil, cmb_errors.SystemError
	}

	var resp resType

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

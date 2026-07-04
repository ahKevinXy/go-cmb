package unit_manager

import (
	"encoding/json"

	"github.com/ahKevinXy/go-cmb/cmb_errors"
	"github.com/ahKevinXy/go-cmb/constants"
	"github.com/ahKevinXy/go-cmb/help"
	"github.com/ahKevinXy/go-cmb/models"
)

// RequestUnitManagerNewRaw sends a 交易管家（新） request and returns the raw response body map.
func RequestUnitManagerNewRaw(userId, sm4PrivateKey, userPrivateKey, funcode string, body interface{}) (*models.UnitManagerNewResponse, error) {
	res, err := requestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, funcode, body)
	if err != nil {
		return nil, err
	}
	var resp models.UnitManagerNewResponse
	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RequestUnitManagerNew sends a 交易管家（新） request and unmarshals into response when response is not nil.
func RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, funcode string, body interface{}, response interface{}) error {
	res, err := requestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, funcode, body)
	if err != nil {
		return err
	}
	if response == nil {
		return nil
	}
	return json.Unmarshal([]byte(res), response)
}

func requestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, funcode string, body interface{}) (string, error) {
	reqData := new(models.UnitManagerNewRequest)
	reqData.Request.Head.Reqid = help.GetReqidString()
	reqData.Request.Head.Funcode = funcode
	reqData.Request.Head.Userid = userId
	reqData.Request.Body = body
	reqData.Signature.Sigtim = help.GetSigtimString()
	reqData.Signature.Sigdat = "__signature_sigdat__"

	req, err := json.Marshal(reqData)
	if err != nil {
		return "", err
	}

	res, err := help.CmbSignRequest(string(req), funcode, userId, userPrivateKey, sm4PrivateKey)
	if err != nil {
		return "", err
	}
	if res == "" {
		return "", cmb_errors.SystemError
	}
	return res, nil
}

// UnitManagerNewNTDMAADD 1.1新增记账子单元NTDMAADD.
func UnitManagerNewNTDMAADD(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewNTDMAADDBody) (*models.UnitManagerNewNTDMAADDResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewNTDMAADDBody{}
	}
	var resp models.UnitManagerNewNTDMAADDResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewNTDMAADD, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewNTDMADLT 1.2关闭记账子单元NTDMADLT.
func UnitManagerNewNTDMADLT(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewNTDMADLTBody) (*models.UnitManagerNewNTDMADLTResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewNTDMADLTBody{}
	}
	var resp models.UnitManagerNewNTDMADLTResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewNTDMADLT, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewNTDMAMNT 1.3修改记账子单元NTDMAMNT.
func UnitManagerNewNTDMAMNT(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewNTDMAMNTBody) (*models.UnitManagerNewNTDMAMNTResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewNTDMAMNTBody{}
	}
	var resp models.UnitManagerNewNTDMAMNTResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewNTDMAMNT, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewNTDMALST 1.4查询记账子单元信息NTDMALST.
func UnitManagerNewNTDMALST(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewNTDMALSTBody) (*models.UnitManagerNewNTDMALSTResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewNTDMALSTBody{}
	}
	var resp models.UnitManagerNewNTDMALSTResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewNTDMALST, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewNTOPRDMR 2.1 记账子单元原路退回NTOPRDMR.
func UnitManagerNewNTOPRDMR(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewNTOPRDMRBody) (*models.UnitManagerNewNTOPRDMRResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewNTOPRDMRBody{}
	}
	var resp models.UnitManagerNewNTOPRDMRResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewNTOPRDMR, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewNTDMATRX 2.2 记账子单元内部调配NTDMATRX.
func UnitManagerNewNTDMATRX(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewNTDMATRXBody) (*models.UnitManagerNewNTDMATRXResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewNTDMATRXBody{}
	}
	var resp models.UnitManagerNewNTDMATRXResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewNTDMATRX, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewNTHKPACT 2.3手工记账NTHKPACT.
func UnitManagerNewNTHKPACT(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewNTHKPACTBody) (*models.UnitManagerNewNTHKPACTResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewNTHKPACTBody{}
	}
	var resp models.UnitManagerNewNTHKPACTResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewNTHKPACT, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewNTHCLACT 2.4记账子单元冲账NTHCLACT.
func UnitManagerNewNTHCLACT(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewNTHCLACTBody) (*models.UnitManagerNewNTHCLACTResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewNTHCLACTBody{}
	}
	var resp models.UnitManagerNewNTHCLACTResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewNTHCLACT, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewBHHDZCTZ 3.1交易查询BHHDZCTZ.
func UnitManagerNewBHHDZCTZ(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewBHHDZCTZBody) (*models.UnitManagerNewBHHDZCTZResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewBHHDZCTZBody{}
	}
	var resp models.UnitManagerNewBHHDZCTZResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewBHHDZCTZ, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewNTDMTQRYPDF 3.2记账子单元交易明细对账单获取NTDMTQRYPDF.
func UnitManagerNewNTDMTQRYPDF(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewNTDMTQRYPDFBody) (*models.UnitManagerNewNTDMTQRYPDFResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewNTDMTQRYPDFBody{}
	}
	var resp models.UnitManagerNewNTDMTQRYPDFResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewNTDMTQRYPDF, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewDCTASKID 3.3记账子单元交易明细对账单处理结果查询请求DCTASKID.
func UnitManagerNewDCTASKID(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewDCTASKIDBody) (*models.UnitManagerNewDCTASKIDResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewDCTASKIDBody{}
	}
	var resp models.UnitManagerNewDCTASKIDResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewDCTASKID, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewBHHDZDBQ 3.4退票交易查询BHHDZDBQ.
func UnitManagerNewBHHDZDBQ(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewBHHDZDBQBody) (*models.UnitManagerNewBHHDZDBQResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewBHHDZDBQBody{}
	}
	var resp models.UnitManagerNewBHHDZDBQResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewBHHDZDBQ, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewNTDMABAL 4.1查询记账子单元余额NTDMABAL.
func UnitManagerNewNTDMABAL(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewNTDMABALBody) (*models.UnitManagerNewNTDMABALResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewNTDMABALBody{}
	}
	var resp models.UnitManagerNewNTDMABALResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewNTDMABAL, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewNTDMAHBD 4.2查询所有记账子单元的某日余额NTDMAHBD.
func UnitManagerNewNTDMAHBD(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewNTDMAHBDBody) (*models.UnitManagerNewNTDMAHBDResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewNTDMAHBDBody{}
	}
	var resp models.UnitManagerNewNTDMAHBDResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewNTDMAHBD, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewNTDMAHAD 4.3查询单个记账子单元的历史余额NTDMAHAD.
func UnitManagerNewNTDMAHAD(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewNTDMAHADBody) (*models.UnitManagerNewNTDMAHADResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewNTDMAHADBody{}
	}
	var resp models.UnitManagerNewNTDMAHADResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewNTDMAHAD, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewNTMTHAUB 4.4查询记账子单元历史月积数NTMTHAUB.
func UnitManagerNewNTMTHAUB(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewNTMTHAUBBody) (*models.UnitManagerNewNTMTHAUBResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewNTMTHAUBBody{}
	}
	var resp models.UnitManagerNewNTMTHAUBResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewNTMTHAUB, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewNTDUMRED 5.1按业务参考号查询结果NTDUMRED.
func UnitManagerNewNTDUMRED(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewNTDUMREDBody) (*models.UnitManagerNewNTDUMREDResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewNTDUMREDBody{}
	}
	var resp models.UnitManagerNewNTDUMREDResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewNTDUMRED, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewNTDUMINF 5.2详情查询NTDUMINF.
func UnitManagerNewNTDUMINF(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewNTDUMINFBody) (*models.UnitManagerNewNTDUMINFResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewNTDUMINFBody{}
	}
	var resp models.UnitManagerNewNTDUMINFResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewNTDUMINF, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewNTDMAWLS 6.1设置结算户入账分类条件NTDMAWLS.
func UnitManagerNewNTDMAWLS(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewNTDMAWLSBody) (*models.UnitManagerNewNTDMAWLSResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewNTDMAWLSBody{}
	}
	var resp models.UnitManagerNewNTDMAWLSResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewNTDMAWLS, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewNTDMAWLD 6.2删除结算户入账分类条件NTDMAWLD.
func UnitManagerNewNTDMAWLD(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewNTDMAWLDBody) (*models.UnitManagerNewNTDMAWLDResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewNTDMAWLDBody{}
	}
	var resp models.UnitManagerNewNTDMAWLDResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewNTDMAWLD, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewBHHDZWQC 6.3 查询结算户入账分类条件BHHDZWQC.
func UnitManagerNewBHHDZWQC(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewBHHDZWQCBody) (*models.UnitManagerNewBHHDZWQCResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewBHHDZWQCBody{}
	}
	var resp models.UnitManagerNewBHHDZWQCResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewBHHDZWQC, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewBHHZWMNT 6.4修改不满足结算户入账分类条件的入账方式BHHZWMNT.
func UnitManagerNewBHHZWMNT(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewBHHZWMNTBody) (*models.UnitManagerNewBHHZWMNTResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewBHHZWMNTBody{}
	}
	var resp models.UnitManagerNewBHHZWMNTResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewBHHZWMNT, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewBHHDNWCA 7.1新增结算户出账分类条件BHHDNWCA.
func UnitManagerNewBHHDNWCA(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewBHHDNWCABody) (*models.UnitManagerNewBHHDNWCAResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewBHHDNWCABody{}
	}
	var resp models.UnitManagerNewBHHDNWCAResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewBHHDNWCA, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewBHHDNWCC 7.2删除结算户出账分类条件BHHDNWCC.
func UnitManagerNewBHHDNWCC(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewBHHDNWCCBody) (*models.UnitManagerNewBHHDNWCCResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewBHHDNWCCBody{}
	}
	var resp models.UnitManagerNewBHHDNWCCResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewBHHDNWCC, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewBHHDNWCQ 7.3查询结算户出账分类条件BHHDNWCQ.
func UnitManagerNewBHHDNWCQ(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewBHHDNWCQBody) (*models.UnitManagerNewBHHDNWCQResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewBHHDNWCQBody{}
	}
	var resp models.UnitManagerNewBHHDNWCQResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewBHHDNWCQ, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewNTDMARLT 8.1设置记账子单元入账条件NTDMARLT.
func UnitManagerNewNTDMARLT(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewNTDMARLTBody) (*models.UnitManagerNewNTDMARLTResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewNTDMARLTBody{}
	}
	var resp models.UnitManagerNewNTDMARLTResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewNTDMARLT, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewNTDMARLD 8.2 删除记账子单元入账条件NTDMARLD.
func UnitManagerNewNTDMARLD(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewNTDMARLDBody) (*models.UnitManagerNewNTDMARLDResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewNTDMARLDBody{}
	}
	var resp models.UnitManagerNewNTDMARLDResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewNTDMARLD, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewNTDMARLQ 8.3查询记账子单元入账条件NTDMARLQ.
func UnitManagerNewNTDMARLQ(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewNTDMARLQBody) (*models.UnitManagerNewNTDMARLQResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewNTDMARLQBody{}
	}
	var resp models.UnitManagerNewNTDMARLQResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewNTDMARLQ, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewNTDMATMN 8.4 修改不满足记账子单元入账条件时的入账方式NTDMATMN.
func UnitManagerNewNTDMATMN(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewNTDMATMNBody) (*models.UnitManagerNewNTDMATMNResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewNTDMATMNBody{}
	}
	var resp models.UnitManagerNewNTDMATMNResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewNTDMATMN, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewBHHDNDSA 9.1新增个人账户白名单BHHDNDSA.
func UnitManagerNewBHHDNDSA(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewBHHDNDSABody) (*models.UnitManagerNewBHHDNDSAResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewBHHDNDSABody{}
	}
	var resp models.UnitManagerNewBHHDNDSAResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewBHHDNDSA, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewBHHDNDSD 9.2删除个人账户白名单BHHDNDSD.
func UnitManagerNewBHHDNDSD(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewBHHDNDSDBody) (*models.UnitManagerNewBHHDNDSDResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewBHHDNDSDBody{}
	}
	var resp models.UnitManagerNewBHHDNDSDResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewBHHDNDSD, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewBHHDNDSQ 9.3查询个人账户白名单BHHDNDSQ.
func UnitManagerNewBHHDNDSQ(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewBHHDNDSQBody) (*models.UnitManagerNewBHHDNDSQResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewBHHDNDSQBody{}
	}
	var resp models.UnitManagerNewBHHDNDSQResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewBHHDNDSQ, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewNTDCDADD 10.1设置记账子单元关联公司卡NTDCDADD.
func UnitManagerNewNTDCDADD(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewNTDCDADDBody) (*models.UnitManagerNewNTDCDADDResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewNTDCDADDBody{}
	}
	var resp models.UnitManagerNewNTDCDADDResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewNTDCDADD, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewNTDCDDLT 10.2解除记账子单元与公司卡关联NTDCDDLT.
func UnitManagerNewNTDCDDLT(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewNTDCDDLTBody) (*models.UnitManagerNewNTDCDDLTResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewNTDCDDLTBody{}
	}
	var resp models.UnitManagerNewNTDCDDLTResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewNTDCDDLT, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewNTDCDRLT 10.3查询记账子单元关联公司卡列表NTDCDRLT.
func UnitManagerNewNTDCDRLT(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewNTDCDRLTBody) (*models.UnitManagerNewNTDCDRLTResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewNTDCDRLTBody{}
	}
	var resp models.UnitManagerNewNTDCDRLTResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewNTDCDRLT, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewNTDUMCTL 11.1设置记账子单元收款/出款限制NTDUMCTL.
func UnitManagerNewNTDUMCTL(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewNTDUMCTLBody) (*models.UnitManagerNewNTDUMCTLResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewNTDUMCTLBody{}
	}
	var resp models.UnitManagerNewNTDUMCTLResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewNTDUMCTL, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewNTDUMCLE 11.2解除记账子单元收款/出款限制NTDUMCLE.
func UnitManagerNewNTDUMCLE(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewNTDUMCLEBody) (*models.UnitManagerNewNTDUMCLEResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewNTDUMCLEBody{}
	}
	var resp models.UnitManagerNewNTDUMCLEResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewNTDUMCLE, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewBHHJINAS 12.1 单笔设置子单元利率BHHJINAS.
func UnitManagerNewBHHJINAS(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewBHHJINASBody) (*models.UnitManagerNewBHHJINASResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewBHHJINASBody{}
	}
	var resp models.UnitManagerNewBHHJINASResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewBHHJINAS, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewBHHJINCS 12.2 单笔删除子单元利率BHHJINCS.
func UnitManagerNewBHHJINCS(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewBHHJINCSBody) (*models.UnitManagerNewBHHJINCSResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewBHHJINCSBody{}
	}
	var resp models.UnitManagerNewBHHJINCSResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewBHHJINCS, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewBHHJINQD 12.3 查询子单元利率BHHJINQD.
func UnitManagerNewBHHJINQD(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewBHHJINQDBody) (*models.UnitManagerNewBHHJINQDResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewBHHJINQDBody{}
	}
	var resp models.UnitManagerNewBHHJINQDResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewBHHJINQD, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewBHHDINRP 12.4 子单元单笔提前结息BHHDINRP.
func UnitManagerNewBHHDINRP(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewBHHDINRPBody) (*models.UnitManagerNewBHHDINRPResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewBHHDINRPBody{}
	}
	var resp models.UnitManagerNewBHHDINRPResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewBHHDINRP, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewBHHJINRS 12.5 子单元单笔手工结息BHHJINRS.
func UnitManagerNewBHHJINRS(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewBHHJINRSBody) (*models.UnitManagerNewBHHJINRSResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewBHHJINRSBody{}
	}
	var resp models.UnitManagerNewBHHJINRSResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewBHHJINRS, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewBHHJINSD 12.6 查询子单元结息结果BHHJINSD.
func UnitManagerNewBHHJINSD(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewBHHJINSDBody) (*models.UnitManagerNewBHHJINSDResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewBHHJINSDBody{}
	}
	var resp models.UnitManagerNewBHHJINSDResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewBHHJINSD, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewBHHDINDT 12.7合并结息明细查询BHHDINDT.
func UnitManagerNewBHHDINDT(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewBHHDINDTBody) (*models.UnitManagerNewBHHDINDTResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewBHHDINDTBody{}
	}
	var resp models.UnitManagerNewBHHDINDTResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewBHHDINDT, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewNTDMAINT 12.8记账子单元利息试算NTDMAINT.
func UnitManagerNewNTDMAINT(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewNTDMAINTBody) (*models.UnitManagerNewNTDMAINTResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewNTDMAINTBody{}
	}
	var resp models.UnitManagerNewNTDMAINTResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewNTDMAINT, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewNTDMIQRY 12.9记账子单元交易利息试算NTDMIQRY.
func UnitManagerNewNTDMIQRY(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewNTDMIQRYBody) (*models.UnitManagerNewNTDMIQRYResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewNTDMIQRYBody{}
	}
	var resp models.UnitManagerNewNTDMIQRYResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewNTDMIQRY, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewBHHDINBB 12.10按指定金额做利息试算（批量）BHHDINBB.
func UnitManagerNewBHHDINBB(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewBHHDINBBBody) (*models.UnitManagerNewBHHDINBBResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewBHHDINBBBody{}
	}
	var resp models.UnitManagerNewBHHDINBBResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewBHHDINBB, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewBHHDLMAD 13.1单笔新增支付限额BHHDLMAD.
func UnitManagerNewBHHDLMAD(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewBHHDLMADBody) (*models.UnitManagerNewBHHDLMADResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewBHHDLMADBody{}
	}
	var resp models.UnitManagerNewBHHDLMADResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewBHHDLMAD, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewBHHDLMMD 13.2 单笔修改支付限额BHHDLMMD.
func UnitManagerNewBHHDLMMD(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewBHHDLMMDBody) (*models.UnitManagerNewBHHDLMMDResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewBHHDLMMDBody{}
	}
	var resp models.UnitManagerNewBHHDLMMDResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewBHHDLMMD, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewBHHDLMCD 13.3 单笔关闭支付限额BHHDLMCD.
func UnitManagerNewBHHDLMCD(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewBHHDLMCDBody) (*models.UnitManagerNewBHHDLMCDResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewBHHDLMCDBody{}
	}
	var resp models.UnitManagerNewBHHDLMCDResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewBHHDLMCD, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewBHHDLMTP 13.4 单笔设置/关闭临时限额BHHDLMTP.
func UnitManagerNewBHHDLMTP(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewBHHDLMTPBody) (*models.UnitManagerNewBHHDLMTPResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewBHHDLMTPBody{}
	}
	var resp models.UnitManagerNewBHHDLMTPResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewBHHDLMTP, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewBHHDLMQY 13.5 查询支付限额BHHDLMQY.
func UnitManagerNewBHHDLMQY(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewBHHDLMQYBody) (*models.UnitManagerNewBHHDLMQYResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewBHHDLMQYBody{}
	}
	var resp models.UnitManagerNewBHHDLMQYResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewBHHDLMQY, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewBHHDGPON 14.1. 开启多层级架构BHHDGPON.
func UnitManagerNewBHHDGPON(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewBHHDGPONBody) (*models.UnitManagerNewBHHDGPONResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewBHHDGPONBody{}
	}
	var resp models.UnitManagerNewBHHDGPONResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewBHHDGPON, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewBHHDGPAD 14.2新增普通分组BHHDGPAD.
func UnitManagerNewBHHDGPAD(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewBHHDGPADBody) (*models.UnitManagerNewBHHDGPADResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewBHHDGPADBody{}
	}
	var resp models.UnitManagerNewBHHDGPADResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewBHHDGPAD, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewBHHDGPMN 14.3 修改普通分组BHHDGPMN.
func UnitManagerNewBHHDGPMN(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewBHHDGPMNBody) (*models.UnitManagerNewBHHDGPMNResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewBHHDGPMNBody{}
	}
	var resp models.UnitManagerNewBHHDGPMNResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewBHHDGPMN, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewBHHDGPSC 14.4关闭普通分组BHHDGPSC.
func UnitManagerNewBHHDGPSC(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewBHHDGPSCBody) (*models.UnitManagerNewBHHDGPSCResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewBHHDGPSCBody{}
	}
	var resp models.UnitManagerNewBHHDGPSCResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewBHHDGPSC, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewBHHDGPBL 14.5分组余额查询BHHDGPBL.
func UnitManagerNewBHHDGPBL(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewBHHDGPBLBody) (*models.UnitManagerNewBHHDGPBLResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewBHHDGPBLBody{}
	}
	var resp models.UnitManagerNewBHHDGPBLResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewBHHDGPBL, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewBHHDZCTG 14.6分组交易查询BHHDZCTG.
func UnitManagerNewBHHDZCTG(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewBHHDZCTGBody) (*models.UnitManagerNewBHHDZCTGResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewBHHDZCTGBody{}
	}
	var resp models.UnitManagerNewBHHDZCTGResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewBHHDZCTG, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewBHHDGNAD 14.7分组下新增下级子单元BHHDGNAD.
func UnitManagerNewBHHDGNAD(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewBHHDGNADBody) (*models.UnitManagerNewBHHDGNADResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewBHHDGNADBody{}
	}
	var resp models.UnitManagerNewBHHDGNADResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewBHHDGNAD, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewBHHDGNCD 14.8分组下删除下级子单元BHHDGNCD.
func UnitManagerNewBHHDGNCD(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewBHHDGNCDBody) (*models.UnitManagerNewBHHDGNCDResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewBHHDGNCDBody{}
	}
	var resp models.UnitManagerNewBHHDGNCDResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewBHHDGNCD, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewBHHDGNAG 14.9分组下新增下级分组BHHDGNAG.
func UnitManagerNewBHHDGNAG(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewBHHDGNAGBody) (*models.UnitManagerNewBHHDGNAGResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewBHHDGNAGBody{}
	}
	var resp models.UnitManagerNewBHHDGNAGResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewBHHDGNAG, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewBHHDGNCG 14.10分组下删除下级分组BHHDGNCG.
func UnitManagerNewBHHDGNCG(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewBHHDGNCGBody) (*models.UnitManagerNewBHHDGNCGResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewBHHDGNCGBody{}
	}
	var resp models.UnitManagerNewBHHDGNCGResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewBHHDGNCG, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewBHHDGNGS 14.11查询下级节点信息BHHDGNGS.
func UnitManagerNewBHHDGNGS(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewBHHDGNGSBody) (*models.UnitManagerNewBHHDGNGSResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewBHHDGNGSBody{}
	}
	var resp models.UnitManagerNewBHHDGNGSResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewBHHDGNGS, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnitManagerNewBHHDGNGT 14.12 查询上级节点信息BHHDGNGT.
func UnitManagerNewBHHDGNGT(userId, sm4PrivateKey, userPrivateKey string, body *models.UnitManagerNewBHHDGNGTBody) (*models.UnitManagerNewBHHDGNGTResponse, error) {
	if body == nil {
		body = &models.UnitManagerNewBHHDGNGTBody{}
	}
	var resp models.UnitManagerNewBHHDGNGTResponse
	if err := RequestUnitManagerNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbUnitManagerNewBHHDGNGT, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

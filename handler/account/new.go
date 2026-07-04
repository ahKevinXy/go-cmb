package account

import (
	"encoding/json"

	"github.com/ahKevinXy/go-cmb/cmb_errors"
	"github.com/ahKevinXy/go-cmb/constants"
	"github.com/ahKevinXy/go-cmb/help"
	"github.com/ahKevinXy/go-cmb/models"
)

// RequestAccountNewRaw sends a 账务查询 request and returns the raw response body map.
func RequestAccountNewRaw(userId, sm4PrivateKey, userPrivateKey, funcode string, body interface{}) (*models.AccountNewResponse, error) {
	res, err := requestAccountNew(userId, sm4PrivateKey, userPrivateKey, funcode, body)
	if err != nil {
		return nil, err
	}
	var resp models.AccountNewResponse
	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RequestAccountNew sends a 账务查询 request and unmarshals into response when response is not nil.
func RequestAccountNew(userId, sm4PrivateKey, userPrivateKey, funcode string, body interface{}, response interface{}) error {
	res, err := requestAccountNew(userId, sm4PrivateKey, userPrivateKey, funcode, body)
	if err != nil {
		return err
	}
	if response == nil {
		return nil
	}
	return json.Unmarshal([]byte(res), response)
}

func requestAccountNew(userId, sm4PrivateKey, userPrivateKey, funcode string, body interface{}) (string, error) {
	reqData := new(models.AccountNewRequest)
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

// AccountNewDCLISMOD 1.可经办业务模式查询DCLISMOD.
func AccountNewDCLISMOD(userId, sm4PrivateKey, userPrivateKey string, body *models.AccountNewDCLISMODBody) (*models.AccountNewDCLISMODResponse, error) {
	if body == nil {
		body = &models.AccountNewDCLISMODBody{}
	}
	var resp models.AccountNewDCLISMODResponse
	if err := RequestAccountNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbAccountNewDCLISMOD, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AccountNewDCLISACC 2.查询可经办的账户列表DCLISACC.
func AccountNewDCLISACC(userId, sm4PrivateKey, userPrivateKey string, body *models.AccountNewDCLISACCBody) (*models.AccountNewDCLISACCResponse, error) {
	if body == nil {
		body = &models.AccountNewDCLISACCBody{}
	}
	var resp models.AccountNewDCLISACCResponse
	if err := RequestAccountNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbAccountNewDCLISACC, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AccountNewNTQACINF 3.结算账户详细信息查询NTQACINF.
func AccountNewNTQACINF(userId, sm4PrivateKey, userPrivateKey string, body *models.AccountNewNTQACINFBody) (*models.AccountNewNTQACINFResponse, error) {
	if body == nil {
		body = &models.AccountNewNTQACINFBody{}
	}
	var resp models.AccountNewNTQACINFResponse
	if err := RequestAccountNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbAccountNewNTQACINF, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AccountNewNTACCBBK 4.查询分行号信息NTACCBBK.
func AccountNewNTACCBBK(userId, sm4PrivateKey, userPrivateKey string, body *models.AccountNewNTACCBBKBody) (*models.AccountNewNTACCBBKResponse, error) {
	if body == nil {
		body = &models.AccountNewNTACCBBKBody{}
	}
	var resp models.AccountNewNTACCBBKResponse
	if err := RequestAccountNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbAccountNewNTACCBBK, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AccountNewQryNonSettleAccounts 5.非结算账户详细信息查询QryNonSettleAccounts.
func AccountNewQryNonSettleAccounts(userId, sm4PrivateKey, userPrivateKey string, body *models.AccountNewQryNonSettleAccountsBody) (*models.AccountNewQryNonSettleAccountsResponse, error) {
	if body == nil {
		body = &models.AccountNewQryNonSettleAccountsBody{}
	}
	var resp models.AccountNewQryNonSettleAccountsResponse
	if err := RequestAccountNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbAccountNewQryNonSettleAccounts, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AccountNewTotalAccountDirect 6.综合账户信息查询TotalAccountDirect.
func AccountNewTotalAccountDirect(userId, sm4PrivateKey, userPrivateKey string, body *models.AccountNewTotalAccountDirectBody) (*models.AccountNewTotalAccountDirectResponse, error) {
	if body == nil {
		body = &models.AccountNewTotalAccountDirectBody{}
	}
	var resp models.AccountNewTotalAccountDirectResponse
	if err := RequestAccountNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbAccountNewTotalAccountDirect, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AccountNewNTQABINF 7.查询账户历史余额NTQABINF.
func AccountNewNTQABINF(userId, sm4PrivateKey, userPrivateKey string, body *models.AccountNewNTQABINFBody) (*models.AccountNewNTQABINFResponse, error) {
	if body == nil {
		body = &models.AccountNewNTQABINFBody{}
	}
	var resp models.AccountNewNTQABINFResponse
	if err := RequestAccountNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbAccountNewNTQABINF, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AccountNewNTQRYHSB 8.历史冻结余额等余额信息查询NTQRYHSB.
func AccountNewNTQRYHSB(userId, sm4PrivateKey, userPrivateKey string, body *models.AccountNewNTQRYHSBBody) (*models.AccountNewNTQRYHSBResponse, error) {
	if body == nil {
		body = &models.AccountNewNTQRYHSBBody{}
	}
	var resp models.AccountNewNTQRYHSBResponse
	if err := RequestAccountNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbAccountNewNTQRYHSB, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AccountNewNTQADINF 9. 批量查询余额NTQADINF.
func AccountNewNTQADINF(userId, sm4PrivateKey, userPrivateKey string, body *models.AccountNewNTQADINFBody) (*models.AccountNewNTQADINFResponse, error) {
	if body == nil {
		body = &models.AccountNewNTQADINFBody{}
	}
	var resp models.AccountNewNTQADINFResponse
	if err := RequestAccountNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbAccountNewNTQADINF, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AccountNewTrsQryByBreakPoint 10.账户交易信息查询trsQryByBreakPoint.
func AccountNewTrsQryByBreakPoint(userId, sm4PrivateKey, userPrivateKey string, body *models.AccountNewTrsQryByBreakPointBody) (*models.AccountNewTrsQryByBreakPointResponse, error) {
	if body == nil {
		body = &models.AccountNewTrsQryByBreakPointBody{}
	}
	var resp models.AccountNewTrsQryByBreakPointResponse
	if err := RequestAccountNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbAccountNewTrsQryByBreakPoint, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AccountNewNonSettleTranOpr 11非结算户交易流水导出申请NonSettleTranOpr.
func AccountNewNonSettleTranOpr(userId, sm4PrivateKey, userPrivateKey string, body *models.AccountNewNonSettleTranOprBody) (*models.AccountNewNonSettleTranOprResponse, error) {
	if body == nil {
		body = &models.AccountNewNonSettleTranOprBody{}
	}
	var resp models.AccountNewNonSettleTranOprResponse
	if err := RequestAccountNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbAccountNewNonSettleTranOpr, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AccountNewNonSettleTranQry 12查询关联非结算户交易流水NonSettleTranQry.
func AccountNewNonSettleTranQry(userId, sm4PrivateKey, userPrivateKey string, body *models.AccountNewNonSettleTranQryBody) (*models.AccountNewNonSettleTranQryResponse, error) {
	if body == nil {
		body = &models.AccountNewNonSettleTranQryBody{}
	}
	var resp models.AccountNewNonSettleTranQryResponse
	if err := RequestAccountNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbAccountNewNonSettleTranQry, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AccountNewDCTRSPDF 13.PDF文件对账单获取DCTRSPDF.
func AccountNewDCTRSPDF(userId, sm4PrivateKey, userPrivateKey string, body *models.AccountNewDCTRSPDFBody) (*models.AccountNewDCTRSPDFResponse, error) {
	if body == nil {
		body = &models.AccountNewDCTRSPDFBody{}
	}
	var resp models.AccountNewDCTRSPDFResponse
	if err := RequestAccountNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbAccountNewDCTRSPDF, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AccountNewIssueBillOfd 14.OFD文件对账单获取issueBillOfd.
func AccountNewIssueBillOfd(userId, sm4PrivateKey, userPrivateKey string, body *models.AccountNewIssueBillOfdBody) (*models.AccountNewIssueBillOfdResponse, error) {
	if body == nil {
		body = &models.AccountNewIssueBillOfdBody{}
	}
	var resp models.AccountNewIssueBillOfdResponse
	if err := RequestAccountNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbAccountNewIssueBillOfd, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AccountNewQueryBillOfd 15.OFD文件对账单获取结果查询queryBillOfd.
func AccountNewQueryBillOfd(userId, sm4PrivateKey, userPrivateKey string, body *models.AccountNewQueryBillOfdBody) (*models.AccountNewQueryBillOfdResponse, error) {
	if body == nil {
		body = &models.AccountNewQueryBillOfdBody{}
	}
	var resp models.AccountNewQueryBillOfdResponse
	if err := RequestAccountNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbAccountNewQueryBillOfd, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AccountNewDCSIGREC 16.单笔回单查询DCSIGREC.
func AccountNewDCSIGREC(userId, sm4PrivateKey, userPrivateKey string, body *models.AccountNewDCSIGRECBody) (*models.AccountNewDCSIGRECResponse, error) {
	if body == nil {
		body = &models.AccountNewDCSIGRECBody{}
	}
	var resp models.AccountNewDCSIGRECResponse
	if err := RequestAccountNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbAccountNewDCSIGREC, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AccountNewASYCALHD 17. 电子回单异步查询ASYCALHD.
func AccountNewASYCALHD(userId, sm4PrivateKey, userPrivateKey string, body *models.AccountNewASYCALHDBody) (*models.AccountNewASYCALHDResponse, error) {
	if body == nil {
		body = &models.AccountNewASYCALHDBody{}
	}
	var resp models.AccountNewASYCALHDResponse
	if err := RequestAccountNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbAccountNewASYCALHD, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AccountNewDCTASKID 18. 异步打印结果查询DCTASKID.
func AccountNewDCTASKID(userId, sm4PrivateKey, userPrivateKey string, body *models.AccountNewDCTASKIDBody) (*models.AccountNewDCTASKIDResponse, error) {
	if body == nil {
		body = &models.AccountNewDCTASKIDBody{}
	}
	var resp models.AccountNewDCTASKIDResponse
	if err := RequestAccountNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbAccountNewDCTASKID, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

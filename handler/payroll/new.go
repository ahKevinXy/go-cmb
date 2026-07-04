package payroll

import (
	"encoding/json"

	"github.com/ahKevinXy/go-cmb/cmb_errors"
	"github.com/ahKevinXy/go-cmb/constants"
	"github.com/ahKevinXy/go-cmb/help"
	"github.com/ahKevinXy/go-cmb/models"
)

// RequestPayrollNewRaw sends a 代发代扣 request and returns the raw response body map.
func RequestPayrollNewRaw(userId, sm4PrivateKey, userPrivateKey, funcode string, body interface{}) (*models.PayrollNewResponse, error) {
	res, err := requestPayrollNew(userId, sm4PrivateKey, userPrivateKey, funcode, body)
	if err != nil {
		return nil, err
	}
	var resp models.PayrollNewResponse
	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RequestPayrollNew sends a 代发代扣 request and unmarshals into response when response is not nil.
func RequestPayrollNew(userId, sm4PrivateKey, userPrivateKey, funcode string, body interface{}, response interface{}) error {
	res, err := requestPayrollNew(userId, sm4PrivateKey, userPrivateKey, funcode, body)
	if err != nil {
		return err
	}
	if response == nil {
		return nil
	}
	return json.Unmarshal([]byte(res), response)
}

func requestPayrollNew(userId, sm4PrivateKey, userPrivateKey, funcode string, body interface{}) (string, error) {
	reqData := new(models.PayrollNewRequest)
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

// PayrollNewCdcMatchCardBin 1.卡BIN匹配识别标准银联卡CdcMatchCardBin.
func PayrollNewCdcMatchCardBin(userId, sm4PrivateKey, userPrivateKey string, body *models.PayrollNewCdcMatchCardBinBody) (*models.PayrollNewCdcMatchCardBinResponse, error) {
	if body == nil {
		body = &models.PayrollNewCdcMatchCardBinBody{}
	}
	var resp models.PayrollNewCdcMatchCardBinResponse
	if err := RequestPayrollNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbPayrollNewCdcMatchCardBin, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PayrollNewBB6BTHHL 2.代发经办BB6BTHHL.
func PayrollNewBB6BTHHL(userId, sm4PrivateKey, userPrivateKey string, body *models.PayrollNewBB6BTHHLBody) (*models.PayrollNewBB6BTHHLResponse, error) {
	if body == nil {
		body = &models.PayrollNewBB6BTHHLBody{}
	}
	var resp models.PayrollNewBB6BTHHLResponse
	if err := RequestPayrollNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbPayrollNewBB6BTHHL, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PayrollNewBB6BPDQY 3.批次与明细查询 BB6BPDQY.
func PayrollNewBB6BPDQY(userId, sm4PrivateKey, userPrivateKey string, body *models.PayrollNewBB6BPDQYBody) (*models.PayrollNewBB6BPDQYResponse, error) {
	if body == nil {
		body = &models.PayrollNewBB6BPDQYBody{}
	}
	var resp models.PayrollNewBB6BPDQYResponse
	if err := RequestPayrollNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbPayrollNewBB6BPDQY, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PayrollNewBB6INTQY 4. 代发代扣综合查询 BB6INTQY.
func PayrollNewBB6INTQY(userId, sm4PrivateKey, userPrivateKey string, body *models.PayrollNewBB6INTQYBody) (*models.PayrollNewBB6INTQYResponse, error) {
	if body == nil {
		body = &models.PayrollNewBB6INTQYBody{}
	}
	var resp models.PayrollNewBB6INTQYResponse
	if err := RequestPayrollNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbPayrollNewBB6INTQY, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PayrollNewBB6BTHQY 5. 代发批次查询 BB6BTHQY.
func PayrollNewBB6BTHQY(userId, sm4PrivateKey, userPrivateKey string, body *models.PayrollNewBB6BTHQYBody) (*models.PayrollNewBB6BTHQYResponse, error) {
	if body == nil {
		body = &models.PayrollNewBB6BTHQYBody{}
	}
	var resp models.PayrollNewBB6BTHQYResponse
	if err := RequestPayrollNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbPayrollNewBB6BTHQY, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PayrollNewBB6DTLQY 6. 代发明细查询 BB6DTLQY.
func PayrollNewBB6DTLQY(userId, sm4PrivateKey, userPrivateKey string, body *models.PayrollNewBB6DTLQYBody) (*models.PayrollNewBB6DTLQYResponse, error) {
	if body == nil {
		body = &models.PayrollNewBB6DTLQYBody{}
	}
	var resp models.PayrollNewBB6DTLQYResponse
	if err := RequestPayrollNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbPayrollNewBB6DTLQY, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PayrollNewBB6RFDQY 7.代发退票查询 BB6RFDQY.
func PayrollNewBB6RFDQY(userId, sm4PrivateKey, userPrivateKey string, body *models.PayrollNewBB6RFDQYBody) (*models.PayrollNewBB6RFDQYResponse, error) {
	if body == nil {
		body = &models.PayrollNewBB6RFDQYBody{}
	}
	var resp models.PayrollNewBB6RFDQYResponse
	if err := RequestPayrollNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbPayrollNewBB6RFDQY, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PayrollNewBB6AGTQY 8.代发类型查询 BB6AGTQY.
func PayrollNewBB6AGTQY(userId, sm4PrivateKey, userPrivateKey string, body *models.PayrollNewBB6AGTQYBody) (*models.PayrollNewBB6AGTQYResponse, error) {
	if body == nil {
		body = &models.PayrollNewBB6AGTQYBody{}
	}
	var resp models.PayrollNewBB6AGTQYResponse
	if err := RequestPayrollNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbPayrollNewBB6AGTQY, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PayrollNewBB6ACLAK 9.代扣经办BB6ACLAK.
func PayrollNewBB6ACLAK(userId, sm4PrivateKey, userPrivateKey string, body *models.PayrollNewBB6ACLAKBody) (*models.PayrollNewBB6ACLAKResponse, error) {
	if body == nil {
		body = &models.PayrollNewBB6ACLAKBody{}
	}
	var resp models.PayrollNewBB6ACLAKResponse
	if err := RequestPayrollNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbPayrollNewBB6ACLAK, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PayrollNewBB6CDCQY 10. 代扣批次查询 BB6CDCQY.
func PayrollNewBB6CDCQY(userId, sm4PrivateKey, userPrivateKey string, body *models.PayrollNewBB6CDCQYBody) (*models.PayrollNewBB6CDCQYResponse, error) {
	if body == nil {
		body = &models.PayrollNewBB6CDCQYBody{}
	}
	var resp models.PayrollNewBB6CDCQYResponse
	if err := RequestPayrollNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbPayrollNewBB6CDCQY, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PayrollNewBB6ACLIF 11.代扣明细查询BB6ACLIF.
func PayrollNewBB6ACLIF(userId, sm4PrivateKey, userPrivateKey string, body *models.PayrollNewBB6ACLIFBody) (*models.PayrollNewBB6ACLIFResponse, error) {
	if body == nil {
		body = &models.PayrollNewBB6ACLIFBody{}
	}
	var resp models.PayrollNewBB6ACLIFResponse
	if err := RequestPayrollNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbPayrollNewBB6ACLIF, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PayrollNewBB6TRSQY 12.代扣类型查询BB6TRSQY.
func PayrollNewBB6TRSQY(userId, sm4PrivateKey, userPrivateKey string, body *models.PayrollNewBB6TRSQYBody) (*models.PayrollNewBB6TRSQYResponse, error) {
	if body == nil {
		body = &models.PayrollNewBB6TRSQYBody{}
	}
	var resp models.PayrollNewBB6TRSQYResponse
	if err := RequestPayrollNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbPayrollNewBB6TRSQY, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PayrollNewDCAGPPDF 13.代发明细对账单查询请求DCAGPPDF.
func PayrollNewDCAGPPDF(userId, sm4PrivateKey, userPrivateKey string, body *models.PayrollNewDCAGPPDFBody) (*models.PayrollNewDCAGPPDFResponse, error) {
	if body == nil {
		body = &models.PayrollNewDCAGPPDFBody{}
	}
	var resp models.PayrollNewDCAGPPDFResponse
	if err := RequestPayrollNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbPayrollNewDCAGPPDF, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PayrollNewDCTASKID 14.代发明细对账单处理结果查询请求DCTASKID.
func PayrollNewDCTASKID(userId, sm4PrivateKey, userPrivateKey string, body *models.PayrollNewDCTASKIDBody) (*models.PayrollNewDCTASKIDResponse, error) {
	if body == nil {
		body = &models.PayrollNewDCTASKIDBody{}
	}
	var resp models.PayrollNewDCTASKIDResponse
	if err := RequestPayrollNew(userId, sm4PrivateKey, userPrivateKey, constants.CmbPayrollNewDCTASKID, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

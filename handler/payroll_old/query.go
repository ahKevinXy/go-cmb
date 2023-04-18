package payroll_old

import (
	"encoding/json"
	"github.com/ahKevinXy/go-cmb/constants"
	"github.com/ahKevinXy/go-cmb/help"
	"github.com/ahKevinXy/go-cmb/models"
	"strconv"
	"time"
)

// QueryPayRollDetail
// @Description:查询交易概要信息
// @Author ahKevinXy
// @Date 2023-02-13 11:37:38
func QueryPayRollDetail(userId, asePrivateKey, userPrivateKey, reqnbr, bthnbr, trxseq, histag string) (*models.QueryPayRollDetailResponse, error) {
	reqData := new(models.PayrollOldQueryTransRequest)
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
	res := help.CmbSignRequest(string(req), constants.CmbPayrollOldQueryTrans, userId, userPrivateKey, asePrivateKey)

	if res == "" {
		return nil, err
	}

	var resp models.QueryPayRollDetailResponse

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		print(err)
	}

	return &resp, err

}

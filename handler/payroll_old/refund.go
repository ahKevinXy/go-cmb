package payroll_old

import (
	"encoding/json"
	"github.com/ahKevinXy/go-cmb/v1/constants"
	"github.com/ahKevinXy/go-cmb/v1/help"
	"github.com/ahKevinXy/go-cmb/v1/models"

	"strconv"
	"time"
)

func Refund(
	userId, sm4PrivateKey, userPrivateKey, accnbr, bbknbr, trstyp, bgDate, endDat string) (*models.GetPayrollPdfResponse, error) {
	reqData := new(models.PayrollOldRefundRequest)
	reqData.Request.Head.Reqid = time.Now().Format("20060102150405000") + strconv.Itoa(time.Now().Nanosecond())
	reqData.Request.Head.Funcode = constants.CmbPayrollOldQueryTransRefund
	reqData.Request.Head.Userid = userId
	reqData.Signature.Sigtim = time.Now().Format("20060102150405")
	reqData.Signature.Sigdat = "__signature_sigdat__"
	reqData.Request.Body.Ntagdrfdy1 = append(reqData.Request.Body.Ntagdrfdy1, &models.Ntagdrfdy1{
		Bbknbr: bbknbr,
		Accnbr: accnbr,
		Trstyp: trstyp,
		Bgndat: bgDate,
		Enddat: endDat,
	})

	req, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	res, err := help.CmbSignRequest(string(req), constants.CmbPayrollOldQueryTransRefund, userId, userPrivateKey, sm4PrivateKey)
	if err != nil {
		return nil, err
	}
	if res == "" {
		return nil, err
	}

	var resp models.GetPayrollPdfResponse

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		print(err)
	}

	return &resp, err

}

package payroll

import (
	"encoding/json"
	"github.com/ahKevinXy/go-cmb/constants"
	"github.com/ahKevinXy/go-cmb/help"
	"github.com/ahKevinXy/go-cmb/models"
	"strconv"
	"time"
)

// QueryRefundList
//  @Description: BB6RFDQY
//  @param userId
//  @param asePrivateKey
//  @param userPrivateKey
//  @param accNbr
//  @param trstyp
//  @param bgndat
//  @param enddat
//  @param cntkey
//  @param reqnbr
//  @return *models.QueryBatchTransListResponse
//  @return error
//  @Author  ahKevinXy
//  @Date  2023-04-14 15:22:17
func QueryRefundList(userId, asePrivateKey, userPrivateKey, accNbr, trstyp, bgndat, enddat, cntkey, reqnbr string) (*models.QueryBatchTransListResponse, error) {
	reqData := new(models.QueryPayrollRefundRequest)
	reqData.Request.Head.Reqid = time.Now().Format("20060102150405000") + strconv.Itoa(time.Now().Nanosecond())
	reqData.Request.Head.Funcode = constants.CmbPayrollQueryBatchList
	reqData.Request.Head.Userid = userId
	reqData.Signature.Sigtim = time.Now().Format("20060102150405")
	reqData.Signature.Sigdat = "__signature_sigdat__"
	reqData.Request.Body.Bb6rfdqyy1 = append(reqData.Request.Body.Bb6rfdqyy1, &models.Bb6rfdqyy1{
		AccNbr: accNbr,
		Trstyp: trstyp,
		BgnDat: bgndat,
		EndDat: enddat,
		CtnKey: cntkey,
		Reqnbr: reqnbr,
	})
	req, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	//  todo
	res := help.CmbSignRequest(string(req), constants.CmbPayrollQueryBatchList, userId, userPrivateKey, asePrivateKey)

	if res == "" {
		return nil, err
	}

	var resp models.QueryBatchTransListResponse

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, err
	}

	return &resp, err
}

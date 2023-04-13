package account

import (
	"encoding/json"
	"github.com/ahKevinXy/go-cmb/constants"
	"github.com/ahKevinXy/go-cmb/help"
	"github.com/ahKevinXy/go-cmb/models"
	"strconv"
	"time"
)

func QueryAccountPaymentTransInfo(userId,
	asePrivateKey, userPrivateKey,
	busCode,
	busMode,
	dtlNbr,
	ctnFlg,
	ctnSts,
	bthNbr string,

	payList []*models.Bb1paybhx1,
) (*models.MainAccountSinglePayResponse, error) {
	reqData := new(models.MainAccountBatchPayRequest)
	reqData.Request.Head.Reqid = time.Now().Format("20060102150405000") + strconv.Itoa(time.Now().Nanosecond())
	reqData.Request.Head.Funcode = constants.CmbAccountBatchPayQueryInfo
	reqData.Request.Head.Userid = userId
	reqData.Signature.Sigtim = time.Now().Format("20060102150405")
	reqData.Signature.Sigdat = "__signature_sigdat__"
	reqData.Request.Body.Bb1bmdbhx1 = append(reqData.Request.Body.Bb1bmdbhx1, &models.Bb1bmdbhx1{
		BusCod: busCode,
		BusMod: busMode,
		DtlNbr: dtlNbr,
		BthNbr: bthNbr,
		CtnFlg: ctnFlg,
		CtnSts: ctnSts,
	})

	reqData.Request.Body.Bb1paybhx1 = payList

	req, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	//  todo
	res := help.CmbSignRequest(string(req), constants.CmbAccountBatchPayQueryInfo, userId, userPrivateKey, asePrivateKey)

	if res == "" {
		return nil, err
	}

	var resp models.MainAccountSinglePayResponse

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

package account

import (
	"encoding/json"
	"github.com/ahKevinXy/go-cmb/constants"
	"github.com/ahKevinXy/go-cmb/help"
	"github.com/ahKevinXy/go-cmb/models"
	"strconv"
	"time"
)

// MainAccountInfo
// @Description:  获取可经办列表 todo 省略 请求参数
// @Author ahKevinXy
// @Date 2023-02-13 13:16:21
func MainAccountInfo(userId, asePrivateKey, userPrivateKey, accnbr, bbknbr string) (*models.AccountInfoResponse, error) {
	reqData := new(models.AccountDetailsRequest)
	reqData.Request.Head.Reqid = time.Now().Format("20060102150405000") + strconv.Itoa(time.Now().Nanosecond())
	reqData.Request.Head.Funcode = constants.CmbAccountInfo
	reqData.Request.Head.Userid = userId
	reqData.Signature.Sigtim = time.Now().Format("20060102150405")
	reqData.Signature.Sigdat = "__signature_sigdat__"
	reqData.Request.Body.Ntqacinfx = append(reqData.Request.Body.Ntqacinfx, &models.Ntqacinfx{
		Accnbr: accnbr,
		Bbknbr: bbknbr,
	})

	req, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	//  todo 优化
	res := help.CmbSignRequest(string(req), constants.CmbAccountInfo, userId, userPrivateKey, asePrivateKey)

	if res == "" {

	}

	var resp models.AccountInfoResponse

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

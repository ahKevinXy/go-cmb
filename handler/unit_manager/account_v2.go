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

// QueryUnitAccountInfoV2
//
//	@Description: 记账子单元查询
//	@param userId
//	@param sm4PrivateKey
//	@param userPrivateKey
//	@param accnbr 账户
//	@param bbknbr 开户行
//	@param danbeg 开始编号 (所有 *)
//	@param danend 结束编号 (所有 *)
//	@param ctnkey 续传key
//	@return *models.AccountUnitInfoResponse
//	@return error
//	@Author  ahKevinXy
//	@Date  2023-05-24 09:52:38
func QueryUnitAccountInfoV2(userId, sm4PrivateKey, userPrivateKey, accnbr, bbknbr, danbeg, danend, ctnkey string) (*models.QueryUnitAccountInfoV2Response, error) {
	reqData := new(models.QueryUnitAccountInfoV2Request)
	reqData.Request.Head.Reqid = time.Now().Format("20060102150405000") + strconv.Itoa(time.Now().Nanosecond())
	reqData.Request.Head.Funcode = constants.CmbUnitManageAccountQueryV2
	reqData.Request.Head.Userid = userId
	reqData.Signature.Sigtim = time.Now().Format("20060102150405")
	reqData.Signature.Sigdat = "__signature_sigdat__"
	reqData.Request.Body.Ntdumqryy1 = append(reqData.Request.Body.Ntdumqryy1, &models.Ntdumqryy1{
		Bbknbr: bbknbr,
		Inbacc: accnbr,
		Danbeg: danbeg,
		Danend: danend,
		Ctnkey: ctnkey,
	})

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

	var resp models.QueryUnitAccountInfoV2Response

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, err
	}

	return &resp, err
}

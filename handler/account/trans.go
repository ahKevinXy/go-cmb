package account

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/ahKevinXy/go-cmb/cmb_errors"
	"github.com/ahKevinXy/go-cmb/constants"
	"github.com/ahKevinXy/go-cmb/help"
	"github.com/ahKevinXy/go-cmb/models"
)

// GetMainAccountTransInfo
//
//	@Description:   获取交易信息
//	@param userId
//	@param sm4Key
//	@param userPrivateKey
//	@param bbknbr 开户行
//	@param accnbr 账户
//	@param trsDat 交易日期
//	@param trsseq
//	@return *models.GetMainAccountTransInfoResponse
//	@return error
//	@Author  ahKevinXy
//	@Date 2023-04-10 14:01:42
func GetMainAccountTransInfo(userId, sm4Key, userPrivateKey,
	bbknbr,
	accnbr,
	trsDat,
	trsseq string) (*models.GetMainAccountTransInfoResponse, error) {
	reqData := new(models.GetMainAccountTransInfoRequest)
	reqData.Request.Head.Reqid = time.Now().Format("20060102150405000") + strconv.Itoa(time.Now().Nanosecond())
	reqData.Request.Head.Funcode = constants.CmbAccountQueryTransInfo
	reqData.Request.Head.Userid = userId
	reqData.Signature.Sigtim = time.Now().Format("20060102150405")
	reqData.Signature.Sigdat = "__signature_sigdat__"
	reqData.Request.Body.Bbknbr = bbknbr
	reqData.Request.Body.Accnbr = accnbr
	reqData.Request.Body.Trsdat = trsDat
	reqData.Request.Body.Trsseq = trsseq

	req, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	//  todo
	res := help.CmbSignRequest(string(req), constants.CmbAccountQueryTransInfo, userId, userPrivateKey, sm4Key)

	if res == "" {
		return nil, cmb_errors.SystemError
	}

	var resp models.GetMainAccountTransInfoResponse

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, err
	}

	return &resp, nil

}

// QueryAccountTransInfo
//
//	@Description:   获取交易信息
//	@param userId
//	@param sm4Key
//	@param userPrivateKey
//	@param bbknbr 开户行
//	@param accnbr 账户
//	@param bgndat 开始时间
//	@param enddat 结束时间
//	@param lowamt 最小金额
//	@param hghamt 最大金额
//	@return *models.QueryAccountTransInfoResponse
//	@return error
//	@Author  ahKevinXy
//	@Date 2023-04-13 15:32:25
func QueryAccountTransInfo(userId, sm4Key, userPrivateKey,
	bbknbr,
	accnbr,
	bgndat,
	enddat,
	lowamt, hghamt string) (*models.QueryAccountTransInfoResponse, error) {
	reqData := new(models.QueryAccountTransInfoRequest)
	reqData.Request.Head.Reqid = time.Now().Format("20060102150405000") + strconv.Itoa(time.Now().Nanosecond())
	reqData.Request.Head.Funcode = constants.CmbAccountGetTransInfo
	reqData.Request.Head.Userid = userId
	reqData.Signature.Sigtim = time.Now().Format("20060102150405")
	reqData.Signature.Sigdat = "__signature_sigdat__"
	reqData.Request.Body.Sdktsinfx = append(reqData.Request.Body.Sdktsinfx, &models.Sdktsinfx{
		Accnbr: accnbr,
		Amtcdr: "",
		Bbknbr: bbknbr,
		Bgndat: bgndat,

		Enddat: enddat,
		Hghamt: hghamt,
		Lowamt: lowamt,
	})

	req, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	//  todo
	res := help.CmbSignRequest(string(req), constants.CmbAccountGetTransInfo, userId, userPrivateKey, sm4Key)

	if res == "" {
		return nil, cmb_errors.SystemError
	}

	var resp models.QueryAccountTransInfoResponse

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, err
	}

	return &resp, nil

}

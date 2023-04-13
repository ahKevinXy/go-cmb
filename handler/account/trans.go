package account

import (
	"encoding/json"
	"github.com/ahKevinXy/go-cmb/constants"
	"github.com/ahKevinXy/go-cmb/help"
	"github.com/ahKevinXy/go-cmb/models"
	"strconv"
	"time"
)

// GetMainAccountTransInfo
//  @Description:   获取交易信息
//  @param userId
//  @param asePrivateKey
//  @param userPrivateKey
//  @param bbknbr
//  @param accnbr
//  @param trsDat
//  @param trsseq
//  @param category
//  @return *models.GetMainAccountTransInfoResponse
//  @return error
//  @Author  ahKevinXy
//  @Date2023-04-10 14:01:42
func GetMainAccountTransInfo(userId, asePrivateKey, userPrivateKey,
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
	res := help.CmbSignRequest(string(req), constants.CmbAccountQueryTransInfo, userId, userPrivateKey, asePrivateKey)

	if res == "" {
		return nil, err
	}

	var resp models.GetMainAccountTransInfoResponse

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, err
	}

	return &resp, nil

}

// QueryAccountTransInfo
//  @Description:   获取交易信息
//  @param userId
//  @param asePrivateKey
//  @param userPrivateKey
//  @param bbknbr
//  @param accnbr
//  @param bgndat
//  @param enddat
//  @param lowamt
//  @param hghamt
//  @return *models.QueryAccountTransInfoResponse
//  @return error
//  @Author  ahKevinXy
//  @Date2023-04-13 15:32:25
func QueryAccountTransInfo(userId, asePrivateKey, userPrivateKey,
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
	res := help.CmbSignRequest(string(req), constants.CmbAccountGetTransInfo, userId, userPrivateKey, asePrivateKey)

	if res == "" {
		return nil, err
	}

	var resp models.QueryAccountTransInfoResponse

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, err
	}

	return &resp, nil

}

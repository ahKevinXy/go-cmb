package unit_manager

import (
	"encoding/json"
	"github.com/ahKevinXy/go-cmb/constants"
	"github.com/ahKevinXy/go-cmb/help"
	"github.com/ahKevinXy/go-cmb/models"
	"strconv"
	"time"
)

// QueryUnitAccountDetail
//  @Description:
//  @param userId
//  @param asePrivateKey
//  @param userPrivateKey
//  @param reqNbr
//  @return *models.QueryUnitAccountDetailResponse
//  @return error
//  @Author  ahKevinXy
//  @Date  2023-04-13 18:13:23
func QueryUnitAccountDetail(
	userId,
	asePrivateKey, userPrivateKey,

	reqNbr string) (*models.QueryUnitAccountDetailResponse, error) {
	reqData := new(models.QueryUnitTransDetailRequest)
	reqData.Request.Head.Reqid = time.Now().Format("20060102150405000") + strconv.Itoa(time.Now().Nanosecond())
	reqData.Request.Head.Funcode = constants.CmbUnitManageAccountTransQueryDetail
	reqData.Request.Head.Userid = userId
	reqData.Signature.Sigtim = time.Now().Format("20060102150405")
	reqData.Signature.Sigdat = "__signature_sigdat__"
	reqData.Request.Body.Ntduminfx1 = append(reqData.Request.Body.Ntduminfx1, &models.Ntduminfx1{
		reqNbr,
	})

	req, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	//  todo
	res := help.CmbSignRequest(string(req), constants.CmbUnitManageAccountTransQueryDetail, userId, userPrivateKey, asePrivateKey)

	if res == "" {
		return nil, err
	}

	var resp models.QueryUnitAccountDetailResponse

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// QueryUnitAccountByBusNo
//  @Description:  通过业务单号获取结果信息
//  @param userId
//  @param asePrivateKey
//  @param userPrivateKey
//  @param yurref
//  @param bgndat
//  @param enddat
//  @return *models.QueryUnitTransByBusNoResponse
//  @return error
//  @Author  ahKevinXy
//  @Date  2023-04-13 18:05:07
func QueryUnitAccountByBusNo(
	userId,
	asePrivateKey, userPrivateKey,
	yurref,
	bgndat,
	enddat string) (*models.QueryUnitTransByBusNoResponse, error) {
	reqData := new(models.QueryUnitTransByBusNoRequest)
	reqData.Request.Head.Reqid = time.Now().Format("20060102150405000") + strconv.Itoa(time.Now().Nanosecond())
	reqData.Request.Head.Funcode = constants.CmbUnitManageAccountTransQueryByBusNo
	reqData.Request.Head.Userid = userId
	reqData.Signature.Sigtim = time.Now().Format("20060102150405")
	reqData.Signature.Sigdat = "__signature_sigdat__"
	reqData.Request.Body.Ntdumredx1 = append(reqData.Request.Body.Ntdumredx1, &models.Ntdumredx1{
		Yurref: yurref,
		Bgndat: bgndat,
		Enddat: enddat,
	})

	req, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	//  todo
	res := help.CmbSignRequest(string(req), constants.CmbUnitManageAccountTransQueryByBusNo, userId, userPrivateKey, asePrivateKey)

	if res == "" {
		return nil, err
	}

	var resp models.QueryUnitTransByBusNoResponse

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

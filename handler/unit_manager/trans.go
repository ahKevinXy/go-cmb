package unit_manager

import (
	"encoding/json"
	"errors"
	"github.com/ahKevinXy/go-cmb/constants"
	"github.com/ahKevinXy/go-cmb/help"
	"github.com/ahKevinXy/go-cmb/models"
	"strconv"
	"time"
)

// GetUnitAccountTransList
//  @Description: 记账子单元当天交易查询
//  @param userId
//  @param asePrivateKey
//  @param userPrivateKey
//  @param accnbr
//  @param dmanbr
//  @param ctnkey
//  @return *models.UnitAccountTransDailyResponse
//  @return error
//  @Author  ahKevinXy
//  @Date  2023-04-13 18:49:18
func GetUnitAccountTransList(userId, asePrivateKey, userPrivateKey, accnbr, dmanbr, ctnkey string) (*models.UnitAccountTransDailyResponse, error) {
	reqData := new(models.AccountAddUnitTransDailyRequest)
	reqData.Request.Head.Reqid = time.Now().Format("20060102150405000") + strconv.Itoa(time.Now().Nanosecond())
	reqData.Request.Head.Funcode = constants.CmbUnitManageAccountTransDaily
	reqData.Request.Head.Userid = userId
	reqData.Signature.Sigtim = time.Now().Format("20060102150405")
	reqData.Signature.Sigdat = "__signature_sigdat__"
	reqData.Request.Body.Ntdmtlsty = append(reqData.Request.Body.Ntdmtlsty, &models.Ntdmtlsty{
		Accnbr: accnbr,
		Dmanbr: dmanbr,
		Ctnkey: ctnkey,
	})

	req, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	//  todo
	res := help.CmbSignRequest(string(req), constants.CmbUnitManageAccountTransDaily, userId, userPrivateKey, asePrivateKey)

	if res == "" {
		return nil, err
	}

	var resp *models.UnitAccountTransDailyResponse

	if err := json.Unmarshal([]byte(res), &resp); err != nil {

		return nil, err
	}

	return resp, nil
}

// GetUnitAccountTransHistoryList
//  @Description:  获取记账单元历史列表
//  @param userId
//  @param asePrivateKey
//  @param userPrivateKey
//  @param accnbr
//  @param dmanbr
//  @param begdat
//  @param enddat
//  @param ctnkey
//  @return *models.UnitAccountTransHistoryResponse
//  @return error
//  @Author  ahKevinXy
//  @Date  2023-04-13 18:53:33
func GetUnitAccountTransHistoryList(userId, asePrivateKey, userPrivateKey, accnbr, dmanbr, begdat, enddat, ctnkey string) (*models.UnitAccountTransHistoryResponse, error) {
	reqData := new(models.AccountUnitTransHistoryRequest)
	reqData.Request.Head.Reqid = time.Now().Format("20060102150405000") + strconv.Itoa(time.Now().Nanosecond())
	reqData.Request.Head.Funcode = constants.CmbUnitManageAccountTransHistory
	reqData.Request.Head.Userid = userId
	reqData.Signature.Sigtim = time.Now().Format("20060102150405")
	reqData.Signature.Sigdat = "__signature_sigdat__"
	reqData.Request.Body.Ntdmthlsy = append(reqData.Request.Body.Ntdmthlsy, &models.Ntdmthlsy{
		Accnbr: accnbr,
		Dmanbr: dmanbr,
		Ctnkey: ctnkey,
		Begdat: begdat,
		Enddat: enddat,
	})

	req, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	//  todo
	res := help.CmbSignRequest(string(req), constants.CmbUnitManageAccountTransHistory, userId, userPrivateKey, asePrivateKey)

	if res == "" {
		return nil, err
	}

	var resp *models.UnitAccountTransHistoryResponse

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, errors.New(res)
	}

	return resp, nil
}

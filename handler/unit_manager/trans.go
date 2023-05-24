package unit_manager

import (
	"encoding/json"
	"errors"
	"github.com/ahKevinXy/go-cmb/cmb_errors"
	"github.com/ahKevinXy/go-cmb/constants"
	"github.com/ahKevinXy/go-cmb/help"
	"github.com/ahKevinXy/go-cmb/models"
	"strconv"
	"time"
)

// GetUnitAccountTransList
//  @Description: 记账子单元当天交易查询
//  @param userId
//  @param sm4PrivateKey
//  @param userPrivateKey
//  @param accnbr 账户
//  @param dmanbr 子单元
//  @param ctnkey  续传key
//  @return *models.UnitAccountTransDailyResponse
//  @return error
//  @Author  ahKevinXy
//  @Date  2023-04-13 18:49:18
func GetUnitAccountTransList(userId, sm4PrivateKey, userPrivateKey, accnbr, dmanbr, ctnkey string) (*models.UnitAccountTransDailyResponse, error) {
	reqData := new(models.AccountUnitTransDailyRequest)
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

	res, err := help.CmbSignRequest(string(req), constants.CmbUnitManageAccountTransDaily, userId, userPrivateKey, sm4PrivateKey)
	if err != nil {
		return nil, err
	}
	if res == "" {
		return nil, cmb_errors.SystemError
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
//  @param sm4PrivateKey
//  @param userPrivateKey
//  @param accnbr 账号
//  @param dmanbr 子单元
//  @param begdat 开始时间
//  @param enddat 结束时间
//  @param ctnkey 续传key
//  @return *models.UnitAccountTransHistoryResponse
//  @return error
//  @Author  ahKevinXy
//  @Date  2023-04-13 18:53:33
func GetUnitAccountTransHistoryList(userId, sm4PrivateKey, userPrivateKey, accnbr, dmanbr, begdat, enddat, ctnkey string) (*models.UnitAccountTransHistoryResponse, error) {
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

	res, err := help.CmbSignRequest(string(req), constants.CmbUnitManageAccountTransHistory, userId, userPrivateKey, sm4PrivateKey)
	if err != nil {
		return nil, err
	}
	if res == "" {
		return nil, cmb_errors.SystemError
	}

	var resp *models.UnitAccountTransHistoryResponse

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, errors.New(res)
	}

	return resp, nil
}

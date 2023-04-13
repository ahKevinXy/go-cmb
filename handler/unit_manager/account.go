package unit_manager

import (
	"encoding/json"
	"github.com/ahKevinXy/go-cmb/constants"
	"github.com/ahKevinXy/go-cmb/help"
	"github.com/ahKevinXy/go-cmb/models"
	"strconv"
	"time"
)

// 记账单元 管理

// AddUnitAccount
//  @Description:  新增记账单元
//  @param userId
//  @param asePrivateKey
//  @param userPrivateKey
//  @param accnbr
//  @param dmanam
//  @param dmanbr
//  @param
//  @return *models.AddUnitAccountResponse
//  @return error
//  @Author  ahKevinXy
//  @Date 2023-04-13 17:28:18
func AddUnitAccount(
	userId, asePrivateKey, userPrivateKey, accnbr, dmanam, dmanbr string,

) (*models.AddUnitAccountResponse, error) {
	reqData := new(models.AccountAddUnitRequest)
	reqData.Request.Head.Reqid = time.Now().Format("20060102150405000") + strconv.Itoa(time.Now().Nanosecond())
	reqData.Request.Head.Funcode = constants.CmbUnitManageAddAccount
	reqData.Request.Head.Userid = userId
	reqData.Signature.Sigtim = time.Now().Format("20060102150405")
	reqData.Signature.Sigdat = "__signature_sigdat__"
	reqData.Request.Body.Ntdmaaddx = append(reqData.Request.Body.Ntdmaaddx, &models.Ntdmaaddx{
		Accnbr: accnbr,
		Dmanam: dmanam,
		Dmanbr: dmanbr,
	})

	req, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	//  todo
	res := help.CmbSignRequest(string(req), constants.CmbUnitManageAddAccount, userId, userPrivateKey, asePrivateKey)

	if res == "" {

	}

	var resp models.AddUnitAccountResponse

	if err := json.Unmarshal([]byte(res), &resp); err != nil {

		return nil, err
	}

	return &resp, nil
}

// CloseUnitAccount
//  @Description:  关闭记账单元
//  @param userId
//  @param asePrivateKey
//  @param userPrivateKey
//  @param accnbr
//  @param dmanbr
//  @param
//  @return *models.CloseUnitAccountResponse
//  @return error
//  @Author  ahKevinXy
//  @Date 2023-04-13 17:46:41
func CloseUnitAccount(
	userId,
	asePrivateKey, userPrivateKey,
	accnbr,
	dmanbr string,

) (*models.CloseUnitAccountResponse, error) {
	reqData := new(models.AccountCloseUnitRequest)
	reqData.Request.Head.Reqid = time.Now().Format("20060102150405000") + strconv.Itoa(time.Now().Nanosecond())
	reqData.Request.Head.Funcode = constants.CmbUnitManageCloseAccount
	reqData.Request.Head.Userid = userId
	reqData.Signature.Sigtim = time.Now().Format("20060102150405")
	reqData.Signature.Sigdat = "__signature_sigdat__"
	reqData.Request.Body.Ntdmadltx1 = append(reqData.Request.Body.Ntdmadltx1, &models.Ntdmadltx1{
		Accnbr: accnbr,
	})
	reqData.Request.Body.Ntdmadltx2 = append(reqData.Request.Body.Ntdmadltx2, &models.Ntdmadltx2{
		Dmanbr: dmanbr,
	})

	req, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	//  todo
	res := help.CmbSignRequest(string(req), constants.CmbUnitManageCloseAccount, userId, userPrivateKey, asePrivateKey)

	if res == "" {

	}

	var resp models.CloseUnitAccountResponse

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		print(err)
		return nil, err
	}
	return &resp, err
}

// QueryUnitAccountInfo
//  @Description:  获取账户信息
//  @param userId
//  @param asePrivateKey
//  @param userPrivateKey
//  @param accnbr
//  @param dmanbr
//  @return *models.AccountUnitInfoResponse
//  @return error
//  @Author  ahKevinXy
//  @Date  2023-04-13 17:51:01
func QueryUnitAccountInfo(userId, asePrivateKey, userPrivateKey, accnbr, dmanbr string) (*models.AccountUnitInfoResponse, error) {
	reqData := new(models.AccountUnitInfoRequest)
	reqData.Request.Head.Reqid = time.Now().Format("20060102150405000") + strconv.Itoa(time.Now().Nanosecond())
	reqData.Request.Head.Funcode = constants.CmbUnitManageAccountQuery
	reqData.Request.Head.Userid = userId
	reqData.Signature.Sigtim = time.Now().Format("20060102150405")
	reqData.Signature.Sigdat = "__signature_sigdat__"
	reqData.Request.Body.Ntdmabalx = append(reqData.Request.Body.Ntdmabalx, &models.Ntdmabalx{
		Accnbr: accnbr,
		Dmanbr: dmanbr,
	})

	req, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	//  todo
	res := help.CmbSignRequest(string(req), constants.CmbUnitManageAccountQuery, userId, userPrivateKey, asePrivateKey)

	if res == "" {
		return nil, err
	}

	var resp models.AccountUnitInfoResponse

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, err
	}
	//fmt.Println(res)
	return &resp, err
}

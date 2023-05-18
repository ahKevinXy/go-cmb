package account

import (
	"encoding/json"
	"fmt"
	"github.com/ahKevinXy/go-cmb/cmb_errors"
	"github.com/ahKevinXy/go-cmb/constants"
	"github.com/ahKevinXy/go-cmb/help"
	"github.com/ahKevinXy/go-cmb/models"
	"strconv"
	"time"
)

// MainAccountInfo
// @Description:  获取账户信息
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

		return nil, cmb_errors.SystemError
	}

	var resp models.AccountInfoResponse

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// MainAccountHistoryBalance
//  @Description:   获取历史余额 todo
//  @param userId
//  @param asePrivateKey
//  @param userPrivateKey
//  @param accnbr 主账户
//  @param bbknbr 子单元
//  @param bgndat 开始时间
//  @param enddat 结束时间
//  @param
//  @return string
//  @Author  ahKevinXy
//  @Date2023-04-10 13:48:04
func MainAccountHistoryBalance(
	userId, asePrivateKey,
	userPrivateKey,
	accnbr,
	bbknbr,
	bgndat,
	enddat string,
) (*models.MainAccountHistoryBalanceResponse, error) {
	reqData := new(models.MainAccountHistoryBalanceRequest)
	reqData.Request.Head.Reqid = time.Now().Format("20060102150405000") + strconv.Itoa(time.Now().Nanosecond())
	reqData.Request.Head.Funcode = constants.CmbAccountHistoryBalance
	reqData.Request.Head.Userid = userId
	reqData.Signature.Sigtim = time.Now().Format("20060102150405")
	reqData.Signature.Sigdat = "__signature_sigdat__"
	reqData.Request.Body.Ntqabinfy = append(reqData.Request.Body.Ntqabinfy, &models.Ntqabinfy{
		Accnbr: accnbr,
		Bbknbr: bbknbr,
		Bgndat: bgndat,
		Enddat: enddat,
	})

	req, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	//  todo
	res := help.CmbSignRequest(string(req), constants.CmbAccountHistoryBalance, userId, userPrivateKey, asePrivateKey)

	if res == "" {
		return nil, cmb_errors.SystemError
	}

	var resp models.MainAccountHistoryBalanceResponse

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, err
	}
	fmt.Println(res)
	return &resp, nil
}

// GetBankLinkNo
//  @Description:   获取银联号
//  @param userId 用户ID
//  @param asePrivateKey  对称秘钥
//  @param userPrivateKey  用户私钥
//  @param accnbr  账号
//  @return *models.AccountBankInfoResponse 请求返回
//  @return error
//  @Author  ahKevinXy
//  @Date2023-04-10 13:51:13
func GetBankLinkNo(userId, asePrivateKey, userPrivateKey, accnbr string) (*models.AccountBankInfoResponse, error) {
	reqData := new(models.AccountBankInfoRequest)
	reqData.Request.Head.Reqid = time.Now().Format("20060102150405000") + strconv.Itoa(time.Now().Nanosecond())
	reqData.Request.Head.Funcode = constants.CmbAccountQueryBankLinkNo
	reqData.Request.Head.Userid = userId
	reqData.Signature.Sigtim = time.Now().Format("20060102150405")
	reqData.Signature.Sigdat = "__signature_sigdat__"
	reqData.Request.Body.Fctval = accnbr

	req, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	//  todo
	res := help.CmbSignRequest(string(req), constants.CmbAccountQueryBankLinkNo, userId, userPrivateKey, asePrivateKey)

	if res == "" {
		return nil, cmb_errors.SystemError
	}

	var resp models.AccountBankInfoResponse

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetMainAccountPayBusList
//  @Description:   获取招商以后支付模式 todo
//  @param userId
//  @param asePrivateKey
//  @param userPrivateKey
//  @param buscode
//  @param category
//  @return *models.QueryAccountTransCodeResponse
//  @return error
//  @Author  ahKevinXy
//  @Date2023-04-10 15:16:34
func GetMainAccountPayBusList(userId, asePrivateKey, userPrivateKey, buscode string) (*models.QueryAccountTransCodeResponse, error) {
	reqData := new(models.QueryAccountTransCodeRequest)
	reqData.Request.Head.Reqid = time.Now().Format("20060102150405000") + strconv.Itoa(time.Now().Nanosecond())
	reqData.Request.Head.Funcode = constants.CmbAccountPayModQuery
	reqData.Request.Head.Userid = userId
	reqData.Signature.Sigtim = time.Now().Format("20060102150405")
	reqData.Signature.Sigdat = "__signature_sigdat__"
	reqData.Request.Body.Buscod = buscode

	req, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	//  todo
	res := help.CmbSignRequest(string(req), constants.CmbAccountPayModQuery, userId, userPrivateKey, asePrivateKey)

	if res == "" {

	}

	var resp models.QueryAccountTransCodeResponse

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, err
	}
	//fmt.Println(res)
	return &resp, nil
}

// QueryBatchAccountBalance
//  @Description:  批量获取余额
//  @param userId
//  @param asePrivateKey
//  @param userPrivateKey
//  @param accnbr 账号
//  @param buscode 分行号
//  @return *models.QueryBatchAccountBalanceResponse
//  @return error
//  @Author  ahKevinXy
//  @Date2023-04-13 13:58:14
func QueryBatchAccountBalance(userId, asePrivateKey, userPrivateKey, accnbr, bbknbr string) (*models.QueryBatchAccountBalanceResponse, error) {
	reqData := new(models.QueryBatchMainAccountBalanceRequest)
	reqData.Request.Head.Reqid = time.Now().Format("20060102150405000") + strconv.Itoa(time.Now().Nanosecond())
	reqData.Request.Head.Funcode = constants.CmbAccountBatchQueryBalance
	reqData.Request.Head.Userid = userId
	reqData.Signature.Sigtim = time.Now().Format("20060102150405")
	reqData.Signature.Sigdat = "__signature_sigdat__"

	reqData.Request.Body.Ntqadinfx = append(reqData.Request.Body.Ntqadinfx, &models.Ntqadinfx{
		Accnbr: accnbr,
		Bbknbr: bbknbr,
	})
	req, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	//  todo
	res := help.CmbSignRequest(string(req), constants.CmbAccountBatchQueryBalance, userId, userPrivateKey, asePrivateKey)

	if res == "" {

	}
	var resp models.QueryBatchAccountBalanceResponse

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

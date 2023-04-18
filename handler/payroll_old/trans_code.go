package payroll_old

import (
	"encoding/json"
	"github.com/ahKevinXy/go-cmb/constants"
	"github.com/ahKevinXy/go-cmb/help"
	"github.com/ahKevinXy/go-cmb/models"
	"strconv"
	"time"
)

// PayMods
//  @Description:   获取交易模式
//  @param userId
//  @param asePrivateKey
//  @param userPrivateKey
//  @param busCode
//  @param accnbr
//  @return *models.QueryPayrollOldTransCodeResponse
//  @return error
//  @Author  ahKevinXy
//  @Date  2023-04-18 09:45:07
func PayMods(userId, asePrivateKey, userPrivateKey, busCode, accnbr string) (*models.QueryPayrollOldTransCodeResponse, error) {

	reqData := new(models.QueryPayrollOldTransCodeRequest)
	reqData.Request.Head.Reqid = time.Now().Format("20060102150405000") + strconv.Itoa(time.Now().Nanosecond())
	reqData.Request.Head.Funcode = constants.CmbPayrollOldTransCode
	reqData.Request.Head.Userid = userId
	reqData.Signature.Sigtim = time.Now().Format("20060102150405")
	reqData.Signature.Sigdat = "__signature_sigdat__"
	reqData.Request.Body.Ntagtls2x = append(reqData.Request.Body.Ntagtls2x, models.Ntagtls2x{
		Buscod: busCode,
		Accnbr: accnbr,
		//Ccynbr: ccynbr,
	})
	req, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}
	res := help.CmbSignRequest(string(req), constants.CmbPayrollOldTransCode, userId, userPrivateKey, asePrivateKey)

	if res == "" {

	}

	var resp models.QueryPayrollOldTransCodeResponse

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		return nil, err
	}
	//fmt.Println(res)
	return &resp, nil
}

package models

// MainAccountPayResultResponse
//  @Description:  支付结果
//  @Author  ahKevinXy
//  @Date2023-04-10 15:12:08
type MainAccountPayResultResponse struct {
	Response struct {
		Body struct {
			Bb1Qrybtz1 []struct {
				BthNbr string `json:"bthNbr,omitempty"`
				BusCod string `json:"busCod,omitempty"`
				BusMod string `json:"busMod,omitempty"`
				DtlAmt string `json:"dtlAmt,omitempty"`
				DtlNum string `json:"dtlNum,omitempty"`
				ReqSts string `json:"reqSts,omitempty"`
				Rsv30Z string `json:"rsv30z,omitempty"`
				RtnFlg string `json:"rtnFlg,omitempty"`
				SucAmt string `json:"sucAmt,omitempty"`
				SucNum string `json:"sucNum,omitempty"`
				TrsDat string `json:"trsDat,omitempty"`
				TrsTim string `json:"trsTim,omitempty"`
			} `json:"bb1qrybtz1,omitempty"`
		} `json:"body,omitempty"`
		Head struct {
			Bizcode    string `json:"bizcode,omitempty"`
			Funcode    string `json:"funcode,omitempty"`
			Reqid      string `json:"reqid,omitempty"`
			Resultcode string `json:"resultcode,omitempty"`
			Resultmsg  string `json:"resultmsg,omitempty"`
			Rspid      string `json:"rspid,omitempty"`
			Userid     string `json:"userid,omitempty"`
		} `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

package models

// MainAccountSinglePayResponse
//  @Description:   单笔经办支付 返回
//  @Author  ahKevinXy
//  @Date2023-04-10 13:54:19
type MainAccountSinglePayResponse struct {
	Response struct {
		Body struct {
			Bb1Payopz1 []struct {
				ErrCod string `json:"errCod,omitempty"`
				EvtIst string `json:"evtIst,omitempty"`
				PrdIst string `json:"prdIst,omitempty"`
				ReqNbr string `json:"reqNbr,omitempty"`
				ReqSts string `json:"reqSts,omitempty"`
				RtnFlg string `json:"rtnFlg,omitempty"`
			} `json:"bb1payopz1,omitempty"`
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

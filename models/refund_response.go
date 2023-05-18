package models

type QueryAccountPaymentRefundResponse struct {
	Response struct {
		Body struct {
			Bb1Payqby1 []struct {
				AccNbr string `json:"accNbr,omitempty"` // 账号
				BbkNbr string `json:"bbkNbr,omitempty"` // 分行号
				BgnDat string `json:"bgnDat,omitempty"` // 开始时间
				CtnKey string `json:"ctnKey,omitempty"` //续传key
				EndDat string `json:"endDat,omitempty"` // 结束时间
				ReqNbr string `json:"reqNbr,omitempty"` // 流程实例号
				Rsv50Z string `json:"rsv50z,omitempty"` //保留字
			} `json:"bb1payqby1,omitempty"`
		} `json:"body,omitempty"`
		Bb1payqbz1 []struct {
			ReqNbr string `json:"reqNbr,omitempty"` // 流程实例号
			YurRef string `json:"yurRef,omitempty"`
			TrsAmt string `json:"trsAmt,omitempty"` // 交易金额
			FeeAmt string `json:"feeAmt,omitempty"` // 费用总额
		} `json:"bb1payqbz1"`
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

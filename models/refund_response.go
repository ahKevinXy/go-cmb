package models

type QueryAccountPaymentRefundResponse struct {
	Response struct {
		Body struct {
			Bb1Payqby1 []struct {
				AccNbr string `json:"accNbr,omitempty"`
				BbkNbr string `json:"bbkNbr,omitempty"`
				BgnDat string `json:"bgnDat,omitempty"`
				CtnKey string `json:"ctnKey,omitempty"`
				EndDat string `json:"endDat,omitempty"`
				ReqNbr string `json:"reqNbr,omitempty"`
				Rsv50Z string `json:"rsv50z,omitempty"`
			} `json:"bb1payqby1,omitempty"`
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

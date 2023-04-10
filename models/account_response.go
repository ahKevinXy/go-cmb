package models

type AccountInfoResponse struct {
	Response struct {
		Body struct {
			Ntqacinfz []struct {
				Accblv string `json:"accblv"`
				Accitm string `json:"accitm"`
				Accnam string `json:"accnam"`
				Accnbr string `json:"accnbr"`
				Avlblv string `json:"avlblv"`
				Bbknbr string `json:"bbknbr"`
				Ccynbr string `json:"ccynbr"`
				Dactyp string `json:"dactyp"`
				Hldblv string `json:"hldblv"`
				Intrat string `json:"intrat"`
				Lmtovr string `json:"lmtovr"`
				Mutdat string `json:"mutdat"`
				Onlblv string `json:"onlblv"`
				Opnbbk string `json:"opnbbk"`
				Opnbrn string `json:"opnbrn"`
				Opndat string `json:"opndat"`
				Relnbr string `json:"relnbr"`
				Stscod string `json:"stscod"`
			} `json:"ntqacinfz"`
		} `json:"body"`
		Head struct {
			Bizcode    string `json:"bizcode"`
			Funcode    string `json:"funcode"`
			Reqid      string `json:"reqid"`
			Resultcode string `json:"resultcode"`
			Resultmsg  string `json:"resultmsg"`
			Rspid      string `json:"rspid"`
			Userid     string `json:"userid"`
		} `json:"head"`
	} `json:"response"`
}

// MainAccountHistoryBalanceResponse
//  @Description:  获取历史余额返回
//  @Author  ahKevinXy
//  @Date2023-04-10 13:46:46
type MainAccountHistoryBalanceResponse struct {
	Response struct {
		Body struct {
			Ntqabinfz []struct {
				Accnbr string `json:"accnbr,omitempty"`
				Balamt string `json:"balamt,omitempty"`
				Bbknbr string `json:"bbknbr,omitempty"`
				Rsv30Z string `json:"rsv30z,omitempty"`
				Trsdat string `json:"trsdat,omitempty"`
			} `json:"ntqabinfz,omitempty"`
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

// AccountBankInfoResponse
//  @Description:   银联号返回
//  @Author  ahKevinXy
//  @Date2023-04-10 13:50:05
type AccountBankInfoResponse struct {
	Response struct {
		Body struct {
			Ntaccbbky []struct {
				Fcttyp string `json:"fcttyp,omitempty"`
				Fctval string `json:"fctval,omitempty"`
			} `json:"ntaccbbky,omitempty"`
			Ntaccbbkz []struct {
				Bbknbr string `json:"bbknbr,omitempty"`
				Fctval string `json:"fctval,omitempty"`
			} `json:"ntaccbbkz,omitempty"`
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

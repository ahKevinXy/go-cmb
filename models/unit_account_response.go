package models

// AddUnitAccountResponse
//  @Description:  添加记账单元
//  @Author  ahKevinXy
//  @Date2023-04-13 17:43:53
type AddUnitAccountResponse struct {
	Response struct {
		Body struct {
			Ntdumaddz1 []struct {
				Dyanbr string `json:"dyanbr,omitempty"`
				Errcod string `json:"errcod,omitempty"`
				Inbacc string `json:"inbacc,omitempty"`
				Reqnbr string `json:"reqnbr,omitempty"`
				Reqsts string `json:"reqsts,omitempty"`
			} `json:"ntdumaddz1,omitempty"`
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

// CloseUnitAccountResponse
//  @Description:  关闭记账单元
//  @Author  ahKevinXy
//  @Date2023-04-13 17:44:38
type CloseUnitAccountResponse struct {
	Response struct {
		Body struct {
			Ntdumdltz1 []struct {
				Dyanbr string `json:"dyanbr,omitempty"`
				Errcod string `json:"errcod,omitempty"`
				Inbacc string `json:"inbacc,omitempty"`
				Reqnbr string `json:"reqnbr,omitempty"`
				Reqsts string `json:"reqsts,omitempty"`
			} `json:"ntdumdltz1,omitempty"`
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

// AccountUnitInfoResponse
//  @Description:  获取余额信息
//  @Author  ahKevinXy
//  @Date  2023-04-13 17:50:47
type AccountUnitInfoResponse struct {
	Response struct {
		Body struct {
			Ntdmabalz []struct {
				Accnbr string `json:"accnbr"`
				Actbal string `json:"actbal"`
				Dmaccy string `json:"dmaccy"`
				Dmanam string `json:"dmanam"`
				Dmanbr string `json:"dmanbr"`
				Eftdat string `json:"eftdat"`
				Stscod string `json:"stscod"`
			} `json:"ntdmabalz"`
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

// QueryUnitTransByBusNoResponse
//  @Description:  获取交易信息
//  @Author  ahKevinXy
//  @Date  2023-04-13 18:04:30
type QueryUnitTransByBusNoResponse struct {
	Response struct {
		Body struct {
			Ntdmaqryz1 []struct {
				Bbknbr string `json:"bbknbr,omitempty"`
				Buscod string `json:"buscod,omitempty"`
				Busmod string `json:"busmod,omitempty"`
				Inbacc string `json:"inbacc,omitempty"`
				Oprdat string `json:"oprdat,omitempty"`
				Oprtyp string `json:"oprtyp,omitempty"`
				Reqnbr string `json:"reqnbr,omitempty"`
				Reqsta string `json:"reqsta,omitempty"`
				Trsamt string `json:"trsamt,omitempty"`
				Yurref string `json:"yurref,omitempty"`
			} `json:"ntdmaqryz1,omitempty"`
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

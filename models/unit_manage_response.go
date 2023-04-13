package models

type UnitAccountPayInResponse struct {
	Response struct {
		Body struct {
			Ntoprrtnz []struct {
				Errcod string `json:"errcod,omitempty"`
				Reqnbr string `json:"reqnbr,omitempty"`
				Reqsts string `json:"reqsts,omitempty"`
			} `json:"ntoprrtnz,omitempty"`
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

// QueryUnitAccountDetailResponse
//  @Description:  详细信息
//  @Author  ahKevinXy
//  @Date  2023-04-13 18:11:53
type QueryUnitAccountDetailResponse struct {
	Response struct {
		Body struct {
			Ntdumadxz1 []struct {
				Bbknbr string `json:"bbknbr,omitempty"`
				Dyanam string `json:"dyanam,omitempty"`
				Dyanbr string `json:"dyanbr,omitempty"`
				Eftdat string `json:"eftdat,omitempty"`
				Enddat string `json:"enddat,omitempty"`
				Inbacc string `json:"inbacc,omitempty"`
				Ovrctl string `json:"ovrctl,omitempty"`
			} `json:"ntdumadxz1,omitempty"`
			Ntduminfz1 []struct {
				Athflg string `json:"athflg,omitempty"`
				Buscod string `json:"buscod,omitempty"`
				Busmod string `json:"busmod,omitempty"`
				Lgnnam string `json:"lgnnam,omitempty"`
				Oprtyp string `json:"oprtyp,omitempty"`
				Reqnbr string `json:"reqnbr,omitempty"`
				Reqsta string `json:"reqsta,omitempty"`
				Usrnam string `json:"usrnam,omitempty"`
				Yurref string `json:"yurref,omitempty"`
			} `json:"ntduminfz1,omitempty"`
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

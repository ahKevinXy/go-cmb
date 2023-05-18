package models

// QueryAccountTransCodeResponse
//  @Description:   交易代码
//  @Author  ahKevinXy
//  @Date  2023-04-26 20:24:34
type QueryAccountTransCodeResponse struct {
	Response struct {
		Body struct {
			Ntqmdlstz []struct {
				Busmod string `json:"busmod"` // 业务模式编号
				Modals string `json:"modals"` // 业务模式名称
			} `json:"ntqmdlstz"`
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

// HeadResponse
//  @Description:   Head 通用回复
//  @Author  ahKevinXy
//  @Date  2023-05-18 09:32:28
type HeadResponse struct {
	Bizcode    string `json:"bizcode"`
	Funcode    string `json:"funcode"`
	Reqid      string `json:"reqid"`
	Resultcode string `json:"resultcode"`
	Resultmsg  string `json:"resultmsg"`
	Rspid      string `json:"rspid"`
	Userid     string `json:"userid"`
}

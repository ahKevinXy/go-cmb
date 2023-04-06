package models

type QueryAccountTransCodeResponse struct {
	Response struct {
		Body struct {
			Ntqmdlstz []struct {
				Busmod string `json:"busmod"`
				Modals string `json:"modals"`
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

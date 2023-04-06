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

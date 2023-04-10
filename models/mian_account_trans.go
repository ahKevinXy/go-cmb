package models

type GetMainAccountTransInfoResponse struct {
	Response struct {
		Body struct {
			Ntqactrsz2 []struct {
				Amtcdr string `json:"amtcdr"`
				Athflg string `json:"athflg,omitempty"`
				Bbknbr string `json:"bbknbr"`
				Busnam string `json:"busnam,omitempty"`
				Etydat string `json:"etydat"`
				Etytim string `json:"etytim"`
				Infflg string `json:"infflg"`
				Narext string `json:"narext"`
				Naryur string `json:"naryur"`
				Refnbr string `json:"refnbr"`
				Refsub string `json:"refsub"`
				Reqnbr string `json:"reqnbr,omitempty"`
				Rpyacc string `json:"rpyacc"`
				Rpyadr string `json:"rpyadr"`
				Rpybnk string `json:"rpybnk"`
				Rpynam string `json:"rpynam"`
				Rsv30Z string `json:"rsv30z"`
				Trsamt string `json:"trsamt"`
				Trsanl string `json:"trsanl"`
				Trsblv string `json:"trsblv"`
				Trscod string `json:"trscod"`
				Vltdat string `json:"vltdat"`
				Yurref string `json:"yurref"`
				Chknbr string `json:"chknbr,omitempty"`
			} `json:"ntqactrsz2"`
			Ntrbptrsz1 []struct {
				Cotflg string `json:"cotflg"`
				Crtamt string `json:"crtamt"`
				Crtnbr string `json:"crtnbr"`
				Dbtamt string `json:"dbtamt"`
				Dbtnbr string `json:"dbtnbr"`
				Trsseq string `json:"trsseq"`
			} `json:"ntrbptrsz1"`
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

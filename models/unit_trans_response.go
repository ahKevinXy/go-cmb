package models

type UnitAccountTransDailyResponse struct {
	Response struct {
		Body struct {
			Ntdmtlstz []struct {
				Accnbr string `json:"accnbr,omitempty"`
				Autflg string `json:"autflg,omitempty"`
				Ccynbr string `json:"ccynbr,omitempty"`
				Dmanam string `json:"dmanam,omitempty"`
				Dmanbr string `json:"dmanbr,omitempty"`
				Rpyacc string `json:"rpyacc,omitempty"`
				Rpynam string `json:"rpynam,omitempty"`
				Trxamt string `json:"trxamt,omitempty"`
				Trxdir string `json:"trxdir,omitempty"`
				Onlbal string `json:"onlbal"`
				Trxnbr string `json:"trxnbr,omitempty"`
				Trxtim string `json:"trxtim,omitempty"`
				Trxtxt string `json:"trxtxt,omitempty"`
				Narinn string `json:"narinn,omitempty"`
			} `json:"ntdmtlstz,omitempty"`
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

type UnitAccountTransHistoryResponse struct {
	Response struct {
		Body struct {
			Ntdmthlsz []struct {
				Accnbr string `json:"accnbr"`
				Autflg string `json:"autflg"`
				Ccynbr string `json:"ccynbr"`
				Dmanam string `json:"dmanam"`
				Dmanbr string `json:"dmanbr"`
				Rpyacc string `json:"rpyacc"`
				Rpynam string `json:"rpynam"`
				Trxamt string `json:"trxamt"`
				Trxdat string `json:"trxdat"`
				Trxdir string `json:"trxdir"`
				Trxnbr string `json:"trxnbr"`
				Trxtim string `json:"trxtim"`
				Trxtxt string `json:"trxtxt"`
				Narinn string `json:"narinn,omitempty"`
				Rltdat string `json:"rltdat,omitempty"`
				Onlbal string `json:"onlbal,omitempty"`
			} `json:"ntdmthlsz"`
			Ntdmthlsy []*Ntdmthlsy `json:"ntdmthlsy,omitempty"`
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

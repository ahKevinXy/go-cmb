package models

type QueryPayrollOldTransCodeResponse struct {
	Response struct {
		Body struct {
			Ntagtls2Z []struct {
				Accnbr  string `json:"accnbr,omitempty"`
				CTrstyp string `json:"c_trstyp,omitempty"`
				Ccynbr  string `json:"ccynbr,omitempty"`
				Cnvnbr  string `json:"cnvnbr,omitempty"`
				Eftdat  string `json:"eftdat,omitempty"`
				Expdat  string `json:"expdat,omitempty"`
				Sgndat  string `json:"sgndat,omitempty"`
				Stscod  string `json:"stscod,omitempty"`
				Trstyp  string `json:"trstyp,omitempty"`
			} `json:"ntagtls2z,omitempty"`
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

type PayrollOldCreditOtherBySupResponse struct {
	Response struct {
		Body struct {
			Ntagcagcz1 []struct {
				Reqnbr string `json:"reqnbr,omitempty"`
				Reqsta string `json:"reqsta,omitempty"`
				Rsv50Z string `json:"rsv50z,omitempty"`
			} `json:"ntagcagcz1,omitempty"`
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

type GetPayrollPdfResponse struct {
	Response struct {
		Body struct {
			Begidx  string `json:"begidx,omitempty"`
			Printid string `json:"printid,omitempty"`
			Total   string `json:"total,omitempty"`
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

type QueryPayRollDetailResponse struct {
	Response struct {
		Body struct {
			Ntagcinny1 []struct {
				Bgndat string `json:"bgndat,omitempty"`
				Buscod string `json:"buscod,omitempty"`
				Busmod string `json:"busmod,omitempty"`
				Datflg string `json:"datflg,omitempty"`
				Enddat string `json:"enddat,omitempty"`
				Ctnkey string `json:"ctnkey,omitempty"`
			} `json:"ntagcinny1,omitempty"`
			Ntagcinqz []struct {
				Accnbr string `json:"accnbr,omitempty"`
				Accnam string `json:"accnam,omitempty"`
				Agctyp string `json:"agctyp,omitempty"`
				Athflg string `json:"athflg,omitempty"`
				Bbknbr string `json:"bbknbr,omitempty"`
				Buscod string `json:"buscod,omitempty"`
				Busmod string `json:"busmod,omitempty"`
				Ccynbr string `json:"ccynbr,omitempty"`
				Eptdat string `json:"eptdat,omitempty"`
				Epttim string `json:"epttim,omitempty"`
				Nusage string `json:"nusage,omitempty"`
				Oprals string `json:"oprals,omitempty"`
				Oprdat string `json:"oprdat,omitempty"`
				Oprsqn string `json:"oprsqn,omitempty"`
				Oprstp string `json:"oprstp,omitempty"`
				Rtnflg string `json:"rtnflg,omitempty"`
				Reqnbr string `json:"reqnbr,omitempty"`
				Reqsta string `json:"reqsta,omitempty"`
				Seqcod string `json:"seqcod,omitempty"`
				Stscod string `json:"stscod,omitempty"`
				Sucamt string `json:"sucamt,omitempty"`
				Sucnum string `json:"sucnum,omitempty"`
				Totamt string `json:"totamt,omitempty"`
				Dmanbr string `json:"dmanbr,omitempty"`
				Trsnum string `json:"trsnum,omitempty"`
				Trstyp string `json:"trstyp,omitempty"`
				Yurref string `json:"yurref,omitempty"`
			} `json:"ntagcinqz,omitempty"`
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

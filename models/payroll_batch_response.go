package models

// QueryBatchTransInfoResponse
//  @Description:   批次信息查询
//  @Author  ahKevinXy
//  @Date  2023-04-14 14:46:26
type QueryBatchTransInfoResponse struct {
	Response struct {
		Body struct {
			Bb6Bpdqyz1 []struct {
				Accnam string `json:"accnam,omitempty"`
				Accnbr string `json:"accnbr,omitempty"`
				Agctyp string `json:"agctyp,omitempty"`
				Athflg string `json:"athflg,omitempty"`
				Bbknbr string `json:"bbknbr,omitempty"`
				Buscod string `json:"buscod,omitempty"`
				Busmod string `json:"busmod,omitempty"`
				Ccynbr string `json:"ccynbr,omitempty"`
				Dmanbr string `json:"dmanbr,omitempty"`
				Eptdat string `json:"eptdat,omitempty"`
				Epttim string `json:"epttim,omitempty"`
				Errdsp string `json:"errdsp,omitempty"`
				Nusage string `json:"nusage,omitempty"`
				Oprdat string `json:"oprdat,omitempty"`
				Oprsqn string `json:"oprsqn,omitempty"`
				Oprstp string `json:"oprstp,omitempty"`
				Reqnbr string `json:"reqnbr,omitempty"`
				Reqsta string `json:"reqsta,omitempty"`
				Rtnflg string `json:"rtnflg,omitempty"`
				Seqcod string `json:"seqcod,omitempty"`
				Stscod string `json:"stscod,omitempty"`
				Sucamt string `json:"sucamt,omitempty"`
				Sucnum string `json:"sucnum,omitempty"`
				Totamt string `json:"totamt,omitempty"`
				Trsnum string `json:"trsnum,omitempty"`
				Trsreq string `json:"trsreq,omitempty"`
				Trsset string `json:"trsset,omitempty"`
				Trstyp string `json:"trstyp,omitempty"`
				Yurref string `json:"yurref,omitempty"`
			} `json:"bb6bpdqyz1,omitempty"`
			Bb6Bpdqyz2 []struct {
				Accnam string `json:"accnam,omitempty"`
				Accnbr string `json:"accnbr,omitempty"`
				Bthnbr string `json:"bthnbr,omitempty"`
				Stscod string `json:"stscod,omitempty"`
				Trsamt string `json:"trsamt,omitempty"`
				Trsdat string `json:"trsdat,omitempty"`
				Trsdsp string `json:"trsdsp,omitempty"`
				Trxseq string `json:"trxseq,omitempty"`
			} `json:"bb6bpdqyz2,omitempty"`
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

type QueryBatchTransListResponse struct {
	Response struct {
		Body struct {
			Bb6Bthqyy1 []struct {
				Bgndat string `json:"bgndat,omitempty"`
				Buscod string `json:"buscod,omitempty"`
				Cntkey string `json:"cntkey,omitempty"`
				Dattyp string `json:"dattyp,omitempty"`
				Enddat string `json:"enddat,omitempty"`
			} `json:"bb6bthqyy1,omitempty"`
			Bb6Bthqyz1 []struct {
				Accnam string `json:"accnam,omitempty"`
				Accnbr string `json:"accnbr,omitempty"`
				Athflg string `json:"athflg,omitempty"`
				Bbknbr string `json:"bbknbr,omitempty"`
				Buscod string `json:"buscod,omitempty"`
				Busmod string `json:"busmod,omitempty"`
				Ccynbr string `json:"ccynbr,omitempty"`
				Dmanbr string `json:"dmanbr,omitempty"`
				Eptdat string `json:"eptdat,omitempty"`
				Epttim string `json:"epttim,omitempty"`
				Errdsp string `json:"errdsp,omitempty"`
				Lgnnam string `json:"lgnnam,omitempty"`
				Nusage string `json:"nusage,omitempty"`
				Oprdat string `json:"oprdat,omitempty"`
				Oprsqn string `json:"oprsqn,omitempty"`
				Reqnbr string `json:"reqnbr,omitempty"`
				Reqsta string `json:"reqsta,omitempty"`
				Rtnflg string `json:"rtnflg,omitempty"`
				Seqcod string `json:"seqcod,omitempty"`
				Sucamt string `json:"sucamt,omitempty"`
				Sucnum string `json:"sucnum,omitempty"`
				Totamt string `json:"totamt,omitempty"`
				Trsnum string `json:"trsnum,omitempty"`
				Trsreq string `json:"trsreq,omitempty"`
				Trsset string `json:"trsset,omitempty"`
				Trstyp string `json:"trstyp,omitempty"`
				Usrnam string `json:"usrnam,omitempty"`
				Yurref string `json:"yurref,omitempty"`
			} `json:"bb6bthqyz1,omitempty"`
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

// QueryPayrollTransDetailResponse
type QueryPayrollTransDetailResponse struct {
	Response struct {
		Body struct {
			Bb6Dtlqyy1 []struct {
				Bthnbr string `json:"bthnbr,omitempty"`
				Reqnbr string `json:"reqnbr,omitempty"`
				Trxseq string `json:"trxseq,omitempty"`
			} `json:"bb6dtlqyy1,omitempty"`
			Bb6Dtlqyz1 []struct {
				Accnam string `json:"accnam,omitempty"`
				Accnbr string `json:"accnbr,omitempty"`
				Bthnbr string `json:"bthnbr,omitempty"`
				Stscod string `json:"stscod,omitempty"`
				Trsamt string `json:"trsamt,omitempty"`
				Trsdat string `json:"trsdat,omitempty"`
				Trsdsp string `json:"trsdsp,omitempty"`
				Trxseq string `json:"trxseq,omitempty"`
			} `json:"bb6dtlqyz1,omitempty"`
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

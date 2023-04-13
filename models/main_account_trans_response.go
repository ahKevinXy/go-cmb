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

// QueryAccountTransInfoResponse
//  @Description: 账户交易信息查询GetTransInfo 返回信息
//  @Author  ahKevinXy
//  @Date2023-04-13 15:29:39
type QueryAccountTransInfoResponse struct {
	Response struct {
		Body struct {
			Ntqtsinfz []struct {
				Amtcdr   string `json:"amtcdr,omitempty"`
				Apdflg   string `json:"apdflg,omitempty"`
				Athflg   string `json:"athflg,omitempty"`
				Bbknbr   string `json:"bbknbr,omitempty"`
				Busnam   string `json:"busnam,omitempty"`
				CAthflg  string `json:"c_athflg,omitempty"`
				CBbknbr  string `json:"c_bbknbr,omitempty"`
				CEtydat  string `json:"c_etydat,omitempty"`
				CGsbbbk  string `json:"c_gsbbbk,omitempty"`
				CRpybbk  string `json:"c_rpybbk,omitempty"`
				CTrsamt  string `json:"c_trsamt,omitempty"`
				CTrsamtc string `json:"c_trsamtc,omitempty"`
				CTrsblv  string `json:"c_trsblv,omitempty"`
				CVltdat  string `json:"c_vltdat,omitempty"`
				Etydat   string `json:"etydat,omitempty"`
				Etytim   string `json:"etytim,omitempty"`
				Narext   string `json:"narext,omitempty"`
				Naryur   string `json:"naryur,omitempty"`
				Refnbr   string `json:"refnbr,omitempty"`
				Refsub   string `json:"refsub,omitempty"`
				Rpyacc   string `json:"rpyacc,omitempty"`
				Rpyadr   string `json:"rpyadr,omitempty"`
				Rpybnk   string `json:"rpybnk,omitempty"`
				Rpynam   string `json:"rpynam,omitempty"`
				Rsv30Z   string `json:"rsv30z,omitempty"`
				Rsv31Z   string `json:"rsv31z,omitempty"`
				Rsvflg   string `json:"rsvflg,omitempty"`
				Trsamt   string `json:"trsamt,omitempty"`
				Trsamtc  string `json:"trsamtc,omitempty"`
				Trsanl   string `json:"trsanl,omitempty"`
				Trsblv   string `json:"trsblv,omitempty"`
				Trscod   string `json:"trscod,omitempty"`
				Vltdat   string `json:"vltdat,omitempty"`
				Yurref   string `json:"yurref,omitempty"`
			} `json:"ntqtsinfz,omitempty"`
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

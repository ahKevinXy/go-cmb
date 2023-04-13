package models

// CommonNotice
//  @Description:  通用 通知消息
//  @Author  ahKevinXy
//  @Date2023-04-11 18:20:50
type CommonNotice struct {
	Sigtim string `json:"sigtim"`
	Sigdat string `json:"sigdat"`
	Notdat string `json:"notdat"`
	Notkey string `json:"notkey"`
	Usrnbr string `json:"usrnbr"`
	Notnbr string `json:"notnbr"`
	Nottyp string `json:"nottyp"`
}

type AccountNotice struct {
	Msgdat struct {
		Chknbr  string `json:"chknbr,omitempty"`
		Infflg  string `json:"infflg,omitempty"`
		Refsub  string `json:"refsub,omitempty"`
		Refnbr  string `json:"refnbr,omitempty"`
		Rpyacc  string `json:"rpyacc,omitempty"`
		Trscod  string `json:"trscod,omitempty"`
		Gsbacc  string `json:"gsbacc,omitempty"`
		Otrnar  string `json:"otrnar,omitempty"`
		Rpynam  string `json:"rpynam,omitempty"`
		Amtcdr  string `json:"amtcdr,omitempty"`
		Naryur  string `json:"naryur,omitempty"`
		Vltdat  string `json:"vltdat,omitempty"`
		Yurref  string `json:"yurref,omitempty"`
		Accnam  string `json:"accnam,omitempty"`
		Gsbnam  string `json:"gsbnam,omitempty"`
		Narext  string `json:"narext,omitempty"`
		Trsanl  string `json:"trsanl,omitempty"`
		Nusage  string `json:"nusage,omitempty"`
		Trsdat  string `json:"trsdat,omitempty"`
		Reqnbr  string `json:"reqnbr,omitempty"`
		Trstim  string `json:"trstim,omitempty"`
		Rpybnk  string `json:"rpybnk,omitempty"`
		Gsbbbk  string `json:"gsbbbk,omitempty"`
		Frmcod  string `json:"frmcod,omitempty"`
		Athflg  string `json:"athflg,omitempty"`
		Rpybbn  string `json:"rpybbn,omitempty"`
		Rsvflg  string `json:"rsvflg,omitempty"`
		Accnbr  string `json:"accnbr,omitempty"`
		Busnam  string `json:"busnam,omitempty"`
		Rpybbk  string `json:"rpybbk,omitempty"`
		CTrsamt string `json:"c_trsamt,omitempty"`
		CCcynbr string `json:"c_ccynbr,omitempty"`
		Busnar  string `json:"busnar,omitempty"`
		Blvamt  string `json:"blvamt,omitempty"`
		Rpyadr  string `json:"rpyadr,omitempty"`
	} `json:"msgdat,omitempty"`
	Msgtyp string `json:"msgtyp,omitempty"`
}

type SatementNotice struct {
	Returl string `json:"returl,omitempty"`
	Retcod string `json:"retcod,omitempty"`
	Taskid string `json:"taskid,omitempty"`
}

// PayResultNotice
//  @Description:   支付结果
//  @Author  ahKevinXy
//  @Date2023-04-11 18:31:45
type PayResultNotice struct {
	Msgdat struct {
		TrsInfo struct {
			BnkFlg string `json:"bnkFlg,omitempty"`
			BusNar string `json:"busNar,omitempty"`
			CcyNbr string `json:"ccyNbr,omitempty"`
			CrtAcc string `json:"crtAcc,omitempty"`
			CrtAdr string `json:"crtAdr,omitempty"`
			CrtBbk string `json:"crtBbk,omitempty"`
			CrtBnk string `json:"crtBnk,omitempty"`
			CrtNam string `json:"crtNam,omitempty"`
			DbtAcc string `json:"dbtAcc,omitempty"`
			DbtBbk string `json:"dbtBbk,omitempty"`
			DbtNam string `json:"dbtNam,omitempty"`
			EptDat string `json:"eptDat,omitempty"`
			EptTim string `json:"eptTim,omitempty"`
			NtfCh1 string `json:"ntfCh1,omitempty"`
			NtfCh2 string `json:"ntfCh2,omitempty"`
			NusAge string `json:"nusAge,omitempty"`
			OprDat string `json:"oprDat,omitempty"`
			ReqNbr string `json:"reqNbr,omitempty"`
			ReqSts string `json:"reqSts,omitempty"`
			RtnFlg string `json:"rtnFlg,omitempty"`
			StlChn string `json:"stlChn,omitempty"`
			TrsAmt string `json:"trsAmt,omitempty"`
			TrxSet string `json:"trxSet,omitempty"`
			YurRef string `json:"yurRef,omitempty"`
		} `json:"trsInfo,omitempty"`
	} `json:"msgdat,omitempty"`
	Msgtyp string `json:"msgtyp,omitempty"`
}

// PayrollResultNotice
//  @Description:  代发通知
//  @Author  ahKevinXy
//  @Date2023-04-13 14:21:11
type PayrollResultNotice struct {
	Msgdat struct {
		DetailInfo []struct {
			Accnam string  `json:"accnam,omitempty"`
			Accnbr string  `json:"accnbr,omitempty"`
			Stscod string  `json:"stscod,omitempty"`
			Trsamt float64 `json:"trsamt,omitempty"`
			Trxseq string  `json:"trxseq,omitempty"`
		} `json:"detailInfo,omitempty"`
		AgcInfo struct {
			Accnbr string  `json:"accnbr,omitempty"`
			Bchnbr string  `json:"bchnbr,omitempty"`
			Buscod string  `json:"buscod,omitempty"`
			Ntbnbr string  `json:"ntbnbr,omitempty"`
			Oprdat string  `json:"oprdat,omitempty"`
			Oprusr string  `json:"oprusr,omitempty"`
			Reqnbr string  `json:"reqnbr,omitempty"`
			Rtnflg string  `json:"rtnflg,omitempty"`
			Sntflg string  `json:"sntflg,omitempty"`
			Sucamt float64 `json:"sucamt,omitempty"`
			Sucnum int     `json:"sucnum,omitempty"`
			Totamt float64 `json:"totamt,omitempty"`
			Trsnum int     `json:"trsnum,omitempty"`
			Yurref string  `json:"yurref,omitempty"`
		} `json:"agcInfo,omitempty"`
	} `json:"msgdat,omitempty"`
	Msgtyp string `json:"msgtyp,omitempty"`
}

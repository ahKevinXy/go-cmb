package models

// CommonNotice
//  @Description:  通用 通知消息
//  @Author  ahKevinXy
//  @Date2023-04-11 18:20:50
type CommonNotice struct {
	Sigtim string `json:"sigtim"` // 签名时间
	Sigdat string `json:"sigdat"` // 签名内容
	Notdat string `json:"notdat"` // 通知内容
	Notkey string `json:"notkey"` // 通知键值
	Usrnbr string `json:"usrnbr"` // 用户编号
	Notnbr string `json:"notnbr"` // 通知编号
	Nottyp string `json:"nottyp"` // 通知类型
}

type AccountNotice struct {
	Msgdat struct {
		Chknbr  string `json:"chknbr,omitempty"` // 票据号
		Infflg  string `json:"infflg,omitempty"` // 信息标志  为空表示 付方账号和子账户 为1 表示收方账号和子账户   3 表示原收方账户和子公司
		Refsub  string `json:"refsub,omitempty"` //由商务支付订单产生
		Refnbr  string `json:"refnbr,omitempty"` // 流程实例号
		Rpyacc  string `json:"rpyacc,omitempty"` // 账号
		Trscod  string `json:"trscod,omitempty"` // 交易类型
		Gsbacc  string `json:"gsbacc,omitempty"` // 母/子账户
		Otrnar  string `json:"otrnar,omitempty"` // 其它摘要
		Rpynam  string `json:"rpynam,omitempty"` // 户名 收/付方的转入或转出帐号
		Amtcdr  string `json:"amtcdr,omitempty"` // 借贷标记  C 贷  D 借
		Naryur  string `json:"naryur,omitempty"` // 摘要
		Vltdat  string `json:"vltdat,omitempty"` // 起息日
		Yurref  string `json:"yurref,omitempty"` // 业务参考号
		Accnam  string `json:"accnam,omitempty"` // 户名
		Gsbnam  string `json:"gsbnam,omitempty"` // 母/子公司名称
		Narext  string `json:"narext,omitempty"` //扩展摘要
		Trsanl  string `json:"trsanl,omitempty"` //交易分析码
		Nusage  string `json:"nusage,omitempty"` //用途
		Trsdat  string `json:"trsdat,omitempty"` //交易日期
		Reqnbr  string `json:"reqnbr,omitempty"` //流程实例号 企业银行交易序号，唯一标示企业银行客户端发起的一笔交易
		Trstim  string `json:"trstim,omitempty"` // 交易时间
		Rpybnk  string `json:"rpybnk,omitempty"` ///收/付方开户地区分行号
		Gsbbbk  string `json:"gsbbbk,omitempty"` // 母/子公司所在地区分行
		Frmcod  string `json:"frmcod,omitempty"`
		Athflg  string `json:"athflg,omitempty"`
		Rpybbn  string `json:"rpybbn,omitempty"`
		Rsvflg  string `json:"rsvflg,omitempty"`
		Accnbr  string `json:"accnbr,omitempty"`
		Busnam  string `json:"busnam,omitempty"`
		Rpybbk  string `json:"rpybbk,omitempty"`
		CTrsamt string `json:"c_trsamt,omitempty"` //交易金额
		CCcynbr string `json:"c_ccynbr,omitempty"`
		Busnar  string `json:"busnar,omitempty"`
		Blvamt  string `json:"blvamt,omitempty"` // 余额
		Rpyadr  string `json:"rpyadr,omitempty"`
	} `json:"msgdat,omitempty"`
	Msgtyp string `json:"msgtyp,omitempty"`
}

type SatementNotice struct {
	Returl string `json:"returl,omitempty"` //文件ur
	Retcod string `json:"retcod,omitempty"` //处理结果
	Taskid string `json:"taskid,omitempty"` //打印任务编号
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

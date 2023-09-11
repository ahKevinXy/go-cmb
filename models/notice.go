package models

// CommonNotice
//
//	@Description:  通用 通知消息
//	@Author  ahKevinXy
//	@Date2023-04-11 18:20:50
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
		Chknbr  string `json:"chknbr,omitempty"`   // 票据号
		Infflg  string `json:"infflg,omitempty"`   // 信息标志  为空表示 付方账号和子账户 为1 表示收方账号和子账户   3 表示原收方账户和子公司
		Refsub  string `json:"refsub,omitempty"`   //由商务支付订单产生
		Refnbr  string `json:"refnbr,omitempty"`   // 流程实例号
		Rpyacc  string `json:"rpyacc,omitempty"`   // 账号
		Trscod  string `json:"trscod,omitempty"`   // 交易类型
		Gsbacc  string `json:"gsbacc,omitempty"`   // 母/子账户
		Otrnar  string `json:"otrnar,omitempty"`   // 其它摘要
		Rpynam  string `json:"rpynam,omitempty"`   // 户名 收/付方的转入或转出帐号
		Amtcdr  string `json:"amtcdr,omitempty"`   // 借贷标记  C 贷  D 借
		Naryur  string `json:"naryur,omitempty"`   // 摘要
		Vltdat  string `json:"vltdat,omitempty"`   // 起息日
		Yurref  string `json:"yurref,omitempty"`   // 业务参考号
		Accnam  string `json:"accnam,omitempty"`   // 户名
		Gsbnam  string `json:"gsbnam,omitempty"`   // 母/子公司名称
		Narext  string `json:"narext,omitempty"`   //扩展摘要
		Trsanl  string `json:"trsanl,omitempty"`   //交易分析码
		Nusage  string `json:"nusage,omitempty"`   //用途
		Trsdat  string `json:"trsdat,omitempty"`   //交易日期
		Reqnbr  string `json:"reqnbr,omitempty"`   //流程实例号 企业银行交易序号，唯一标示企业银行客户端发起的一笔交易
		Trstim  string `json:"trstim,omitempty"`   // 交易时间
		Rpybnk  string `json:"rpybnk,omitempty"`   ///收/付方开户地区分行号
		Gsbbbk  string `json:"gsbbbk,omitempty"`   // 母/子公司所在地区分行
		Frmcod  string `json:"frmcod,omitempty"`   //企业识别码
		Athflg  string `json:"athflg,omitempty"`   //有否附件信息标志
		Rpybbn  string `json:"rpybbn,omitempty"`   //收/付方开户行行号 收/付方帐号的开户行的行号
		Rsvflg  string `json:"rsvflg,omitempty"`   // 冲帐标志
		Accnbr  string `json:"accnbr,omitempty"`   //账号
		Busnam  string `json:"busnam,omitempty"`   //业务名称
		Rpybbk  string `json:"rpybbk,omitempty"`   //收/付方开户地区分行号
		CTrsamt string `json:"c_trsamt,omitempty"` //交易金额
		CCcynbr string `json:"c_ccynbr,omitempty"` //币种
		Busnar  string `json:"busnar,omitempty"`   //业务摘要
		Blvamt  string `json:"blvamt,omitempty"`   // 余额
		Rpyadr  string `json:"rpyadr,omitempty"`   //收/付方开户行地址
	} `json:"msgdat,omitempty"`
	Msgtyp string `json:"msgtyp,omitempty"`
}

type StatementNotice struct {
	Returl string `json:"returl,omitempty"` //文件ur
	Retcod string `json:"retcod,omitempty"` //处理结果
	Taskid string `json:"taskid,omitempty"` //打印任务编号
}

// PayResultNotice
//
//	@Description:   支付结果
//	@Author  ahKevinXy
//	@Date2023-04-11 18:31:45
type PayResultNotice struct {
	Msgdat struct {
		TrsInfo struct {
			BnkFlg string `json:"bnkFlg,omitempty"` //系统内外标志 Y 系统内  N 系统内
			BusNar string `json:"busNar,omitempty"` // 业务摘要
			CcyNbr string `json:"ccyNbr,omitempty"` // 币种
			CrtAcc string `json:"crtAcc,omitempty"` // 收方账户
			CrtAdr string `json:"crtAdr,omitempty"` // 收方行地址
			CrtBbk string `json:"crtBbk,omitempty"` // 收方分行号
			CrtBnk string `json:"crtBnk,omitempty"` // 收方行名称
			CrtNam string `json:"crtNam,omitempty"` // 付方账号名
			DbtAcc string `json:"dbtAcc,omitempty"` // 付方账号
			DbtBbk string `json:"dbtBbk,omitempty"` // 转出分行号
			DbtNam string `json:"dbtNam,omitempty"` // 付方账号名称
			EptDat string `json:"eptDat,omitempty"` // 期望日
			EptTim string `json:"eptTim,omitempty"` // 期望时间
			NtfCh1 string `json:"ntfCh1,omitempty"` // 通知方式一
			NtfCh2 string `json:"ntfCh2,omitempty"` // 通知方式二
			NusAge string `json:"nusAge,omitempty"` // 用途
			OprDat string `json:"oprDat,omitempty"` // 经办日期
			ReqNbr string `json:"reqNbr,omitempty"` // 流程实例号
			ReqSts string `json:"reqSts,omitempty"` // 请求状态
			RtnFlg string `json:"rtnFlg,omitempty"` // 业务处理结果  S 成功  F 失败 B  退票  R 否决  D  过期 C 撤销
			StlChn string `json:"stlChn,omitempty"` // 结算通路 G  普通  Q 快速 R 实时 超网
			TrsAmt string `json:"trsAmt,omitempty"` // 交易金额
			TrxSet string `json:"trxSet,omitempty"` // 账务套号
			YurRef string `json:"yurRef,omitempty"` // 业务参考号
		} `json:"trsInfo,omitempty"`
	} `json:"msgdat,omitempty"`
	Msgtyp string `json:"msgtyp,omitempty"`
}

// PayrollResultNotice
//
//	@Description:  代发通知
//	@Author  ahKevinXy
//	@Date2023-04-13 14:21:11
type PayrollResultNotice struct {
	Msgdat struct {
		DetailInfo []struct {
			Accnam string  `json:"accnam,omitempty"` // 账户名称
			Accnbr string  `json:"accnbr,omitempty"` //  账户
			Stscod string  `json:"stscod,omitempty"` // 交易状态 E 失败  A 登记  S  成功
			Trsamt float64 `json:"trsamt,omitempty"` // 明细交易金额
			Trxseq string  `json:"trxseq,omitempty"` // 交易序号
		} `json:"detailInfo,omitempty"`
		AgcInfo struct {
			Accnbr string  `json:"accnbr,omitempty"` // 账号
			Bchnbr string  `json:"bchnbr,omitempty"` // 批次号码
			Buscod string  `json:"buscod,omitempty"` // 业务代码
			Ntbnbr string  `json:"ntbnbr,omitempty"` // 企银编号
			Oprdat string  `json:"oprdat,omitempty"` // 经办时间
			Oprusr string  `json:"oprusr,omitempty"` // 经办用户
			Reqnbr string  `json:"reqnbr,omitempty"` // 流程实例号
			Rtnflg string  `json:"rtnflg,omitempty"` // 请求结果 S 成功 F 失败  C 撤销 D 过期 R 否决
			Sntflg string  `json:"sntflg,omitempty"` // 超网标志
			Sucamt float64 `json:"sucamt,omitempty"` // 成功金额
			Sucnum int     `json:"sucnum,omitempty"` // 成功数目
			Totamt float64 `json:"totamt,omitempty"` // 总金额
			Trsnum int     `json:"trsnum,omitempty"` // 交易笔数
			Yurref string  `json:"yurref,omitempty"` // 业务参考号
		} `json:"agcInfo,omitempty"`
	} `json:"msgdat,omitempty"`
	Msgtyp string `json:"msgtyp,omitempty"`
}

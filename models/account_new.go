package models

import "encoding/json"

// Code generated from CMB official 账务查询 documentation. DO NOT EDIT BY HAND.
// Source: https://openbiz.cmbchina.com/doc-detail?bizkey=DCCT20201214145038074&subclass=1&treeID=2026062909220930543
// Version: 1.0.0, updated: 2026-07-02 08:30:14

type AccountNewRequest struct {
	Request   AccountNewRequestData `json:"request"`
	Signature Signature             `json:"signature"`
}

type AccountNewRequestData struct {
	Body interface{} `json:"body,omitempty"`
	Head Head        `json:"head"`
}

type AccountNewResponse struct {
	Response AccountNewResponseData `json:"response,omitempty"`
}

type AccountNewResponseData struct {
	Body AccountNewRawBody `json:"body,omitempty"`
	Head HeadResponse      `json:"head,omitempty"`
}

type AccountNewRawBody map[string]json.RawMessage

type AccountNewBody map[string]interface{}

// AccountNewNtqmdlstz 接口 ntqmdlstz.
type AccountNewNtqmdlstz struct {
	// 业务模式编号 string(5) required:Y
	Busmod string `json:"busmod,omitempty"`
	// 业务模式名称 Z(62) required:Y
	Modals string `json:"modals,omitempty"`
}

// AccountNewNtqaclstz 接口 ntqaclstz.
type AccountNewNtqaclstz struct {
	// 户名 Z(62) required:Y
	Accnam string `json:"accnam,omitempty"`
	// 帐号 string(35) required:Y
	Accnbr string `json:"accnbr,omitempty"`
	// A.1 招商分行 招商银行分行代码（代码对照表请参照附录） string(2) required:Y
	Bbknbr string `json:"bbknbr,omitempty"`
	// A.3 货币代码表 帐号币种代码（代码对照表请参照附录） string(2) required:Y
	Ccynbr string `json:"ccynbr,omitempty"`
	// 集团网银内部成员编号 string(10)
	Relnbr  string `json:"relnbr,omitempty"`
	CRelnbr string `json:"c_relnbr,omitempty"`
	MAccnbr string `json:"m_accnbr,omitempty"`
}

// AccountNewNtqacinfx 接口 ntqacinfx.
type AccountNewNtqacinfx struct {
	// 账号 string(35) required:Y
	Accnbr string `json:"accnbr,omitempty"`
	// A.1 招商分行 string(2)
	Bbknbr string `json:"bbknbr,omitempty"`
	// 币种 string(2) required:N
	Ccynbr string `json:"ccynbr,omitempty"`
}

// AccountNewNtqacinfz 接口 ntqacinfz.
type AccountNewNtqacinfz struct {
	// 当intcod ='S'时，这个字段的值显示为"头寸额度（集团支付子公司余额）"是子公司的虚拟余额 M required:Y
	Accblv string `json:"accblv,omitempty"`
	// 科目代码 string(5) required:Y
	Accitm string `json:"accitm,omitempty"`
	// 一般为户名 Z(62) required:Y
	Accnam string `json:"accnam,omitempty"`
	// 帐号 string(35) required:Y
	Accnbr string `json:"accnbr,omitempty"`
	// 账户实际可以用来付款的金额，可用余额=联机余额-冻结金额-预期金额+透支额度 M required:N
	Avlblv string `json:"avlblv,omitempty"`
	// A.1 招商分行 string(2) required:Y
	Bbknbr string `json:"bbknbr,omitempty"`
	// A.3 货币代码表 帐号币种代码（代码对照表请参照附录） string(2) required:Y
	Ccynbr string `json:"ccynbr,omitempty"`
	Dactyp string `json:"dactyp,omitempty"`
	// 包含司法冻结余额、银行冻结余额等被冻结的金额加总 M required:N
	Hldblv string `json:"hldblv,omitempty"`
	// 年利率 F(11,7) required:N
	Intrat string `json:"intrat,omitempty"`
	// 透支额度 M required:N
	Lmtovr string `json:"lmtovr,omitempty"`
	// 8位数字 D required:N
	Mutdat string `json:"mutdat,omitempty"`
	// 账户实际的资金 M required:Y
	Onlblv string `json:"onlblv,omitempty"`
	Opnbbk string `json:"opnbbk,omitempty"`
	Opnbrn string `json:"opnbrn,omitempty"`
	// 8位数字 D required:Y
	Opndat string `json:"opndat,omitempty"`
	Relnbr string `json:"relnbr,omitempty"`
	// A=活动，B=冻结，C=关户 string(1) required:Y
	Stscod string `json:"stscod,omitempty"`
	// S=子公司虚拟余额 string(1) required:N
	Intcod string `json:"intcod,omitempty"`
	// ZZZ=不计息 TD2=定期不滚积数 TD1=定期滚积数 D99=手工计息 B01=凭证式国债 D12=挂牌对公活期保证金关户结息 D11=挂牌对公活期保证金按季结息 C01=协议计息 D32=离岸挂牌个人活期存款 D03=挂牌同业活期存款 D02=在岸挂牌个人活期存款 D31=离岸挂牌对公活期存款 D01=在岸挂牌对公活期存款 string(3)
	Inttyp string `json:"inttyp,omitempty"`
	// 定期时，取值： 一天 七天 一个月 三个月 六个月 一年 二年 三年 四年 五年 Z(12)
	Dpstxt string `json:"dpstxt,omitempty"`
}

// AccountNewNtaccbbky 接口 ntaccbbky.
type AccountNewNtaccbbky struct {
	Fcttyp string `json:"fcttyp,omitempty"`
	Fctval string `json:"fctval,omitempty"`
}

// AccountNewNtaccbbkz 接口 ntaccbbkz.
type AccountNewNtaccbbkz struct {
	// 分行号 string(2) required:Y
	Bbknbr string `json:"bbknbr,omitempty"`
	// 账号 string(35) required:Y
	Fctval string `json:"fctval,omitempty"`
}

// AccountNewQueryNoSettlementAccountsX1 接口 QueryNoSettlementAccountsX1.
type AccountNewQueryNoSettlementAccountsX1 struct {
	// 不输入户口号时查询所有非结算账户 String(35)
	CardNbr string `json:"cardNbr,omitempty"`
	// 不输入币种时，可能存在不同币种账号返回 String(2)
	CurrencyNbr string `json:"currencyNbr,omitempty"`
}

// AccountNewQueryNoSettlementAccountsY1 接口 QueryNoSettlementAccountsY1.
type AccountNewQueryNoSettlementAccountsY1 struct {
	// 枚举值Y|N String(1)
	CtnFlag string `json:"ctnFlag,omitempty"`
	// String(10)
	CtnKey string `json:"ctnKey,omitempty"`
}

// AccountNewQueryNoSettlementAccountsZ1 接口 QueryNoSettlementAccountsZ1.
type AccountNewQueryNoSettlementAccountsZ1 struct {
	// String(5)
	AccountCode string `json:"accountCode,omitempty"`
	AccountType string `json:"accountType,omitempty"`
	// 账户实际可以用来付款的金额，可用余额=联机余额-冻结金额-预期金额+透支额度 F(20,2)
	AvailableBalance string `json:"availableBalance,omitempty"`
	// String(2)
	BankNbr string `json:"bankNbr,omitempty"`
	// String(8)
	CardCode string `json:"cardCode,omitempty"`
	// Z(400)
	CardName string `json:"cardName,omitempty"`
	// String(35)
	CardNbr  string `json:"cardNbr,omitempty"`
	ClientId string `json:"clientId,omitempty"`
	// String(2)
	CurrencyNbr string `json:"currencyNbr,omitempty"`
	// 包含司法冻结余额、银行冻结余额等被冻结的金额加总 F(20,2)
	FreezeBalance string `json:"freezeBalance,omitempty"`
	// F (11,7)
	IntRate string `json:"intRate,omitempty"`
	// String(3)
	IntRateType string `json:"intRateType,omitempty"`
	// 账户实际的资金 F(20,2)
	OnlineBalance string `json:"onlineBalance,omitempty"`
	// D
	OpenDate string `json:"openDate,omitempty"`
	// F(20,2)
	OverdraftBalance string `json:"overdraftBalance,omitempty"`
	// F(20,2)
	PostedBalance string `json:"postedBalance,omitempty"`
	// A正常B冻结 String(1)
	Status string `json:"status,omitempty"`
	// Z(20)
	TermCode string `json:"termCode,omitempty"`
	// D
	TermMatureDate string `json:"termMatureDate,omitempty"`
	TermValueDate  string `json:"termValueDate,omitempty"`
	// String(1)
	IntCod string `json:"intCod,omitempty"`
}

// AccountNewTotalAccountDirectReq 接口 TotalAccountDirectReq.
type AccountNewTotalAccountDirectReq struct {
	// 否 客户编号与账号择一个输入 String(10)
	ClientNbr string `json:"clientNbr,omitempty"`
	// 否 客户编号与账号择一个输入 String(35)
	CardNbr string `json:"cardNbr,omitempty"`
	// 否 A:仅活期户 B 仅非结算户 C 所有 仅按客户编号查询时该字段有效 默认为C 所有 String(1)
	QueryType string `json:"queryType,omitempty"`
	// 否 见招行两位币别编码，空则不筛选 String(2)
	Currency string `json:"currency,omitempty"`
	// 否 默认N String(1)
	QueryWarnStatusFlag string `json:"queryWarnStatusFlag,omitempty"`
}

// AccountNewTotalAccountDirectRsp 接口 TotalAccountDirectRsp.
type AccountNewTotalAccountDirectRsp struct {
	// 罗列以下状态，若有多个异常状态，用字母代替，逗号隔开： A-中止非柜面、B-中止全部、C-停止支付 、D-久悬、E-账户冻结、F-网银冻结、G-销户 String(20)
	AccountAlarm string `json:"accountAlarm,omitempty"`
	// String(5)
	AccountCode      string `json:"accountCode,omitempty"`
	AvailableBalance string `json:"availableBalance,omitempty"`
	// 见附录分行编码 String(2)
	BranchNbr string `json:"branchNbr,omitempty"`
	// Z(200)
	CardName string `json:"cardName,omitempty"`
	// String(35)
	CardNbr string `json:"cardNbr,omitempty"`
	// LI: 活期结算户 CD: 通知存款 TD: 普通定期 DT: 保证金 SD: 结构性存款 LAD: 大额存单 OT: 其他 Strng(3)
	CardType string `json:"cardType,omitempty"`
	// 若未关户，为空 D
	ClosedDate  string `json:"closedDate,omitempty"`
	CurrencyNbr string `json:"currencyNbr,omitempty"`
	// 包含司法冻结余额、银行冻结余额等被冻结的金额加总 F(20,2)
	FreezeBalance string `json:"freezeBalance,omitempty"`
	// F(11,7)
	IntRate string `json:"intRate,omitempty"`
	// String(1)
	IntRateType   string `json:"intRateType,omitempty"`
	OnlineBalance string `json:"onlineBalance,omitempty"`
	// F(20,2)
	OpenAmount string `json:"openAmount,omitempty"`
	// yyyyMMdd格式 D
	OpenDate string `json:"openDate,omitempty"`
	// F(20,2)
	OverDraftLimit string `json:"overDraftLimit,omitempty"`
	// 集团网银成员的集团关系编号 普通网银为空 String(10)
	RelationNbr string `json:"relationNbr,omitempty"`
	// A:活动 B:冻结 C:关户 String(1)
	Status string `json:"status,omitempty"`
	// D
	TermMatureDate string `json:"termMatureDate,omitempty"`
	// String(2)
	Currency string `json:"currency,omitempty"`
	// 账户实际的资金 F(20,2)
	OnlieBalance string `json:"onlieBalance,omitempty"`
	// 账户实际可以用来付款的金额，可用余额=联机余额-冻结金额-预期金额+透支额度 F(20,2)
	AvailabelBalance string `json:"availabelBalance,omitempty"`
	// String(20)
	Term string `json:"term,omitempty"`
	// 多个预警状态会通过半角逗号进行拼接，示例如下：“即将久悬,即将终止全部,即将终止非柜面”。如因通信异常等特殊情况，导致预警状态查询失败，返回“-”，稍后重试查询即可。 String(200)
	WarnStatus string `json:"warnStatus,omitempty"`
}

// AccountNewTotalAccountDirectSum 接口 TotalAccountDirectSum.
type AccountNewTotalAccountDirectSum struct {
	SumAvailableBalance string `json:"sumAvailableBalance,omitempty"`
	// F(20,2)
	SumFreezeBalance string `json:"sumFreezeBalance,omitempty"`
	// F(20,2)
	SumOnlineBalance string `json:"sumOnlineBalance,omitempty"`
	// F(20,2)
	SumOverDraftLimit string `json:"sumOverDraftLimit,omitempty"`
	// F(20,2)
	SumAvailabelBalance string `json:"sumAvailabelBalance,omitempty"`
}

// AccountNewNtqabinfy 接口 ntqabinfy.
type AccountNewNtqabinfy struct {
	// 账号 string(35) required:Y
	Accnbr string `json:"accnbr,omitempty"`
	// A.1 招商分行 string(2)
	Bbknbr string `json:"bbknbr,omitempty"`
	// 格式：YYYYMMDD,开始日期与结束日期的间隔不能超过31天 D required:Y
	Bgndat string `json:"bgndat,omitempty"`
	// 格式：YYYYMMDD D required:Y
	Enddat string `json:"enddat,omitempty"`
	// 币种 String(2) required:N
	Ccynbr string `json:"ccynbr,omitempty"`
}

// AccountNewNtqabinfz 接口 ntqabinfz.
type AccountNewNtqabinfz struct {
	// 帐号 string(35) required:Y
	Accnbr string `json:"accnbr,omitempty"`
	// 联机余额 M required:Y
	Balamt string `json:"balamt,omitempty"`
	// A.1 招商分行 string(2) required:Y
	Bbknbr string `json:"bbknbr,omitempty"`
	// 该字段为“币种编码”字段 string(30)
	Rsv30z string `json:"rsv30z,omitempty"`
	// 交易日期 D required:Y
	Trsdat  string `json:"trsdat,omitempty"`
	CCcynbr string `json:"c_ccynbr,omitempty"`
}

// AccountNewNtqryhsbx1 接口 ntqryhsbx1.
type AccountNewNtqryhsbx1 struct {
	// 账号 String(35) required:Y
	Accnbr string `json:"accnbr,omitempty"`
	// 开始日期与结束日期的间隔不能超过31天 D required:Y
	Bgndat string `json:"bgndat,omitempty"`
	// 开始日期和结束日期必须早于当日 D required:Y
	Enddat string `json:"enddat,omitempty"`
	// 空查账号下所有账户的历史余额，指定币种则查该币种账户，如账号下没有该指定币种，则报错 String(2)
	Ccynbr string `json:"ccynbr,omitempty"`
}

// AccountNewNtqryhsbz1 接口 ntqryhsbz1.
type AccountNewNtqryhsbz1 struct {
	// String(35)
	Accnbr string `json:"accnbr,omitempty"`
	// M
	Avlblv string `json:"avlblv,omitempty"`
	// String(2)
	Ccynbr string `json:"ccynbr,omitempty"`
	// M
	Hldblv string `json:"hldblv,omitempty"`
	// M
	Lmtovr string `json:"lmtovr,omitempty"`
	// M
	Onlblv string `json:"onlblv,omitempty"`
	// D
	Trsdat string `json:"trsdat,omitempty"`
}

// AccountNewNtqadinfx 接口 ntqadinfx.
type AccountNewNtqadinfx struct {
	// 账号 string(35) required:Y
	Accnbr string `json:"accnbr,omitempty"`
	// A.1 招商分行 string(2)
	Bbknbr string `json:"bbknbr,omitempty"`
	// 币种 String(2) required:N
	Ccynbr string `json:"ccynbr,omitempty"`
	// 保留字 String(30) required:N
	Rsv30z string `json:"rsv30z,omitempty"`
}

// AccountNewNtqadinfz 接口 ntqadinfz.
type AccountNewNtqadinfz struct {
	// 上日余额=联机余额-当日账户的金融交易发生额； 当intcod='S'时，这个字段的值显示为"头寸额度（集团支付子公司余额）"是子公司的虚拟余额 M required:Y
	Accblv string `json:"accblv,omitempty"`
	// 科目 string(5) required:Y
	Accitm string `json:"accitm,omitempty"`
	// 一般为户名 Z(62) required:Y
	Accnam string `json:"accnam,omitempty"`
	// 帐号 string(35) required:Y
	Accnbr string `json:"accnbr,omitempty"`
	// 账户实际可以用来付款的金额，可用余额=联机余额-冻结金额-预期金额+透支额度 M required:N
	Avlblv string `json:"avlblv,omitempty"`
	// A.1 招商分行 string(2) required:Y
	Bbknbr string `json:"bbknbr,omitempty"`
	// A.3 货币代码表 string(2) required:Y
	Ccynbr string `json:"ccynbr,omitempty"`
	// 错误码 string(7) required:N
	Errcod string `json:"errcod,omitempty"`
	// 包含司法冻结余额、银行冻结余额等被冻结的金额加总 M required:N
	Hldblv string `json:"hldblv,omitempty"`
	// S=子公司虚拟余额 string(1) required:N
	Intcod string `json:"intcod,omitempty"`
	// 年利率 F(11,7) required:N
	Intrat string `json:"intrat,omitempty"`
	// 透支额度 M required:N
	Lmtovr string `json:"lmtovr,omitempty"`
	// 8位数字 D required:N
	Mutdat string `json:"mutdat,omitempty"`
	// 账户实际的资金 M required:N
	Onlblv string `json:"onlblv,omitempty"`
	// 8位数字 D required:Y
	Opndat string `json:"opndat,omitempty"`
	// 客户关系号 string(10) required:N
	Relnbr string `json:"relnbr,omitempty"`
	// A=活动，B=冻结，C=关户 string(1) required:Y
	Stscod string `json:"stscod,omitempty"`
	// ZZZ=不计息 TD2=定期不滚积数 TD1=定期滚积数 D99=手工计息 B01=凭证式国债 D12=挂牌对公活期保证金关户结息 D11=挂牌对公活期保证金按季结息 C01=协议计息 D32=离岸挂牌个人活期存款 D03=挂牌同业活期存款 D02=在岸挂牌个人活期存款 D31=离岸挂牌对公活期存款 D01=在岸挂牌对公活期存款 string(3) required:Y
	Inttyp string `json:"inttyp,omitempty"`
	// 定期时，取值： 一天 七天 一个月 三个月 六个月 一年 二年 三年 四年 五年 Z(12) required:Y
	Dpstxt string `json:"dpstxt,omitempty"`
	// 错误说明 Z(92) required:N
	Errtxt string `json:"errtxt,omitempty"`
}

// AccountNewTRANSQUERYBYBREAKPOINTX1 接口 TRANSQUERYBYBREAKPOINT_X1.
type AccountNewTRANSQUERYBYBREAKPOINTX1 struct {
	// 户口号 String(35) required:Y
	CardNbr string `json:"cardNbr,omitempty"`
	// 查询交易时间段的开始日期 D required:Y
	BeginDate string `json:"beginDate,omitempty"`
	// 查询交易时间段的结束日期 D required:Y
	EndDate string `json:"endDate,omitempty"`
	// 仅在 不传入 续传键值 时有效，表示初始查询从第几笔开始查询，默认从第1笔开始查询。 String(9)
	TransactionSequence string `json:"transactionSequence,omitempty"`
	// 可为空，但建议使用时传该字段 String(2)
	CurrencyCode string `json:"currencyCode,omitempty"`
	// 首次查询时，无须传该字段。本接口支持最多一次查询200条记录。若仍有记录未查询完，response中Z1实体会返回 未传完标记 为 Y ，请调用方在下次查询时将response中的Z1实体中 继续查询账号 字段赋值到入参X1的本字段，并携带response中的Y1进行查询，会继续查询接下来200条的记录。 详细见下方请求报文和续传示例。 String(200)
	QueryAcctNbr string `json:"queryAcctNbr,omitempty"`
	// 保留字段，字段长度未限制 String(200)
	Reserve string `json:"reserve,omitempty"`
	// C贷方，D借方 String(1)
	LoanCode string `json:"loanCode,omitempty"`
}

// AccountNewTRANSQUERYBYBREAKPOINTY1 接口 TRANSQUERYBYBREAKPOINT_Y1.
type AccountNewTRANSQUERYBYBREAKPOINTY1 struct {
	// 账号 String(200) required:Y
	AcctNbr string `json:"acctNbr,omitempty"`
	// 当前查询最后一笔交易日期 D required:Y
	TransDate string `json:"transDate,omitempty"`
	// 期望下一笔记账的序号 String(9) required:Y
	ExpectNextSequence string `json:"expectNextSequence,omitempty"`
}

// AccountNewTRANSQUERYBYBREAKPOINTZ1 接口 TRANSQUERYBYBREAKPOINT_Z1.
type AccountNewTRANSQUERYBYBREAKPOINTZ1 struct {
	// 贷方金额 M required:Y
	CreditAmount string `json:"creditAmount,omitempty"`
	// 贷方笔数 N required:Y
	CreditNums string `json:"creditNums,omitempty"`
	// Y|N ，Y表示还有记录需要查询，N表示当前已查询完毕 String(1) required:Y
	CtnFlag string `json:"ctnFlag,omitempty"`
	// 当 ctnFlag 为Y时，下一次请求时携带该值进行下一次查询 String(200) required:Y
	QueryAcctNbr string `json:"queryAcctNbr,omitempty"`
	// 借方金额 M required:Y
	DebitAmount string `json:"debitAmount,omitempty"`
	// 借方笔数 N required:Y
	DebitNums string `json:"debitNums,omitempty"`
	// String(200)
	Reserve string `json:"reserve,omitempty"`
}

// AccountNewTRANSQUERYBYBREAKPOINTZ2 接口 TRANSQUERYBYBREAKPOINT_Z2.
type AccountNewTRANSQUERYBYBREAKPOINTZ2 struct {
	// 余额 M required:Y
	AcctOnlineBal string `json:"acctOnlineBal,omitempty"`
	// String(20)
	BillNumber string `json:"billNumber,omitempty"`
	// Z(200)
	CtpAcctName string `json:"ctpAcctName,omitempty"`
	// String(35)
	CtpAcctNbr string `json:"ctpAcctNbr,omitempty"`
	// Z(200)
	CtpBankAddress string `json:"ctpBankAddress,omitempty"`
	// Z(400)
	CtpBankName string `json:"ctpBankName,omitempty"`
	// 币种 String(2) required:Y
	CurrencyNbr string `json:"currencyNbr,omitempty"`
	// 有效位数为16 Z(34)
	ExtendedRemark string `json:"extendedRemark,omitempty"`
	// String(35)
	FatOrSonAccount string `json:"fatOrSonAccount,omitempty"`
	// Z(200)
	FatOrSonBankAddress string `json:"fatOrSonBankAddress,omitempty"`
	// Z(200)
	FatOrSonBankName string `json:"fatOrSonBankName,omitempty"`
	// Z(200)
	FatOrSonCompanyName string `json:"fatOrSonCompanyName,omitempty"`
	// 用于标识收/付方帐号和母/子公司的信息。 为空表示付方帐号和子公司； 为“1”表示收方帐号和子公司； 为“2”表示收方帐号和母公司； 为“3”表示原收方帐号和子公司； String(1)
	InfoFlag string `json:"infoFlag,omitempty"`
	// C贷方，D借方 String(1)
	LoanCode string `json:"loanCode,omitempty"`
	// 若为企业银行客户端经办的交易，则该字段为用途信息（4.0版代发代扣业务除外）若为其它渠道经办的交易，则该字段为交易的简单说明和注解。 Z(200)
	RemarkTextClt string `json:"remarkTextClt,omitempty"`
	// String
	Reserve string `json:"reserve,omitempty"`
	// *为冲帐，X为补帐 （冲账交易与原交易借贷相反） String(1)
	ReversalFlag string `json:"reversalFlag,omitempty"`
	// 见附录A.9 String(12)
	TextCode string `json:"textCode,omitempty"`
	// 交易金额 M required:Y
	TransAmount string `json:"transAmount,omitempty"`
	// 交易日 D required:Y
	TransDate string `json:"transDate,omitempty"`
	// 流水号 String(15) required:Y
	TransSequenceIdn string `json:"transSequenceIdn,omitempty"`
	// String(6)
	TransTime string `json:"transTime,omitempty"`
	// D
	ValueDate string `json:"valueDate,omitempty"`
	// Z(60)
	BusinessName string `json:"businessName,omitempty"`
	// 对业务的简单说明或注解。企业银行客户端录入的摘要信息 Z(400)
	BusinessText string `json:"businessText,omitempty"`
	// 请注意该字段为String类型 String(10)
	RequestNbr string `json:"requestNbr,omitempty"`
	// 支付和代发的业务参考号会记录在此字段 String(30)
	YurRef string `json:"yurRef,omitempty"`
	// 子单元编号，需开通交易管家-收款识别功能。 String(16)
	VirtualNbr string `json:"virtualNbr,omitempty"`
	// 由商务支付订单产生 String(50)
	MchOrderNbr string `json:"mchOrderNbr,omitempty"`
	// 仅当查询输入X1中cardNbr和实际记账不一致的情况下会有值，如公司卡号、131农民工监管子户户口号 String(35)
	TransCardNbr string `json:"transCardNbr,omitempty"`
}

// AccountNewNonSettleTranExportTaskRegX1 接口 NonSettleTranExportTaskRegX1.
type AccountNewNonSettleTranExportTaskRegX1 struct {
	// 账号 string(35) required:Y
	CardNbr string `json:"cardNbr,omitempty"`
	// yyyyMMdd 开始日期必须大于五年前的今天 D required:Y
	BeginDate string `json:"beginDate,omitempty"`
	// yyyyMMdd，必须小于当天。 D required:Y
	EndDate string `json:"endDate,omitempty"`
	// 1 定期 2保证金 3 通知存款 4 大额存单 5 结构性存款 *全部 string(1) required:Y
	NonSettleType string `json:"nonSettleType,omitempty"`
	// 空不筛选 C贷 D借 string(1)
	Direction string `json:"direction,omitempty"`
	// M
	LowAmount string `json:"lowAmount,omitempty"`
	// M
	HighAmount string `json:"highAmount,omitempty"`
}

// AccountNewNonSettleTranExportTaskRegZ1 接口 NonSettleTranExportTaskRegZ1.
type AccountNewNonSettleTranExportTaskRegZ1 struct {
	// string(32)
	TaskNbr string `json:"taskNbr,omitempty"`
}

// AccountNewQueryNonSettleTranListByTaskY1 接口 QueryNonSettleTranListByTaskY1.
type AccountNewQueryNonSettleTranListByTaskY1 struct {
	// 账号 string(35) required:Y
	CardNbr string `json:"cardNbr,omitempty"`
	// 任务编号 string(32) required:Y
	TaskNbr string `json:"taskNbr,omitempty"`
	// String
	ResultText string `json:"resultText,omitempty"`
	// S: 成功 E: 出错 W: 等待执行 I: 处理中 String
	Status string `json:"status,omitempty"`
	// 首次传空，当返回报文中包含此Y1接口且续传键值字段非空时，需要续传查询 String
	ContinueKey string `json:"continueKey,omitempty"`
}

// AccountNewQueryNonSettleTranListByTaskZ1 接口 QueryNonSettleTranListByTaskZ1.
type AccountNewQueryNonSettleTranListByTaskZ1 struct {
	// String (35)
	TransCardNbr string `json:"transCardNbr,omitempty"`
	// 1 定期 2保证金 3 通知存款 4 大额存单 5 结构性存款 string(1)
	NonSettleType string `json:"nonSettleType,omitempty"`
	// String(15)
	TransSequenceIdn string `json:"transSequenceIdn,omitempty"`
	// String (6)
	TransTime string `json:"transTime,omitempty"`
	// D
	TransDate string `json:"transDate,omitempty"`
	// D
	ValueDate string `json:"valueDate,omitempty"`
	// M
	AcctTransSeq string `json:"acctTransSeq,omitempty"`
	// C: 贷 D: 借 String(1)
	Direction string `json:"direction,omitempty"`
	// 两位币种代码 String(2)
	CurrencyNbr string `json:"currencyNbr,omitempty"`
	// M
	TransAmount string `json:"transAmount,omitempty"`
	// M
	AcctOnlineBal string `json:"acctOnlineBal,omitempty"`
	// *为冲帐，X为补帐 N或者空， 代表非冲非补的常规交易 String（1）
	ReversalFlag string `json:"reversalFlag,omitempty"`
	// String (4)
	VoucherType string `json:"voucherType,omitempty"`
	// String(20)
	VoucherNbr string `json:"voucherNbr,omitempty"`
	// 见附录A.9 String (8)
	TextCode string `json:"textCode,omitempty"`
	// String (200)
	RemarkTextClt string `json:"remarkTextClt,omitempty"`
	// String (35)
	CtpAcctNbr string `json:"ctpAcctNbr,omitempty"`
	// String (200)
	CtpAcctName string `json:"ctpAcctName,omitempty"`
	// String (20)
	CtpAcctBankId string `json:"ctpAcctBankId,omitempty"`
	// String (400)
	CtpAcctBankName string `json:"ctpAcctBankName,omitempty"`
	// String (200)
	RemarkTextRcvAdr string `json:"remarkTextRcvAdr,omitempty"`
	// String (2)
	RemarkTextBusSys string `json:"remarkTextBusSys,omitempty"`
	// String (16)
	RemarkTextNarInn string `json:"remarkTextNarInn,omitempty"`
	// String (60)
	RemarkTextTrsRef string `json:"remarkTextTrsRef,omitempty"`
	// String (20)
	RemarkTextMchNbr string `json:"remarkTextMchNbr,omitempty"`
	// String (20)
	RemarkTextBusNbr string `json:"remarkTextBusNbr,omitempty"`
	// String(20)
	RemarkTextVchNbr string `json:"remarkTextVchNbr,omitempty"`
}

// AccountNewSdktsinfx 接口 sdktsinfx.
type AccountNewSdktsinfx struct {
	// 账号 String(35) required:Y
	Accnbr string `json:"accnbr,omitempty"`
	// C:贷；D:借 String(1)
	Amtcdr string `json:"amtcdr,omitempty"`
	// N(2)
	Bbknbr string `json:"bbknbr,omitempty"`
	// 起始日期 D required:Y
	Bgndat string `json:"bgndat,omitempty"`
	// 与结束日期的间隔不能超过100天 D required:Y
	Enddat string `json:"enddat,omitempty"`
	// 默认9999999999999.99 M
	Hghamt string `json:"hghamt,omitempty"`
	// 默认0.00 M
	Lowamt string `json:"lowamt,omitempty"`
	// 如果使用主动通知该值有效。 true或者空：生成的url链接为互联网可访问 false：生成的url链接为专线访问地址，（请专线接入客户该值传false） String
	Genble string `json:"genble,omitempty"`
}

// AccountNewNtqacctny 接口 ntqacctny.
type AccountNewNtqacctny struct {
	// 首次传空， 当 返回接口有ntqacctny并且ntqacctny的续传字段(ctndta)有值时，将返回的ntqacctny放入请求报文body中继续查询，直到返回的不返回ntqacctny或者ntqacctny的续传字段为空，续传结束。 String(60)
	Ctndta string `json:"ctndta,omitempty"`
	// 首次传空 String(2)
	Pagnbr string `json:"pagnbr,omitempty"`
}

// AccountNewDctrspdfz 接口 dctrspdfz.
type AccountNewDctrspdfz struct {
	// 返回的任务ID，请根据该ID查询PDF文件。 String
	Printid string `json:"printid,omitempty"`
	// 如果使用主动通知，可以和通知结果关联 String
	Taskid string `json:"taskid,omitempty"`
}

// AccountNewCprDirectIssueBillOfdX1 接口 CprDirectIssueBillOfdX1.
type AccountNewCprDirectIssueBillOfdX1 struct {
	// 账号 string(35) required:Y
	Accnbr string `json:"accnbr,omitempty"`
	// 币种 string(2) required:Y
	Ccynbr string `json:"ccynbr,omitempty"`
	// 如：2023 string(4) required:Y
	BillYear string `json:"billYear,omitempty"`
	// 如：02 string(2) required:Y
	BillMonth string `json:"billMonth,omitempty"`
	// 生产内网地址，专线接入时输入Y； N或者空时生成互联网地址 。当urlType为Y时，返回的输出接口为NSTF_CprDirectIssueBillOfdZ1，当urlType为N或者空时，返回的输出接口为CprDirectIssueBillOfdZ1 string(1) required:N
	UrlType string `json:"urlType,omitempty"`
}

// AccountNewCprDirectIssueBillOfdZ1 接口 CprDirectIssueBillOfdZ1.
type AccountNewCprDirectIssueBillOfdZ1 struct {
	// string(200)
	FileName string `json:"fileName,omitempty"`
	// string(100)
	IssueKey string `json:"issueKey,omitempty"`
	SpecCode string `json:"specCode,omitempty"`
}

// AccountNewNSTFCprDirectIssueBillOfdZ1 接口 NSTF_CprDirectIssueBillOfdZ1.
type AccountNewNSTFCprDirectIssueBillOfdZ1 struct {
	// string(100)
	IssueKey string `json:"issueKey,omitempty"`
	// string(200)
	FileName string `json:"fileName,omitempty"`
}

// AccountNewCprDirectQueryBillOfdX1 接口 CprDirectQueryBillOfdX1.
type AccountNewCprDirectQueryBillOfdX1 struct {
	// 开具账单键值 string(100) required:Y
	IssueKey string `json:"issueKey,omitempty"`
}

// AccountNewCprDirectQueryBillOfdZ1 接口 CprDirectQueryBillOfdZ1.
type AccountNewCprDirectQueryBillOfdZ1 struct {
	// Y-完成、N-未完成 string()
	FinishFlag string `json:"finishFlag,omitempty"`
	SpecCode   string `json:"specCode,omitempty"`
}

// AccountNewCprDirectQueryBillOfdZ2 接口 CprDirectQueryBillOfdZ2.
type AccountNewCprDirectQueryBillOfdZ2 struct {
	// String()
	FileKey string `json:"fileKey,omitempty"`
	// string()
	FileSeq string `json:"fileSeq,omitempty"`
	// 有效期24小时 string()
	FileUrl  string `json:"fileUrl,omitempty"`
	SpecCode string `json:"specCode,omitempty"`
}

// AccountNewAsycalhdz1 接口 asycalhdz1.
type AccountNewAsycalhdz1 struct {
	// String(20)
	Rtncod string `json:"rtncod,omitempty"`
	// String(100)
	Rtnmsg string `json:"rtnmsg,omitempty"`
	// 返回的任务ID，请根据该ID查询PDF文件。 String(100)
	Rtndat string `json:"rtndat,omitempty"`
}

// AccountNewCtnkeyz2 接口 ctnkeyz2.
type AccountNewCtnkeyz2 struct {
	Begamt string `json:"begamt,omitempty"`
	Begdat string `json:"begdat,omitempty"`
	Daltag string `json:"daltag,omitempty"`
	Eacnbr string `json:"eacnbr,omitempty"`
	Endamt string `json:"endamt,omitempty"`
	Enddat string `json:"enddat,omitempty"`
	Nxtdat string `json:"nxtdat,omitempty"`
	Nxtnbr string `json:"nxtnbr,omitempty"`
	Nxttim string `json:"nxttim,omitempty"`
	Oprtyp string `json:"oprtyp,omitempty"`
	Pagcnt string `json:"pagcnt,omitempty"`
	Pattyp string `json:"pattyp,omitempty"`
	Predat string `json:"predat,omitempty"`
	Prenbr string `json:"prenbr,omitempty"`
	Pretim string `json:"pretim,omitempty"`
	Rrccod string `json:"rrccod,omitempty"`
	Rrcflg string `json:"rrcflg,omitempty"`
	Spc100 string `json:"spc100,omitempty"`
}

// AccountNewDCLISMODBody 1.可经办业务模式查询DCLISMOD request body.
type AccountNewDCLISMODBody struct {
	// 附录A.4 string(6) required:Y
	Buscod string `json:"buscod,omitempty"`
}

// AccountNewDCLISMODResponseBody 1.可经办业务模式查询DCLISMOD response body.
type AccountNewDCLISMODResponseBody struct {
	Ntqmdlstz []*AccountNewNtqmdlstz `json:"ntqmdlstz,omitempty"`
}

type AccountNewDCLISMODResponse struct {
	Response struct {
		Body AccountNewDCLISMODResponseBody `json:"body,omitempty"`
		Head HeadResponse                   `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// AccountNewDCLISACCBody 2.查询可经办的账户列表DCLISACC request body.
type AccountNewDCLISACCBody struct {
	// A4.业务代码 string(6) required:Y
	Buscod string `json:"buscod,omitempty"`
	// 某个业务有哪些可经办的业务模式，可通过可经办业务模式查询(DCLISMOD) 获得。账务查询（"BUSCOD=N01010"）时忽略该项 string(5) required:Y
	Busmod string `json:"busmod,omitempty"`
}

// AccountNewDCLISACCResponseBody 2.查询可经办的账户列表DCLISACC response body.
type AccountNewDCLISACCResponseBody struct {
	Ntqaclstz []*AccountNewNtqaclstz `json:"ntqaclstz,omitempty"`
}

type AccountNewDCLISACCResponse struct {
	Response struct {
		Body AccountNewDCLISACCResponseBody `json:"body,omitempty"`
		Head HeadResponse                   `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// AccountNewNTQACINFBody 3.结算账户详细信息查询NTQACINF request body.
type AccountNewNTQACINFBody struct {
	Ntqacinfx []*AccountNewNtqacinfx `json:"ntqacinfx,omitempty"`
}

// AccountNewNTQACINFResponseBody 3.结算账户详细信息查询NTQACINF response body.
type AccountNewNTQACINFResponseBody struct {
	Ntqacinfz []*AccountNewNtqacinfz `json:"ntqacinfz,omitempty"`
}

type AccountNewNTQACINFResponse struct {
	Response struct {
		Body AccountNewNTQACINFResponseBody `json:"body,omitempty"`
		Head HeadResponse                   `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// AccountNewNTACCBBKBody 4.查询分行号信息NTACCBBK request body.
type AccountNewNTACCBBKBody struct {
	// 空表示查询该用户有权查询的所有账号 string(35)
	Fctval string `json:"fctval,omitempty"`
}

// AccountNewNTACCBBKResponseBody 4.查询分行号信息NTACCBBK response body.
type AccountNewNTACCBBKResponseBody struct {
	Ntaccbbky []*AccountNewNtaccbbky `json:"ntaccbbky,omitempty"`
	Ntaccbbkz []*AccountNewNtaccbbkz `json:"ntaccbbkz,omitempty"`
}

type AccountNewNTACCBBKResponse struct {
	Response struct {
		Body AccountNewNTACCBBKResponseBody `json:"body,omitempty"`
		Head HeadResponse                   `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// AccountNewQryNonSettleAccountsBody 5.非结算账户详细信息查询QryNonSettleAccounts request body.
type AccountNewQryNonSettleAccountsBody struct {
	QueryNoSettlementAccountsX1 []*AccountNewQueryNoSettlementAccountsX1 `json:"QueryNoSettlementAccountsX1,omitempty"`
	QueryNoSettlementAccountsY1 []*AccountNewQueryNoSettlementAccountsY1 `json:"QueryNoSettlementAccountsY1,omitempty"`
}

// AccountNewQryNonSettleAccountsResponseBody 5.非结算账户详细信息查询QryNonSettleAccounts response body.
type AccountNewQryNonSettleAccountsResponseBody struct {
	QueryNoSettlementAccountsY1 []*AccountNewQueryNoSettlementAccountsY1 `json:"QueryNoSettlementAccountsY1,omitempty"`
	QueryNoSettlementAccountsZ1 []*AccountNewQueryNoSettlementAccountsZ1 `json:"QueryNoSettlementAccountsZ1,omitempty"`
}

type AccountNewQryNonSettleAccountsResponse struct {
	Response struct {
		Body AccountNewQryNonSettleAccountsResponseBody `json:"body,omitempty"`
		Head HeadResponse                               `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// AccountNewTotalAccountDirectBody 6.综合账户信息查询TotalAccountDirect request body.
type AccountNewTotalAccountDirectBody struct {
	TotalAccountDirectReq []*AccountNewTotalAccountDirectReq `json:"TotalAccountDirectReq,omitempty"`
}

// AccountNewTotalAccountDirectResponseBody 6.综合账户信息查询TotalAccountDirect response body.
type AccountNewTotalAccountDirectResponseBody struct {
	TotalAccountDirectRsp []*AccountNewTotalAccountDirectRsp `json:"TotalAccountDirectRsp,omitempty"`
	TotalAccountDirectSum []*AccountNewTotalAccountDirectSum `json:"TotalAccountDirectSum,omitempty"`
}

type AccountNewTotalAccountDirectResponse struct {
	Response struct {
		Body AccountNewTotalAccountDirectResponseBody `json:"body,omitempty"`
		Head HeadResponse                             `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// AccountNewNTQABINFBody 7.查询账户历史余额NTQABINF request body.
type AccountNewNTQABINFBody struct {
	Ntqabinfy []*AccountNewNtqabinfy `json:"ntqabinfy,omitempty"`
}

// AccountNewNTQABINFResponseBody 7.查询账户历史余额NTQABINF response body.
type AccountNewNTQABINFResponseBody struct {
	Ntqabinfz []*AccountNewNtqabinfz `json:"ntqabinfz,omitempty"`
}

type AccountNewNTQABINFResponse struct {
	Response struct {
		Body AccountNewNTQABINFResponseBody `json:"body,omitempty"`
		Head HeadResponse                   `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// AccountNewNTQRYHSBBody 8.历史冻结余额等余额信息查询NTQRYHSB request body.
type AccountNewNTQRYHSBBody struct {
	Ntqryhsbx1 []*AccountNewNtqryhsbx1 `json:"ntqryhsbx1,omitempty"`
}

// AccountNewNTQRYHSBResponseBody 8.历史冻结余额等余额信息查询NTQRYHSB response body.
type AccountNewNTQRYHSBResponseBody struct {
	Ntqryhsbz1 []*AccountNewNtqryhsbz1 `json:"ntqryhsbz1,omitempty"`
}

type AccountNewNTQRYHSBResponse struct {
	Response struct {
		Body AccountNewNTQRYHSBResponseBody `json:"body,omitempty"`
		Head HeadResponse                   `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// AccountNewNTQADINFBody 9. 批量查询余额NTQADINF request body.
type AccountNewNTQADINFBody struct {
	Ntqadinfx []*AccountNewNtqadinfx `json:"ntqadinfx,omitempty"`
}

// AccountNewNTQADINFResponseBody 9. 批量查询余额NTQADINF response body.
type AccountNewNTQADINFResponseBody struct {
	Ntqadinfz []*AccountNewNtqadinfz `json:"ntqadinfz,omitempty"`
}

type AccountNewNTQADINFResponse struct {
	Response struct {
		Body AccountNewNTQADINFResponseBody `json:"body,omitempty"`
		Head HeadResponse                   `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// AccountNewTrsQryByBreakPointBody 10.账户交易信息查询trsQryByBreakPoint request body.
type AccountNewTrsQryByBreakPointBody struct {
	TRANSQUERYBYBREAKPOINTX1 []*AccountNewTRANSQUERYBYBREAKPOINTX1 `json:"TRANSQUERYBYBREAKPOINT_X1,omitempty"`
	TRANSQUERYBYBREAKPOINTY1 []*AccountNewTRANSQUERYBYBREAKPOINTY1 `json:"TRANSQUERYBYBREAKPOINT_Y1,omitempty"`
}

// AccountNewTrsQryByBreakPointResponseBody 10.账户交易信息查询trsQryByBreakPoint response body.
type AccountNewTrsQryByBreakPointResponseBody struct {
	TRANSQUERYBYBREAKPOINTY1 []*AccountNewTRANSQUERYBYBREAKPOINTY1 `json:"TRANSQUERYBYBREAKPOINT_Y1,omitempty"`
	TRANSQUERYBYBREAKPOINTZ1 []*AccountNewTRANSQUERYBYBREAKPOINTZ1 `json:"TRANSQUERYBYBREAKPOINT_Z1,omitempty"`
	TRANSQUERYBYBREAKPOINTZ2 []*AccountNewTRANSQUERYBYBREAKPOINTZ2 `json:"TRANSQUERYBYBREAKPOINT_Z2,omitempty"`
}

type AccountNewTrsQryByBreakPointResponse struct {
	Response struct {
		Body AccountNewTrsQryByBreakPointResponseBody `json:"body,omitempty"`
		Head HeadResponse                             `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// AccountNewNonSettleTranOprBody 11非结算户交易流水导出申请NonSettleTranOpr request body.
type AccountNewNonSettleTranOprBody struct {
	NonSettleTranExportTaskRegX1 []*AccountNewNonSettleTranExportTaskRegX1 `json:"NonSettleTranExportTaskRegX1,omitempty"`
}

// AccountNewNonSettleTranOprResponseBody 11非结算户交易流水导出申请NonSettleTranOpr response body.
type AccountNewNonSettleTranOprResponseBody struct {
	NonSettleTranExportTaskRegZ1 []*AccountNewNonSettleTranExportTaskRegZ1 `json:"NonSettleTranExportTaskRegZ1,omitempty"`
}

type AccountNewNonSettleTranOprResponse struct {
	Response struct {
		Body AccountNewNonSettleTranOprResponseBody `json:"body,omitempty"`
		Head HeadResponse                           `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// AccountNewNonSettleTranQryBody 12查询关联非结算户交易流水NonSettleTranQry request body.
type AccountNewNonSettleTranQryBody struct {
	QueryNonSettleTranListByTaskY1 []*AccountNewQueryNonSettleTranListByTaskY1 `json:"QueryNonSettleTranListByTaskY1,omitempty"`
}

// AccountNewNonSettleTranQryResponseBody 12查询关联非结算户交易流水NonSettleTranQry response body.
type AccountNewNonSettleTranQryResponseBody struct {
	QueryNonSettleTranListByTaskY1 []*AccountNewQueryNonSettleTranListByTaskY1 `json:"QueryNonSettleTranListByTaskY1,omitempty"`
	QueryNonSettleTranListByTaskZ1 []*AccountNewQueryNonSettleTranListByTaskZ1 `json:"QueryNonSettleTranListByTaskZ1,omitempty"`
}

type AccountNewNonSettleTranQryResponse struct {
	Response struct {
		Body AccountNewNonSettleTranQryResponseBody `json:"body,omitempty"`
		Head HeadResponse                           `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// AccountNewDCTRSPDFBody 13.PDF文件对账单获取DCTRSPDF request body.
type AccountNewDCTRSPDFBody struct {
	Sdktsinfx []*AccountNewSdktsinfx `json:"sdktsinfx,omitempty"`
	Ntqacctny []*AccountNewNtqacctny `json:"ntqacctny,omitempty"`
}

// AccountNewDCTRSPDFResponseBody 13.PDF文件对账单获取DCTRSPDF response body.
type AccountNewDCTRSPDFResponseBody struct {
	Dctrspdfz []*AccountNewDctrspdfz `json:"dctrspdfz,omitempty"`
	Ntqacctny []*AccountNewNtqacctny `json:"ntqacctny,omitempty"`
}

type AccountNewDCTRSPDFResponse struct {
	Response struct {
		Body AccountNewDCTRSPDFResponseBody `json:"body,omitempty"`
		Head HeadResponse                   `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// AccountNewIssueBillOfdBody 14.OFD文件对账单获取issueBillOfd request body.
type AccountNewIssueBillOfdBody struct {
	CprDirectIssueBillOfdX1 []*AccountNewCprDirectIssueBillOfdX1 `json:"CprDirectIssueBillOfdX1,omitempty"`
}

// AccountNewIssueBillOfdResponseBody 14.OFD文件对账单获取issueBillOfd response body.
type AccountNewIssueBillOfdResponseBody struct {
	CprDirectIssueBillOfdZ1     []*AccountNewCprDirectIssueBillOfdZ1     `json:"CprDirectIssueBillOfdZ1,omitempty"`
	NSTFCprDirectIssueBillOfdZ1 []*AccountNewNSTFCprDirectIssueBillOfdZ1 `json:"NSTF_CprDirectIssueBillOfdZ1,omitempty"`
}

type AccountNewIssueBillOfdResponse struct {
	Response struct {
		Body AccountNewIssueBillOfdResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// AccountNewQueryBillOfdBody 15.OFD文件对账单获取结果查询queryBillOfd request body.
type AccountNewQueryBillOfdBody struct {
	CprDirectQueryBillOfdX1 []*AccountNewCprDirectQueryBillOfdX1 `json:"CprDirectQueryBillOfdX1,omitempty"`
}

// AccountNewQueryBillOfdResponseBody 15.OFD文件对账单获取结果查询queryBillOfd response body.
type AccountNewQueryBillOfdResponseBody struct {
	CprDirectQueryBillOfdZ1 []*AccountNewCprDirectQueryBillOfdZ1 `json:"CprDirectQueryBillOfdZ1,omitempty"`
	CprDirectQueryBillOfdZ2 []*AccountNewCprDirectQueryBillOfdZ2 `json:"CprDirectQueryBillOfdZ2,omitempty"`
}

type AccountNewQueryBillOfdResponse struct {
	Response struct {
		Body AccountNewQueryBillOfdResponseBody `json:"body,omitempty"`
		Head HeadResponse                       `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// AccountNewDCSIGRECBody 16.单笔回单查询DCSIGREC request body.
type AccountNewDCSIGRECBody struct {
	// 账号 String(35) required:Y
	Eacnbr string `json:"eacnbr,omitempty"`
	// yyyy-MM-dd格式 String(10) required:Y
	Quedat string `json:"quedat,omitempty"`
	// 交易流水号 String(15) required:Y
	Trsseq string `json:"trsseq,omitempty"`
	// 空时默认PDF. PDF 或者 OFDEX Srting(3)
	Primod string `json:"primod,omitempty"`
}

// AccountNewDCSIGRECResponseBody 16.单笔回单查询DCSIGREC response body.
type AccountNewDCSIGRECResponseBody struct {
	// String
	Checod string `json:"checod,omitempty"`
	// 返回Base64后的数据，Base64解码后写入文件可得到PDF或OFD文件 String
	Fildat string `json:"fildat,omitempty"`
	// String
	Istnbr string `json:"istnbr,omitempty"`
}

type AccountNewDCSIGRECResponse struct {
	Response struct {
		Body AccountNewDCSIGRECResponseBody `json:"body,omitempty"`
		Head HeadResponse                   `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// AccountNewASYCALHDBody 17. 电子回单异步查询ASYCALHD request body.
type AccountNewASYCALHDBody struct {
	// 账号 String(35) required:Y
	Eacnbr string `json:"eacnbr,omitempty"`
	// yyyy-MM-dd格式 String(10) required:Y
	Begdat string `json:"begdat,omitempty"`
	// yyyy-MM-dd格式 String(10) required:Y
	Enddat string `json:"enddat,omitempty"`
	// 1:未打印 2:已打印 传值则根据条件筛选 String(1) required:N
	Rrcflg string `json:"rrcflg,omitempty"`
	// 13位整数 + 1位小数点 + 2位小数点后面数字 M
	Begamt string `json:"begamt,omitempty"`
	// 13位整数 + 1位小数点 + 2位小数点后面数字 M
	Endamt string `json:"endamt,omitempty"`
	// String(8)
	Rrccod string `json:"rrccod,omitempty"`
	// PDF 或者 OFD String(3) required:Y
	Primod string `json:"primod,omitempty"`
}

// AccountNewASYCALHDResponseBody 17. 电子回单异步查询ASYCALHD response body.
type AccountNewASYCALHDResponseBody struct {
	Asycalhdz1 *AccountNewAsycalhdz1 `json:"asycalhdz1,omitempty"`
	Ctnkeyz2   *AccountNewCtnkeyz2   `json:"ctnkeyz2,omitempty"`
}

type AccountNewASYCALHDResponse struct {
	Response struct {
		Body AccountNewASYCALHDResponseBody `json:"body,omitempty"`
		Head HeadResponse                   `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// AccountNewDCTASKIDBody 18. 异步打印结果查询DCTASKID request body.
type AccountNewDCTASKIDBody struct {
	// ASYCALHD请求返回报文中的rtndat字段 String(40) required:Y
	Taskid string `json:"taskid,omitempty"`
	// true或者空：生成的url链接为互联网可访问 false：生成的url链接为专线访问地址，（请专线接入客户该值传false） String
	Qwenab string `json:"qwenab,omitempty"`
}

// AccountNewDCTASKIDResponseBody 18. 异步打印结果查询DCTASKID response body.
type AccountNewDCTASKIDResponseBody struct {
	// 生成完成后，获取地址链接，请注意该链接有效期1天，获取到后请立即下载保存，以防过期失效。 String(300)
	Fileurl string `json:"fileurl,omitempty"`
	// 文件生成完成时间 String(20)
	Fintim string `json:"fintim,omitempty"`
}

type AccountNewDCTASKIDResponse struct {
	Response struct {
		Body AccountNewDCTASKIDResponseBody `json:"body,omitempty"`
		Head HeadResponse                   `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

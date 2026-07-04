package models

import "encoding/json"

// Code generated from CMB official 代发代扣 documentation. DO NOT EDIT BY HAND.
// Source: https://openbiz.cmbchina.com/doc-detail?bizkey=DCCT20220309170929467&subclass=1&treeID=2026062909291024364
// Version: 1.0.0, updated: 2026-07-02 08:29:55

type PayrollNewRequest struct {
	Request   PayrollNewRequestData `json:"request"`
	Signature Signature             `json:"signature"`
}

type PayrollNewRequestData struct {
	Body interface{} `json:"body,omitempty"`
	Head Head        `json:"head"`
}

type PayrollNewResponse struct {
	Response PayrollNewResponseData `json:"response,omitempty"`
}

type PayrollNewResponseData struct {
	Body PayrollNewRawBody `json:"body,omitempty"`
	Head HeadResponse      `json:"head,omitempty"`
}

type PayrollNewRawBody map[string]json.RawMessage

type PayrollNewBody map[string]interface{}

// PayrollNewBb6aclakx1 输入接口 bb6aclakx1.
type PayrollNewBb6aclakx1 struct {
	// 批次开始标志 String (1) required:Y 非空 首次传输赋值为’Y’，非首次固定赋值为’N’
	Begtag string `json:"begtag,omitempty"`
	// 批次结束标志 String (1) required:Y 非空 最后一次传输赋值为’Y’，其他批次固定赋值为’N’
	Endtag string `json:"endtag,omitempty"`
	// 流程实例号 String (10) 第一次上传时必须为空；续传时不能为空， 所有续传次数流程实例号必须为同一个；续传时的值来源于第一次的系统返回值，该值为招行系统生成
	Reqnbr string `json:"reqnbr,omitempty"`
	// 总金额 M required:Y 非空。大于0，13位整数，2位小数
	Ttlamt string `json:"ttlamt,omitempty"`
	// 总笔数 N(8) required:Y 非空，大于0，8位整数。
	Ttlcnt string `json:"ttlcnt,omitempty"`
	// 总次数 N(3) required:Y 非空，大于0，3位整数。该批次数据计划分多少次上传完
	Ttlnum string `json:"ttlnum,omitempty"`
	// 本次金额 M required:Y 非空，大于0，13位整数，2位小数
	Curamt string `json:"curamt,omitempty"`
	// 本次笔数 N(8) required:Y 非空，大于0，8位整数。
	Curcnt string `json:"curcnt,omitempty"`
	// 合作方协议号 String (6) 合作方协议号和代扣类型不能同时为空
	Cnvnbr string `json:"cnvnbr,omitempty"`
	// 交易货币 N(2) required:Y 非空。两位数字，代扣只支持人民币“10”
	Ccynbr string `json:"ccynbr,omitempty"`
	// 个性化短信 Z(22) 支持中文 ，不超过10个汉字。在收方入账短信里展示
	Ntfinf string `json:"ntfinf,omitempty"`
	// 分行号 String (2) 可空
	Bbknbr string `json:"bbknbr,omitempty"`
	// 账号 String (35) required:Y 非空
	Accnbr string `json:"accnbr,omitempty"`
	// 货币市场 String (1) 预留，不输入 0 不分钞汇 1 现钞 2 现汇
	Ccymkt string `json:"ccymkt,omitempty"`
	// 代扣类型 String (4) 合作方协议号和代扣类型不能同时为空，可通过“代扣类型查询- BB6TRSQY”接口查询 AKTI=代扣出租车规费 AKTL=代扣电话费 AKTV=代扣有线电视费 AKTW=代扣水电费 AKTX=代扣税费 AKWD=代扣业务 AKWT=代扣水费等
	Trstyp string `json:"trstyp,omitempty"`
	// 用途 Z(42) required:Y 非空，支持中文，展示在入账回单的交易摘要 不超过20个汉字
	Nusage string `json:"nusage,omitempty"`
	// 期望日 D required:Y 非空，不得早于经办日期 格式：YYYYMMDD
	Eptdat string `json:"eptdat,omitempty"`
	// 期望时间 String (6) HHMMSS,为空的话，默认值为000000
	Epttim string `json:"epttim,omitempty"`
	// 业务参考号 String (30) required:Y 非空，允许数字或者字母 不允许中文 业务参考号（YURREF）是每笔交易的唯一编号，是防止重复提交的重要手段，如需重复发送请求，请务必保证同一笔交易的业务参考号不变，否则会存在重复提交风险。
	Yurref string `json:"yurref,omitempty"`
	// 记账子单元 String (20) 预留，代扣暂不支持交易管家（记账）虚拟户
	Dmanbr string `json:"dmanbr,omitempty"`
	// 网银审批标志 String (1) 默认为N，直联无审批模式；Y：直联经办，网银审批模式。
	Grtflg string `json:"grtflg,omitempty"`
	// 保留字 String (100)
	Rsv100 string `json:"rsv100,omitempty"`
}

// PayrollNewBb6aclaky1 输入接口 bb6aclaky1.
type PayrollNewBb6aclaky1 struct {
	// 交易序号 String (8) required:Y 非空 需要客户自行保证批次范围内的序号唯一性，代发代扣系统要求格式为全数字，如 ‘00000001’,’00000002’,不允许“00000000” 每次最多接收1000笔明细
	Trxseq string `json:"trxseq,omitempty"`
	// 账号 String (35) required:Y 非空
	Accnbr string `json:"accnbr,omitempty"`
	// 户名 Z(62) required:Y 非空（不超过30个汉字）
	Accnam string `json:"accnam,omitempty"`
	// 金额 M required:Y 非空, 大于0，13位整数，2位小数
	Trsamt string `json:"trsamt,omitempty"`
	// 注释 Z(42) 支持中文。展示在出款回单的交易摘要（当此字段值为空时，默认为“网上代发代扣”） 不超过20个汉字
	Trsdsp string `json:"trsdsp,omitempty"`
	// 系统内标志 String (1) required:Y 非空，代扣只支持“Y” Y: 开户行是招商银行; N：开户行是他行。
	Bnkflg string `json:"bnkflg,omitempty"`
	// 他行快速标志 String (1) 新增，空值
	Fstflg string `json:"fstflg,omitempty"`
	// 他行户口联行号 String (12) 新增，空值
	Rcvbnk string `json:"rcvbnk,omitempty"`
	// 他行户口开户行 Z(62) 预留，当前为空 不超过30个汉字
	Eacbnk string `json:"eacbnk,omitempty"`
	// 他行户口开户地 Z(62) 预留，当前为空 不超过30个汉字
	Eaccty string `json:"eaccty,omitempty"`
	// 客户代码 String (20) 以前代扣将合同方账号填到注释字段里，现在可以改为填到这个字段；代发可空
	Cpract string `json:"cpract,omitempty"`
	// 合作方流水号 String (20)
	Cprref string `json:"cprref,omitempty"`
	// 保留字段 String (30)
	Rsv30z string `json:"rsv30z,omitempty"`
}

// PayrollNewBb6aclakz1 输出接口 bb6aclakz1.
type PayrollNewBb6aclakz1 struct {
	// 流程实例号 String (10) 该值为招行系统生成
	Reqnbr string `json:"reqnbr,omitempty"`
	// 请求状态 String (3) OPR 数据接收中 AUT 等待审批 NTE 终审完毕 APW 银行人工审批 WRF 可疑 BNK 银行处理中 FIN 完成
	Reqsta string `json:"reqsta,omitempty"`
	// 保留字 String (50)
	Rsv50z string `json:"rsv50z,omitempty"`
}

// PayrollNewBb6aclify1 输入/输出接口 bb6aclify1.
type PayrollNewBb6aclify1 struct {
	// 流程实例号 String (10) required:Y 该值为招行系统生成 请注意该字段为String类型
	Reqnbr string `json:"reqnbr,omitempty"`
	// 续传批次号 Strin g (30) 续传时输入
	Bthnbr string `json:"bthnbr,omitempty"`
	// 续传明细序号 String (8) 续传时输入
	Trxseq string `json:"trxseq,omitempty"`
	// 续传历史查询标志 String (1) 续传时输入
	Histag string `json:"histag,omitempty"`
	// 保留字 String (50)
	Rsv50z string `json:"rsv50z,omitempty"`
}

// PayrollNewBb6aclifz1 输出接口 bb6aclifz1.
type PayrollNewBb6aclifz1 struct {
	// 交易序号 String (8)
	Trxseq string `json:"trxseq,omitempty"`
	// 批次号 String (30)
	Bthnbr string `json:"bthnbr,omitempty"`
	// 账号 String (35)
	Accnbr string `json:"accnbr,omitempty"`
	// 户名 Z(62) 不超过30个汉字
	Accnam string `json:"accnam,omitempty"`
	// 金额 M 大于0，13位整数，2位小数
	Trsamt string `json:"trsamt,omitempty"`
	// 实际代扣金额 M 13位整数，2位小数
	Lgramt string `json:"lgramt,omitempty"`
	// 注释 Z(42) 批次信息表中的摘要 不超过99个汉字
	Trsdsp string `json:"trsdsp,omitempty"`
	// 交易日期 D 新增,YYYYMMDD,经办日期
	Trsdat string `json:"trsdat,omitempty"`
	// 记录状态 String (1) S：成功； E：失败； A：登记
	Stscod string `json:"stscod,omitempty"`
	// 错误码 String (7)
	Errcod string `json:"errcod,omitempty"`
	// 错误信息 Z(192) 不超过95个汉字
	Errtxt string `json:"errtxt,omitempty"`
	// 系统内标志 String (1)
	Bnkflg string `json:"bnkflg,omitempty"`
	// 他行户口开户行 Z(62) 不超过30个汉字
	Eacbnk string `json:"eacbnk,omitempty"`
	// 他行户口开户地 Z(62) 不超过30个汉字
	Eaccty string `json:"eaccty,omitempty"`
	// 他行快速标志 String (1)
	Fstflg string `json:"fstflg,omitempty"`
	// 他行户口联行号 String (12)
	Rcvbnk string `json:"rcvbnk,omitempty"`
	// 客户代码 String (20)
	Cpract string `json:"cpract,omitempty"`
	// 合作方流水号 String (20)
	Cprref string `json:"cprref,omitempty"`
	// 交易流水 String (30)
	Trsreq string `json:"trsreq,omitempty"`
	// 交易套号 String (21)
	Trsset string `json:"trsset,omitempty"`
	// 对接状态 String(1) 预留，代扣目前不支持他行，目前返回空值
	Othsts string `json:"othsts,omitempty"`
	// 保留字段 String (30)
	Rsv30z string `json:"rsv30z,omitempty"`
}

// PayrollNewBb6bpdqyy1 输入接口 bb6bpdqyy1.
type PayrollNewBb6bpdqyy1 struct {
	// 业务编码 String (6) required:Y N03010 代发工资 N03020 代发 N03030 代扣 N03040 钱包代发
	Buscod string `json:"buscod,omitempty"`
	// 业务参考号 String (30) required:Y
	Yurref string `json:"yurref,omitempty"`
	// 续传批次号 String(10) 首次传空，当返回报文中有bb6bpdqyy1则需要续传，将返回的bb6bpdqyy1替换请求报文的bb6bpdqyy1接口，再次查询，直到返回报文中没有bb6bpdqyy1接口
	Bthnbr string `json:"bthnbr,omitempty"`
	// 续传明细序号 String(8)
	Trxseq string `json:"trxseq,omitempty"`
	// 续传历史查询标志 String (1)
	Histag string `json:"histag,omitempty"`
	// 保留字 String (50)
	Rsv50z string `json:"rsv50z,omitempty"`
}

// PayrollNewBb6bpdqyz1 输出接口 bb6bpdqyz1.
type PayrollNewBb6bpdqyz1 struct {
	// 流程实例号 String(10) 该值为招行系统生成 格式含数字和字母 请注意该字段为String类型
	Reqnbr string `json:"reqnbr,omitempty"`
	// 业务类型 String (6) N03010 代发工资 N03020 代发其他 N03030 代扣 N03040 钱包代发
	Buscod string `json:"buscod,omitempty"`
	// 业务模式 String(5) 5位数字，例00001
	Busmod string `json:"busmod,omitempty"`
	// 经办日 D 格式YYYYMMDD
	Oprdat string `json:"oprdat,omitempty"`
	// 期望日 D 格式YYYYMMDD
	Eptdat string `json:"eptdat,omitempty"`
	// 期望时间 String (6) 格式HHMMSS
	Epttim string `json:"epttim,omitempty"`
	// 业务请求状态 String (3) OPR 数据接收中 AUT 等待审批 NTE 终审完毕 APW 银行人工审批 WRF 可疑 BNK 银行处理中 FIN 完成
	Reqsta string `json:"reqsta,omitempty"`
	// 业务请求结果 String (1) S 成功 F 失败 C 撤销 D 过期 企业过期不审批 R 否决 企业审批否决
	Rtnflg string `json:"rtnflg,omitempty"`
	// 待处理操作步骤 String (6)
	Oprstp string `json:"oprstp,omitempty"`
	// 待处理操作序列 String (3)
	Oprsqn string `json:"oprsqn,omitempty"`
	// 业务受理分行号 String (2)
	Bbknbr string `json:"bbknbr,omitempty"`
	// 账号 String (35) 出款方为账户时返回
	Accnbr string `json:"accnbr,omitempty"`
	// 户名 String (400) 出款方为账户时返回。 不超过100个汉字。账号户名
	Accnam string `json:"accnam,omitempty"`
	// 钱包ID String (16) 出款方为钱包时返回
	Waltid string `json:"waltid,omitempty"`
	// 公司名称 String (600) 出款方为钱包时返回
	Fctnam string `json:"fctnam,omitempty"`
	// 交易笔数 N(6)
	Trsnum string `json:"trsnum,omitempty"`
	// 总金额 M 13位整数，2位小数
	Totamt string `json:"totamt,omitempty"`
	// 成功笔数 N(6)
	Sucnum string `json:"sucnum,omitempty"`
	// 成功金额 M 13位整数，2位小数
	Sucamt string `json:"sucamt,omitempty"`
	// 币种 String (2) 两位数字，默认币种赋值为10， 例如10 人民币
	Ccynbr string `json:"ccynbr,omitempty"`
	// 货币市场 String (1) 0 不分钞汇 1 现钞 2 现汇
	Ccymkt string `json:"ccymkt,omitempty"`
	// 代发类型 String (4) BYBC 代发劳务收入 BYBD 代发投资本益 BYBE 代发债产权转让 等
	Trstyp string `json:"trstyp,omitempty"`
	// 用途 String (42) 最多不超过20个汉字
	Nusage string `json:"nusage,omitempty"`
	// 业务参考号 String (30)
	Yurref string `json:"yurref,omitempty"`
	// 记录状态 String(1) T : 未处理 A: 银行处理中 E:完成 C: 撤销、否决 、过期
	Stscod string `json:"stscod,omitempty"`
	// 错误码 String (7)
	Errcod string `json:"errcod,omitempty"`
	// 错误描述 String (192) 最多不超过95个汉字
	Errdsp string `json:"errdsp,omitempty"`
	// 处理批次号 String (30)
	Seqcod string `json:"seqcod,omitempty"`
	// 附件标志 String (1) Y/N，默认为N，当前直联不支持附件
	Athflg string `json:"athflg,omitempty"`
	// 记账子单元 String (20) 预留，代扣暂不支持交易管家（记账）虚拟户
	Dmanbr string `json:"dmanbr,omitempty"`
	// 渠道标志 String (1) Y 超网 N 大小额 M 智能路由（超网和大小额混发） A 数币代发(账户代发到钱包) W 钱包代发(钱包代发到钱包 )
	Agctyp string `json:"agctyp,omitempty"`
	// 交易流水 String (30)
	Trsreq string `json:"trsreq,omitempty"`
	// 交易套号 String (21)
	Trsset string `json:"trsset,omitempty"`
	// 记账方式 String (1) 不填默认为汇总记账 Y 单笔记账 ， N或不填 汇总记账 单笔记账（依据代发明细生成多笔流水及回单，将展示具体收款人信息，请评估薪资保密需求） 汇总记账（仅生成一笔汇总交易流水及回单） 使用单笔记账前需要联系客户经理申请
	Sigacc string `json:"sigacc,omitempty"`
	// 保留字 String (100) 第一位为P 表示部分成功
	Rsv100 string `json:"rsv100,omitempty"`
}

// PayrollNewBb6bpdqyz2 输出接口 bb6bpdqyz2.
type PayrollNewBb6bpdqyz2 struct {
	// 批次号码 String (30)
	Bthnbr string `json:"bthnbr,omitempty"`
	// 交易序号 String (8)
	Trxseq string `json:"trxseq,omitempty"`
	// 账号 String (35) 收方为账户时有值
	Accnbr string `json:"accnbr,omitempty"`
	// 钱包ID String (16) 收方为钱包时有值
	Waltid string `json:"waltid,omitempty"`
	// 钱包昵称 String (300) 经办时没传可能没值。行内钱包昵称可以自动补充
	Walnam string `json:"walnam,omitempty"`
	// 户名/ 钱包真实姓名 String (400) 不超过200个汉字
	Accnam string `json:"accnam,omitempty"`
	// 数币标记 String(1) Y 钱包 N 账户
	Digflg string `json:"digflg,omitempty"`
	// 交易金额 M 13位整数，2位小数
	Trsamt string `json:"trsamt,omitempty"`
	// 实际代扣金额 M 13位整数，2位小数
	Lgramt string `json:"lgramt,omitempty"`
	// 交易日期 D 格式YYYYMMDD
	Trsdat string `json:"trsdat,omitempty"`
	// 注释 String (42) 不超过20个汉字
	Trsdsp string `json:"trsdsp,omitempty"`
	// 交易状态 String (1) E：失败 A: 登记 S: 成功
	Stscod string `json:"stscod,omitempty"`
	// 失败代码 String (7)
	Errcod string `json:"errcod,omitempty"`
	// 失败原因 String (192) 不超过95个汉字
	Errtxt string `json:"errtxt,omitempty"`
	// 系统内标志 String (1) Y：开户行为招行； N：开户行为他行 当前为空，后续将升级
	Bnkflg string `json:"bnkflg,omitempty"`
	// 他行户口开户行 String (62) 非超网有值 不超过30个汉字
	Eacbnk string `json:"eacbnk,omitempty"`
	// 他行户口开户地 String (62) 非超网有值 不超过30个汉字
	Eaccty string `json:"eaccty,omitempty"`
	// 他行快速标志 String(1) 非超网从代发系统取值，但返回值为空
	Fstflg string `json:"fstflg,omitempty"`
	// 他行户口联行号 String (12) 非超网有值；
	Rcvbnk string `json:"rcvbnk,omitempty"`
	// 他行钱包运营机构号 String(14) 收方为行外钱包时可能有值
	Bnkorg string `json:"bnkorg,omitempty"`
	// 他行钱包运营机构名称 String(400) 收方为行外钱包时可能有值
	Orgnam string `json:"orgnam,omitempty"`
	// 客户代码 String (20) 赋值为空
	Cpract string `json:"cpract,omitempty"`
	// 合作方流水号 String (20) 暂无用，预留
	Cprref string `json:"cprref,omitempty"`
	// 交易流水号 String (30)
	Trssqn string `json:"trssqn,omitempty"`
	// 交易套号 String (21)
	Trsset string `json:"trsset,omitempty"`
	// 对接代理行状态 String(1) 当前为空，后续将支持
	Othsts string `json:"othsts,omitempty"`
	// 付方交易套号 String (30) 单笔记账对账使用
	Payseq string `json:"payseq,omitempty"`
	// 付方交易流水 String (21) 单笔记账对账使用
	Paysqn string `json:"paysqn,omitempty"`
	// 退票附言 String(200) 对手行返回的退票附言
	Rtntxt string `json:"rtntxt,omitempty"`
	// 保留字 String (50)
	Rsv30z string `json:"rsv30z,omitempty"`
}

// PayrollNewBb6bthqyy1 输入/输出接口 bb6bthqyy1.
type PayrollNewBb6bthqyy1 struct {
	// 业务编码 String (6) required:Y 非空 N03010 代发工资 N03020 代发 N03040 钱包代发
	Buscod string `json:"buscod,omitempty"`
	// 业务参考号 String (30) 如果输入业务参考号，则查询指定业务参考号批次，则日期字段不需要输入,且会忽略日期的限制
	Yurref string `json:"yurref,omitempty"`
	// 起始日期 D 业务参考号为空，起始日期非空 最大180天区间，只支持近13个月的数据
	Bgndat string `json:"bgndat,omitempty"`
	// 结束日期 D 业务参考号为空，结束日期非空 最大180天区间，只支持近13个月的数据
	Enddat string `json:"enddat,omitempty"`
	// 日期类型 String (1) A：经办日期； B：期望日期； 默认为A
	Dattyp string `json:"dattyp,omitempty"`
	// 续传键值 String (80) 当返回报文有bb6bthqyy1接口返回且cntkey字段不为空时标明需要续传，请将返回报文的bb6bthqyy1内容放入下次请求继续查询
	Cntkey string `json:"cntkey,omitempty"`
	// 保留字 String (50)
	Rsv50z string `json:"rsv50z,omitempty"`
}

// PayrollNewBb6bthqyz1 输出接口 bb6bthqyz1.
type PayrollNewBb6bthqyz1 struct {
	// 流程实例号 String (10) 该值为招行系统生成 格式含数字和字母 请注意该字段为String类型
	Reqnbr string `json:"reqnbr,omitempty"`
	// 业务类型 String (6) N03010 代发工资 N03020 代发其他 N03040 钱包代发
	Buscod string `json:"buscod,omitempty"`
	// 业务模式 String(5) 5位数字，例00001
	Busmod string `json:"busmod,omitempty"`
	// 经办日 D 格式YYYYMMDD
	Oprdat string `json:"oprdat,omitempty"`
	// 期望日 D 格式YYYYMMDD
	Eptdat string `json:"eptdat,omitempty"`
	// 期望时间 String (6) 格式HHMMSS
	Epttim string `json:"epttim,omitempty"`
	// 请求状态 String (3) OPR 数据接收中 AUT 等待审批 NTE 终审完毕 APW 银行人工审批 WRF 可疑 BNK 银行处理中 FIN 完成
	Reqsta string `json:"reqsta,omitempty"`
	// 请求结果 String (1) S 成功 F 失败 C 撤销 D 过期 企业过期不审批 R 否决 企业审批否决
	Rtnflg string `json:"rtnflg,omitempty"`
	// 记录状态 String(1) T : 未处理 A: 银行处理中 E:完成 C: 撤销、否决 、过期
	Stscod string `json:"stscod,omitempty"`
	// 待处理步骤 String (6)
	Oprstp string `json:"oprstp,omitempty"`
	// 待处理序列 String (3)
	Oprsqn string `json:"oprsqn,omitempty"`
	// 操作别名 String (32) 不超过15个汉字
	Oprals string `json:"oprals,omitempty"`
	// 用户登录名 String (30) 不超过14个汉字
	Lgnnam string `json:"lgnnam,omitempty"`
	// 用户姓名 String (20) 不超过9个汉字
	Usrnam string `json:"usrnam,omitempty"`
	// 业务受理城市号 String (6)
	Ctynbr string `json:"ctynbr,omitempty"`
	// 业务受理部门号 String (4)
	Brnnbr string `json:"brnnbr,omitempty"`
	// 业务受理分行号 String (2)
	Bbknbr string `json:"bbknbr,omitempty"`
	// 账号 String (35) 付方为账号时返回
	Accnbr string `json:"accnbr,omitempty"`
	// 钱包ID String(16) 付方为钱包时返回
	Waltid string `json:"waltid,omitempty"`
	// 户名 String (400) 不超过133个汉字。付方为账户时返回户名
	Accnam string `json:"accnam,omitempty"`
	// 公司名称 String(600) 付方为钱包时返回公司名称
	Fctnam string `json:"fctnam,omitempty"`
	// 交易笔数 N(6) 大于等于0
	Trsnum string `json:"trsnum,omitempty"`
	// 总金额 M 13位整数，2位小数
	Totamt string `json:"totamt,omitempty"`
	// 成功笔数 N(6) 大于等于0
	Sucnum string `json:"sucnum,omitempty"`
	// 成功金额 M 13位整数，2位小数
	Sucamt string `json:"sucamt,omitempty"`
	// 币种 N(2) 两位数字，默认币种赋值为10， 例如10 人民币
	Ccynbr string `json:"ccynbr,omitempty"`
	// 货币市场 String (1) 0 不分钞汇 1 现钞 2 现汇
	Ccymkt string `json:"ccymkt,omitempty"`
	// 交易类型 String (4) BYBC 代发劳务收入 BYBD 代发投资本益 BYBE 代发债产权转让 等
	Trstyp string `json:"trstyp,omitempty"`
	// 用途 String (42) 显示在交易流水中、电子回单中、客户审批可见，收款员工不可见，不超过20个汉字
	Nusage string `json:"nusage,omitempty"`
	// 业务参考号 String (30)
	Yurref string `json:"yurref,omitempty"`
	// 错误码 String (7)
	Errcod string `json:"errcod,omitempty"`
	// 错误描述 String (192) 不超过95个汉字
	Errdsp string `json:"errdsp,omitempty"`
	// 处理批次号 String (30)
	Seqcod string `json:"seqcod,omitempty"`
	// 附件标志 String (1) Y/N，默认为N，当前直联不支持附件
	Athflg string `json:"athflg,omitempty"`
	// 记账子单元 String (20) 预留，代扣暂不支持交易管家（记账）虚拟户
	Dmanbr string `json:"dmanbr,omitempty"`
	// 渠道标志 String (1) Y 超网 N 大小额 M 智能路由（超网和大小额混发） A 数币代发（账户代发到钱包） W钱包代发（钱包代发到钱包）
	Agctyp string `json:"agctyp,omitempty"`
	// 交易流水 String (30)
	Trsreq string `json:"trsreq,omitempty"`
	// 交易套号 String (21)
	Trsset string `json:"trsset,omitempty"`
	// 记账方式 String (1) 不填默认为汇总记账 Y 单笔记账 ， N或不填 汇总记账 单笔记账（依据代发明细生成多笔流水及回单，将展示具体收款人信息，请评估薪资保密需求） 汇总记账（仅生成一笔汇总交易流水及回单） 使用单笔记账前需要联系客户经理申请
	Sigacc string `json:"sigacc,omitempty"`
	// 保留字 String (100) 第1位为P表示部分成功
	Rsv100 string `json:"rsv100,omitempty"`
}

// PayrollNewBb6busmod 模式接口 bb6busmod.
type PayrollNewBb6busmod struct {
	// 业务类型 String(6) required:Y 非空 N03010 代发工资 N03020 代发其他 N03040 钱包代发（从钱包出款，选择此业务时，结算通道赋值W）
	Buscod string `json:"buscod,omitempty"`
	// 业务模式 String(5) required:Y 非空
	Busmod string `json:"busmod,omitempty"`
	// 保留字 String(50)
	Rsv50z string `json:"rsv50z,omitempty"`
}

// PayrollNewBb6cdcbhx1 汇总接口 bb6cdcbhx1.
type PayrollNewBb6cdcbhx1 struct {
	// 业务受理分行号 String(2)
	Bbknbr string `json:"bbknbr,omitempty"`
	// 批次开始标志 String(1) required:Y Y或N 首次传输Y，续传批次N
	Begtag string `json:"begtag,omitempty"`
	// 批次结束标志 String (1) required:Y Y或N 结束批次Y，非结束批次N
	Endtag string `json:"endtag,omitempty"`
	// 流程实例号 String (10) 第一次上传时必须为空； 续传时不能为空，所有续传次数流程实例号必须为同一个；该值为招行系统生成 请注意该字段为String类型
	Reqnbr string `json:"reqnbr,omitempty"`
	// 账号 String (35) 账号出款时非空，其它情况不用赋值
	Accnbr string `json:"accnbr,omitempty"`
	// 户名 String (400) 账户出款时使用填入账户户名，其它情况不用赋值
	Accnam string `json:"accnam,omitempty"`
	// 钱包ID String (16) 钱包出款时非空，其它情况不用赋值
	Waltid string `json:"waltid,omitempty"`
	// 公司名称 String(600) 钱包出款时使用填入公司名称，其它情况不用赋值。传空时后台自动获取
	Fctnam string `json:"fctnam,omitempty"`
	// 总金额 M required:Y 非空，大于0，13位整数，2位小数
	Ttlamt string `json:"ttlamt,omitempty"`
	// 总笔数 N(8) required:Y 非空，大于0
	Ttlcnt string `json:"ttlcnt,omitempty"`
	// 总次数 N(3) required:Y 非空，大于0
	Ttlnum string `json:"ttlnum,omitempty"`
	// 本次金额 M required:Y 非空，大于0，13位整数，2位小数，总笔数少于1000笔，本次金额为总金额；总笔数大于1000笔，将会进行续传，本次金额为本批次总金额
	Curamt string `json:"curamt,omitempty"`
	// 本次笔数 N(8) required:Y 非空，大于0 ，总笔数少于1000笔，本次笔数为总笔数；总笔数大于1000笔，将会进行续传，本次笔数按照1000笔进行分割
	Curcnt string `json:"curcnt,omitempty"`
	// 合作方协议号 String (6) 合作方协议号和交易类型不能同时为空
	Cnvnbr string `json:"cnvnbr,omitempty"`
	// 币种 N(2) required:Y 两位数字，默认币种赋值为10， 例如10 人民币
	Ccynbr string `json:"ccynbr,omitempty"`
	// 个性化短信 String (38) 最大12个中文
	Ntfinf string `json:"ntfinf,omitempty"`
	// 交易类型 String (4) 可以通过代发类型查询 BB6AGTQY接口获取当前账号可经办的交易类型，所有交易类型请参考附录A.45
	Trstyp string `json:"trstyp,omitempty"`
	// 用途 String (100) required:Y 最大33个中文
	Nusage string `json:"nusage,omitempty"`
	// 期望日期 D required:Y 非空 格式：YYYYMMDD
	Eptdat string `json:"eptdat,omitempty"`
	// 期望时间 String (6) HHMMSS,为空的话，默认值为000000
	Epttim string `json:"epttim,omitempty"`
	// 业务参考号 String (30) required:Y 非空，允许数字或者字母，不允许中文， 业务参考号（YURREF）是每笔交易的唯一编号，是防止重复提交的重要手段，如需重复发送请求，请务必保证同一笔交易的业务参考号不变，否则会存在重复提交风险。 同一个业务参考号，只校验企业近13个月内的批次
	Yurref string `json:"yurref,omitempty"`
	// 记账子单元 String (20)
	Dmanbr string `json:"dmanbr,omitempty"`
	// 结算通道 String (1) required:Y 非空， Y 超网 N 大小额 M 智能路由（超网和大小额混发） A 数币代发（账户代发到钱包） W 钱包代发（钱包代发到钱包，当选择W时，BB6BUSMOD中的buscod必须为N03040）
	Chlflg string `json:"chlflg,omitempty"`
	// 记账方式 String (1) 不填默认为汇总记账 Y 单笔记账 ， N或不填 汇总记账 单笔记账（依据代发明细生成多笔流水及回单，将展示具体收款人信息，请评估薪资保密需求） 汇总记账（仅生成一笔汇总交易流水及回单） 默认为汇总记账，使用单笔记账前需要联系客户经理提无纸化申请
	Sigacc string `json:"sigacc,omitempty"`
	// 保留字 String (100)
	Rsv100 string `json:"rsv100,omitempty"`
}

// PayrollNewBb6cdcbhz1 输出接口 bb6cdcbhz1.
type PayrollNewBb6cdcbhz1 struct {
	// 流程实例号 String (10) 该值为招行系统生成 格式含数字和字母
	Reqnbr string `json:"reqnbr,omitempty"`
	// 请求状态 String (3) OPR 数据接收中 AUT 等待审批 NTE 终审完毕 APW 银行人工审批 WRF 可疑 BNK 银行处理中 FIN 完成
	Reqsta string `json:"reqsta,omitempty"`
	// 保留字 String (50)
	Rsv50z string `json:"rsv50z,omitempty"`
}

// PayrollNewBb6cdcdlx1 明细接口 bb6cdcdlx1.
type PayrollNewBb6cdcdlx1 struct {
	// 交易序号 String (8) required:Y 非空 需要客户自行保证批次范围内的序号唯一性，代发代扣系统要求格式为全数字，如 ‘00000001’,’00000002’,不允许“00000000” 每次最多接收1000笔明细
	Trxseq string `json:"trxseq,omitempty"`
	// 账号 String (35) 收方为账户时，必输. 收方为钱包时，传空
	Accnbr string `json:"accnbr,omitempty"`
	// 户名 /员工姓名 String (400) required:Y 非空，账户户名、员工姓名 最大133个中文
	Accnam string `json:"accnam,omitempty"`
	// 数币标记 String(1) 为空时默认N Y：收方为钱包； N：收方是账号
	Digflg string `json:"digflg,omitempty"`
	// 钱包ID String (16) 收方为钱包时，必输. 收方为账户时，传空。
	Waltid string `json:"waltid,omitempty"`
	// 钱包昵称 String (300) 可不传，如果传了会校验
	Walnam string `json:"walnam,omitempty"`
	// 交易金额 M required:Y 非空，大于0，13位整数，2位小数
	Trsamt string `json:"trsamt,omitempty"`
	// 注释 String (200) 最大66个中文
	Trsdsp string `json:"trsdsp,omitempty"`
	// 系统内标志 String (1) 已不再使用该字段，会忽略客户的传值
	Bnkflg string `json:"bnkflg,omitempty"`
	// 他行户口开户行 String (62) 不超过30个汉字 1. 行内卡传入空即可（包括标准银联卡和非标准银联卡） 2. 标准银联卡传入空即可。（可以调用CdcMatchCardBin判断是否银联卡） 3. 行外且非标准银联卡需同时传入正确的开户行开户地，只有1个字段有值会视为无效输入
	Eacbnk string `json:"eacbnk,omitempty"`
	// 他行户口开户地 String (62) 不超过30个汉字 传值同开户行的描述
	Eaccty string `json:"eaccty,omitempty"`
	// 合作方流水号 String(20)
	Cprref string `json:"cprref,omitempty"`
	// 保留字 String (45)
	Rsv45z string `json:"rsv45z,omitempty"`
}

// PayrollNewBb6cdcqyx1 输入接口 bb6cdcqyx1.
type PayrollNewBb6cdcqyx1 struct {
	// 业务参考号 String (30) 1.当传入为空时，按照业务代码和日期查询 2. 非空时，日期字段不需要输入,且会忽略日期的限制
	Yurref string `json:"yurref,omitempty"`
	// 业务类型 String (6) required:Y 非空，N03030
	Buscod string `json:"buscod,omitempty"`
	// 起始日期 D 最大180天区间，只支持近13个月的数据
	Bgndat string `json:"bgndat,omitempty"`
	// 结束日期 D 最大180天区间，只支持近13个月的数据
	Enddat string `json:"enddat,omitempty"`
	// 日期类型 String (1) 空时默认为A,按经办日查 B：按期望日查
	Datflg string `json:"datflg,omitempty"`
}

// PayrollNewBb6cdcqyy1 输入输出接口 bb6cdcqyy1.
type PayrollNewBb6cdcqyy1 struct {
	// 续传键值 String (80) 当返回报文有bb6cdcqyy1接口返回且 ctnkey 字段不为空时标明需要续传，请将返回报文的bb6cdcqyy1内容放入下次请求继续查询
	Ctnkey string `json:"ctnkey,omitempty"`
	// 保留字 String (50)
	Rsv50z string `json:"rsv50z,omitempty"`
}

// PayrollNewBb6cdcqyz1 输出接口 bb6cdcqyz1.
type PayrollNewBb6cdcqyz1 struct {
	// 流程实例号 String (10) 该值为招行系统生成
	Reqnbr string `json:"reqnbr,omitempty"`
	// 业务类型 String (6) N03030
	Buscod string `json:"buscod,omitempty"`
	// 业务模式 String (5) 00001
	Busmod string `json:"busmod,omitempty"`
	// 经办日 D 格式YYYYMMDD
	Oprdat string `json:"oprdat,omitempty"`
	// 期望日 D 格式YYYYMMDD
	Eptdat string `json:"eptdat,omitempty"`
	// 期望时间 String (6) HHMMSS
	Epttim string `json:"epttim,omitempty"`
	// 业务请求状态 String (3) OPR 数据接收中 AUT 等待审批 NTE 终审完毕 APW 银行人工审批 WRF 可疑 BNK 银行处理中 FIN 完成
	Reqsta string `json:"reqsta,omitempty"`
	// 业务请求结果 String (1) S 成功 F 失败 C 撤销 D 过期 R 否决
	Rtnflg string `json:"rtnflg,omitempty"`
	// 待处理操作步骤 String (6)
	Oprstp string `json:"oprstp,omitempty"`
	// 待处理操作序列 String (3)
	Oprsqn string `json:"oprsqn,omitempty"`
	// 操作别名 Z(32) 不超过30个汉字
	Oprals string `json:"oprals,omitempty"`
	// 用户登录名 Z(30) 不超过14个汉字
	Lgnnam string `json:"lgnnam,omitempty"`
	// 用户姓名 Z(20) 不超过9个汉字
	Usrnam string `json:"usrnam,omitempty"`
	// 业务受理城市号 String (6)
	Ctynbr string `json:"ctynbr,omitempty"`
	// 业务受理部门号 String (4)
	Brnnbr string `json:"brnnbr,omitempty"`
	// 业务受理分行号 String (2)
	Bbknbr string `json:"bbknbr,omitempty"`
	// 账号 String (35)
	Accnbr string `json:"accnbr,omitempty"`
	// 户名 Z(62) 不超过30个汉字 ，固定空值
	Accnam string `json:"accnam,omitempty"`
	// 交易笔数 N(6)
	Trsnum string `json:"trsnum,omitempty"`
	// 总金额 M 大于0，13位整数，2位小数
	Totamt string `json:"totamt,omitempty"`
	// 币种 N(2) 两位数字，默认币种赋值为10， 例如10 人民币
	Ccynbr string `json:"ccynbr,omitempty"`
	// 货币市场 String (1) 空值
	Ccymkt string `json:"ccymkt,omitempty"`
	// 代扣类型 String (4) AKTI=代扣出租车规费 AKTL=代扣电话费 AKTV=代扣有线电视费 AKTW=代扣水电费 AKTX=代扣税费 AKWD=代扣业务 AKWT=代扣水费等
	Trstyp string `json:"trstyp,omitempty"`
	// 成功笔数 N(6)
	Sucnum string `json:"sucnum,omitempty"`
	// 成功金额 M 13位整数，2位小数
	Sucamt string `json:"sucamt,omitempty"`
	// 用途 Z(42) 不超过20个汉字，显示在交易流水中、电子回单中、客户审批可见，收款员工不可见
	Nusage string `json:"nusage,omitempty"`
	// 业务参考号 String (30)
	Yurref string `json:"yurref,omitempty"`
	// 记录状态 String (1) T : 未处理 A: 银行处理中 E:完成 C: 撤销、否决 、过期
	Stscod string `json:"stscod,omitempty"`
	// 错误码 String (7)
	Errcod string `json:"errcod,omitempty"`
	// 错误描述 Z(92) 不超过45个汉字
	Errdsp string `json:"errdsp,omitempty"`
	// 处理批次号 String (30)
	Seqcod string `json:"seqcod,omitempty"`
	// 附件标志 String (1) Y/N，默认为N，当前直联不支持附件
	Athflg string `json:"athflg,omitempty"`
	// 记账子单元编号 String (20) 预留，代扣暂不支持交易管家（记账）虚拟户
	Dmanbr string `json:"dmanbr,omitempty"`
	// 渠道标志 String (1) 新增，Y是超网，N是大小额
	Agctyp string `json:"agctyp,omitempty"`
	// 交易流水 String (30)
	Trsreq string `json:"trsreq,omitempty"`
	// 交易套号 String (21)
	Trsset string `json:"trsset,omitempty"`
	// 保留字 String (100) 第一位为P时，表示部分成功
	Rsv100 string `json:"rsv100,omitempty"`
}

// PayrollNewBb6dtlqyy1 输入/输出接口 bb6dtlqyy1.
type PayrollNewBb6dtlqyy1 struct {
	// 流程实例号 String (10) required:Y 该值为招行系统生成， 格式含数字和字母 请注意该字段为String类型
	Reqnbr string `json:"reqnbr,omitempty"`
	// 续传批次号 String (30) 续传时，将接收值原封不动传回
	Bthnbr string `json:"bthnbr,omitempty"`
	// 续传明细序号 String (8) 当返回报文有bb6dtlqyy1接口返回且trxseq字段不为空时标明需要续传，请将返回报文的bb6dtlqyy1内容放入下次请求继续查询
	Trxseq string `json:"trxseq,omitempty"`
	// 续传历史查询标志 String (1) 续传时，将接收值原封不动传回
	Histag string `json:"histag,omitempty"`
	// 保留字 String (50)
	Rsv50z string `json:"rsv50z,omitempty"`
}

// PayrollNewBb6dtlqyz1 输出接口 bb6dtlqyz1.
type PayrollNewBb6dtlqyz1 struct {
	// 批次号码 String (30)
	Bthnbr string `json:"bthnbr,omitempty"`
	// 交易序号 String (8)
	Trxseq string `json:"trxseq,omitempty"`
	// 账号 String (35) 收方为账户时返回
	Accnbr string `json:"accnbr,omitempty"`
	// 钱包ID String(16) 收方为钱包时有值
	Waltid string `json:"waltid,omitempty"`
	// 户名/真实姓名 Z(400) 不超过133个汉字
	Accnam string `json:"accnam,omitempty"`
	// 钱包昵称 Z(300) 经办时传入才返回
	Walnam string `json:"walnam,omitempty"`
	// 数币标记 String(1) Y 钱包 N 账户
	Digflg string `json:"digflg,omitempty"`
	// 交易金额 M 13位整数，2位小数
	Trsamt string `json:"trsamt,omitempty"`
	// 交易日期 D 格式YYYYMMDD
	Trsdat string `json:"trsdat,omitempty"`
	// 注释 Z(42) 不超过20个汉字
	Trsdsp string `json:"trsdsp,omitempty"`
	// 交易状态 String (1) E：失败 A: 登记 S: 成功
	Stscod string `json:"stscod,omitempty"`
	// 失败代码 String (7)
	Errcod string `json:"errcod,omitempty"`
	// 失败原因 Z(192) 不超过95个汉字
	Errtxt string `json:"errtxt,omitempty"`
	// 系统内标志 String (1) Y：开户行为招行； N：开户行为他行
	Bnkflg string `json:"bnkflg,omitempty"`
	// 他行户口开户行 Z(62) 非超网有值, 不超过30个汉字
	Eacbnk string `json:"eacbnk,omitempty"`
	// 他行户口开户地 Z(62) 非超网有值 不超过30个汉字
	Eaccty string `json:"eaccty,omitempty"`
	// 他行快速标志 String (1) 赋值为空
	Fstflg string `json:"fstflg,omitempty"`
	// 他行户口联行号 String (12) 非超网有值
	Rcvbnk string `json:"rcvbnk,omitempty"`
	// 他行钱包运营机构号 String (14) 收方为行外钱包时可能有值
	Bnkorg string `json:"bnkorg,omitempty"`
	// 他行钱包运营机构名称 String (400) 收方为行外钱包时可能有值
	Orgnam string `json:"orgnam,omitempty"`
	// 客户代码 String (20) 赋值为空
	Cpract string `json:"cpract,omitempty"`
	// 合作方流水号 String (20) 非超网有值
	Cprref string `json:"cprref,omitempty"`
	// 交易流水号 String (30)
	Trssqn string `json:"trssqn,omitempty"`
	// 交易套号 String (21)
	Trsset string `json:"trsset,omitempty"`
	// 对接代理行状态 String(1) 当前为空，后续升级支持
	Othsts string `json:"othsts,omitempty"`
	// 退票附言 String(200) 对手行返回的退票附言
	RtnTxt string `json:"rtnTxt,omitempty"`
	// 付方交易套号 String (30) 单笔记账对账使用
	Payseq string `json:"payseq,omitempty"`
	// 付方交易流水 String (21) 单笔记账对账使用
	Paysqn string `json:"paysqn,omitempty"`
	// 保留字 String (50)
	Rsv30z string `json:"rsv30z,omitempty"`
}

// PayrollNewBb6intqyy1 输入/输出接口 bb6intqyy1.
type PayrollNewBb6intqyy1 struct {
	// 业务编码 String (6) N03010 代发工资 N03020 代发 N03030 代扣 N03040 钱包代发 ****** 查全部(按时间区间查时可以查全部，当按业务参考号查时不能输入******）
	Buscod string `json:"buscod,omitempty"`
	// 业务参考号 String (30) 如果输入业务参考号，则查询指定业务参考号批次，则日期字段不需要输入,且会忽略日期的限制
	Yurref string `json:"yurref,omitempty"`
	// 起始日期 D 业务参考号为空，起始日期非空 最大180天区间，只支持近13个月的数据
	Bgndat string `json:"bgndat,omitempty"`
	// 结束日期 D 业务参考号为空，结束日期非空 最大180天区间，只支持近13个月的数据
	Enddat string `json:"enddat,omitempty"`
	// 日期类型 String (1) A：经办日期； B：期望日期； 默认为A
	Dattyp string `json:"dattyp,omitempty"`
	// 经办用户 String (30) 不超过14个汉字
	Oprlgn string `json:"oprlgn,omitempty"`
	// 续传键值 String (80) 当返回报文有bb6intqyy1接口返回且cntkey字段不为空时标明需要续传，请将返回报文的bb6intqyy1内容放入下次请求继续查询
	Cntkey string `json:"cntkey,omitempty"`
	// 保留字 String (50)
	Rsv50z string `json:"rsv50z,omitempty"`
}

// PayrollNewBb6intqyz1 输出接口 bb6intqyz1.
type PayrollNewBb6intqyz1 struct {
	// 流程实例号 String(10) 该值为招行系统生成 格式含数字和字母 请注意该字段为String类型
	Reqnbr string `json:"reqnbr,omitempty"`
	// 业务类型 String (6) N03010 代发工资 N03020 代发其他 N03030 代扣 N03040 钱包代发
	Buscod string `json:"buscod,omitempty"`
	// 业务模式 String(5) 5位数字，例00001
	Busmod string `json:"busmod,omitempty"`
	// 经办日 D 格式YYYYMMDD
	Oprdat string `json:"oprdat,omitempty"`
	// 期望日 D 格式YYYYMMDD
	Eptdat string `json:"eptdat,omitempty"`
	// 期望时间 String (6) 格式HHMMSS
	Epttim string `json:"epttim,omitempty"`
	// 业务请求状态 String (3) OPR 数据接收中 AUT 等待审批 NTE 终审完毕 APW 银行人工审批 WRF 可疑 BNK 银行处理中 FIN 完成
	Reqsta string `json:"reqsta,omitempty"`
	// 业务请求结果 String (1) S 成功 F 失败 C 撤销 D 过期 企业过期不审批 R 否决 企业审批否决
	Rtnflg string `json:"rtnflg,omitempty"`
	// 待处理操作步骤 String (6)
	Oprstp string `json:"oprstp,omitempty"`
	// 待处理操作序列 String (3)
	Oprsqn string `json:"oprsqn,omitempty"`
	// 操作别名 String (32) 不超过15个汉字
	Oprals string `json:"oprals,omitempty"`
	// 用户登录名 String (30) 不超过14个汉字
	Lgnnam string `json:"lgnnam,omitempty"`
	// 用户姓名 String (20) 不超过9个汉字
	Usrnam string `json:"usrnam,omitempty"`
	// 业务受理城市号 String (6) 固定为空
	Ctynbr string `json:"ctynbr,omitempty"`
	// 业务受理部门号 String (4) 固定为空
	Brnnbr string `json:"brnnbr,omitempty"`
	// 业务受理分行号 String (2)
	Bbknbr string `json:"bbknbr,omitempty"`
	// 账号 String (35) 付方为账号时返回
	Accnbr string `json:"accnbr,omitempty"`
	// 钱包ID String(16) 付方为钱包时返回
	Waltid string `json:"waltid,omitempty"`
	// 户名 String (400) 不超过133个汉字。 付方为账户时返回户名
	Accnam string `json:"accnam,omitempty"`
	// 公司名称 String(600) 付方为钱包时返回公司名称
	Fctnam string `json:"fctnam,omitempty"`
	// 交易笔数 N(6)
	Trsnum string `json:"trsnum,omitempty"`
	// 总金额 M 13位整数，2位小数
	Totamt string `json:"totamt,omitempty"`
	// 币种 N(2) 两位数字，默认币种赋值为10， 例如10 人民币
	Ccynbr string `json:"ccynbr,omitempty"`
	// 货币市场 String (1) 固定为空
	Ccymkt string `json:"ccymkt,omitempty"`
	// 代发代扣类型 String (4) BYBC 代发劳务收入 BYBD 代发投资本益 BYBE 代发债产权转让 AKTI=代扣出租车规费 AKTL=代扣电话费 AKTV=代扣有线电视费 AKTW=代扣水电费 AKTX=代扣税费 等
	Trstyp string `json:"trstyp,omitempty"`
	// 成功笔数 N(6)
	Sucnum string `json:"sucnum,omitempty"`
	// 成功金额 M 13位整数，2位小数
	Sucamt string `json:"sucamt,omitempty"`
	// 用途 String (42) 不超过20个汉字
	Nusage string `json:"nusage,omitempty"`
	// 业务参考号 String (30)
	Yurref string `json:"yurref,omitempty"`
	// 记录状态 String (1) T : 未处理 A: 银行处理中 E:完成 C: 撤销、否决 、过期
	Stscod string `json:"stscod,omitempty"`
	// 错误码 String (7)
	Errcod string `json:"errcod,omitempty"`
	// 错误描述 String (92) 不超过45个汉字
	Errdsp string `json:"errdsp,omitempty"`
	// 处理批号 String (30)
	Seqcod string `json:"seqcod,omitempty"`
	// 附件标志 String (1) 固定值N，当前直联不支持附件
	Athflg string `json:"athflg,omitempty"`
	// 渠道标志 String (1) Y 超网 N 大小额 M 智能路由（超网和大小额混发） A 数币代发(账户代发到钱包) W钱包代发(钱包代发到钱包 )
	Agctyp string `json:"agctyp,omitempty"`
	// 记账子单元 String (20) 预留，代扣暂不支持交易管家（记账）虚拟户
	Dmanbr string `json:"dmanbr,omitempty"`
	// 交易流水 String (30)
	Trsreq string `json:"trsreq,omitempty"`
	// 交易套号 String (21)
	Trsset string `json:"trsset,omitempty"`
	// 记账方式 String (1) 不填默认为汇总记账 Y 单笔记账 ， N或不填 汇总记账 单笔记账（依据代发明细生成多笔流水及回单，将展示具体收款人信息，请评估薪资保密需求） 汇总记账（仅生成一笔汇总交易流水及回单） 使用单笔记账前需要联系客户经理申请
	Sigacc string `json:"sigacc,omitempty"`
	// 保留字 String (100) 第1位为P表示部分成功
	Rsv100 string `json:"rsv100,omitempty"`
}

// PayrollNewBb6rfdqyy1 输入/输出接口 bb6rfdqyy1.
type PayrollNewBb6rfdqyy1 struct {
	// 账号 String (35) 账号和钱包ID不能同时为空，也不能同时有值
	Accnbr string `json:"accnbr,omitempty"`
	// 钱包ID String（16） 当查询钱包出款的退票时，输入钱包ID。账号和钱包ID不能同时为空，也不能同时有值
	Waltid string `json:"waltid,omitempty"`
	// 代发种类 String (4) required:Y 非空 BYBC 代发劳务收入 BYBD 代发投资本益 BYBE 代发债产权转让 等
	Trstyp string `json:"trstyp,omitempty"`
	// 交易开始日期 D required:Y 非空，格式YYYYMMDD, 查询时间范围不超过31天，并且不超过13个月之前
	Bgndat string `json:"bgndat,omitempty"`
	// 交易结束日期 D required:Y 非空，格式YYYYMMDD, 查询时间范围不超过31天，并且不超过13个月之前
	Enddat string `json:"enddat,omitempty"`
	// 流程实例号 String (10) 该值为招行系统生成，当前实际长度为10位，后续升级可能会拓展到20位 请注意该字段为String类型
	Reqnbr string `json:"reqnbr,omitempty"`
	// 续传键值 String (80) 当返回报文有bb6rfdqyy1接口返回且 ctnkey 字段不为空时标明需要续传，请将返回报文的bb6rfdqyy1内容放入下次请求继续查询
	Ctnkey string `json:"ctnkey,omitempty"`
	// 保留字段 String (50)
	Rsv50z string `json:"rsv50z,omitempty"`
}

// PayrollNewBb6rfdqyz1 输出接口 bb6rfdqyz1.
type PayrollNewBb6rfdqyz1 struct {
	// 流程实例号 String (10) 该值为招行系统生成， 格式含数字和字母 请注意该字段为String类型
	Reqnbr string `json:"reqnbr,omitempty"`
	// 业务参考号 String (30) 原代发的业务参考号
	Yurref string `json:"yurref,omitempty"`
	// 明细序号 String (8)
	Bthdtl string `json:"bthdtl,omitempty"`
	// 明细账号 String (35)
	Eacnbr string `json:"eacnbr,omitempty"`
	// 明细户名 Z(400) 不超过100个汉字
	Eacnam string `json:"eacnam,omitempty"`
	// 交易金额 M 13位整数，2位小数
	Trsamt string `json:"trsamt,omitempty"`
	// 客户在商户账号 String (20)
	Cpract string `json:"cpract,omitempty"`
	// 客户摘要 Z(42) 不超过20个汉字
	Txtclt string `json:"txtclt,omitempty"`
	// 明细账户开户行 Z(62) 不超过30个汉字
	Eacbnm string `json:"eacbnm,omitempty"`
	// 明细账户开户地 Z(62) 不超过30个汉字
	Eaccty string `json:"eaccty,omitempty"`
	// 商户协议号 String (6)
	Cprcnv string `json:"cprcnv,omitempty"`
	// 交易日期 D 格式YYYYMMDD
	Trsdat string `json:"trsdat,omitempty"`
	// 交易状态 String (1) E：失败
	Stscod string `json:"stscod,omitempty"`
	// 退票失败代码 String (5)
	Rtncod string `json:"rtncod,omitempty"`
	// 退票附言 Z(100) 不超过49个汉字 收方银行返回的退票附言，对于该值有疑问请直接联系对方行
	Rtntxt string `json:"rtntxt,omitempty"`
	// 退票交易流水号 String (30)
	Trsreq string `json:"trsreq,omitempty"`
	// 退票交易套号 String (21)
	Trsset string `json:"trsset,omitempty"`
}

// PayrollNewBb6trsqyx1 输入接口 bb6trsqyx1.
type PayrollNewBb6trsqyx1 struct {
	// 业务代码 String (6) required:Y 非空 N03030 代扣
	Buscod string `json:"buscod,omitempty"`
	// 账号 String (35) required:Y 非空
	Accnbr string `json:"accnbr,omitempty"`
	// 币种 N(2) 两位数字，默认币种赋值为10， 例如10 人民币
	Ccynbr string `json:"ccynbr,omitempty"`
	// 协议状态 String (1) 可空，默认为‘A’ A 有效 C 关闭
	Stscod string `json:"stscod,omitempty"`
	// 保留字 String (50)
	Rsv50z string `json:"rsv50z,omitempty"`
}

// PayrollNewBb6trsqyz1 输出接口 bb6trsqyz1.
type PayrollNewBb6trsqyz1 struct {
	// 账户 String (35)
	Accnbr string `json:"accnbr,omitempty"`
	// 协议号 String (6)
	Cnvnbr string `json:"cnvnbr,omitempty"`
	// 代扣类型 String (4) AKTI=代扣出租车规费 AKTL=代扣电话费 AKTV=代扣有线电视费 AKTW=代扣水电费 AKTX=代扣税费 AKWD=代扣业务 AKWT=代扣水费等
	Trstyp string `json:"trstyp,omitempty"`
	// 币种 N(2) 两位数字，默认币种赋值为10， 例如10 人民币
	Ccynbr string `json:"ccynbr,omitempty"`
	// 生效日期 D 格式YYYYMMDD
	Eftdat string `json:"eftdat,omitempty"`
	// 失效日期 D 格式YYYYMMDD
	Expdat string `json:"expdat,omitempty"`
	// 签约日期 D 格式YYYYMMDD
	Sgndat string `json:"sgndat,omitempty"`
	// 三方协议标志 String (1) Y/N
	Cltcnv string `json:"cltcnv,omitempty"`
	// 协议状态 String (1) A 有效 C 关闭
	Stscod string `json:"stscod,omitempty"`
	// 保留字 30 String (30)
	Rsv30z string `json:"rsv30z,omitempty"`
}

// PayrollNewBb6typqyx1 输入接口 bb6typqyx1.
type PayrollNewBb6typqyx1 struct {
	// 业务类型 String (6) required:Y 非空 N03010 代发工资 N03020 代发其他 N03040 钱包代发
	Buscod string `json:"buscod,omitempty"`
	// 钱包id String(5) 当从钱包出款代发时，非空
	Waltid string `json:"waltid,omitempty"`
	// 账号 String (35) 查询账户协议时非空。 查询钱包协议时，该字段输入空
	Accnbr string `json:"accnbr,omitempty"`
	// 币种 N(2) 两位数字，默认币种赋值为10， 例如10 人民币
	Ccynbr string `json:"ccynbr,omitempty"`
	// 协议状态 String (1) A 有效 C 关闭 为空默认赋值为A
	Stscod string `json:"stscod,omitempty"`
}

// PayrollNewBb6typqyz1 输出接口 bb6typqyz1.
type PayrollNewBb6typqyz1 struct {
	// 账户 String (35) 付款为账户时才返回
	Accnbr string `json:"accnbr,omitempty"`
	// 协议号 String (6)
	Cnvnbr string `json:"cnvnbr,omitempty"`
	// 代发类型 String (4) BYBC 代发劳务收入 BYBD 代发投资本益 BYBE 代发债产权转让 等
	Trstyp string `json:"trstyp,omitempty"`
	// 币种 N(2) 两位数字，默认币种赋值为10，例如10为人民币
	Ccynbr string `json:"ccynbr,omitempty"`
	// 生效日期 D 格式YYYYMMDD
	Eftdat string `json:"eftdat,omitempty"`
	// 失效日期 D 格式YYYYMMDD
	Expdat string `json:"expdat,omitempty"`
	// 签约日期 D 格式YYYYMMDD
	Sgndat string `json:"sgndat,omitempty"`
	// 三方协议标志 String (1)
	Cltcnv string `json:"cltcnv,omitempty"`
	// 协议状态 String (1) A 有效 C 关闭
	Stscod string `json:"stscod,omitempty"`
	// 单笔记账标记 String(1) 协议默认为汇总记账，使用单笔记账前需要联系客户经理提无纸化申请。N或空 只允许汇总记账（默认方式），Y 允许单笔记账/汇总记账 。单笔记账（依据代发明细生成多笔流水及回单，将展示具体收款人信息，请评估薪资保密需求）汇总记账（仅生成一笔汇总交易流水及回单）
	Sigflg string `json:"sigflg,omitempty"`
}

// PayrollNewCdcMatchCardBinX1 输入接口 CdcMatchCardBinX1.
type PayrollNewCdcMatchCardBinX1 struct {
	// 账号 string(35) required:Y
	AccNbr string `json:"accNbr,omitempty"`
	// 保留字段 String(50)
	Rsv50z string `json:"rsv50z,omitempty"`
}

// PayrollNewCdcMatchCardBinZ1 输出接口 CdcMatchCardBinZ1.
type PayrollNewCdcMatchCardBinZ1 struct {
	// 账号 String(35)
	AccNbr string `json:"accNbr,omitempty"`
	// 是否标准银联卡 String(1) Y:标准银联卡 N:非标准银联卡
	UniFlg string `json:"uniFlg,omitempty"`
	// 行内外标志 String(1) Y:我行 N:他行 标准银联卡才返回这个字段。非标准银联卡无法判断行内行外，用户需自行确认
	BnkFlg string `json:"bnkFlg,omitempty"`
	// 保留字段 String(50)
	Rsv50z string `json:"rsv50z,omitempty"`
	// 是否支持超网代发 String(1) Y：支持 N：不支持（走大小额）
	SntFlg string `json:"sntFlg,omitempty"`
}

// PayrollNewCdcMatchCardBinBody 1.卡BIN匹配识别标准银联卡CdcMatchCardBin request body.
type PayrollNewCdcMatchCardBinBody struct {
	CdcMatchCardBinX1 []*PayrollNewCdcMatchCardBinX1 `json:"CdcMatchCardBinX1,omitempty"`
}

// PayrollNewCdcMatchCardBinResponseBody 1.卡BIN匹配识别标准银联卡CdcMatchCardBin response body.
type PayrollNewCdcMatchCardBinResponseBody struct {
	CdcMatchCardBinZ1 []*PayrollNewCdcMatchCardBinZ1 `json:"CdcMatchCardBinZ1,omitempty"`
}

type PayrollNewCdcMatchCardBinResponse struct {
	Response struct {
		Body PayrollNewCdcMatchCardBinResponseBody `json:"body,omitempty"`
		Head HeadResponse                          `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// PayrollNewBB6BTHHLBody 2.代发经办BB6BTHHL request body.
type PayrollNewBB6BTHHLBody struct {
	Bb6busmod  []*PayrollNewBb6busmod  `json:"bb6busmod,omitempty"`
	Bb6cdcbhx1 []*PayrollNewBb6cdcbhx1 `json:"bb6cdcbhx1,omitempty"`
	Bb6cdcdlx1 []*PayrollNewBb6cdcdlx1 `json:"bb6cdcdlx1,omitempty"`
}

// PayrollNewBB6BTHHLResponseBody 2.代发经办BB6BTHHL response body.
type PayrollNewBB6BTHHLResponseBody struct {
	Bb6cdcbhz1 []*PayrollNewBb6cdcbhz1 `json:"bb6cdcbhz1,omitempty"`
}

type PayrollNewBB6BTHHLResponse struct {
	Response struct {
		Body PayrollNewBB6BTHHLResponseBody `json:"body,omitempty"`
		Head HeadResponse                   `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// PayrollNewBB6BPDQYBody 3.批次与明细查询 BB6BPDQY request body.
type PayrollNewBB6BPDQYBody struct {
	Bb6bpdqyy1 []*PayrollNewBb6bpdqyy1 `json:"bb6bpdqyy1,omitempty"`
}

// PayrollNewBB6BPDQYResponseBody 3.批次与明细查询 BB6BPDQY response body.
type PayrollNewBB6BPDQYResponseBody struct {
	Bb6bpdqyy1 []*PayrollNewBb6bpdqyy1 `json:"bb6bpdqyy1,omitempty"`
	Bb6bpdqyz1 []*PayrollNewBb6bpdqyz1 `json:"bb6bpdqyz1,omitempty"`
	Bb6bpdqyz2 []*PayrollNewBb6bpdqyz2 `json:"bb6bpdqyz2,omitempty"`
}

type PayrollNewBB6BPDQYResponse struct {
	Response struct {
		Body PayrollNewBB6BPDQYResponseBody `json:"body,omitempty"`
		Head HeadResponse                   `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// PayrollNewBB6INTQYBody 4. 代发代扣综合查询 BB6INTQY request body.
type PayrollNewBB6INTQYBody struct {
	Bb6intqyy1 []*PayrollNewBb6intqyy1 `json:"bb6intqyy1,omitempty"`
}

// PayrollNewBB6INTQYResponseBody 4. 代发代扣综合查询 BB6INTQY response body.
type PayrollNewBB6INTQYResponseBody struct {
	Bb6intqyz1 []*PayrollNewBb6intqyz1 `json:"bb6intqyz1,omitempty"`
}

type PayrollNewBB6INTQYResponse struct {
	Response struct {
		Body PayrollNewBB6INTQYResponseBody `json:"body,omitempty"`
		Head HeadResponse                   `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// PayrollNewBB6BTHQYBody 5. 代发批次查询 BB6BTHQY request body.
type PayrollNewBB6BTHQYBody struct {
	Bb6bthqyy1 []*PayrollNewBb6bthqyy1 `json:"bb6bthqyy1,omitempty"`
}

// PayrollNewBB6BTHQYResponseBody 5. 代发批次查询 BB6BTHQY response body.
type PayrollNewBB6BTHQYResponseBody struct {
	Bb6bthqyz1 []*PayrollNewBb6bthqyz1 `json:"bb6bthqyz1,omitempty"`
}

type PayrollNewBB6BTHQYResponse struct {
	Response struct {
		Body PayrollNewBB6BTHQYResponseBody `json:"body,omitempty"`
		Head HeadResponse                   `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// PayrollNewBB6DTLQYBody 6. 代发明细查询 BB6DTLQY request body.
type PayrollNewBB6DTLQYBody struct {
	Bb6dtlqyy1 []*PayrollNewBb6dtlqyy1 `json:"bb6dtlqyy1,omitempty"`
}

// PayrollNewBB6DTLQYResponseBody 6. 代发明细查询 BB6DTLQY response body.
type PayrollNewBB6DTLQYResponseBody struct {
	Bb6dtlqyz1 []*PayrollNewBb6dtlqyz1 `json:"bb6dtlqyz1,omitempty"`
}

type PayrollNewBB6DTLQYResponse struct {
	Response struct {
		Body PayrollNewBB6DTLQYResponseBody `json:"body,omitempty"`
		Head HeadResponse                   `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// PayrollNewBB6RFDQYBody 7.代发退票查询 BB6RFDQY request body.
type PayrollNewBB6RFDQYBody struct {
	Bb6rfdqyy1 []*PayrollNewBb6rfdqyy1 `json:"bb6rfdqyy1,omitempty"`
}

// PayrollNewBB6RFDQYResponseBody 7.代发退票查询 BB6RFDQY response body.
type PayrollNewBB6RFDQYResponseBody struct {
	Bb6rfdqyz1 []*PayrollNewBb6rfdqyz1 `json:"bb6rfdqyz1,omitempty"`
}

type PayrollNewBB6RFDQYResponse struct {
	Response struct {
		Body PayrollNewBB6RFDQYResponseBody `json:"body,omitempty"`
		Head HeadResponse                   `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// PayrollNewBB6AGTQYBody 8.代发类型查询 BB6AGTQY request body.
type PayrollNewBB6AGTQYBody struct {
	Bb6typqyx1 []*PayrollNewBb6typqyx1 `json:"bb6typqyx1,omitempty"`
}

// PayrollNewBB6AGTQYResponseBody 8.代发类型查询 BB6AGTQY response body.
type PayrollNewBB6AGTQYResponseBody struct {
	Bb6typqyz1 []*PayrollNewBb6typqyz1 `json:"bb6typqyz1,omitempty"`
}

type PayrollNewBB6AGTQYResponse struct {
	Response struct {
		Body PayrollNewBB6AGTQYResponseBody `json:"body,omitempty"`
		Head HeadResponse                   `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// PayrollNewBB6ACLAKBody 9.代扣经办BB6ACLAK request body.
type PayrollNewBB6ACLAKBody struct {
	Bb6busmod  []*PayrollNewBb6busmod  `json:"bb6busmod,omitempty"`
	Bb6aclakx1 []*PayrollNewBb6aclakx1 `json:"bb6aclakx1,omitempty"`
	Bb6aclaky1 []*PayrollNewBb6aclaky1 `json:"bb6aclaky1,omitempty"`
}

// PayrollNewBB6ACLAKResponseBody 9.代扣经办BB6ACLAK response body.
type PayrollNewBB6ACLAKResponseBody struct {
	Bb6aclakz1 []*PayrollNewBb6aclakz1 `json:"bb6aclakz1,omitempty"`
}

type PayrollNewBB6ACLAKResponse struct {
	Response struct {
		Body PayrollNewBB6ACLAKResponseBody `json:"body,omitempty"`
		Head HeadResponse                   `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// PayrollNewBB6CDCQYBody 10. 代扣批次查询 BB6CDCQY request body.
type PayrollNewBB6CDCQYBody struct {
	Bb6cdcqyx1 []*PayrollNewBb6cdcqyx1 `json:"bb6cdcqyx1,omitempty"`
	Bb6cdcqyy1 []*PayrollNewBb6cdcqyy1 `json:"bb6cdcqyy1,omitempty"`
}

// PayrollNewBB6CDCQYResponseBody 10. 代扣批次查询 BB6CDCQY response body.
type PayrollNewBB6CDCQYResponseBody struct {
	Bb6cdcqyz1 []*PayrollNewBb6cdcqyz1 `json:"bb6cdcqyz1,omitempty"`
}

type PayrollNewBB6CDCQYResponse struct {
	Response struct {
		Body PayrollNewBB6CDCQYResponseBody `json:"body,omitempty"`
		Head HeadResponse                   `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// PayrollNewBB6ACLIFBody 11.代扣明细查询BB6ACLIF request body.
type PayrollNewBB6ACLIFBody struct {
	Bb6aclify1 []*PayrollNewBb6aclify1 `json:"bb6aclify1,omitempty"`
}

// PayrollNewBB6ACLIFResponseBody 11.代扣明细查询BB6ACLIF response body.
type PayrollNewBB6ACLIFResponseBody struct {
	Bb6aclifz1 []*PayrollNewBb6aclifz1 `json:"bb6aclifz1,omitempty"`
}

type PayrollNewBB6ACLIFResponse struct {
	Response struct {
		Body PayrollNewBB6ACLIFResponseBody `json:"body,omitempty"`
		Head HeadResponse                   `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// PayrollNewBB6TRSQYBody 12.代扣类型查询BB6TRSQY request body.
type PayrollNewBB6TRSQYBody struct {
	Bb6trsqyx1 []*PayrollNewBb6trsqyx1 `json:"bb6trsqyx1,omitempty"`
}

// PayrollNewBB6TRSQYResponseBody 12.代扣类型查询BB6TRSQY response body.
type PayrollNewBB6TRSQYResponseBody struct {
	Bb6trsqyz1 []*PayrollNewBb6trsqyz1 `json:"bb6trsqyz1,omitempty"`
}

type PayrollNewBB6TRSQYResponse struct {
	Response struct {
		Body PayrollNewBB6TRSQYResponseBody `json:"body,omitempty"`
		Head HeadResponse                   `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// PayrollNewDCAGPPDFBody 13.代发明细对账单查询请求DCAGPPDF request body.
type PayrollNewDCAGPPDFBody struct {
	// 付款账户 string (35) required:Y
	Payeac string `json:"payeac,omitempty"`
	// 开始日期 D required:Y YYYY-MM-DD格式
	Begdat string `json:"begdat,omitempty"`
	// 结束日期 D required:Y YYYY-MM-DD格式
	Enddat string `json:"enddat,omitempty"`
	// 业务类型 string (10) required:Y 代发工资:N03010,其他：N03020
	Buscod string `json:"buscod,omitempty"`
	// 业务模式 string (10) required:Y 对应业务类型的业务模式，查询的账号需要在该业务模式的可经办账号列表里面
	Busmod string `json:"busmod,omitempty"`
	// 收方户名 string (100)
	Eacnam string `json:"eacnam,omitempty"`
	// 收方账号 string (100)
	Eacnbr string `json:"eacnbr,omitempty"`
	// 业务参考号 string (30)
	Ptyref string `json:"ptyref,omitempty"`
	// 流水号 string (30)
	Trxsrl string `json:"trxsrl,omitempty"`
	// 最小金额 M
	Minamt string `json:"minamt,omitempty"`
	// 最大金额 M
	Maxamt string `json:"maxamt,omitempty"`
	// 打印模式 string (1) required:Y S：逐笔打印：每笔明细生成一个PDF文件 M：批量打印：所有明细生成生成到一个PDF文件里面
	Prtmod string `json:"prtmod,omitempty"`
	// 查询标记 string (4) required:Y 首次查询请输入0，续传时请将返回报文中的begidx值放入该字段继续请求。
	Begidx string `json:"begidx,omitempty"`
	// 每批数量 string (4) 默认1000，最大限制1000，打印模式为M时表示一个文件里面包含的明细的条数，打印模式为S时表示一次可以查询明细的条数。
	Pagsiz string `json:"pagsiz,omitempty"`
	// 续传键值 String 首次查询为空，当返回报文里面有cntkey返回时表明需要续传，请将返回报文中的cntkey值放入该字段继续请求。
	Cntkey string `json:"cntkey,omitempty"`
}

// PayrollNewDCAGPPDFResponseBody 13.代发明细对账单查询请求DCAGPPDF response body.
type PayrollNewDCAGPPDFResponseBody struct {
	// 续传键值 String 当该字段有返回时表明需要续传，请将该值放入下次请求继续查询。
	Cntkey string `json:"cntkey,omitempty"`
	// 查询标记 string(10) 续传时请将该值放入下次请求继续查询。
	Begidx string `json:"begidx,omitempty"`
	// 打印任务编号 string(100) 返回的任务ID，请根据该ID查询PDF文件。
	Printid string `json:"printid,omitempty"`
	// 总笔数 string(10) 满足条件的总数
	Total string `json:"total,omitempty"`
}

type PayrollNewDCAGPPDFResponse struct {
	Response struct {
		Body PayrollNewDCAGPPDFResponseBody `json:"body,omitempty"`
		Head HeadResponse                   `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

// PayrollNewDCTASKIDBody 14.代发明细对账单处理结果查询请求DCTASKID request body.
type PayrollNewDCTASKIDBody struct {
	// 查询ID string (100) required:Y 查询请求DCAGPPDF时返回报文中的printid字段
	Taskid string `json:"taskid,omitempty"`
	// 文件地址类型 String true或者空：生成的url链接为互联网可访问 false：生成的url链接为专线访问地址，（请专线接入客户该值传false）
	Qwenab string `json:"qwenab,omitempty"`
}

// PayrollNewDCTASKIDResponseBody 14.代发明细对账单处理结果查询请求DCTASKID response body.
type PayrollNewDCTASKIDResponseBody struct {
	// 文件获取地址 string(300) 生成完成后，代发明细对账单获取地址链接，请注意该链接有效期1天，获取到后请立即下载保存，以防过期失效。
	Fileurl string `json:"fileurl,omitempty"`
	// 文件生成时间 string(20) 后台PDF文件生成完成时间
	Fintim string `json:"fintim,omitempty"`
}

type PayrollNewDCTASKIDResponse struct {
	Response struct {
		Body PayrollNewDCTASKIDResponseBody `json:"body,omitempty"`
		Head HeadResponse                   `json:"head,omitempty"`
	} `json:"response,omitempty"`
}

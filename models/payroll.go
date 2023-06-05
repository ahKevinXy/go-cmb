package models

type UnitPayrollPaymentRequest struct {
	Request   UnitPayrollPaymentData `json:"request"`
	Signature Signature              `json:"signature"`
}

type UnitPayrollPaymentData struct {
	Body UnitPayrollPaymentBody `json:"body,omitempty"`
	Head Head                   `json:"head"`
}
type UnitPayrollPaymentBody struct {
	Bb6cdcbhx1 []*Bb6cdcbhx1 `json:"bb6cdcbhx1,omitempty"`
	Bb6cdcdlx1 []*Bb6cdcdlx1 `json:"bb6cdcdlx1,omitempty"`
	Bb6Busmod  []*Bb6Busmod  `json:"bb6busmod,omitempty"`
}

// Bb6cdcbhx1
//  @Description: 汇总接口
//  @Author  ahKevinXy
//  @Date  2023-06-05 10:20:20
type Bb6cdcbhx1 struct {
	Bbknbr string `json:"bbknbr,omitempty"` // 业务受理分行号
	Begtag string `json:"begtag,omitempty"` // 批次开始标志
	Endtag string `json:"endtag,omitempty"` // 批次结束标志
	Reqnbr string `json:"reqnbr,omitempty"` // 流程实例号
	Accnbr string `json:"accnbr,omitempty"` // 账号
	Accnam string `json:"accnam,omitempty"` // 名称
	Waltid string `json:"waltid,omitempty"` // 账号
	Fctnam string `json:"fctnam,omitempty"` // 公司名称
	Ttlamt string `json:"ttlamt,omitempty"` // 总金额
	Ttlcnt string `json:"ttlcnt,omitempty"` // 总笔数
	Curamt string `json:"curamt,omitempty"` // 本次金额
	Curcnt string `json:"curcnt,omitempty"` // 本次笔数
	Cnvnbr string `json:"cnvnbr,omitempty"` // 合作方协议号和交易类型不能同时为空
	Ccynbr string `json:"ccynbr,omitempty"` // 币种  默认赋值 10 (人民币)
	Ntfinf string `json:"ntfinf,omitempty"` // 个性化短信
	Trstyp string `json:"trstyp,omitempty"` // 个性化短信
	Nusage string `json:"nusage,omitempty"` // 用途
	Eptdat string `json:"eptdat,omitempty"` // 期望日期 格式：YYYYMMDD
	Epttim string `json:"epttim,omitempty"` // 期望时间 HHMMSS,为空的话，默认值为000000
	Yurref string `json:"yurref,omitempty"` // 业务参考号
	Dmanbr string `json:"dmanbr,omitempty"` // 记账单元
	Chlflg string `json:"chlflg,omitempty"` // 结算通道  Y 超网  N 大小额  M 智能路由  A 数币代发 W 钱包代发
	Rsv100 string `json:"rsv100,omitempty"` //保留字

}

// Bb6cdcdlx1
//  @Description:  明细接口
//  @Author  ahKevinXy
//  @Date  2023-06-05 10:20:44
type Bb6cdcdlx1 struct {
	Trxseq string `json:"trxseq,omitempty"` // 交易序号
	Accnbr string `json:"accnbr,omitempty"` // 账号
	Accnam string `json:"accnam,omitempty"` // 姓名
	Digflg string `json:"digflg,omitempty"` // 货币标志 N 收方是账号  Y 收方是钱包
	Waltid string `json:"waltid,omitempty"` // 钱包ID 收方是钱包
	Walnam string `json:"walnam,omitempty"` //  钱包昵称
	Trsamt string `json:"trsamt,omitempty"` //  交易金额
	Trsdsp string `json:"trsdsp,omitempty"` //   注释
	Bnkflg string `json:"bnkflg,omitempty"` //   系统内标志
	Eacbnk string `json:"eacbnk,omitempty"` //   他行户口开户行
	Eaccty string `json:"eaccty,omitempty"` //   他行户口开户行地
	Cprref string `json:"cprref,omitempty"` //   合作方流水号
	Rsv45z string `json:"rsv45z,omitempty"` //   保留字

}
type Bb6Aclaky1 struct {
	Trxseq string `json:"trxseq,omitempty"` // 交易序号
	Accnbr string `json:"accnbr,omitempty"` // 账号
	Accnam string `json:"accnam,omitempty"` // 名称
	Trsamt string `json:"trsamt,omitempty"` // 金额
	Trsdsp string `json:"trsdsp,omitempty"` // 备注
	Bnkflg string `json:"bnkflg,omitempty"` // 系统标志 (可不传 自动识别)
}

type Bb6Aclakx1 struct {
	Begtag string `json:"begtag,omitempty"` // 批次开始标志 首次传输 y  建议 1000内 begtag endtag 都设置y
	Endtag string `json:"endtag,omitempty"` //批次结束标志
	Ttlamt string `json:"ttlamt,omitempty"` // 总金额
	Ttlcnt string `json:"ttlcnt,omitempty"` // 总笔数
	Ttlnum string `json:"ttlnum,omitempty"` // 总次数
	Curamt string `json:"curamt,omitempty"` // 本次金额
	Curcnt string `json:"curcnt,omitempty"` // 本次笔数
	Ccynbr string `json:"ccynbr,omitempty"` // 币种 默认 10 人民币
	Accnbr string `json:"accnbr,omitempty"` // 账号
	Trstyp string `json:"trstyp,omitempty"` // 交易类型
	Nusage string `json:"nusage,omitempty"` // 摘要
	Yurref string `json:"yurref,omitempty"` //业务参考号
	Eptdat string `json:"eptdat,omitempty"` // 期望时间
	Epttim string `json:"epttim,omitempty"` // 期望时间
	Dmanbr string `json:"dmanbr,omitempty"` // 记账单元
	Chlflg string `json:"chlflg,omitempty"` // 结算通道 Y 超网 N 大小额 M 智能路由  A 数币代发 W 签包代发
}

type Bb6Busmod struct {
	Buscod string `json:"buscod,omitempty"` //业务类型 N03010 代发工资 N03020 代发其他 N03040 钱包代发
	Busmod string `json:"busmod,omitempty"` // 业务模式 通过接口获取
	Rsv50z string `json:"rsv50z,omitempty"` // 业务模式 通过接口获取
}

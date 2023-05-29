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
	Bb6Aclaky1 []*Bb6Aclaky1 `json:"bb6aclaky1,omitempty"`
	Bb6Aclakx1 []*Bb6Aclakx1 `json:"bb6aclakx1,omitempty"`
	Bb6Busmod  []*Bb6Busmod  `json:"bb6busmod,omitempty"`
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
}

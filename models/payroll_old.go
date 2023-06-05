package models

// QueryPayrollOldTransCodeRequest   获取交易代码
type QueryPayrollOldTransCodeRequest struct {
	Request   QueryPayrollOldTransCodeData `json:"request"`
	Signature Signature                    `json:"signature"`
}

type QueryPayrollOldTransCodeData struct {
	Body QueryPayrollOldTransCodeBody `json:"body,omitempty"`
	Head Head                         `json:"head"`
}

type QueryPayrollOldTransCodeBody struct {
	Ntagtls2x []Ntagtls2x `json:"ntagtls2x"`
}

type Ntagtls2x struct {
	Accnbr string `json:"accnbr"`
	Buscod string `json:"buscod"`

	//Ccynbr string `json:"ccynbr"`
}

// PayrollOldCreditOtherBySupRequest   超网代发其他
type PayrollOldCreditOtherBySupRequest struct {
	Request   PayrollOldCreditOtherBySupData `json:"request"`
	Signature Signature                      `json:"signature"`
}

type PayrollOldCreditOtherBySupData struct {
	Body PayrollOldCreditOtherBySupBody `json:"body"`
	Head Head                           `json:"head"`
}
type PayrollOldCreditOtherBySupBody struct {
	Ntbusmody  []*Ntbusmody  `json:"ntbusmody"`
	Ntagcsaix1 []*Ntagcsaix1 `json:"ntagcsaix1"`
	Ntagcsaix2 []*Ntagcsaix2 `json:"ntagcsaix2"`
}

//type Ntbusmody struct {
//	Busmod string `json:"busmod,omitempty"`
//}

type Ntagcsaix1 struct {
	Bbknbr string `json:"bbknbr,omitempty"`
	Begtag string `json:"begtag,omitempty"`
	Bustya string `json:"bustya,omitempty"`
	Curamt string `json:"curamt,omitempty"`
	Curcnt string `json:"curcnt,omitempty"`
	Endtag string `json:"endtag,omitempty"`
	Eptdat string `json:"eptdat,omitempty"`
	Epttim string `json:"epttim,omitempty"`
	Grtflg string `json:"grtflg,omitempty"`
	Jzbnbr string `json:"jzbnbr,omitempty"`
	Ntfinf string `json:"ntfinf,omitempty"`
	Payacc string `json:"payacc,omitempty"`
	Reqnbr string `json:"reqnbr,omitempty"`
	Trxrmk string `json:"trxrmk,omitempty"`
	Ttlamt string `json:"ttlamt,omitempty"`
	Ttlcnt string `json:"ttlcnt,omitempty"`
	Ttlnum string `json:"ttlnum,omitempty"`
	Yurref string `json:"yurref,omitempty"`
}
type Ntagcsaix2 struct {
	Cpract string `json:"cpract,omitempty"`
	Cprref string `json:"cprref,omitempty"`
	Eacnam string `json:"eacnam,omitempty"`
	Eacnbr string `json:"eacnbr,omitempty"`
	Rcvamt string `json:"rcvamt,omitempty"`
	Trxseq string `json:"trxseq,omitempty"`
	Trxtxt string `json:"trxtxt,omitempty"`
}

// QueryOldPayRollOrderRequest    概要查询
type QueryOldPayRollOrderRequest struct {
	Request   QueryOldPayRollOrderData `json:"request"`
	Signature Signature                `json:"signature"`
}

type QueryOldPayRollOrderData struct {
	Body Ntagcinny1Body `json:"body,omitempty"`
	Head Head           `json:"head"`
}

type Ntagcinny1Body struct {
	Ntagcinny1 []*Ntagcinny1 `json:"ntagcinny1,omitempty"`
}

type Ntagcinny1 struct {
	Busmod string `json:"busmod,omitempty"` //业务类型 N03010 代发工资 N03020 代发 N03030 代扣
	Buscod string `json:"buscod,omitempty"` //业务模式
	Bgndat string `json:"bgndat,omitempty"` //开始时间
	Enddat string `json:"enddat,omitempty"` //结束时间
	Datflg string `json:"datflg,omitempty"` //时间类型 A 经办日期 B 期望日期
	Ctnkey string `json:"ctnkey,omitempty"` // 续传接口

}

// PayrollOldPdfFileRequest   查询交易明细信息
type PayrollOldPdfFileRequest struct {
	Request   PayrollPdfFileData `json:"request"`
	Signature Signature          `json:"signature"`
}

type PayrollPdfFileData struct {
	Body PayrollPdfFileBody `json:"body,omitempty"`
	Head Head               `json:"head"`
}
type PayrollPdfFileBody struct {
	Taskid string `json:"taskid"`
}

// QueryOldPayRollTransDetailRecordRequest    明细查询
type QueryOldPayRollOrderDetailRequest struct {
	Request   QueryOldPayRollOrderDetailData `json:"request"`
	Signature Signature                      `json:"signature"`
}

type QueryOldPayRollOrderDetailData struct {
	Body Ntagdinfy1Body `json:"body,omitempty"`
	Head Head           `json:"head"`
}

type Ntagdinfy1Body struct {
	Ntagdinfy1 []*Ntagdinfy1 `json:"ntagdinfy1,omitempty"`
}

//type Ntagdinfy1 struct {
//	Reqnbr string `json:"reqnbr,omitempty"` //流程实例号
//	Bthnbr string `json:"bthnbr,omitempty"`
//	Trxseq string `json:"trxseq,omitempty"`
//	Histag string `json:"histag,omitempty"`
//}

type Ntagdinfy1 struct {
	Reqnbr string `json:"reqnbr,omitempty"` //流程实例号
	Bthnbr string `json:"bthnbr,omitempty"` // 续传标志
	Trxseq string `json:"trxseq,omitempty"` //  续传明细序号
	Histag string `json:"histag,omitempty"` //  续传使用
}

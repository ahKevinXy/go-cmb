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
	Trxseq string `json:"trxseq,omitempty"`
	Accnbr string `json:"accnbr,omitempty"`
	Accnam string `json:"accnam,omitempty"`
	Trsamt string `json:"trsamt,omitempty"`
	Trsdsp string `json:"trsdsp,omitempty"`
	Bnkflg string `json:"bnkflg,omitempty"`
}

type Bb6Aclakx1 struct {
	Begtag string `json:"begtag,omitempty"`
	Endtag string `json:"endtag,omitempty"`
	Ttlamt string `json:"ttlamt,omitempty"`
	Ttlcnt string `json:"ttlcnt,omitempty"`
	Ttlnum string `json:"ttlnum,omitempty"`
	Curamt string `json:"curamt,omitempty"`
	Curcnt string `json:"curcnt,omitempty"`
	Ccynbr string `json:"ccynbr,omitempty"`
	Accnbr string `json:"accnbr,omitempty"`
	Trstyp string `json:"trstyp,omitempty"`
	Nusage string `json:"nusage,omitempty"`
	Yurref string `json:"yurref,omitempty"`
	Eptdat string `json:"eptdat,omitempty"`
	Epttim string `json:"epttim,omitempty"`
}

type Bb6Busmod struct {
	Buscod string `json:"buscod,omitempty"`
	Busmod string `json:"busmod,omitempty"`
}

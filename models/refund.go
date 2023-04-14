package models

type QueryAccountPaymentRefundRequest struct {
	Request   QueryAccountPaymentRefundData `json:"request"`
	Signature Signature                     `json:"signature"`
}

type QueryAccountPaymentRefundData struct {
	Body QueryAccountPaymentRefundBody `json:"body,omitempty"`
	Head Head                          `json:"head"`
}
type QueryAccountPaymentRefundBody struct {
	Bb1payqby1 []*Bb1payqby1 `json:"bb1payqby1,omitempty"`
}

type Bb1payqby1 struct {
	BbkNbr string `json:"bbkNbr,omitempty"`
	AccNbr string `json:"accNbr,omitempty"`
	BgnDat string `json:"bgnDat,omitempty"`
	CtnKey string `json:"ctnKey,omitempty"`
	EndDat string `json:"endDat,omitempty"`
	ReqNbr string `json:"reqNbr,omitempty"`
	Rsv50z string `json:"rsv50z,omitempty"`
}

type QueryPayrollRefundRequest struct {
	Request   QueryPayrollRefundData `json:"request"`
	Signature Signature              `json:"signature"`
}

type QueryPayrollRefundData struct {
	Body QueryPayrollRefundBody `json:"body,omitempty"`
	Head Head                   `json:"head"`
}
type QueryPayrollRefundBody struct {
	Bb6rfdqyy1 []*Bb6rfdqyy1 `json:"bb6rfdqyy1,omitempty"`
}

type Bb6rfdqyy1 struct {
	AccNbr string `json:"accnbr,omitempty"`
	Trstyp string `json:"trstyp,omitempty"`
	BgnDat string `json:"bgndat,omitempty"`
	EndDat string `json:"enddat,omitempty"`
	Reqnbr string `json:"reqnbr,omitempty"`
	CtnKey string `json:"ctnKey,omitempty"`
	Rsv50z string `json:"rsv50z,omitempty"`
}

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

package models

// AccountUnitTransInRequest    内部调账
type AccountUnitTransInRequest struct {
	Request   AccountUnitTransInData `json:"request"`
	Signature Signature              `json:"signature"`
}

type AccountUnitTransInData struct {
	Body AccountUnitTransInDataBody `json:"body,omitempty"`
	Head Head                       `json:"head"`
}

type AccountUnitTransInDataBody struct {
	Ntdmatrxx1 []*Ntdmatrxx1 `json:"ntdmatrxx1,omitempty"`
}

type Ntdmatrxx1 struct {
	Accnbr string `json:"accnbr"`
	Dmacrt string `json:"dmacrt"`
	Dmadbt string `json:"dmadbt"`
	Trxamt string `json:"trxamt"`
	Trxtxt string `json:"trxtxt"`
}

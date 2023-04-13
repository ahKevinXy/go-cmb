package models

// QueryUnitTransDetailRequest    获取交易详情
type QueryUnitTransDetailRequest struct {
	Request   QueryUnitTransDetailData `json:"request"`
	Signature Signature                `json:"signature"`
}

type QueryUnitTransDetailData struct {
	Body Ntduminfx1Body `json:"body,omitempty"`
	Head Head           `json:"head"`
}

type Ntduminfx1Body struct {
	Ntduminfx1 []*Ntduminfx1 `json:"ntduminfx1,omitempty"`
}
type Ntduminfx1 struct {
	Reqnbr string `json:"reqnbr,omitempty"`
}

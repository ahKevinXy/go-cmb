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

// AccountUnitTransDailyRequest   获取当天交易信息
type AccountUnitTransDailyRequest struct {
	Request   AccountAddUnitTransDailyData `json:"request"`
	Signature Signature                    `json:"signature"`
}

type AccountAddUnitTransDailyData struct {
	Body NtdmtlstyBody `json:"body,omitempty"`
	Head Head          `json:"head"`
}
type NtdmtlstyBody struct {
	Ntdmtlsty []*Ntdmtlsty `json:"ntdmtlsty,omitempty"`
}

type Ntdmtlsty struct {
	Accnbr string `json:"accnbr"`
	Ctnkey string `json:"ctnkey"`
	Dmanbr string `json:"dmanbr"`
}

// AccountUnitTransHistoryRequest   获取历史
type AccountUnitTransHistoryRequest struct {
	Request   AccountUnitTransHistoryData `json:"request"`
	Signature Signature                   `json:"signature"`
}

type AccountUnitTransHistoryData struct {
	Body NtdmthlsyBody `json:"body,omitempty"`
	Head Head          `json:"head"`
}
type NtdmthlsyBody struct {
	Ntdmthlsy []*Ntdmthlsy `json:"ntdmthlsy,omitempty"`
}

type Ntdmthlsy struct {
	Accnbr string `json:"accnbr"`
	Begdat string `json:"begdat"`
	Ctnkey string `json:"ctnkey"`
	Dmanbr string `json:"dmanbr"`
	Enddat string `json:"enddat"`
}

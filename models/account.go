package models

// AccountDetailsRequest   查询交易明细信息
type AccountDetailsRequest struct {
	Request   AccountDetailsData `json:"request"`
	Signature Signature          `json:"signature"`
}

type AccountDetailsData struct {
	Body NtqacinfxBody `json:"body,omitempty"`
	Head Head          `json:"head"`
}

type NtqacinfxBody struct {
	Ntqacinfx []*Ntqacinfx `json:"ntqacinfx"`
}

type Ntqacinfx struct {
	Accnbr string `json:"accnbr,omitempty"`
	Bbknbr string `json:"bbknbr,omitempty"`
}

// Signature
//  @Description:  签名
//  @Author  ahKevinXy
//  @Date2023-04-06 19:29:12
type Signature struct {
	Sigdat string `json:"sigdat"`
	Sigtim string `json:"sigtim"`
}

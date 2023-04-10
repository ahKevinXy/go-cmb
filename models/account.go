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

// MainAccountHistoryBalanceRequest   查询概要信息
type MainAccountHistoryBalanceRequest struct {
	Request   MainAccountHistoryBalanceData `json:"request"`
	Signature Signature                     `json:"signature"`
}

type MainAccountHistoryBalanceData struct {
	Body NtqabinfyBody `json:"body,omitempty"`
	Head Head          `json:"head"`
}

type NtqabinfyBody struct {
	Ntqabinfy []*Ntqabinfy `json:"ntqabinfy,omitempty"`
}

type Ntqabinfy struct {
	Accnbr string `json:"accnbr"`
	Bbknbr string `json:"bbknbr"`
	Bgndat string `json:"bgndat"`
	Enddat string `json:"enddat"`
}

// AccountBankInfoRequest   获取银联号请求 todo
type AccountBankInfoRequest struct {
	Request   AccountBankInfoData `json:"request"`
	Signature Signature           `json:"signature"`
}

type AccountBankInfoData struct {
	Body AccountBankInfoBody `json:"body,omitempty"`
	Head Head                `json:"head"`
}
type AccountBankInfoBody struct {
	Fctval string `json:"fctval,omitempty"`
}

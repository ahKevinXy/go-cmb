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

// QueryBatchMainAccountBalanceRequest   查询概要信息
type QueryBatchMainAccountBalanceRequest struct {
	Request   QueryBatchMainAccountBalanceData `json:"request"`
	Signature Signature                        `json:"signature"`
}

type QueryBatchMainAccountBalanceData struct {
	Body NtqadinfxBody `json:"body,omitempty"`
	Head Head          `json:"head"`
}

type NtqadinfxBody struct {
	Ntqadinfx []*Ntqadinfx `json:"ntqadinfx,omitempty"`
}

type Ntqadinfx struct {
	Accnbr string `json:"accnbr"` //账号

	Bbknbr string `json:"bbknbr"` //分行号

}

// MainAccountUsersRequest    获取账户列表
type MainAccountUsersRequest struct {
	Request   MainAccountUsersData `json:"request"`
	Signature Signature            `json:"signature"`
}

type MainAccountUsersData struct {
	Body MainAccountUsersBody `json:"body,omitempty"`
	Head Head                 `json:"head"`
}

type MainAccountUsersBody struct {
	Buscode string `json:"buscod,omitempty"` //业务代码
	Busmod  string `json:"busmod,omitempty"` //业务模式
}

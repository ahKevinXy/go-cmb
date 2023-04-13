package models

// GetMainAccountTransInfoRequest
//  @Description:   获取交易流水
//  @Author  ahKevinXy
//  @Date2023-04-10 14:03:20
type GetMainAccountTransInfoRequest struct {
	Request   GetMainAccountTransInfoData `json:"request"`
	Signature Signature                   `json:"signature"`
}

type GetMainAccountTransInfoData struct {
	Body GetMainAccountTransInfoBody `json:"body,omitempty"`
	Head Head                        `json:"head"`
}
type GetMainAccountTransInfoBody struct {
	Bbknbr string `json:"bbknbr,omitempty"`
	Accnbr string `json:"accnbr,omitempty"`
	Trsdat string `json:"trsdat,omitempty"`
	Trsseq string `json:"trsseq,omitempty"`
}

// QueryAccountTransInfoRequest
//  @Description:   获取交易流水
//  @Author  ahKevinXy
//  @Date2023-04-10 14:03:20
type QueryAccountTransInfoRequest struct {
	Request   QueryAccountTransInfoData `json:"request"`
	Signature Signature                 `json:"signature"`
}

type QueryAccountTransInfoData struct {
	Body QueryAccountTransInfoBody `json:"body,omitempty"`
	Head Head                      `json:"head"`
}
type QueryAccountTransInfoBody struct {
	Sdktsinfx []*Sdktsinfx `json:"sdktsinfx,omitempty"`
	Ntqacctny []*Ntqacctny `json:"ntqacctny,omitempty"`
}

type Sdktsinfx struct {
	Accnbr  string `json:"accnbr,omitempty"`
	Amtcdr  string `json:"amtcdr,omitempty"`
	Bbknbr  string `json:"bbknbr,omitempty"`
	Bgndat  string `json:"bgndat,omitempty"`
	CBbknbr string `json:"c_bbknbr,omitempty"`
	Enddat  string `json:"enddat,omitempty"`
	Hghamt  string `json:"hghamt,omitempty"`
	Lowamt  string `json:"lowamt,omitempty"`
}
type Ntqacctny struct {
	Ctndta string `json:"ctndta,omitempty"`
	Rsv50z string `json:"rsv50z,omitempty"`
}

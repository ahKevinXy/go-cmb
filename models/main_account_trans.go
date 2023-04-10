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

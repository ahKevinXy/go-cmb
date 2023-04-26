package models

// QueryAccountTransCodeRequest   获取交易代码
type QueryAccountTransCodeRequest struct {
	Request   QueryAccountTransCodeData `json:"request"`
	Signature Signature                 `json:"signature"`
}

type QueryAccountTransCodeData struct {
	Body QueryAccountTransCodeBody `json:"body,omitempty"`
	Head Head                      `json:"head"`
}

type QueryAccountTransCodeBody struct {
	Buscod string `json:"buscod,omitempty"` // 模式
}

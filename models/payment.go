package models

// MainAccountPayResultRequest   7. 账户交易信息查询DCTRSINF
type MainAccountPayResultRequest struct {
	Request   MainAccountPayResultData `json:"request"`
	Signature Signature                `json:"signature"`
}

type MainAccountPayResultData struct {
	Body MainAccountPayResultBody `json:"body,omitempty"`
	Head Head                     `json:"head"`
}
type MainAccountPayResultBody struct {
	Bb1qrybtx1 []*Bb1qrybtx1 `json:"bb1qrybtx1"`
}
type Bb1qrybtx1 struct {
	BegDat string `json:"begDat,omitempty"`
	EndDat string `json:"endDat,omitempty"`
	AutStr string `json:"autStr,omitempty"`
	RtnStr string `json:"rtnStr,omitempty"`
}

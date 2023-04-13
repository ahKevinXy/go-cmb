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

type QueryAccountPaymentTransInfoRequest struct {
	Request   MainAccountPayResultData `json:"request"`
	Signature Signature                `json:"signature"`
}

type QueryAccountPaymentTransInfoData struct {
	Body MainAccountPayResultBody `json:"body,omitempty"`
	Head Head                     `json:"head"`
}
type QueryAccountPaymentTransInfoBody struct {
	Bb1qrybtx1 []*Bb1qrybtx1 `json:"bb1qrybtx1"`
}

type Bb1qrybtx1 struct {
	BegDat string `json:"begDat,omitempty"`
	EndDat string `json:"endDat,omitempty"`
	AutStr string `json:"autStr,omitempty"`
	RtnStr string `json:"rtnStr,omitempty"`
}

//type Bb1qrybtx1 struct {
//	BegDat string `json:"begDat,omitempty"`
//	EndDat string `json:"endDat,omitempty"`
//	AutStr string `json:"autStr,omitempty"`
//	RtnStr string `json:"rtnStr,omitempty"`
//}

// QueryAccountPaymentDetailRequest
//  @Description:  获取交易明细
//  @Author  ahKevinXy
//  @Date2023-04-13 16:48:34
type QueryAccountPaymentDetailRequest struct {
	Request   QueryAccountPaymentDetailData `json:"request"`
	Signature Signature                     `json:"signature"`
}

type QueryAccountPaymentDetailData struct {
	Body QueryAccountPaymentDetailBody `json:"body,omitempty"`
	Head Head                          `json:"head"`
}
type QueryAccountPaymentDetailBody struct {
	Bb1qrybdy1 []*Bb1qrybdy1 `json:"bb1qrybdy1,omitempty"`
}

type Bb1qrybdy1 struct {
	BthNbr string `json:"bthNbr,omitempty"`
	CtnKey string `json:"ctnKey,omitempty"`
	AutStr string `json:"autStr,omitempty"`
	RtnStr string `json:"rtnStr,omitempty"`
}

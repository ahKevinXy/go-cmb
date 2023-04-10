package models

// MainAccountSinglePayRequest   7. 账户交易信息查询DCTRSINF
type MainAccountSinglePayRequest struct {
	Request   MainAccountSinglePayData `json:"request"`
	Signature Signature                `json:"signature"`
}

type MainAccountSinglePayData struct {
	Body MainAccountSinglePayBody `json:"body,omitempty"`
	Head Head                     `json:"head"`
}
type MainAccountSinglePayBody struct {
	Bb1Paybmx1 []*Bb1Paybmx1 `json:"bb1paybmx1"`
	Bb1Payopx1 []*Bb1Payopx1 `json:"bb1payopx1,omitempty"`
}
type Bb1Paybmx1 struct {
	BusCod string `json:"busCod,omitempty"`
	BusMod string `json:"busMod,omitempty"`
}
type Bb1Payopx1 struct {
	CcyNbr string `json:"ccyNbr,omitempty"`
	CrtAcc string `json:"crtAcc,omitempty"`
	CrtNam string `json:"crtNam,omitempty"`
	DbtAcc string `json:"dbtAcc,omitempty"`
	NusAge string `json:"nusAge,omitempty"`
	BnkFlg string `json:"bnkFlg,omitempty"`
	TrsAmt string `json:"trsAmt,omitempty"`
	YurRef string `json:"yurRef,omitempty"`
	RcvChk string `json:"rcvChk,omitempty"`
	CrtBnk string `json:"crtBnk,omitempty"`
	CrtAdr string `json:"crtAdr,omitempty"`
	BrdNbr string `json:"brdNbr,omitempty"`
	DmaNbr string `json:"dmaNbr,omitempty"`
	TrsTyp string `json:"trsTyp,omitempty"`
	BusNar string `json:"busNar,omitempty"`
}

// MainAccountBatchPayRequest
//  @Description:  批量支付
//  @Author  ahKevinXy
//  @Date2023-04-10 13:59:44
type MainAccountBatchPayRequest struct {
	Request   MainAccountPayData `json:"request"`
	Signature Signature          `json:"signature"`
}

type MainAccountPayData struct {
	Body MainAccountPayBody `json:"body,omitempty"`
	Head Head               `json:"head"`
}

// MainAccountPayBody
//  @Description:   批量支付
//  @Author  ahKevinXy
//  @Date2023-04-10 13:59:18
type MainAccountPayBody struct {
	Bb1bmdbhx1 []*Bb1bmdbhx1 `json:"bb1bmdbhx1"`
	Bb1paybhx1 []*Bb1paybhx1 `json:"bb1paybhx1,omitempty"`
}
type Bb1bmdbhx1 struct {
	BusMod string `json:"busMod,omitempty"`
	BusCod string `json:"busCod,omitempty"`
	BthNbr string `json:"bthNbr,omitempty"`
	DtlNbr string `json:"dtlNbr,omitempty"`
	CtnFlg string `json:"ctnFlg,omitempty"`
	CtnSts string `json:"ctnSts,omitempty"`
}

type Bb1paybhx1 struct {
	DbtAcc string `json:"dbtAcc,omitempty"`
	DmaNbr string `json:"dmaNbr,omitempty"`
	CrtAcc string `json:"crtAcc,omitempty"`
	CrtNam string `json:"crtNam,omitempty"`
	CrtBnk string `json:"crtBnk,omitempty"`
	CrtAdr string `json:"crtAdr,omitempty"`
	BrdNbr string `json:"brdNbr,omitempty"`
	CcyNbr string `json:"ccyNbr,omitempty"`
	TrsAmt string `json:"trsAmt,omitempty"`
	BnkFlg string `json:"bnkFlg,omitempty"`
	StlChn string `json:"stlChn,omitempty"`
	NusAge string `json:"nusAge,omitempty"`
	YurRef string `json:"yurRef,omitempty"`
	BusNar string `json:"busNar,omitempty"`
	TrsTyp string `json:"trsTyp,omitempty"`
	RcvChk string `json:"rcvChk,omitempty"`
	TrxAmt string `json:"trxAmt,omitempty"`
}

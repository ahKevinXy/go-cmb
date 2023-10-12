package models

type SetUnitAccountRelationRequest struct {
	Request   SetUnitAccountRelationData `json:"request"`
	Signature Signature                  `json:"signature"`
}

type SetUnitAccountRelationData struct {
	Body Ntdmarltx1Body `json:"body,omitempty"`
	Head Head           `json:"head"`
}

type Ntdmarltx1Body struct {
	Ntbusmody  []*Ntbusmody  `json:"ntbusmody,omitempty"`
	Ntdmarltx1 []*Ntdmarltx1 `json:"ntdmarltx1,omitempty"`
}

type Ntdmarltx1 struct {
	Bbknbr string `json:"bbknbr"`
	Accnbr string `json:"accnbr"`
	Dmanbr string `json:"dmanbr"`
	Tlyopr string `json:"tlyopr"` // Y：非关联收款入账付方指定子单元 N：非关联收款入账默认子单元 R：非关联收款拒绝入账
	Dbtacc string `json:"dbtacc"`
	Yurref string `json:"yurref"` // string(30) 己方提供的索引号，需要保证唯一
}

// DelUnitAccountRelationRequest
//
//	@Description:  删除 关联账号
//	@Author  ahKevinXy
//	@Date  2023-10-12 12:28:18
type DelUnitAccountRelationRequest struct {
	Request   DelUnitAccountRelationData `json:"request"`
	Signature Signature                  `json:"signature"`
}

type DelUnitAccountRelationData struct {
	Body Ntdmarldx1Body `json:"body,omitempty"`
	Head Head           `json:"head"`
}

type Ntdmarldx1Body struct {
	Ntdmarldx1 []*Ntdmarldx1 `json:"ntdmarldx1,omitempty"`
}

type Ntdmarldx1 struct {
	Bbknbr string `json:"bbknbr,omitempty"`
	Accnbr string `json:"accnbr,omitempty"`
	Dmanbr string `json:"dmanbr,omitempty"`
	Dbtacc string `json:"dbtacc,omitempty"`
}

// UpdateUnitAccountRelationRequest
//
//	@Description:  更新记账单元模式
//	@Author  ahKevinXy
//	@Date  2023-10-12 12:45:56
type UpdateUnitAccountRelationRequest struct {
	Request   UpdateUnitAccountRelationData `json:"request"`
	Signature Signature                     `json:"signature"`
}

type UpdateUnitAccountRelationData struct {
	Body Ntdmatmnx1Body `json:"body,omitempty"`
	Head Head           `json:"head"`
}

type Ntdmatmnx1Body struct {
	Ntdmatmnx1 []*Ntdmatmnx1 `json:"ntdmatmnx1,omitempty"`
}

type Ntdmatmnx1 struct {
	Bbknbr string `json:"bbknbr,omitempty"`
	Accnbr string `json:"accnbr,omitempty"`
	Dmanbr string `json:"dmanbr,omitempty"`
	Tlyopr string `json:"tlyopr,omitempty"`
}

type QueryUnitAccountRelationRequest struct {
	Request   QueryUnitAccountRelationData `json:"request"`
	Signature Signature                    `json:"signature"`
}

type QueryUnitAccountRelationData struct {
	Body Ntdmarlqy1Body `json:"body,omitempty"`
	Head Head           `json:"head"`
}

type Ntdmarlqy1Body struct {
	Ntdmarlqy1 []*Ntdmarlqy1 `json:"ntdmarlqy1,omitempty"`
}

type Ntdmarlqy1 struct {
	Bbknbr string `json:"bbknbr,omitempty"`
	Accnbr string `json:"accnbr,omitempty"`
	Dmanbr string `json:"dmanbr,omitempty"`
	Cntkey string `json:"cntkey,omitempty"`
}

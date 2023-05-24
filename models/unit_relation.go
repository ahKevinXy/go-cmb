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

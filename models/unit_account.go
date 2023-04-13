package models

// AccountAddUnitRequest 添加子单元
type AccountAddUnitRequest struct {
	Request   AccountAddUnitData `json:"request"`
	Signature Signature          `json:"signature"`
}

type AccountAddUnitData struct {
	Body NtdmaaddxBody `json:"body,omitempty"`
	Head Head          `json:"head"`
}

type NtdmaaddxBody struct {
	Ntdmaaddx []*Ntdmaaddx `json:"ntdmaaddx,omitempty"`
}

type Ntdmaaddx struct {
	Accnbr string `json:"accnbr"`
	//Bcktyp string `json:"bcktyp"`
	//Clstyp string `json:"clstyp"`
	Dmanam string `json:"dmanam"`
	Dmanbr string `json:"dmanbr"`
	//Ovrctl string `json:"ovrctl"`
}

// AccountCloseUnitRequest 关闭子单元
type AccountCloseUnitRequest struct {
	Request   AccountCloseUnitData `json:"request"`
	Signature Signature            `json:"signature"`
}

type AccountCloseUnitData struct {
	Body NtbusmodyBody `json:"body,omitempty"`
	Head Head          `json:"head"`
}

type NtbusmodyBody struct {
	Ntbusmody  []*Ntbusmody  `json:"ntdmaaddx,omitempty"`
	Ntdmadltx1 []*Ntdmadltx1 `json:"ntdmadltx1,omitempty"`
	Ntdmadltx2 []*Ntdmadltx2 `json:"ntdmadltx2,omitempty"`
}
type Ntdmadltx1 struct {
	Accnbr string `json:"accnbr"`
}
type Ntdmadltx2 struct {
	Dmanbr string `json:"dmanbr"`
}
type Ntbusmody struct {
	Busmod string `json:"busmod,omitempty"`
}

// AccountUnitInfoRequest    获取金额
type AccountUnitInfoRequest struct {
	Request   AccountUnitBalanceData `json:"request"`
	Signature Signature              `json:"signature"`
}

type AccountUnitBalanceData struct {
	Body NtdmabalxBody `json:"body,omitempty"`
	Head Head          `json:"head"`
}

type NtdmabalxBody struct {
	Ntdmabalx []*Ntdmabalx `json:"ntdmabalx,omitempty"`
}
type Ntdmabalx struct {
	Accnbr string `json:"accnbr"`
	Dmanbr string `json:"dmanbr"`
}

//QueryUnitTransByBusNoRequest

// QueryUnitTransByBusNoRequest    按照交易流水获取
type QueryUnitTransByBusNoRequest struct {
	Request   QueryUnitTransByBusNoData `json:"request"`
	Signature Signature                 `json:"signature"`
}

type QueryUnitTransByBusNoData struct {
	Body Ntdumredx1Body `json:"body,omitempty"`
	Head Head           `json:"head"`
}

type Ntdumredx1Body struct {
	Ntdumredx1 []*Ntdumredx1 `json:"ntdumredx1,omitempty"`
}
type Ntdumredx1 struct {
	Yurref string `json:"yurref,omitempty"`
	Bgndat string `json:"bgndat,omitempty"`
	Enddat string `json:"enddat,omitempty"`
}

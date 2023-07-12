package models

// AccountAddUnitV1Request 添加子单元
type AccountAddUnitV1Request struct {
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
	Ntbusmody  []*Ntbusmody  `json:"ntbusmody,omitempty"`
	Ntdumdltx1 []*Ntdumdltx1 `json:"ntdumdltx1,omitempty"`
	Ntdumdltx2 []*Ntdumdltx2 `json:"ntdumdltx2,omitempty"`
}
type Ntdumdltx1 struct {
	Bbknbr string `json:"bbknbr,omitempty"`
	Inbacc string `json:"inbacc,omitempty"`
}
type Ntdumdltx2 struct {
	Dyanbr string `json:"dyanbr,omitempty"`
	Yurref string `json:"yurref,omitempty"`
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

// AccountCloseUnitV1Request 关闭子单元
type AccountCloseUnitV1Request struct {
	Request   AccountCloseUnitDataV1 `json:"request"`
	Signature Signature              `json:"signature"`
}
type AccountCloseUnitDataV1 struct {
	Body Ntdmadltx1Body `json:"body,omitempty"`
	Head Head           `json:"head"`
}

type Ntdmadltx1Body struct {
	Ntdmadltx1 []*Ntdmadltx1 `json:"ntdmadltx1,omitempty"`
	Ntdmadltx2 []*Ntdmadltx2 `json:"ntdmadltx2,omitempty"`
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

// QueryUnitAccountBalanceHistoryRequest    获取所有子单元账户历史余额
type QueryUnitAccountBalanceHistoryRequest struct {
	Request   QueryUnitAccountBalanceHistoryData `json:"request"`
	Signature Signature                          `json:"signature"`
}

type QueryUnitAccountBalanceHistoryData struct {
	Body Ntdmahbdz1Body `json:"body,omitempty"`
	Head Head           `json:"head"`
}

type Ntdmahbdz1Body struct {
	Ntdmahbdx1 []*Ntdmahbdx1 `json:"ntdmahbdx1,omitempty"`
}
type Ntdmahbdx1 struct {
	Bbknbr string `json:"bbknbr,omitempty"`
	Accnbr string `json:"accnbr,omitempty"`
	Qrydat string `json:"qrydat,omitempty"`
	Dmanbr string `json:"dmanbr,omitempty"`
}

// QueryUnitAccountSingleBalanceHistoryRequest    获取单个交易历史余额
type QueryUnitAccountSingleBalanceHistoryRequest struct {
	Request   QueryUnitAccountSingleBalanceHistoryData `json:"request"`
	Signature Signature                                `json:"signature"`
}

type QueryUnitAccountSingleBalanceHistoryData struct {
	Body Ntdmahadx1Body `json:"body,omitempty"`
	Head Head           `json:"head"`
}
type Ntdmahadx1Body struct {
	Ntdmahadx1 []*Ntdmahadx1 `json:"ntdmahadx1,omitempty"`
}

type Ntdmahadx1 struct {
	Bbknbr string `json:"bbknbr,omitempty"` // 分行号
	Accnbr string `json:"accnbr,omitempty"` //账户
	Dmanbr string `json:"dmanbr,omitempty"` // 记账单元
	Begdat string `json:"begdat,omitempty"` // 开始时间
	Enddat string `json:"enddat,omitempty"` // 结束时间
}

// UpdateUnitAccountV1Request    获取单个交易历史余额
type UpdateUnitAccountV1Request struct {
	Request   UpdateUnitAccountV1Data `json:"request"`
	Signature Signature               `json:"signature"`
}

type UpdateUnitAccountV1Data struct {
	Body UpdateUnitAccountV1Body `json:"body,omitempty"`
	Head Head                    `json:"head"`
}
type UpdateUnitAccountV1Body struct {
	Ntbusmody  []*Ntbusmody  `json:"ntbusmody,omitempty"`
	Ntdmamntx1 []*Ntdmamntx1 `json:"ntdmamntx1,omitempty"`
}

type Ntdmamntx1 struct {
	Bbknbr string `json:"bbknbr,omitempty"` //分行号
	Accnbr string `json:"accnbr,omitempty"` // 账号
	Dmanbr string `json:"dmanbr,omitempty"` // 记账单元编号
	Dmanam string `json:"dmanam,omitempty"` // 记账单元名称
	Ovrctl string `json:"ovrctl,omitempty"` // 额度控制标志 (y 允许透支 n 不允许透支)
	Clstyp string `json:"clstyp,omitempty"` // 余额为零是否关闭
	Yurref string `json:"yurref,omitempty"` // 交易编号 保证唯一性
	Ballmt string `json:"BALLMT,omitempty"` // 余额上限
	Lmtflg string `json:"LMTFLG,omitempty"` // 额度标志 n 不设置 y 设置
	Bcktyp string `json:"bcktyp,omitempty"` // 退款处理方式 Y 退回原记账单元  N 退回结算户
}

// QueryUnitAccountInfoV2Request    获取单个交易历史余额
type QueryUnitAccountInfoV2Request struct {
	Request   QueryUnitAccountInfoV2Data `json:"request"`
	Signature Signature                  `json:"signature"`
}

type QueryUnitAccountInfoV2Data struct {
	Body QueryUnitAccountInfoV2Body `json:"body,omitempty"`
	Head Head                       `json:"head"`
}
type QueryUnitAccountInfoV2Body struct {
	//Ntbusmody  []*Ntbusmody  `json:"ntbusmody,omitempty"`
	Ntdumqryy1 []*Ntdumqryy1 `json:"ntdumqryy1,omitempty"`
}

type Ntdumqryy1 struct {
	Bbknbr string `json:"bbknbr,omitempty"` //分行号
	Inbacc string `json:"inbacc,omitempty"` // 账号
	Danbeg string `json:"danbeg,omitempty"` // 记账单元开始(*)
	Danend string `json:"danend,omitempty"` // 记账单元结束(*)
	Ctnkey string `json:"ctnkey,omitempty"` // 额度控制标志 (y 允许透支 n 不允许透支)

}

type Ntdumrlax1 struct {
	Bbknbr string `json:"bbknbr,omitempty"` //分行号
	Accnbr string `json:"accnbr,omitempty"` // 账号
	Dumnbr string `json:"dumnbr,omitempty"` // 记账子单元编号
	Rltnam string `json:"rltnam,omitempty"` // 关联户名
	Rltacc string `json:"rltacc,omitempty"` // 关联账户
	Chktyp string `json:"chktyp,omitempty"` // 白名单校验类型 1:帐号 2:户名
	Rcvtyp string `json:"rcvtyp,omitempty"` // 入账方式 N：非关联收款入账默认子单元 R：非关联收款拒绝入账
	Yurref string `json:"yurref,omitempty"` // 批次号
}

type AccountSetWhitePayBody struct {
	Ntbusmody  []*Ntbusmody  `json:"ntbusmody,omitempty"`
	Ntdumrlax1 []*Ntdumrlax1 `json:"ntdumrlax1,omitempty"`
}

type AccountSetWhitePayData struct {
	Body AccountSetWhitePayBody `json:"body,omitempty"`
	Head Head                   `json:"head"`
}

// AccountSetWhitePayRequest    设置白名单
type AccountSetWhitePayRequest struct {
	Request   AccountSetWhitePayData `json:"request"`
	Signature Signature              `json:"signature"`
}

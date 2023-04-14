package models

// QueryBatchTransInfoRequest
//  @Description:   查询批次
//  @Author  ahKevinXy
//  @Date  2023-04-14 14:42:25
type QueryBatchTransInfoRequest struct {
	Request   QueryBatchTransInfoData `json:"request"`
	Signature Signature               `json:"signature"`
}

type QueryBatchTransInfoData struct {
	Body QueryBatchTransInfoBody `json:"body,omitempty"`
	Head Head                    `json:"head"`
}
type QueryBatchTransInfoBody struct {
	Bb6bpdqyy1 []*Bb6bpdqyy1 `json:"bb6bpdqyy1,omitempty"`
}

type Bb6bpdqyy1 struct {
	Buscod string `json:"buscod,omitempty"`
	Yurref string `json:"yurref,omitempty"`
	Bthnbr string `json:"bthnbr,omitempty"`
	Trxseq string `json:"trxseq,omitempty"`
	Histag string `json:"histag,omitempty"`
	Rsv50z string `json:"rsv50z,omitempty"`
}

// QueryBatchTransListRequest
//  @Description:  代发批次查询
//  @Author  ahKevinXy
//  @Date  2023-04-14 14:54:46
type QueryBatchTransListRequest struct {
	Request   QueryBatchTransListData `json:"request"`
	Signature Signature               `json:"signature"`
}

type QueryBatchTransListData struct {
	Body QueryBatchTransListBody `json:"body,omitempty"`
	Head Head                    `json:"head"`
}
type QueryBatchTransListBody struct {
	Bb6bthqyy1 []*Bb6bthqyy1 `json:"bb6bthqyy1,omitempty"`
}

type Bb6bthqyy1 struct {
	Buscod string `json:"buscod,omitempty"` //业务编码
	Yurref string `json:"yurref,omitempty"` //业务参考号
	Bgndat string `json:"bgndat,omitempty"` //起始日期
	Enddat string `json:"enddat,omitempty"` //结束日期
	Dattyp string `json:"dattyp,omitempty"` //日期类型
	Cntkey string `json:"cntkey,omitempty"` //续传键值
	Rsv50z string `json:"rsv50z,omitempty"` //保留字
}

type QueryPayrollTransDetailRequest struct {
	Request   QueryPayrollTransDetailData `json:"request"`
	Signature Signature                   `json:"signature"`
}

type QueryPayrollTransDetailData struct {
	Body QueryPayrollTransDetailBody `json:"body,omitempty"`
	Head Head                        `json:"head"`
}
type QueryPayrollTransDetailBody struct {
	Bb6dtlqyy1 []*Bb6dtlqyy1 `json:"bb6dtlqyy1,omitempty"`
}

type Bb6dtlqyy1 struct {
	Reqnbr string `json:"reqnbr,omitempty"` //流程实例号
	Bthnbr string `json:"bthnbr,omitempty"` //业务参考号
	Trxseq string `json:"trxseq,omitempty"` //起始日期
	Histag string `json:"histag,omitempty"` //结束日期
	Rsv50z string `json:"rsv50z,omitempty"` //保留字
}

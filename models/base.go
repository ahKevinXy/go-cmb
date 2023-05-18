package models

type ReqV1 struct {
	Request     RequestV1   `json:"request"`   // 请求参数
	SignatureV1 SignatureV1 `json:"signature"` // 签名
}
type SignatureV1 struct {
	Sigdat     string `json:"sigdat"` //
	Sigtim     string `json:"sigtim"`
	Paltsigdat string `json:"paltsigdat,omitempty"` //
}
type RequestV1 struct {
	Body interface{} `json:"body"`
	Head Head        `json:"head"`
}

type Head struct {
	Funcode string `json:"funcode"` //交易码
	Reqid   string `json:"reqid"`   //请求ID
	Userid  string `json:"userid"`  // 用户ID
}

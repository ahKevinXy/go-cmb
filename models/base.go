package models

type ReqV1 struct {
	Request     RequestV1   `json:"request"`
	SignatureV1 SignatureV1 `json:"signature"`
}
type SignatureV1 struct {
	Sigdat     string `json:"sigdat"`
	Sigtim     string `json:"sigtim"`
	Paltsigdat string `json:"paltsigdat"`
}
type RequestV1 struct {
	Body interface{} `json:"body"`
	Head Head        `json:"head"`
}

type Head struct {
	Funcode string `json:"funcode"`
	Reqid   string `json:"reqid"`
	Userid  string `json:"userid"`
}

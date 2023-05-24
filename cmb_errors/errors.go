package cmb_errors

import "errors"

var (
	SystemError        = errors.New("系统错误")
	JsonUnmarshalError = errors.New("json 序列化失败")

	DecodePrivateKeyError = errors.New("解析私钥失败")

	CreateSignError = errors.New("生成签名失败")
)

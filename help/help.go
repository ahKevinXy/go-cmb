package help

import (
	"bytes"
	"crypto/cipher"
	"encoding/json"
	"github.com/tjfoc/gmsm/sm4"
)

// GetJson 根据ascii获取jsonStr  利用map的排序方式 实现 ascII排序
func GetJson(reqAccount string) string {

	mp := make(map[string]interface{})
	err := json.Unmarshal([]byte(reqAccount), &mp)
	if err != nil {
		return ""
	}
	resJson, _ := json.Marshal(mp)
	return string(resJson)
}

func Sm4Encrypt(key, iv, plainText []byte) ([]byte, error) {
	block, err := sm4.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData := pkcs5Padding(plainText, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	cryted := make([]byte, len(origData))
	blockMode.CryptBlocks(cryted, origData)
	return cryted, nil
}

// pkcs5填充
func pkcs5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

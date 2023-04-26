package help

import (
	"bytes"
	"crypto/cipher"
	"encoding/json"
	"errors"
	"github.com/ahKevinXy/go-cmb/config"
	"github.com/tjfoc/gmsm/sm4"
	"os/exec"
	"strings"
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

// SignSM2 参数加密算法
func SignSM2(signContent, privateKey, userId string) (string, error) {
	//fmt.Println("-----请求加密文件 -----")

	if config.Settings.CmbPay.CmbSign != "local" {
		return SignSm2ByRequest(signContent, privateKey, userId)
	} else {
		return SignSm2ByCmd(signContent, privateKey, userId)
	}
}

func SignSm2ByCmd(signContent, privateKey, userId string) (string, error) {
	signContent = strings.ReplaceAll(signContent, `"`, `\"`)

	command := config.Settings.CmbPay.CmbJavaBin + "  -jar " + config.Settings.CmbPay.Sm2Jar + " \"" + signContent + "\" " + userId + " " + privateKey

	//fmt.Println("请求加密信息-----------", command)
	cmd := exec.Command("/bin/sh", "-c", command)
	buf, err := cmd.Output()
	if err != nil {
		return "", err
	}
	data := strings.ReplaceAll(string(buf), "\n", "")
	data = strings.ReplaceAll(data, " ", "")

	//fmt.Println("------------加密数据", data)
	return data, nil
}

func SignSm2ByRequest(signContent, privateKey, userId string) (string, error) {
	mp := make(map[string]interface{})
	//signContent = strings.ReplaceAll(signContent, `"`, `\"`)
	//signContent = strings.ReplaceAll(signContent, " ", "")
	mp["sign_content"] = signContent
	mp["private_key"] = privateKey
	mp["user_id"] = userId
	//fmt.Println("加密数据::::----", mp)  config.Settings.CmbPay
	//r, code, _ := MakeHttpRequest("POST", "http://localhost:8080/api/sign", mp, nil)
	r, code, _ := MakeHttpRequest("POST", config.Settings.CmbPay.CmbSignUrl, mp, nil)

	if code != 200 {
		return "", errors.New("sign error")
	}
	//fmt.Println(r, "")
	return r, nil
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

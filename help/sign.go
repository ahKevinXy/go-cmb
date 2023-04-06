package help

import (
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/ahKevinXy/go-cmb/config"
	"github.com/ahKevinXy/go-cmb/models"
	"github.com/tjfoc/gmsm/sm4"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func CmbSignRequest(reqStr string, funCode, uid, userKey, AESKey string) string {

	return SignatureDataSM(reqStr, funCode, uid, userKey, AESKey)

}

func SignatureDataSM(
	reqStr, funCode, uid, userKey, AESKey string) string {

	reqStr = GetJson(reqStr)
	var reqV1 models.ReqV1

	if err := json.Unmarshal([]byte(reqStr), &reqV1); err != nil {
		fmt.Println(err)
		return ""
	}

	reqStr = strings.ReplaceAll(reqStr, "\n", "")
	reqStr = strings.ReplaceAll(reqStr, "\r", "")
	reqStr = strings.ReplaceAll(reqStr, " ", "")

	//用户RSA
	reqSign, err := SignSM2(reqStr, userKey, uid)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	//平台方RSA
	reqSignSaas, err := SignSM2(reqSign, config.Settings.CmbPay.CmbSaasPrivateKey, uid)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	signatureV1 := models.SignatureV1{Sigtim: reqV1.SignatureV1.Sigtim, Sigdat: reqSign, Paltsigdat: reqSignSaas}
	reqV1.SignatureV1 = signatureV1
	reqV1Json, err := json.Marshal(reqV1)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	userId := uid + "000000"
	reqNewAccountAes, err := Sm4Encrypt([]byte(AESKey), []byte(userId), reqV1Json)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	u := url.Values{}

	u.Set("ALG", "SM")
	u.Set("DATA", base64.StdEncoding.EncodeToString(reqNewAccountAes))
	u.Set("INSPLAT", config.Settings.CmbPay.CmbSaasName)
	u.Set("UID", uid)
	u.Set("FUNCODE", funCode)

	u.Encode()

	resp, err := http.PostForm(config.Settings.CmbPay.CmbUrl, u)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)

	var dataStr string
	if !strings.Contains(string(respBody), "ErrMsg") {
		//dataByte, err := sm4.Sm4Cbc([]byte(AESKey), respBody, false)
		//if err != nil {
		//	fmt.Println(err)
		//	return ""
		//}
		//dataStr = string(dataByte)

		respBody64, err := base64.StdEncoding.DecodeString(string(respBody))
		dataByte, err := sm4Decrypt([]byte(AESKey), []byte(userId), respBody64)
		if err != nil {

			fmt.Println(string(reqV1Json), "++++错误信息++++", string(respBody))
			return ""
		}
		dataStr = string(dataByte)
	} else {
		dataStr = string(respBody)
	}

	return dataStr
}

func sm4Decrypt(key, iv, cipherText []byte) ([]byte, error) {
	block, err := sm4.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(cipherText))
	blockMode.CryptBlocks(origData, cipherText)
	origData = pkcs5UnPadding(origData)
	return origData, nil
}

func pkcs5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}

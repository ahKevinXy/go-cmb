package help

import (
	"bytes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"net/url"
	"strings"

	"github.com/ahKevinXy/go-cmb/config"
	"github.com/ahKevinXy/go-cmb/models"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/sm4"
)

// CmbSignRequest
//
//	@Description:  招商银行 统一请求
//	@param reqStr 请求参数
//	@param funCode 请求代码
//	@param uid   用户ID
//	@param userKey 用户秘钥
//	@param AESKey 用户对称秘钥
//	@return string 结果返回
//	@Author  ahKevinXy
//	@Date2023-04-10 13:41:37
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

	//用户签名

	key, _ := base64.StdEncoding.DecodeString(userKey)

	priv, err := FormatPri([]byte(key))
	if err != nil {
		fmt.Println("decode private key fail")
		return ""
	}

	sm2uid := uid + "0000000000000000"
	reqSign, err := SM3WithSM2Sign(priv, reqStr, []byte(sm2uid[0:16]))

	if err != nil {
		fmt.Println(err)
		return ""
	}

	//平台方签名

	var reqSignSaas = ""

	if len(config.Settings.CmbPay.CmbSaasPrivateKey) > 0 {

		saaskey, _ := base64.StdEncoding.DecodeString(config.Settings.CmbPay.CmbSaasPrivateKey)

		saasPriv, err := FormatPri([]byte(saaskey))
		if err != nil {
			fmt.Println("decode saas private key fail")
			return ""
		}
		reqSignSaas, err = SM3WithSM2Sign(saasPriv, reqSign, []byte(sm2uid[0:16]))
		if err != nil {
			fmt.Println(err)
			return ""
		}
		fmt.Println(reqSignSaas)

	}
	fmt.Println("用户签名", reqSign)
	fmt.Println("平台签名", reqSignSaas)
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

	fmt.Println("请求返回内容:", string(respBody))
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

// SM3WithSM2Sign SM3WithSM2签名 Hex ToUpper
func SM3WithSM2Sign(privateKey *sm2.PrivateKey, forSignStr string, uid []byte) (string, error) {

	r, s, err := sm2.Sm2Sign(privateKey, []byte(forSignStr), uid, rand.Reader)
	if err != nil {
		return "", err
	}

	rByte := r.Bytes()
	sByte := s.Bytes()
	if len(rByte) < 32 {
		rByte = append([]byte{0}, rByte...)
	}
	if len(sByte) < 32 {
		sByte = append([]byte{0}, sByte...)
	}
	var buffer bytes.Buffer
	buffer.Write(rByte)
	buffer.Write(sByte)

	return base64.StdEncoding.EncodeToString(buffer.Bytes()), nil

}

func FormatPri(priByte []byte) (*sm2.PrivateKey, error) {
	c := sm2.P256Sm2()
	k := new(big.Int).SetBytes(priByte)
	priv := new(sm2.PrivateKey)
	priv.PublicKey.Curve = c
	priv.D = k
	priv.PublicKey.X, priv.PublicKey.Y = c.ScalarBaseMult(k.Bytes())
	return priv, nil
}

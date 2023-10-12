package help

import (
	"bytes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/ahKevinXy/go-cmb/cmb_errors"
	"io"
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
//	@Description:  招商银行 统一请求 todo  返回错误信息
//	@param reqStr 请求参数
//	@param funCode 请求代码
//	@param uid   用户ID
//	@param userKey 用户秘钥
//	@param AESKey 用户对称秘钥
//	@return string 结果返回
//	@Author  ahKevinXy
//	@Date 2023-04-10 13:41:37
func CmbSignRequest(
	reqStr,
	funCode,
	uid,
	userKey,
	AESKey string) (string, error) {

	reqStr = GetJson(reqStr)
	var reqV1 models.ReqV1

	if err := json.Unmarshal([]byte(reqStr), &reqV1); err != nil {

		return "", cmb_errors.JsonUnmarshalError
	}

	reqStr = strings.ReplaceAll(reqStr, "\n", "")
	reqStr = strings.ReplaceAll(reqStr, "\r", "")
	reqStr = strings.ReplaceAll(reqStr, " ", "")

	//用户签名
	key, _ := base64.StdEncoding.DecodeString(userKey)

	priv, err := FormatPri([]byte(key))
	if err != nil {
		return "", cmb_errors.DecodePrivateKeyError
	}

	sm2uid := uid + "0000000000000000"
	reqSign, err := SM3WithSM2Sign(priv, reqStr, []byte(sm2uid[0:16]))

	if err != nil {
		return "", cmb_errors.CreateSignError
	}

	//平台方签名
	var reqSignSaas = ""

	if len(config.Settings.CmbPay.CmbSaasPrivateKey) > 0 {

		saasKey, _ := base64.StdEncoding.DecodeString(config.Settings.CmbPay.CmbSaasPrivateKey)

		saasPriv, err := FormatPri(saasKey)
		if err != nil {

			return "", err
		}
		reqSignSaas, err = SM3WithSM2Sign(saasPriv, reqSign, []byte(sm2uid[0:16]))
		if err != nil {

			return "", err
		}

	}
	signatureV1 := models.SignatureV1{Sigtim: reqV1.SignatureV1.Sigtim, Sigdat: reqSign, Paltsigdat: reqSignSaas}
	reqV1.SignatureV1 = signatureV1
	reqV1Json, err := json.Marshal(reqV1)

	if err != nil {
		return "", err
	}

	userId := uid + "000000"
	reqNewAccountAes, err := Sm4Encrypt([]byte(AESKey), []byte(userId), reqV1Json)
	if err != nil {

		return "", err
	}

	u := url.Values{}

	u.Set("ALG", "SM")
	u.Set("DATA", base64.StdEncoding.EncodeToString(reqNewAccountAes))
	u.Set("INSPLAT", config.Settings.CmbPay.CmbSaasName)
	u.Set("UID", uid)
	u.Set("FUNCODE", funCode)

	u.Encode()
	if config.Settings.IsDebug {
		fmt.Println("请求参数:" + "\n" + reqStr)
		// 用于测试报告打印参数
		fmt.Println("请求参数带签名:"+"\n", string(reqV1Json))
		fmt.Println("用户签名--->"+"\n", reqSign)
		fmt.Println("SaaS平台签名--->"+"\n", reqSignSaas)
	}
	resp, err := http.PostForm(config.Settings.CmbPay.CmbUrl, u)
	if err != nil {

		return "", err
	}
	defer resp.Body.Close()

	// 读取数据
	respBody, _ := io.ReadAll(resp.Body)
	// 返回数据处理
	var dataStr string
	if !strings.Contains(string(respBody), "ErrMsg") {
		respBody64, err := base64.StdEncoding.DecodeString(string(respBody))
		dataByte, err := sm4Decrypt([]byte(AESKey), []byte(userId), respBody64)
		if err != nil {

			return "", err
		}
		dataStr = string(dataByte)
	} else {
		dataStr = string(respBody)
	}
	if config.Settings.IsDebug {
		fmt.Println("返回结果加密结果:" + "\n" + string(respBody))
		fmt.Println("返回结果:" + "\n" + dataStr)
	}
	return dataStr, nil
}

// sm4Decrypt
//
//	@Description:  sm4解密
//	@param key
//	@param iv
//	@param cipherText
//	@return []byte
//	@return error
//	@Author  ahKevinXy
//	@Date  2023-10-11 16:25:15
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

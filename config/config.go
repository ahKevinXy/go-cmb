package config

import (
	"errors"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

// InitConfigByFilePath 初始化配置
func InitConfigByFilePath(path string) {
	if len(path) == 0 {
		panic(errors.New("configuration path is not provided"))
	}

	configPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	if _, err := toml.DecodeFile(configPath, Settings); err != nil {
		panic(err)
	}
}

// IniConfigByImport
//  @Description:  配置导入
//  @param cmbConfig
//  @Author  ahKevinXy
//  @Date2023-04-13 15:04:18
func IniConfigByImport(sassName, url, prKey, signDefault string) {
	Settings = &Config{
		CmbPay{
			CmbSaasName:       sassName,
			CmbUrl:            url,
			CmbSaasPrivateKey: prKey,
			CmbSigdatDefult:   signDefault,
		},
	}
}

// Settings 全局配置
var Settings *Config = &Config{}

type Config struct {
	CmbPay `toml:"cmb_pay"`
}

type CmbPay struct {
	CmbSaasName       string `toml:"cmb_saas_name"`        //Sass名称
	CmbUrl            string `toml:"cmb_url"`              //招商请求地址
	CmbSaasPrivateKey string `toml:"cmb_saas_private_key"` //私钥
	CmbSigdatDefult   string `toml:"cmb_sigdat_defult"`    //默认签名

}

package config

import (
	"errors"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

// InitConfig 初始化配置
func InitConfig(path string) {
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

// Settings 全局配置
var Settings *Config = &Config{}

type Config struct {
	CmbPay `toml:"cmb_pay"`
}

type CmbPay struct {
	CmbSaasName       string `toml:"cmb_saas_name"`
	CmbUrl            string `toml:"cmb_url"`
	CmbSaasPrivateKey string `toml:"cmb_saas_private_key"`
	CmbSigdatDefult   string `toml:"cmb_sigdat_defult"`

	CmbBusCode string `toml:"cmb_bus_code"`
	CmbBusNAME string `toml:"cmb_bus_name"`

	CmbAccountUrl string `toml:"cmb_account_url"`
	CmbSign       string `toml:"cmb_sign"`

	ZipLocalPath       string `toml:"zip_local_path"`
	FileLocalPath      string `toml:"file_local_path"`
	CmbPayIntervalTime string `toml:"cmb_pay_interval_time"`
	CmbPayUpper        string `toml:"cmb_pay_upper"`
	RefundTicket       int    `toml:"refund_ticket_query_interval"`
	Sm2Jar             string `toml:"sm2_jar"`
}

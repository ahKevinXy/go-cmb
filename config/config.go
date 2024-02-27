package config

// InitConfigByFilePath 初始化配置
//func InitConfigByFilePath(path string) {
//	if len(path) == 0 {
//		panic(errors.New("configuration path is not provided"))
//	}
//
//	configPath, err := filepath.Abs(path)
//	if err != nil {
//		panic(err)
//	}
//
//	if _, err := toml.DecodeFile(configPath, Settings); err != nil {
//		panic(err)
//	}
//}

// InitConfig
//
//	@Description:  初始化配置文件
//	@param cmbConfig
//	@Author  ahKevinXy
//	@Date2023-04-13 15:04:18
func InitConfig(sassName, url, prKey, signDefault string, isDebug bool) {
	Settings = &Config{
		CmbPay{
			CmbSaasName:       sassName,
			CmbUrl:            url,
			CmbSaasPrivateKey: prKey,
			CmbSigdatDefult:   signDefault,
			IsDebug:           isDebug,
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
	IsDebug           bool   `toml:"is_debug"`             // 是否开启debug 模式
}

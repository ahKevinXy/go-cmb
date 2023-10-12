package cmb

import (
	"github.com/ahKevinXy/go-cmb/v1/config"
	"github.com/ahKevinXy/go-cmb/v1/testdata"
)

// init
//
//	@Description:   初始化
//	@Author  ahKevinXy
//	@Date  2023-05-22 14:36:31
func init() {
	config.InitConfig(testdata.CmbSassName,
		testdata.CmbUrl, testdata.CmbSassPrivateKey,
		testdata.CmbSigdatDefault, testdata.IsDebug)
}

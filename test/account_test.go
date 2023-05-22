package cmb

import (
	"github.com/ahKevinXy/go-cmb/handler/account"
	"github.com/ahKevinXy/go-cmb/testdata"
	"testing"
)

// 获取账务信息测试完成
func TestMainAccountInfo(t *testing.T) {

	info, err := account.MainAccountInfo(testdata.UserId, testdata.AseKey, testdata.UserKey, testdata.Account, testdata.BankLinkNo)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(info)
}

// 获取账户历史余额完成
func TestMainAccountHistoryBalance(t *testing.T) {
	balance, err := account.MainAccountHistoryBalance(testdata.UserId, testdata.AseKey, testdata.UserKey, testdata.Account, testdata.BankLinkNo, "20220504", "20220512")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(balance)
}

// 获取银联号测试完成
func TestGetBankLinkNo(t *testing.T) {
	no, err := account.GetBankLinkNo(testdata.UserId, testdata.AseKey, testdata.UserKey, testdata.Account)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(no)
}

// 获取支付模式 测试完成
func TestPayMods(t *testing.T) {
	busList, err := account.PayMods(testdata.UserId, testdata.AseKey, testdata.UserKey, "N03020")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", busList)
}

// 批量获取余额 测试完成
func TestQueryBatchAccountBalance(t *testing.T) {
	balance, err := account.QueryBatchAccountBalance(testdata.UserId, testdata.AseKey, testdata.UserKey, testdata.Account, testdata.BankLinkNo)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", balance)
}

// 获取可经办用户列表 测试完成
func TestMainAccountUsers(t *testing.T) {
	users, err := account.MainAccountUsers(testdata.UserId, testdata.AseKey, testdata.UserKey, "N02030", "00001")
	if err != nil {
		t.Error(err)

		return
	}
	t.Logf("%+v", users)
}

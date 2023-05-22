package cmb

import (
	"github.com/ahKevinXy/go-cmb/handler/unit_manager"
	"github.com/ahKevinXy/go-cmb/testdata"
	"testing"
)

// 添加 记账单元 测试完成
func TestAddUnitAccount(t *testing.T) {

	account, err := unit_manager.AddUnitAccount(testdata.UserId, testdata.AseKey, testdata.UserKey, testdata.Account, "测试添加记账单元0001", "0001000111")
	if err != nil {
		t.Errorf("%+v", err)
		return

	}
	t.Logf("%+v", account)
}

// 关闭子单元 测试完成
func TestCloseUnitAccount(t *testing.T) {
	account, err := unit_manager.CloseUnitAccount(testdata.UserId, testdata.AseKey, testdata.UserKey, testdata.Account, testdata.BankLinkNo, "0001000111", "00001", "close_111")
	if err != nil {
		t.Errorf("%+v", err)
		return
	}
	t.Logf("%+v", account)
}

// 获取所有子单元余额 测试完成
func TestQueryUnitAccountInfo(t *testing.T) {
	info, err := unit_manager.QueryUnitAccountInfo(testdata.UserId, testdata.AseKey, testdata.UserKey, testdata.Account, "")
	if err != nil {
		t.Errorf("%+v", err)
		return
	}
	t.Logf("%+v", info)
}

// 记账单元历史余额 测试完成
func TestQueryUnitAccountBalanceHistory(t *testing.T) {
	history, err := unit_manager.QueryUnitAccountBalanceHistory(testdata.UserId, testdata.AseKey, testdata.UserKey, testdata.Account, testdata.BankLinkNo, "20230102", "")
	if err != nil {
		t.Errorf("%+v", err)
		return
	}
	t.Logf("%+v", history)
}

// 获取单个机子单元历史余额 测试完成
func TestQueryUnitAccountSingleBalanceHistory(t *testing.T) {

	history, err := unit_manager.QueryUnitAccountSingleBalanceHistory(testdata.UserId, testdata.AseKey, testdata.UserKey, testdata.Account, testdata.BankLinkNo, "0000000000", "20230102", "20230122")
	if err != nil {
		t.Errorf("%+v", err)
		return
	}
	t.Logf("%+v", history)
}

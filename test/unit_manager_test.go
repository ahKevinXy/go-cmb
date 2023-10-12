package cmb

import (
	"testing"

	"github.com/ahKevinXy/go-cmb/handler/unit_manager"
	"github.com/ahKevinXy/go-cmb/help"
	"github.com/ahKevinXy/go-cmb/testdata"
)

// 添加 记账单元 测试完成
func TestAddUnitAccountV(t *testing.T) {

	account, err := unit_manager.AddUnitAccountV1(testdata.UserId, testdata.AseKey, testdata.UserKey, testdata.Account, "结转账户", "6666666666")
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

// 设置记账子单元关联付款方信息
func TestSetUnitAccountRelation(t *testing.T) {
	setRes, err := unit_manager.SetUnitAccountRelation(
		testdata.UserId,
		testdata.AseKey,
		testdata.UserKey,
		"00001",
		testdata.BankLinkNo,
		testdata.Account,
		"0000000004",
		"R",
		"102100099996",
		help.GenYurref(),
	)
	if err != nil {
		t.Errorf("%+v", err)
		return
	}
	t.Logf("%+v", setRes)
}

func TestGetUnitAccountTransDailyList(t *testing.T) {
	dailyRes, err := unit_manager.GetUnitAccountTransList(
		testdata.UserId,
		testdata.AseKey,
		testdata.UserKey,
		testdata.Account,
		"",
		"",
	)
	if err != nil {
		t.Errorf("%+v", err)
		return
	}
	t.Logf("%+v", dailyRes)
}

func TestAccountSetWhitePay(t *testing.T) {
	dailyRes, err := unit_manager.AccountSetWhitePay(
		testdata.UserId,
		testdata.AseKey,
		testdata.UserKey,
		testdata.Account,
		"0001",
		"0000000001",
		"",
		"",
		"",
		"",
		"",
		"11111",
	)
	if err != nil {
		t.Errorf("%+v", err)
		return
	}
	t.Logf("%+v", dailyRes)
}

func TestDelUnitAccountRelation(t *testing.T) {
	dailyRes, err := unit_manager.DelUnitAccountRelation(
		testdata.UserId,
		testdata.AseKey,
		testdata.UserKey,
		testdata.BankLinkNo,
		testdata.Account,
		"0000000001",
		"102100099996",
	)
	if err != nil {
		t.Errorf("%+v", err)
		return
	}
	t.Logf("%+v", dailyRes)
}

func TestUpdateUnitAccountRelation(t *testing.T) {
	dailyRes, err := unit_manager.UpdateUnitAccountRelation(
		testdata.UserId,
		testdata.AseKey,
		testdata.UserKey,
		testdata.BankLinkNo,
		testdata.Account,
		"0000000001",
		"Y",
	)
	if err != nil {
		t.Errorf("%+v", err)
		return
	}
	t.Logf("%+v", dailyRes)
}

func TestQueryUnitAccountRelation(t *testing.T) {
	dailyRes, err := unit_manager.QueryUnitAccountRelation(
		testdata.UserId,
		testdata.AseKey,
		testdata.UserKey,
		testdata.BankLinkNo,
		testdata.Account,
		"",
		"",
	)
	if err != nil {
		t.Errorf("%+v", err)
		return
	}
	t.Logf("%+v", dailyRes)
}

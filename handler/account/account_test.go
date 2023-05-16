package account

import (
	"github.com/ahKevinXy/go-cmb/config"
	"github.com/ahKevinXy/go-cmb/models"
	"reflect"
	"testing"
)

func init() {
	// 初始化配置文件
	config.InitConfigByFilePath("../../configs/cmb_dev.conf")
}
func TestGetBankLinkNo(t *testing.T) {
	type args struct {
		userId         string
		asePrivateKey  string
		userPrivateKey string
		accnbr         string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.AccountBankInfoResponse
		wantErr bool
	}{
		// TODO: Add test cases.

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetBankLinkNo(tt.args.userId, tt.args.asePrivateKey, tt.args.userPrivateKey, tt.args.accnbr)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBankLinkNo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBankLinkNo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMainAccountPayBusList(t *testing.T) {
	type args struct {
		userId         string
		asePrivateKey  string
		userPrivateKey string
		buscode        string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.QueryAccountTransCodeResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetMainAccountPayBusList(tt.args.userId, tt.args.asePrivateKey, tt.args.userPrivateKey, tt.args.buscode)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMainAccountPayBusList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMainAccountPayBusList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMainAccountHistoryBalance(t *testing.T) {
	type args struct {
		userId         string
		asePrivateKey  string
		userPrivateKey string
		accnbr         string
		bbknbr         string
		bgndat         string
		enddat         string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MainAccountHistoryBalance(tt.args.userId, tt.args.asePrivateKey, tt.args.userPrivateKey, tt.args.accnbr, tt.args.bbknbr, tt.args.bgndat, tt.args.enddat); got != tt.want {
				t.Errorf("MainAccountHistoryBalance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMainAccountInfo(t *testing.T) {
	type args struct {
		userId         string
		asePrivateKey  string
		userPrivateKey string
		accnbr         string
		bbknbr         string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.AccountInfoResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MainAccountInfo(tt.args.userId, tt.args.asePrivateKey, tt.args.userPrivateKey, tt.args.accnbr, tt.args.bbknbr)
			if (err != nil) != tt.wantErr {
				t.Errorf("MainAccountInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MainAccountInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryBatchAccountBalance(t *testing.T) {
	type args struct {
		userId         string
		asePrivateKey  string
		userPrivateKey string
		accnbr         string
		buscode        string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.QueryBatchAccountBalanceResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := QueryBatchAccountBalance(tt.args.userId, tt.args.asePrivateKey, tt.args.userPrivateKey, tt.args.accnbr, tt.args.buscode)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryBatchAccountBalance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryBatchAccountBalance() got = %v, want %v", got, tt.want)
			}
		})
	}
}

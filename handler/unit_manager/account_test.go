package unit_manager

import (
	"github.com/ahKevinXy/go-cmb/config"
	"github.com/ahKevinXy/go-cmb/models"
	"github.com/ahKevinXy/go-cmb/testdata"
	"reflect"
	"testing"
)

func init() {

	config.InitConfigByFilePath("../../configs/cmb_dev.conf")
}

func TestAddUnitAccount(t *testing.T) {
	type args struct {
		userId         string
		asePrivateKey  string
		userPrivateKey string
		accnbr         string
		dmanam         string
		dmanbr         string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.AddUnitAccountResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AddUnitAccount(tt.args.userId, tt.args.asePrivateKey, tt.args.userPrivateKey, tt.args.accnbr, tt.args.dmanam, tt.args.dmanbr)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddUnitAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddUnitAccount() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCloseUnitAccount(t *testing.T) {
	type args struct {
		userId         string
		asePrivateKey  string
		userPrivateKey string
		accnbr         string
		dmanbr         string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.CloseUnitAccountResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CloseUnitAccount(tt.args.userId, tt.args.asePrivateKey, tt.args.userPrivateKey, tt.args.accnbr, tt.args.dmanbr)
			if (err != nil) != tt.wantErr {
				t.Errorf("CloseUnitAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CloseUnitAccount() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryUnitAccountBalanceHistory(t *testing.T) {
	type args struct {
		userId         string
		asePrivateKey  string
		userPrivateKey string
		accnbr         string
		bbknbr         string
		qrydat         string
		dmanbr         string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.QueryUnitAccountBalanceHistoryResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "获取测试模式",
			args: args{
				userId:         testdata.UserId,
				asePrivateKey:  testdata.AseKey,
				userPrivateKey: testdata.UserKey,
				accnbr:         testdata.Account,
				bbknbr:         "75",
				qrydat:         "20230112",
				dmanbr:         "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := QueryUnitAccountBalanceHistory(tt.args.userId, tt.args.asePrivateKey, tt.args.userPrivateKey, tt.args.accnbr, tt.args.bbknbr, tt.args.qrydat, tt.args.dmanbr)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryUnitAccountBalanceHistory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryUnitAccountBalanceHistory() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryUnitAccountInfo(t *testing.T) {
	type args struct {
		userId         string
		asePrivateKey  string
		userPrivateKey string
		accnbr         string
		dmanbr         string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.AccountUnitInfoResponse
		wantErr bool
	}{
		// TODO: Add test cases.

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := QueryUnitAccountInfo(tt.args.userId, tt.args.asePrivateKey, tt.args.userPrivateKey, tt.args.accnbr, tt.args.dmanbr)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryUnitAccountInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryUnitAccountInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

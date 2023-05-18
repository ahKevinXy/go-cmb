package account

import (
	"github.com/ahKevinXy/go-cmb/models"
	"github.com/ahKevinXy/go-cmb/testdata"
	"reflect"
	"testing"
)

// 获取银联号 测试完成
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
		{
			name: "获取测试模式",
			args: args{
				userId:         testdata.UserId,
				asePrivateKey:  testdata.AseKey,
				userPrivateKey: testdata.UserKey,
				accnbr:         testdata.Account,
			},
		},
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

//todo
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

// 批量获取账户历史余额
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
		name    string
		args    args
		want    *models.MainAccountHistoryBalanceResponse
		wantErr bool
	}{

		{
			name: "获取测试模式",
			args: args{
				userId:         testdata.UserId,
				asePrivateKey:  testdata.AseKey,
				userPrivateKey: testdata.UserKey,
				accnbr:         testdata.Account,
				bbknbr:         "75",
				bgndat:         "20230112",
				enddat:         "20230112",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MainAccountHistoryBalance(tt.args.userId, tt.args.asePrivateKey, tt.args.userPrivateKey, tt.args.accnbr, tt.args.bbknbr, tt.args.bgndat, tt.args.enddat)
			if (err != nil) != tt.wantErr {
				t.Errorf("MainAccountHistoryBalance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MainAccountHistoryBalance() got = %v, want %v", got, tt.want)
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

		{
			name: "获取测试模式",
			args: args{
				userId:         testdata.UserId,
				asePrivateKey:  testdata.AseKey,
				userPrivateKey: testdata.UserKey,
				accnbr:         testdata.Account,
				bbknbr:         "75",
			},
		},
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

// 批量获取
func TestQueryBatchAccountBalance(t *testing.T) {
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
		want    *models.QueryBatchAccountBalanceResponse
		wantErr bool
	}{
		{
			name: "获取测试模式",
			args: args{
				userId:         testdata.UserId,
				asePrivateKey:  testdata.AseKey,
				userPrivateKey: testdata.UserKey,
				accnbr:         testdata.Account,
				bbknbr:         "75",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := QueryBatchAccountBalance(tt.args.userId, tt.args.asePrivateKey, tt.args.userPrivateKey, tt.args.accnbr, tt.args.bbknbr)
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

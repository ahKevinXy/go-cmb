package account

import (
	"reflect"
	"testing"

	"github.com/ahKevinXy/go-cmb/config"
	"github.com/ahKevinXy/go-cmb/models"
	"github.com/ahKevinXy/go-cmb/testdata"
)

func init() {

	config.InitConfigByFilePath("../../configs/cmb_dev.conf")
}

// 测试获取支付模式
func TestPayMods(t *testing.T) {
	type args struct {
		userId         string
		sm4Key         string
		userPrivateKey string
		busCode        string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.QueryAccountTransCodeResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "获取测试模式",
			args: args{
				userId:         testdata.UserId,
				sm4Key:         testdata.AseKey,
				userPrivateKey: testdata.UserKey,
				busCode:        "N02030",
			},
			want: &models.QueryAccountTransCodeResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PayMods(tt.args.userId, tt.args.sm4Key, tt.args.userPrivateKey, tt.args.busCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("PayMods() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PayMods() got = %v, want %v", got, tt.want)
			}
		})
	}
}

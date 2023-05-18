package account

import (
	"github.com/ahKevinXy/go-cmb/models"
	"github.com/ahKevinXy/go-cmb/testdata"
	"reflect"
	"testing"
)

// 测试可经办业务
func TestMainAccountUsers(t *testing.T) {
	type args struct {
		userId         string
		asePrivateKey  string
		userPrivateKey string
		buscod         string
		busmod         string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.MainAccountUsersResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "获取测试模式",
			args: args{
				userId:         testdata.UserId,
				asePrivateKey:  testdata.AseKey,
				userPrivateKey: testdata.UserKey,
				buscod:         "N02030",
				busmod:         "00004",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MainAccountUsers(tt.args.userId, tt.args.asePrivateKey, tt.args.userPrivateKey, tt.args.buscod, tt.args.busmod)
			if (err != nil) != tt.wantErr {
				t.Errorf("MainAccountUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MainAccountUsers() got = %v, want %v", got, tt.want)
			}
		})
	}
}

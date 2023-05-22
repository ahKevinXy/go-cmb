package payroll_old

import (
	"reflect"
	"testing"

	"github.com/ahKevinXy/go-cmb/models"
)

func TestPayMods(t *testing.T) {
	type args struct {
		userId         string
		sm4Key         string
		userPrivateKey string
		busCode        string
		accnbr         string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.QueryPayrollOldTransCodeResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PayMods(tt.args.userId, tt.args.sm4Key, tt.args.userPrivateKey, tt.args.busCode, tt.args.accnbr)
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

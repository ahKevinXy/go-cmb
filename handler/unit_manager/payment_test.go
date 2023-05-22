package unit_manager

import (
	"reflect"
	"testing"

	"github.com/ahKevinXy/go-cmb/models"
)

func TestUnitAccountPayIn(t *testing.T) {
	type args struct {
		userId         string
		sm4Key         string
		userPrivateKey string
		accnbr         string
		dmadbt         string
		dmacrt         string
		trxamt         string
		trxtxt         string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.UnitAccountPayInResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnitAccountPayIn(tt.args.userId, tt.args.sm4Key, tt.args.userPrivateKey, tt.args.accnbr, tt.args.dmadbt, tt.args.dmacrt, tt.args.trxamt, tt.args.trxtxt)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnitAccountPayIn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnitAccountPayIn() got = %v, want %v", got, tt.want)
			}
		})
	}
}

package unit_manager

import (
	"github.com/ahKevinXy/go-cmb/models"
	"reflect"
	"testing"
)

func TestGetUnitAccountTransHistoryList(t *testing.T) {
	type args struct {
		userId         string
		asePrivateKey  string
		userPrivateKey string
		accnbr         string
		dmanbr         string
		begdat         string
		enddat         string
		ctnkey         string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.UnitAccountTransHistoryResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUnitAccountTransHistoryList(tt.args.userId, tt.args.asePrivateKey, tt.args.userPrivateKey, tt.args.accnbr, tt.args.dmanbr, tt.args.begdat, tt.args.enddat, tt.args.ctnkey)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUnitAccountTransHistoryList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUnitAccountTransHistoryList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUnitAccountTransList(t *testing.T) {
	type args struct {
		userId         string
		asePrivateKey  string
		userPrivateKey string
		accnbr         string
		dmanbr         string
		ctnkey         string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.UnitAccountTransDailyResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUnitAccountTransList(tt.args.userId, tt.args.asePrivateKey, tt.args.userPrivateKey, tt.args.accnbr, tt.args.dmanbr, tt.args.ctnkey)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUnitAccountTransList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUnitAccountTransList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

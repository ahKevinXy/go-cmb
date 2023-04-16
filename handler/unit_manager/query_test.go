package unit_manager

import (
	"github.com/ahKevinXy/go-cmb/models"
	"reflect"
	"testing"
)

func TestQueryUnitAccountByBusNo(t *testing.T) {
	type args struct {
		userId         string
		asePrivateKey  string
		userPrivateKey string
		yurref         string
		bgndat         string
		enddat         string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.QueryUnitTransByBusNoResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := QueryUnitAccountByBusNo(tt.args.userId, tt.args.asePrivateKey, tt.args.userPrivateKey, tt.args.yurref, tt.args.bgndat, tt.args.enddat)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryUnitAccountByBusNo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryUnitAccountByBusNo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryUnitAccountDetail(t *testing.T) {
	type args struct {
		userId         string
		asePrivateKey  string
		userPrivateKey string
		reqNbr         string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.QueryUnitAccountDetailResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := QueryUnitAccountDetail(tt.args.userId, tt.args.asePrivateKey, tt.args.userPrivateKey, tt.args.reqNbr)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryUnitAccountDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryUnitAccountDetail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

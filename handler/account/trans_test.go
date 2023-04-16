package account

import (
	"github.com/ahKevinXy/go-cmb/models"
	"reflect"
	"testing"
)

func TestGetMainAccountTransInfo(t *testing.T) {
	type args struct {
		userId         string
		asePrivateKey  string
		userPrivateKey string
		bbknbr         string
		accnbr         string
		trsDat         string
		trsseq         string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.GetMainAccountTransInfoResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetMainAccountTransInfo(tt.args.userId, tt.args.asePrivateKey, tt.args.userPrivateKey, tt.args.bbknbr, tt.args.accnbr, tt.args.trsDat, tt.args.trsseq)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMainAccountTransInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMainAccountTransInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryAccountTransInfo(t *testing.T) {
	type args struct {
		userId         string
		asePrivateKey  string
		userPrivateKey string
		bbknbr         string
		accnbr         string
		bgndat         string
		enddat         string
		lowamt         string
		hghamt         string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.QueryAccountTransInfoResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := QueryAccountTransInfo(tt.args.userId, tt.args.asePrivateKey, tt.args.userPrivateKey, tt.args.bbknbr, tt.args.accnbr, tt.args.bgndat, tt.args.enddat, tt.args.lowamt, tt.args.hghamt)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryAccountTransInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryAccountTransInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

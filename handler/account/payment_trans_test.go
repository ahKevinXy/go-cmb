package account

import (
	"reflect"
	"testing"

	"github.com/ahKevinXy/go-cmb/models"
)

func TestQueryAccountPaymentDetail(t *testing.T) {
	type args struct {
		userId         string
		sm4Key         string
		userPrivateKey string
		bthNbr         string
		autStr         string
		rtnStr         string
		ctnKey         string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.QueryAccountPaymentDetailResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := QueryAccountPaymentDetail(tt.args.userId, tt.args.sm4Key, tt.args.userPrivateKey, tt.args.bthNbr, tt.args.autStr, tt.args.rtnStr, tt.args.ctnKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryAccountPaymentDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryAccountPaymentDetail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryAccountPaymentTransInfo(t *testing.T) {
	type args struct {
		userId         string
		sm4Key         string
		userPrivateKey string
		begDat         string
		endDat         string
		autStr         string
		rtnStr         string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.MainAccountSinglePayResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := QueryAccountPaymentTransInfo(tt.args.userId, tt.args.sm4Key, tt.args.userPrivateKey, tt.args.begDat, tt.args.endDat, tt.args.autStr, tt.args.rtnStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryAccountPaymentTransInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryAccountPaymentTransInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

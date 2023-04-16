package account

import (
	"github.com/ahKevinXy/go-cmb/models"
	"reflect"
	"testing"
)

func TestMainAccountBatchPay(t *testing.T) {
	type args struct {
		userId         string
		asePrivateKey  string
		userPrivateKey string
		busCode        string
		busMode        string
		dtlNbr         string
		ctnFlg         string
		ctnSts         string
		bthNbr         string
		payList        []*models.Bb1paybhx1
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
			got, err := MainAccountBatchPay(tt.args.userId, tt.args.asePrivateKey, tt.args.userPrivateKey, tt.args.busCode, tt.args.busMode, tt.args.dtlNbr, tt.args.ctnFlg, tt.args.ctnSts, tt.args.bthNbr, tt.args.payList)
			if (err != nil) != tt.wantErr {
				t.Errorf("MainAccountBatchPay() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MainAccountBatchPay() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMainAccountPaySingle(t *testing.T) {
	type args struct {
		userId         string
		asePrivateKey  string
		userPrivateKey string
		dbAcc          string
		buscode        string
		busMode        string
		dmaNbr         string
		crtAcc         string
		crtNam         string
		crtBnk         string
		crtAdr         string
		bnkFlg         string
		trsAmt         string
		brdNbr         string
		nusAge         string
		yurRef         string
		trsTyp         string
		busNar         string
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
			got, err := MainAccountPaySingle(tt.args.userId, tt.args.asePrivateKey, tt.args.userPrivateKey, tt.args.dbAcc, tt.args.buscode, tt.args.busMode, tt.args.dmaNbr, tt.args.crtAcc, tt.args.crtNam, tt.args.crtBnk, tt.args.crtAdr, tt.args.bnkFlg, tt.args.trsAmt, tt.args.brdNbr, tt.args.nusAge, tt.args.yurRef, tt.args.trsTyp, tt.args.busNar)
			if (err != nil) != tt.wantErr {
				t.Errorf("MainAccountPaySingle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MainAccountPaySingle() got = %v, want %v", got, tt.want)
			}
		})
	}
}

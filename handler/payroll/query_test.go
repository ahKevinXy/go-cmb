package payroll

import (
	"github.com/ahKevinXy/go-cmb/models"
	"reflect"
	"testing"
)

func TestQueryBatchTransInfo(t *testing.T) {
	type args struct {
		userId         string
		asePrivateKey  string
		userPrivateKey string
		buscode        string
		yurref         string
		bthnbr         string
		trxseq         string
		histga         string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.QueryBatchTransInfoResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := QueryBatchTransInfo(tt.args.userId, tt.args.asePrivateKey, tt.args.userPrivateKey, tt.args.buscode, tt.args.yurref, tt.args.bthnbr, tt.args.trxseq, tt.args.histga)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryBatchTransInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryBatchTransInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryBatchTransList(t *testing.T) {
	type args struct {
		userId         string
		asePrivateKey  string
		userPrivateKey string
		buscode        string
		yurref         string
		bgndat         string
		enddat         string
		cntkey         string
		dattyp         string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.QueryBatchTransListResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := QueryBatchTransList(tt.args.userId, tt.args.asePrivateKey, tt.args.userPrivateKey, tt.args.buscode, tt.args.yurref, tt.args.bgndat, tt.args.enddat, tt.args.cntkey, tt.args.dattyp)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryBatchTransList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryBatchTransList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryPayrollTransDetail(t *testing.T) {
	type args struct {
		userId         string
		asePrivateKey  string
		userPrivateKey string
		reqnbr         string
		bthnbr         string
		trxseq         string
		histag         string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.QueryPayrollTransDetailResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := QueryPayrollTransDetail(tt.args.userId, tt.args.asePrivateKey, tt.args.userPrivateKey, tt.args.reqnbr, tt.args.bthnbr, tt.args.trxseq, tt.args.histag)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryPayrollTransDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryPayrollTransDetail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

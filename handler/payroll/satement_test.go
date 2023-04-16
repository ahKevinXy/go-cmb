package payroll

import (
	"github.com/ahKevinXy/go-cmb/models"
	"reflect"
	"testing"
)

func TestQueryPayrollStatement(t *testing.T) {
	type args struct {
		userId         string
		asePrivateKey  string
		userPrivateKey string
		payeac         string
		begdat         string
		enddat         string
		buscod         string
		busmod         string
		eacnam         string
		ptyref         string
		trxsrl         string
		minamt         string
		maxamt         string
		prtmod         string
		begidx         string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.QueryPayrollStatementResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := QueryPayrollStatement(tt.args.userId, tt.args.asePrivateKey, tt.args.userPrivateKey, tt.args.payeac, tt.args.begdat, tt.args.enddat, tt.args.buscod, tt.args.busmod, tt.args.eacnam, tt.args.ptyref, tt.args.trxsrl, tt.args.minamt, tt.args.maxamt, tt.args.prtmod, tt.args.begidx)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryPayrollStatement() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryPayrollStatement() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryPayrollStatementDownloadUrl(t *testing.T) {
	type args struct {
		userId         string
		asePrivateKey  string
		userPrivateKey string
		taskid         string
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
			got, err := QueryPayrollStatementDownloadUrl(tt.args.userId, tt.args.asePrivateKey, tt.args.userPrivateKey, tt.args.taskid)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryPayrollStatementDownloadUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryPayrollStatementDownloadUrl() got = %v, want %v", got, tt.want)
			}
		})
	}
}

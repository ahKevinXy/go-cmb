package account

import (
	"reflect"
	"testing"

	"github.com/ahKevinXy/go-cmb/models"
)

func TestAsyncStatement(t *testing.T) {
	type args struct {
		userId         string
		sm4Key         string
		userPrivateKey string
		eacnbr         string
		begdat         string
		enddat         string
		begamt         string
		endamt         string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.QueryAccountCallbackAsyncResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AsyncStatement(tt.args.userId, tt.args.sm4Key, tt.args.userPrivateKey, tt.args.eacnbr, tt.args.begdat, tt.args.enddat, tt.args.begamt, tt.args.endamt)
			if (err != nil) != tt.wantErr {
				t.Errorf("AsyncStatement() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AsyncStatement() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBatchStatementQuery(t *testing.T) {
	type args struct {
		userId         string
		sm4Key         string
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
		want    *models.BatchStatementQueryResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BatchStatementQuery(tt.args.userId, tt.args.sm4Key, tt.args.userPrivateKey, tt.args.bbknbr, tt.args.accnbr, tt.args.bgndat, tt.args.enddat, tt.args.lowamt, tt.args.hghamt)
			if (err != nil) != tt.wantErr {
				t.Errorf("BatchStatementQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BatchStatementQuery() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetStatementPdf(t *testing.T) {
	type args struct {
		userId         string
		sm4Key         string
		userPrivateKey string
		taskid         string
		qwenab         string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.QueryAccountCallbackDownloadPdfResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetStatementPdf(tt.args.userId, tt.args.sm4Key, tt.args.userPrivateKey, tt.args.taskid, tt.args.qwenab)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStatementPdf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStatementPdf() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSingleStatementQuery(t *testing.T) {
	type args struct {
		userId         string
		sm4Key         string
		userPrivateKey string
		eacnbr         string
		quedat         string
		trsseq         string
		primod         string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.SingleCallBackPdfResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SingleStatementQuery(tt.args.userId, tt.args.sm4Key, tt.args.userPrivateKey, tt.args.eacnbr, tt.args.quedat, tt.args.trsseq, tt.args.primod)
			if (err != nil) != tt.wantErr {
				t.Errorf("SingleStatementQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SingleStatementQuery() got = %v, want %v", got, tt.want)
			}
		})
	}
}

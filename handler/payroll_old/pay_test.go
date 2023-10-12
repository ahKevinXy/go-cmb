package payroll_old

import (
	"github.com/ahKevinXy/go-cmb/v1/models"
	"reflect"
	"testing"
)

func TestCreditHandleOtherBySup(t *testing.T) {
	type args struct {
		userId         string
		sm4PrivateKey  string
		userPrivateKey string
		busmod         string
		total          []*models.Ntagcsaix1
		detail         []*models.Ntagcsaix2
	}
	tests := []struct {
		name    string
		args    args
		want    *models.PayrollOldCreditOtherBySupResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreditHandleOtherBySup(tt.args.userId, tt.args.sm4PrivateKey, tt.args.userPrivateKey, tt.args.busmod, tt.args.total, tt.args.detail)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreditHandleOtherBySup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreditHandleOtherBySup() got = %v, want %v", got, tt.want)
			}
		})
	}
}

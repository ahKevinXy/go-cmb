package payroll

import (
	"github.com/ahKevinXy/go-cmb/models"
	"reflect"
	"testing"
)

func TestUnitPayrollPayment(t *testing.T) {
	type args struct {
		userId         string
		sm2PrivateKey  string
		userPrivateKey string
		payMod         []*models.Bb6Busmod
		totalInfo      []*models.Bb6Aclakx1
		payList        []*models.Bb6Aclaky1
	}
	tests := []struct {
		name    string
		args    args
		want    *models.UnitPayrollPaymentResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnitPayrollPayment(tt.args.userId, tt.args.sm2PrivateKey, tt.args.userPrivateKey, tt.args.payMod, tt.args.totalInfo, tt.args.payList)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnitPayrollPayment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnitPayrollPayment() got = %v, want %v", got, tt.want)
			}
		})
	}
}

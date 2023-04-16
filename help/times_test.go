package help

import (
	"reflect"
	"testing"
	"time"
)

func TestGetCmbOpTime(t *testing.T) {
	type args struct {
		d string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCmbOpTime(tt.args.d)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCmbOpTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCmbOpTime() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCmbTransTime(t *testing.T) {
	type args struct {
		d string
		t string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCmbTransTime(tt.args.d, tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCmbTransTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCmbTransTime() got = %v, want %v", got, tt.want)
			}
		})
	}
}

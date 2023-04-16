package help

import (
	"reflect"
	"testing"
)

func TestCmbSignRequest(t *testing.T) {
	type args struct {
		reqStr  string
		funCode string
		uid     string
		userKey string
		AESKey  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CmbSignRequest(tt.args.reqStr, tt.args.funCode, tt.args.uid, tt.args.userKey, tt.args.AESKey); got != tt.want {
				t.Errorf("CmbSignRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSignatureDataSM(t *testing.T) {
	type args struct {
		reqStr  string
		funCode string
		uid     string
		userKey string
		AESKey  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SignatureDataSM(tt.args.reqStr, tt.args.funCode, tt.args.uid, tt.args.userKey, tt.args.AESKey); got != tt.want {
				t.Errorf("SignatureDataSM() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pkcs5UnPadding(t *testing.T) {
	type args struct {
		src []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pkcs5UnPadding(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pkcs5UnPadding() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sm4Decrypt(t *testing.T) {
	type args struct {
		key        []byte
		iv         []byte
		cipherText []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sm4Decrypt(tt.args.key, tt.args.iv, tt.args.cipherText)
			if (err != nil) != tt.wantErr {
				t.Errorf("sm4Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sm4Decrypt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

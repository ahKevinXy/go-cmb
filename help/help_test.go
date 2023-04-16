package help

import (
	"reflect"
	"testing"
)

func TestGetJson(t *testing.T) {
	type args struct {
		reqAccount string
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
			if got := GetJson(tt.args.reqAccount); got != tt.want {
				t.Errorf("GetJson() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSignSM2(t *testing.T) {
	type args struct {
		signContent string
		privateKey  string
		userId      string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SignSM2(tt.args.signContent, tt.args.privateKey, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignSM2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SignSM2() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSignSm2ByCmd(t *testing.T) {
	type args struct {
		signContent string
		privateKey  string
		userId      string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SignSm2ByCmd(tt.args.signContent, tt.args.privateKey, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignSm2ByCmd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SignSm2ByCmd() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSignSm2ByRequest(t *testing.T) {
	type args struct {
		signContent string
		privateKey  string
		userId      string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SignSm2ByRequest(tt.args.signContent, tt.args.privateKey, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignSm2ByRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SignSm2ByRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSm4Encrypt(t *testing.T) {
	type args struct {
		key       []byte
		iv        []byte
		plainText []byte
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
			got, err := Sm4Encrypt(tt.args.key, tt.args.iv, tt.args.plainText)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sm4Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sm4Encrypt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pkcs5Padding(t *testing.T) {
	type args struct {
		src       []byte
		blockSize int
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
			if got := pkcs5Padding(tt.args.src, tt.args.blockSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pkcs5Padding() = %v, want %v", got, tt.want)
			}
		})
	}
}

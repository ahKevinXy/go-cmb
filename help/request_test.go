package help

import (
	"net/http/cookiejar"
	"testing"
)

func TestMakeHttpRequest(t *testing.T) {
	type args struct {
		method string
		url    string
		entity map[string]interface{}
		jar    *cookiejar.Jar
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := MakeHttpRequest(tt.args.method, tt.args.url, tt.args.entity, tt.args.jar)
			if (err != nil) != tt.wantErr {
				t.Errorf("MakeHttpRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MakeHttpRequest() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MakeHttpRequest() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

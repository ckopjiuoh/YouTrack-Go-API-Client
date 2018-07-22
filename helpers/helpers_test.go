package helpers

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

func TestReadBody(t *testing.T) {
	type args struct {
		r *http.Response
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Positive test",
			args: struct{ r *http.Response }{r: &http.Response{
				Body: ioutil.NopCloser(bytes.NewReader([]byte("foo"))),
			}},
			want:    []byte("foo"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadBody(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadBody(%v) error = %v, wantErr %v", tt.args.r, err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadBody(%v) = %v, want %v", tt.args.r, got, tt.want)
			}
		})
	}
}

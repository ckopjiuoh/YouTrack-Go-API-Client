package ytapi

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestNewYouTrackClient_Negative(t *testing.T) {
	_, err := NewYouTrackAPI(
		"some",
		"http://127.0.0.1",
		true,
	)

	if err == nil {
		t.Error(err.Error())
		t.Fail()
	}
}

func TestNewYouTrackClient_Positive(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		fmt.Fprintln(w, "Hello, client")
	}))

	defer ts.Close()

	_, err := NewYouTrackAPI(
		"some",
		ts.URL,
		true,
	)

	if err != nil {
		t.Error(err.Error())
		t.Fail()
	}
}

func TestNewYouTrackClient_Negative_WrongStatusCode(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(403)
		fmt.Fprintln(w, "Hello, client")
	}))

	defer ts.Close()

	_, err := NewYouTrackAPI(
		"some",
		ts.URL,
		true,
	)

	if err == nil {
		t.Error(err.Error())
		t.Fail()
	}
}

func TestNewYouTrackAPI(t *testing.T) {
	type args struct {
		token string
		addr  string
		debug bool
	}
	tests := []struct {
		name    string
		args    args
		want    *YouTrackAPI
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewYouTrackAPI(tt.args.token, tt.args.addr, tt.args.debug)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewYouTrackAPI(%v, %v, %v) error = %v, wantErr %v", tt.args.token, tt.args.addr, tt.args.debug, err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewYouTrackAPI(%v, %v, %v) = %v, want %v", tt.args.token, tt.args.addr, tt.args.debug, got, tt.want)
			}
		})
	}
}

func TestYouTrackAPI_MakeRequest(t *testing.T) {
	type fields struct {
		Token  string
		Addr   string
		Debug  bool
		Client *http.Client
	}
	type args struct {
		request *Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *http.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &YouTrackAPI{
				Token:  tt.fields.Token,
				Addr:   tt.fields.Addr,
				Debug:  tt.fields.Debug,
				Client: tt.fields.Client,
			}
			got, err := api.MakeRequest(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("YouTrackAPI.MakeRequest(%v) error = %v, wantErr %v", tt.args.request, err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("YouTrackAPI.MakeRequest(%v) = %v, want %v", tt.args.request, got, tt.want)
			}
		})
	}
}

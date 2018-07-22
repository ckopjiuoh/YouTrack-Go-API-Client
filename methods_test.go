package ytapi

import (
	"net/http"
	"testing"
)

func TestYouTrackAPI_HealthCheck(t *testing.T) {
	type fields struct {
		Token  string
		Addr   string
		Debug  bool
		Client *http.Client
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &YouTrackAPI{
				Token:  tt.fields.Token,
				Addr:   tt.fields.Addr,
				Debug:  tt.fields.Debug,
				Client: tt.fields.Client,
			}
			if err := c.HealthCheck(); (err != nil) != tt.wantErr {
				t.Errorf("YouTrackAPI.HealthCheck() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

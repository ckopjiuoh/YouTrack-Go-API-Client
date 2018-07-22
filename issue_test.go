package ytapi

import (
	"net/http"
	"reflect"
	"testing"
)

func TestIssue_GetValueByFieldName(t *testing.T) {
	type fields struct {
		ID       string
		EntityID string
		JiraID   string
		Field    []Field
		Comment  []interface{}
		Tag      []Tag
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := &Issue{
				ID:       tt.fields.ID,
				EntityID: tt.fields.EntityID,
				JiraID:   tt.fields.JiraID,
				Field:    tt.fields.Field,
				Comment:  tt.fields.Comment,
				Tag:      tt.fields.Tag,
			}
			got, err := is.GetValueByFieldName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Issue.GetValueByFieldName(%v) error = %v, wantErr %v", tt.args.name, err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Issue.GetValueByFieldName(%v) = %v, want %v", tt.args.name, got, tt.want)
			}
		})
	}
}

func TestIssue_GetID(t *testing.T) {
	type fields struct {
		ID       string
		EntityID string
		JiraID   string
		Field    []Field
		Comment  []interface{}
		Tag      []Tag
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := &Issue{
				ID:       tt.fields.ID,
				EntityID: tt.fields.EntityID,
				JiraID:   tt.fields.JiraID,
				Field:    tt.fields.Field,
				Comment:  tt.fields.Comment,
				Tag:      tt.fields.Tag,
			}
			if got := is.GetID(); got != tt.want {
				t.Errorf("Issue.GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIssue_GetEntityID(t *testing.T) {
	type fields struct {
		ID       string
		EntityID string
		JiraID   string
		Field    []Field
		Comment  []interface{}
		Tag      []Tag
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := &Issue{
				ID:       tt.fields.ID,
				EntityID: tt.fields.EntityID,
				JiraID:   tt.fields.JiraID,
				Field:    tt.fields.Field,
				Comment:  tt.fields.Comment,
				Tag:      tt.fields.Tag,
			}
			if got := is.GetEntityID(); got != tt.want {
				t.Errorf("Issue.GetEntityID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIssue_GetProjectShortName(t *testing.T) {
	type fields struct {
		ID       string
		EntityID string
		JiraID   string
		Field    []Field
		Comment  []interface{}
		Tag      []Tag
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := &Issue{
				ID:       tt.fields.ID,
				EntityID: tt.fields.EntityID,
				JiraID:   tt.fields.JiraID,
				Field:    tt.fields.Field,
				Comment:  tt.fields.Comment,
				Tag:      tt.fields.Tag,
			}
			got, err := is.GetProjectShortName()
			if (err != nil) != tt.wantErr {
				t.Errorf("Issue.GetProjectShortName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Issue.GetProjectShortName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIssue_GetNumberInProject(t *testing.T) {
	type fields struct {
		ID       string
		EntityID string
		JiraID   string
		Field    []Field
		Comment  []interface{}
		Tag      []Tag
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := &Issue{
				ID:       tt.fields.ID,
				EntityID: tt.fields.EntityID,
				JiraID:   tt.fields.JiraID,
				Field:    tt.fields.Field,
				Comment:  tt.fields.Comment,
				Tag:      tt.fields.Tag,
			}
			got, err := is.GetNumberInProject()
			if (err != nil) != tt.wantErr {
				t.Errorf("Issue.GetNumberInProject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Issue.GetNumberInProject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIssue_GetSummary(t *testing.T) {
	type fields struct {
		ID       string
		EntityID string
		JiraID   string
		Field    []Field
		Comment  []interface{}
		Tag      []Tag
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := &Issue{
				ID:       tt.fields.ID,
				EntityID: tt.fields.EntityID,
				JiraID:   tt.fields.JiraID,
				Field:    tt.fields.Field,
				Comment:  tt.fields.Comment,
				Tag:      tt.fields.Tag,
			}
			got, err := is.GetSummary()
			if (err != nil) != tt.wantErr {
				t.Errorf("Issue.GetSummary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Issue.GetSummary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIssue_GetDescription(t *testing.T) {
	type fields struct {
		ID       string
		EntityID string
		JiraID   string
		Field    []Field
		Comment  []interface{}
		Tag      []Tag
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := &Issue{
				ID:       tt.fields.ID,
				EntityID: tt.fields.EntityID,
				JiraID:   tt.fields.JiraID,
				Field:    tt.fields.Field,
				Comment:  tt.fields.Comment,
				Tag:      tt.fields.Tag,
			}
			got, err := is.GetDescription()
			if (err != nil) != tt.wantErr {
				t.Errorf("Issue.GetDescription() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Issue.GetDescription() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYouTrackAPI_GetAnIssue(t *testing.T) {
	type fields struct {
		Token  string
		Addr   string
		Debug  bool
		Client *http.Client
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Issue
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
			got, err := api.GetAnIssue(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("YouTrackAPI.GetAnIssue(%v) error = %v, wantErr %v", tt.args.name, err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("YouTrackAPI.GetAnIssue(%v) = %v, want %v", tt.args.name, got, tt.want)
			}
		})
	}
}

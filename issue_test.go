package ytapi

import (
	"net/http"
	"os"
	"reflect"
	"testing"
)

var TestIssue = &Issue{
	ID:       "Issue",
	EntityID: "Issue",
	JiraID:   "Issue",
	Field: []Field{
		Field{
			Name:  "Field",
			Value: "01",
		},
	},
	Comment: nil,
	Tag: []Tag{
		Tag{
			Value:    "Tag",
			CSSClass: "Tag",
		},
	},
}

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
		{name: "Positive test",
			fields: struct {
				ID       string
				EntityID string
				JiraID   string
				Field    []Field
				Comment  []interface{}
				Tag      []Tag
			}{ID: "Issue",
				EntityID: "Issue",
				JiraID:   "Issue",
				Field: []Field{
					Field{
						Name:  "Field",
						Value: "01",
					},
				},
				Comment: nil,
				Tag: []Tag{
					Tag{
						Value:    "Tag",
						CSSClass: "Tag",
					},
				}}, args: struct{ name string }{name: "Field"}, want: "01", wantErr: false},
		{name: "Negative test",
			fields: struct {
				ID       string
				EntityID string
				JiraID   string
				Field    []Field
				Comment  []interface{}
				Tag      []Tag
			}{ID: "Issue",
				EntityID: "Issue",
				JiraID:   "Issue",
				Field: []Field{
					Field{
						Name:  "Field",
						Value: "01",
					},
				},
				Comment: nil,
				Tag: []Tag{
					Tag{
						Value:    "Tag",
						CSSClass: "Tag",
					},
				}}, args: struct{ name string }{name: "FieldTwo"}, wantErr: true},
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
		{name: "Positive test",
			fields: struct {
				ID       string
				EntityID string
				JiraID   string
				Field    []Field
				Comment  []interface{}
				Tag      []Tag
			}{ID: "YouTrack issue",
				EntityID: "Issue",
				JiraID:   "Issue",
				Field: []Field{
					Field{
						Name:  "Field",
						Value: "01",
					},
				},
				Comment: nil,
				Tag: []Tag{
					Tag{
						Value:    "Tag",
						CSSClass: "Tag",
					},
				}}, want: "YouTrack issue"},
		{name: "Positive test empty field",
			fields: struct {
				ID       string
				EntityID string
				JiraID   string
				Field    []Field
				Comment  []interface{}
				Tag      []Tag
			}{ID: "",
				EntityID: "Issue",
				JiraID:   "Issue",
				Field: []Field{
					Field{
						Name:  "Field",
						Value: "01",
					},
				},
				Comment: nil,
				Tag: []Tag{
					Tag{
						Value:    "Tag",
						CSSClass: "Tag",
					},
				}}, want: ""},
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
		{name: "Positive test",
			fields: struct {
				ID       string
				EntityID string
				JiraID   string
				Field    []Field
				Comment  []interface{}
				Tag      []Tag
			}{ID: "YouTrack issue",
				EntityID: "YouTrack issue",
				JiraID:   "Issue",
				Field: []Field{
					Field{
						Name:  "Field",
						Value: "01",
					},
				},
				Comment: nil,
				Tag: []Tag{
					Tag{
						Value:    "Tag",
						CSSClass: "Tag",
					},
				}}, want: "YouTrack issue"},
		{name: "Positive test empty field",
			fields: struct {
				ID       string
				EntityID string
				JiraID   string
				Field    []Field
				Comment  []interface{}
				Tag      []Tag
			}{ID: "YouTrack issue",
				EntityID: "",
				JiraID:   "Issue",
				Field: []Field{
					Field{
						Name:  "Field",
						Value: "01",
					},
				},
				Comment: nil,
				Tag: []Tag{
					Tag{
						Value:    "Tag",
						CSSClass: "Tag",
					},
				}}, want: ""},
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
		{name: "Positive test",
			fields: struct {
				ID       string
				EntityID string
				JiraID   string
				Field    []Field
				Comment  []interface{}
				Tag      []Tag
			}{ID: "YouTrack issue",
				EntityID: "YouTrack issue",
				JiraID:   "Issue",
				Field: []Field{
					Field{
						Name:  "projectShortName",
						Value: "project",
					},
				},
				Comment: nil,
				Tag: []Tag{
					Tag{
						Value:    "Tag",
						CSSClass: "Tag",
					},
				}}, want: "project", wantErr: false},
		{name: "Positive test EmptyField",
			fields: struct {
				ID       string
				EntityID string
				JiraID   string
				Field    []Field
				Comment  []interface{}
				Tag      []Tag
			}{ID: "YouTrack issue",
				EntityID: "YouTrack issue",
				JiraID:   "Issue",
				Field: []Field{
					Field{
						Name:  "projectShortName",
						Value: "01",
					},
				},
				Comment: nil,
				Tag: []Tag{
					Tag{
						Value:    "Tag",
						CSSClass: "Tag",
					},
				}}, want: "01", wantErr: false},
		{name: "Negative test",
			fields: struct {
				ID       string
				EntityID string
				JiraID   string
				Field    []Field
				Comment  []interface{}
				Tag      []Tag
			}{ID: "YouTrack issue",
				EntityID: "YouTrack issue",
				JiraID:   "Issue",
				Field: []Field{
					Field{
						Name:  "Field",
						Value: "project",
					},
				},
				Comment: nil,
				Tag: []Tag{
					Tag{
						Value:    "Tag",
						CSSClass: "Tag",
					},
				}}, wantErr: true},
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
		{name: "Positive test",
			fields: struct {
				ID       string
				EntityID string
				JiraID   string
				Field    []Field
				Comment  []interface{}
				Tag      []Tag
			}{ID: "YouTrack issue",
				EntityID: "YouTrack issue",
				JiraID:   "Issue",
				Field: []Field{
					Field{
						Name:  "numberInProject",
						Value: "01",
					},
				},
				Comment: nil,
				Tag: []Tag{
					Tag{
						Value:    "Tag",
						CSSClass: "Tag",
					},
				}}, want: 1},
		{name: "Negative test",
			fields: struct {
				ID       string
				EntityID string
				JiraID   string
				Field    []Field
				Comment  []interface{}
				Tag      []Tag
			}{ID: "YouTrack issue",
				EntityID: "YouTrack issue",
				JiraID:   "Issue",
				Field: []Field{
					Field{
						Name:  "numberInProject",
						Value: "something",
					},
				},
				Comment: nil,
				Tag: []Tag{
					Tag{
						Value:    "Tag",
						CSSClass: "Tag",
					},
				}}, wantErr: true},
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
		{name: "Positive test",
			fields: struct {
				ID       string
				EntityID string
				JiraID   string
				Field    []Field
				Comment  []interface{}
				Tag      []Tag
			}{ID: "YouTrack issue",
				EntityID: "YouTrack issue",
				JiraID:   "Issue",
				Field: []Field{
					Field{
						Name:  "summary",
						Value: "01",
					},
				},
				Comment: nil,
				Tag: []Tag{
					Tag{
						Value:    "Tag",
						CSSClass: "Tag",
					},
				}}, want: "01"},
		{name: "Negative test",
			fields: struct {
				ID       string
				EntityID string
				JiraID   string
				Field    []Field
				Comment  []interface{}
				Tag      []Tag
			}{ID: "YouTrack issue",
				EntityID: "YouTrack issue",
				JiraID:   "Issue",
				Field: []Field{
					Field{
						Name:  "sumary",
						Value: "01",
					},
				},
				Comment: nil,
				Tag: []Tag{
					Tag{
						Value:    "Tag",
						CSSClass: "Tag",
					},
				}}, wantErr: true},
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
		{name: "Positive test",
			fields: struct {
				ID       string
				EntityID string
				JiraID   string
				Field    []Field
				Comment  []interface{}
				Tag      []Tag
			}{ID: "YouTrack issue",
				EntityID: "YouTrack issue",
				JiraID:   "Issue",
				Field: []Field{
					Field{
						Name:  "description",
						Value: "01",
					},
				},
				Comment: nil,
				Tag: []Tag{
					Tag{
						Value:    "Tag",
						CSSClass: "Tag",
					},
				}}, want: "01"},
		{name: "Negative test",
			fields: struct {
				ID       string
				EntityID string
				JiraID   string
				Field    []Field
				Comment  []interface{}
				Tag      []Tag
			}{ID: "YouTrack issue",
				EntityID: "YouTrack issue",
				JiraID:   "Issue",
				Field: []Field{
					Field{
						Name:  "summary",
						Value: "01",
					},
				},
				Comment: nil,
				Tag: []Tag{
					Tag{
						Value:    "Tag",
						CSSClass: "Tag",
					},
				}}, wantErr: true},
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

func TestYouTrackAPI_CreateIssue(t *testing.T) {
	type fields struct {
		Token  string
		Addr   string
		Debug  bool
		Client *http.Client
	}
	type args struct {
		project        string
		summary        string
		description    string
		permittedGroup string
		attachments    []os.File
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
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
			if err := api.CreateIssue(tt.args.project, tt.args.summary, tt.args.description, tt.args.permittedGroup, tt.args.attachments); (err != nil) != tt.wantErr {
				t.Errorf("YouTrackAPI.CreateIssue(%v, %v, %v, %v, %v) error = %v, wantErr %v", tt.args.project, tt.args.summary, tt.args.description, tt.args.permittedGroup, tt.args.attachments, err, tt.wantErr)
			}
		})
	}
}

func TestYouTrackAPI_GetIssueHistory(t *testing.T) {
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
		want    []*Issue
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
			got, err := api.GetIssueHistory(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("YouTrackAPI.GetIssueHistory(%v) error = %v, wantErr %v", tt.args.name, err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("YouTrackAPI.GetIssueHistory(%v) = %v, want %v", tt.args.name, got, tt.want)
			}
		})
	}
}

func TestYouTrackAPI_CheckIssueExist(t *testing.T) {
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
		want    bool
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
			got, err := api.CheckIssueExist(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("YouTrackAPI.CheckIssueExist(%v) error = %v, wantErr %v", tt.args.name, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("YouTrackAPI.CheckIssueExist(%v) = %v, want %v", tt.args.name, got, tt.want)
			}
		})
	}
}

func TestYouTrackAPI_GetListOfIssues(t *testing.T) {
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
		want    bool
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
			got, err := api.GetListOfIssues(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("YouTrackAPI.GetListOfIssues(%v) error = %v, wantErr %v", tt.args.name, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("YouTrackAPI.GetListOfIssues(%v) = %v, want %v", tt.args.name, got, tt.want)
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

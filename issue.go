package ytapi

import (
	"fmt"
	"net/url"
	"github.com/ckopjiuoh/YouTrack-Go-API-Client/helpers"
	"encoding/json"
)

type Issue struct {
	ID       string        `json:"id"`
	EntityID string        `json:"entityId"`
	JiraID   interface{}   `json:"jiraId"`
	Field    []*Field       `json:"field"`
	Comment  []interface{} `json:"comment"`
	Tag      []*Tag         `json:"tag"`
}

type Tag struct {
	Value    string `json:"value"`
	CSSClass string `json:"cssClass"`
}

type Color struct {
	Bg string `json:"bg"`
	Fg string `json:"fg"`
}

type Field struct {
	Name    string   `json:"name"`
	Value   string   `json:"value"`
	ValueID []string `json:"valueId,omitempty"`
	Color   Color    `json:"color,omitempty"`
}

func (api *YouTrackAPI) GetAnIssue(name string) (*Issue, error) {
	r, e := api.MakeRequest(
		&Request{
			&url.Values{},
			"/rest/issue/" + name,
			GET,
			nil,
			nil,
		},
	)

	if e != nil {
		return nil, e
	}

	if r.StatusCode == 404 {
		return nil, fmt.Errorf("issue not found")
	}

	b, e := helpers.ReadBody(r)

	if e != nil {
		return nil, e
	}

	issue := &Issue{}

	e = json.Unmarshal(b, issue)

	if e != nil {
		return nil, e
	}

	return issue, nil
}

package ytapi

import (
	"encoding/json"
	"fmt"
	"net/url"

	"strings"

	"strconv"

	"github.com/ckopjiuoh/go-youtrack-api-client/helpers"
)

type Issue struct {
	ID       string        `json:"id"`
	EntityID string        `json:"entityId"`
	JiraID   string        `json:"jiraId, omitempty"`
	Field    []Field       `json:"field"`
	Comment  []interface{} `json:"comment"`
	Tag      []Tag         `json:"tag"`
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
	Name    string      `json:"name"`
	Value   interface{} `json:"value"`
	ValueID []string    `json:"valueId,omitempty"`
	Color   Color       `json:"color,omitempty"`
}

func (is *Issue) GetValueByFieldName(name string) (interface{}, error) {
	for _, v := range is.Field {
		if strings.ToLower(v.Name) == strings.ToLower(name) {
			return v.Value, nil
		}
	}
	return nil, fmt.Errorf("can't find field: %s", name)
}

func (is *Issue) GetID() string {
	return is.ID
}

func (is *Issue) GetEntityID() string {
	return is.EntityID
}

func (is *Issue) GetProjectShortName() (string, error) {
	p, err := is.GetValueByFieldName("ProjectShortName")
	if err != nil {
		return "", err
	}
	return p.(string), nil
}

func (is *Issue) GetNumberInProject() (int, error) {
	p, err := is.GetValueByFieldName("NumberInProject")
	if err != nil {
		return -1, err
	}

	d, err := strconv.Atoi(p.(string))
	return d, err
}

func (is *Issue) GetSummary() (string, error) {
	p, err := is.GetValueByFieldName("ValueByFieldName")
	if err != nil {
		return "", err
	}
	return p.(string), nil
}

func (is *Issue) GetDescription() (string, error) {
	p, err := is.GetValueByFieldName("Description")
	if err != nil {
		return "", err
	}
	return p.(string), nil
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

	issue := Issue{}

	e = json.Unmarshal(b, &issue)

	if e != nil {
		println(string(b))
		return nil, e
	}

	return &issue, nil
}

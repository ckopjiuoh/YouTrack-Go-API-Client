package ytapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

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
	return nil, fmt.Errorf("field not found: %s", name)
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
	p, err := is.GetValueByFieldName("Summary")
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
		http.StatusOK,
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

func (api *YouTrackAPI) CreateIssue(project, summary, description, permittedGroup string, attachments []os.File) error {
	query := &url.Values{}
	query.Add("project", project)
	query.Add("summary", summary)
	query.Add("description", description)

	if permittedGroup != "" {
		query.Set("permittedGroup", permittedGroup)
	}

	//todo add attchments multipart
	_, e := api.MakeRequest(
		&Request{
			query,
			"/rest/issue?",
			PUT,
			nil,
			nil,
		},
		http.StatusCreated,
	)

	if e != nil {
		return e
	}

	return nil
}

func (api *YouTrackAPI) GetIssueHistory(name string) ([]*Issue, error) {
	r, e := api.MakeRequest(
		&Request{
			&url.Values{},
			fmt.Sprintf("/rest/issue/%s/history", strings.ToUpper(name)),
			GET,
			nil,
			nil,
		},
		http.StatusOK,
	)

	if r != nil {
		switch r.StatusCode {
		case http.StatusOK:
			b, _ := helpers.ReadBody(r)

			result := []*Issue{}

			json.Unmarshal(b, &result)
		case http.StatusNotFound:
			return nil, NotFound
		}
	}

	return nil, e
}

func (api *YouTrackAPI) CheckIssueExist(name string) (bool, error) {
	res, e := api.MakeRequest(
		&Request{
			&url.Values{},
			fmt.Sprintf("/rest/issue/%s/exists", strings.ToUpper(name)),
			GET,
			nil,
			nil,
		},
		http.StatusOK,
	)

	if res != nil {
		switch res.StatusCode {
		case http.StatusOK:
			return true, nil
		case http.StatusNotFound:
			return false, NotFound
		}
	}

	return false, e
}

func (api *YouTrackAPI) GetListOfIssues(name string) (bool, error) {
	res, e := api.MakeRequest(
		&Request{
			&url.Values{},
			fmt.Sprintf("/rest/issue/%s/exists", strings.ToUpper(name)),
			GET,
			nil,
			nil,
		},
		http.StatusOK,
	)

	if res != nil {
		switch res.StatusCode {
		case http.StatusOK:
			return true, nil
		case http.StatusNotFound:
			return false, NotFound
		}
	}

	return false, e
}

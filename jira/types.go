package jira

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

// HTTPDoer definition goes here?
type HTTPDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// The Client struct contains fields necessary to connect to (and authenticate against) Jira
type Client struct {
	hostname   string
	username   string
	password   string
	httpClient HTTPDoer
}

// New do thing
func New(host string, user string, pass string, client *http.Client) Client {
	return Client{host, user, pass, client}
}

// User definition goes here
type User = string

// Issue definition goes here
type Issue struct {
	JiraID     string `db:"jira_id"`
	JiraKey    string `db:"jira_key"`
	ExternalID string `db:"external_id"`
}

// Worklog definition goes here
type Worklog struct {
	JiraID      string
	JiraIssueID string
	Author      string
	Comment     string
	DateLogged  time.Time
	TimeLogged  time.Duration
}

// Scrollable definition goes here
type Scrollable struct {
	StartAt    int `json:"startAt"`
	MaxResults int `json:"maxResults"`
	Total      int `json:"total"`
}

// SearchResponse definition goes here
type SearchResponse struct {
	Issues []Issue `json:"issues"`
	Scrollable
}

// WorklogResponse definition goes here
type WorklogResponse struct {
	Worklogs []Worklog `json:"worklogs"`
	Scrollable
}

// UnmarshalJSON definition goes here
func (i *Issue) UnmarshalJSON(b []byte) error {
	var f interface{}
	json.Unmarshal(b, &f)

	m := f.(map[string]interface{})
	fieldmap := m["fields"]
	v := fieldmap.(map[string]interface{})

	i.JiraID, _ = m["id"].(string)
	i.JiraKey = m["key"].(string)
	i.ExternalID = v["customfield_10101"].(string)

	return nil
}

// UnmarshalJSON definition goes here
func (w *Worklog) UnmarshalJSON(b []byte) error {
	var f interface{}
	json.Unmarshal(b, &f)

	m := f.(map[string]interface{})
	authormap := m["author"]
	a := authormap.(map[string]interface{})

	w.JiraID = m["id"].(string)
	w.JiraIssueID = m["issueId"].(string)
	w.Author = a["name"].(string)
	w.Comment = m["comment"].(string)

	if date, err := time.Parse("2006-01-02T15:04:05.000-0500", m["created"].(string)); err == nil {
		w.DateLogged = date
	} else {
		return err
	}

	trimmedTime := strings.Replace(m["timeSpent"].(string), " ", "", -1)
	if time, err := time.ParseDuration(trimmedTime); err == nil {
		w.TimeLogged = time
	} else {
		return err
	}

	return nil
}

package jira

import (
	"encoding/json"
	"net/http"
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

// SearchResponse definition goes here
type SearchResponse struct {
	StartAt    int     `json:"startAt"`
	MaxResults int     `json:"maxResults"`
	Total      int     `json:"total"`
	Issues     []Issue `json:"issues"`
}

// User definition goes here
type User = string

// Issue definition goes here
type Issue struct {
	JiraID     string
	JiraKey    string
	ExternalID string
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

// New do thing
func New(host string, user string, pass string, client *http.Client) Client {
	return Client{host, user, pass, client}
}

// Worklog definition goes here
type Worklog struct {
	JiraID     int
	Author     string
	Comment    string
	DateLogged string // date would be better
	TimeLogged int    // time would be better
}

// func (w *Worklog) UnmarshalJSON(b []byte) error {
// 	return nil
// }

// User definition goes here

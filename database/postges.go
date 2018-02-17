package database

import (
	"github.com/michaelmosher/wio/jira"
)

// LoadJiraUsers functions
func (db *Client) LoadJiraUsers() (users []jira.User, err error) {
	smt := "SELECT USERNAME FROM jira_users order by id;"

	err = db.Select(&users, smt)
	return
}

// SaveJiraIssue function
func (db *Client) SaveJiraIssue(i jira.Issue) (err error) {
	iSmt := `INSERT INTO jira_issues (jira_id, jira_key, external_id) VALUES ($1, $2, $3)
			 ON CONFLICT (jira_key) DO
				UPDATE SET external_id = EXCLUDED.external_id;`

	_, err = db.Exec(iSmt, i.JiraID, i.JiraKey, i.ExternalID)
	return
}

// LoadJiraIssues function
func (db *Client) LoadJiraIssues() (issues []jira.Issue, err error) {
	smt := "SELECT jira_id, jira_key, external_id from jira_issues order by id;"

	err = db.Select(&issues, smt)
	return
}

// SaveWorklog function
func (db *Client) SaveWorklog(w jira.Worklog) (new bool, err error) {
	return
}

// LoadWorklogs

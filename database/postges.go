package database

import "github.com/michaelmosher/wio/jira"

// func (db *DBClient) LoadJiraUsers() (users []jira.User) {

// }

// LoadUsers

// SaveIssue function
func (db *Client) SaveIssue(i jira.Issue) (err error) {
	iSmt := `INSERT INTO issues (jira_id, jira_key, external_id)
	         VALUES (:jiraid, :jirakey, :externalid)
			 ON CONFLICT (jira_key) DO
				UPDATE SET external_id = EXCLUDED.external_id;`
	_, err = db.NamedExec(iSmt, i)
	return
}

// LoadIssues
// SaveWorklog
// LoadWorklogs

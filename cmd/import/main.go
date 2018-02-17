package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/michaelmosher/wio/database"
	"github.com/michaelmosher/wio/jira"
)

func main() {
	cfg := readConfig("wio.toml")
	// connect to redis
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	jira := jira.New(cfg.Jira.Hostname, cfg.Jira.Username, cfg.Jira.Password, client)
	db := database.New(cfg.DB.User, cfg.DB.Pass, cfg.DB.Host, cfg.DB.Port, cfg.DB.Name)

	importIssues(jira, db)
	importWorklogs(jira, db)
	// fetch worklogs -> save worklogs
	// publish worklogs to redis
}

func importIssues(j jira.Client, db database.Client) {
	users, _ := db.LoadJiraUsers()

	issueChan := make(chan jira.Issue)

	go func() {
		for _, u := range users {
			j.Issues(u, issueChan)
		}
		close(issueChan)
	}()

	for i := range issueChan {
		err := db.SaveJiraIssue(i)

		if err != nil {
			fmt.Println(err)
		}
	}
}

func importWorklogs(j jira.Client, db database.Client) {
	issues, _ := db.LoadJiraIssues()

	worklogChan := make(chan jira.Worklog, 2)

	go func() {
		for _, i := range issues {
			j.Worklogs(i, worklogChan)
		}
		close(worklogChan)
	}()

	for w := range worklogChan {
		fmt.Println(w)
	}
}

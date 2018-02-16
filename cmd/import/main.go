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
	// fetch worklogs -> save worklogs
	// publish worklogs to redis
}

func importIssues(j jira.Client, db database.Client) {
	users := [1]jira.User{"michaelm"} // this will come from DB soon

	issueChan := make(chan jira.Issue)
	for _, u := range users {
		go j.Issues(u, issueChan)
	}

	for i := range issueChan {
		err := db.SaveIssue(i)

		if err != nil {
			fmt.Println(err)
		}
		// save issue
	}
}

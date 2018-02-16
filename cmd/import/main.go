package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/michaelmosher/wio/jira"
)

func main() {
	cfg := readConfig("wio.toml")
	// connect to mysql
	// connect to redis
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	j := jira.New(cfg.Jira.Hostname, cfg.Jira.Username, cfg.Jira.Password, client)

	importIssues(j)
	// fetch worklogs -> save worklogs
	// publish worklogs to redis
}

func importIssues(j jira.Client /*, database //coming soon */) {
	users := [1]string{"michaelm"} // this will come from DB soon

	issueChan := make(chan jira.Issue)
	for _, u := range users {
		go j.Issues(u, issueChan)
	}

	for i := range issueChan {
		fmt.Println(i)
		// save issue
	}
}

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
	// create http client?
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	j := jira.New(cfg.Jira.Hostname, cfg.Jira.Username, cfg.Jira.Password, client)

	userChan := make(chan string, 5)
	issueChan := make(chan jira.Issue, 30)

	userChan <- "michaelm"
	go handleUsers(j, userChan, issueChan)
	handleIssues(j, issueChan)

	close(userChan)
	// fetch worklogs -> save worklogs
	// publish worklogs to redis
}

func handleUsers(j jira.Client, users chan string, issues chan jira.Issue) {
	for u := range users {
		j.Issues(u, issues)
	}
}

func handleIssues(_ jira.Client, issues chan jira.Issue) {
	for i := range issues {
		fmt.Println(i)
		// save issue
		// pass issue into getWorklog channel
	}
}

/*
Package jira does some things with Jira.
This is a block-style comment describing it.
*/
package jira

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// Issues function
func (c Client) Issues(user string, issueC chan Issue) {
	res := c.IssueSearch(user, 0)

	done := sendIssues(issueC, res)

	for !done {
		res = c.IssueSearch(user, res.StartAt+res.MaxResults)
		done = sendIssues(issueC, res)
	}
	close(issueC)
}

// IssueSearch function
func (c Client) IssueSearch(user string, startAt int) SearchResponse {
	req := c.issueSearchRequest(user, startAt)
	resp, err := c.httpClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	return issueSearchResponse(resp)
}

// In previous versions of JiraTrack,
// this would fetch issues that were open OR recently closed.
// given the change to constantly polling and storing issues,
// I now believe this is an unneccessary complication.
func (c Client) issueSearchRequest(user string, startAt int) *http.Request {
	belongingToUser := fmt.Sprintf("assignee=%s", user)
	beingUnresolved := "resolution = Unresolved"
	withOpenAirData := "\"External issue ID\" is not EMPTY"

	req := c.requestBuilder("/rest/api/2/search")

	q := req.URL.Query()
	q.Add("startAt", strconv.Itoa(startAt))
	q.Add("maxResults", "200")
	q.Add("fields", "issues,customfield_10101")
	q.Add("jql", fmt.Sprintf("%s AND %s AND %s", belongingToUser, beingUnresolved, withOpenAirData))
	req.URL.RawQuery = q.Encode()

	return req
}

func issueSearchResponse(resp *http.Response) SearchResponse {
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var r SearchResponse
	json.Unmarshal(bodyText, &r)
	return r
}

func sendIssues(issueC chan Issue, r SearchResponse) (done bool) {
	for _, i := range r.Issues {
		issueC <- i
	}

	return r.Total <= r.StartAt+r.MaxResults
}

// Issue.GetWorklogs
// User.GetOpenIssues function (and someway to get recently closed issues?)
// GetWorklogs function
// Issue.getWorklogs function

func (c Client) requestBuilder(location string) *http.Request {
	url := c.hostname + location
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Accept", "application/json")
	req.SetBasicAuth(c.username, c.password)

	return req
}

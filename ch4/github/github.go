package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// IssuesURL is the starting address for api get
// Exported type should have comment or unexported
// 暴露给别的包的东西必须给注释
const IssuesURL = "https://api.github.com/search/issues"

// IssuesSearchResult is the return total count + issue
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}


//Issue contains all return values
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

// User contains Login and url
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

//SearchIssues queries the Github issue tracker
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q) // Here I missed a "="
	if err != nil {
		return nil, err
	}

	// We must close resp.Body on all execution paths
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

// Exercis e 4.11: Build a tool that lets users cre ate, read, update, and delete GitHu b issues fro m
// the command line, inv oking their preferred text editor when subst ant ial text input is required.
package issuesutil

import "time"

const SearchIssueURL = "https://api.github.com/search/issues"
const ReadIssueURL = "https://api.github.com/repos/:owner/:repo/issues/:number"
const CreateIssueURL = "https://api.github.com/repos/:owner/:repo/issues"
const UpdateIssueURL = "https://api.github.com/repos/:owner/:repo/issues/:number"

type IssueBodyRequest struct {
	Title     string
	Body      string
	assignee  string
	Milestone int
	State     string
	Labels    []string
	Assignees []string
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at`
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

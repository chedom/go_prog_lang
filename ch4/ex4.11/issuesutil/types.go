// Exercis e 4.11: Build a tool that lets users cre ate, read, update, and delete GitHu b issues fro m
// the command line, inv oking their preferred text editor when subst ant ial text input is required.
package issuesutil

import "time"

const SearchIssueURL = "https://api.github.com/search/issues"
const CreateIssueURL = "https://api.github.com/repos/:owner/:repo/issues"
const ReadIssueURL = "https://api.github.com/repos/:owner/:repo/issues/:number"
const UpdateIssueURL = "https://api.github.com/repos/:owner/:repo/issues/:number"
const DefaultEditor = "vim"

type IssueBodyRequest struct {
	Title     string   `json:"title"`
	Body      string   `json:"body,omitempty"`
	Milestone int      `json:"milestone,omitempty"`
	Labels    []string `json:"labels,omitempty"`
	Assignees []string `json:"assignees,omitempty"`
	State     string   `json:"state,omitempty"`
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

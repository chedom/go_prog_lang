package main

import "time"

const ISSUES_URL = "https://api.github.com/repos/%s/%s/issues"

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type IssueCache struct {
	Dict  map[int]*Issue
	Items []*Issue
}

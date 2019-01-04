package main

import (
	"flag"

	"github.com/chedom/go_prog_lang/ch4/ex4.11/issuesutil"
)

func main() {
	var operation string
	var token, owner, repo, issueNumber string
	flag.StringVar(&operation, "op", "create", "type of operation")
	flag.StringVar(&token, "t", "", "token for authorization")
	flag.StringVar(&owner, "o", "", "owner of repository")
	flag.StringVar(&repo, "r", "", "name of repository")
	flag.StringVar(&issueNumber, "i", "", "issue id")
	flag.Parse()
	switch operation {
	case "create":
		issuesutil.CreateAnIssue(token, owner, repo)
	case "read":
		issuesutil.ReadAnIssue(token, owner, repo, issueNumber)
	case "update":
		issuesutil.UpdateAnIssue(token, owner, repo, issueNumber)
	case "delete":
		issuesutil.DeleteAnIssue(token, owner, repo, issueNumber)
	}
}

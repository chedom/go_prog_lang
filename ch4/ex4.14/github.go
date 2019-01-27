package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Get resourse failed: %s", resp.Status)
	}
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return res, nil
}

func getIssues(owner, repo string) ([]*Issue, error) {
	resp, err := get(fmt.Sprintf(ISSUES_URL, owner, repo))
	var result []*Issue

	if err = json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func createIssueCache(issues []*Issue) IssueCache {
	dict := make(map[int]*Issue)
	for _, i := range issues {
		dict[i.Number] = i
	}
	cache := IssueCache{
		Dict:  dict,
		Items: issues,
	}
	return cache
}

func createIssueLink(n int) string {
	return fmt.Sprintf("/issues/%d", n)
}

var issueList = template.Must(template.New("issueslist").
	Funcs(template.FuncMap{"issueLink": createIssueLink}).
	Parse(`
<h1>{{.Items | len}} issues</h1>
<table>
<tr style='text-align:left'>
	<th>#</th>
	<th>State</th>
	<th>User</th>
	<th>Title</th>
</tr>
{{range .Items}}
<tr>
	<td><a href='{{.Number | issueLink}}'>{{.Number}}</a></td>
	<td>{{.State}}</td>
	<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
	<td><a href='{{.Number | issueLink}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

var issueTemplate = template.Must(template.New("issue").Parse(`
<h1>{{.Title}}</h1>
<dl>
	<dt>user</dt>
	<dd><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></dd>
	<dt>state</dt>
	<dd>{{.State}}</dd>
</dl>
<p>{{.Body}}</p>
`))

func main() {
	owner := os.Args[1]
	repo := os.Args[2]
	issues, err := getIssues(owner, repo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cant get issues")
	}
	cache := createIssueCache(issues)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) < 3 || parts[2] == "" {
			issueList.Execute(w, cache)
			return
		}
		num, err := strconv.Atoi(parts[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cant parse issue url: %s", parts[2])
			return
		}
		issue, ok := cache.Dict[num]
		if !ok {
			fmt.Fprintf(os.Stderr, "There isnt issue: %d", num)
		}
		issueTemplate.Execute(w, issue)
	})
	http.ListenAndServe(":8000", nil)

}

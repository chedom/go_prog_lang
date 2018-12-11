package issuesutil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func replaceUrlParts(str string, replacement map[string]string) string {
	for k, v := range replacement {
		pattern := ":" + k
		str = strings.Replace(str, pattern, v, -1)
	}
	return str
}

func GetIssue(owner, repo, number string) (*Issue, error) {
	replacement := map[string]string{
		"owner":  owner,
		"repo":   repo,
		"number": number,
	}
	url := replaceUrlParts(ReadIssueURL, replacement)
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get issue failed %s", resp.Status)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func CreateAnIssue(owner, repo string, body *IssueBodyRequest) (*Issue, error) {
	replacement := map[string]string{
		"owner": owner,
		"repo":  repo,
	}
	url := replaceUrlParts(CreateIssueURL, replacement)
	serialized, _ := json.Marshal(body)
	resp, err := http.Post(url, "application/json", bytes.NewReader(serialized))
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get issue failed %s", resp.Status)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

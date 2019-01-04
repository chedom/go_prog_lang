package issuesutil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func replaceUrlParts(str string, replacement map[string]string) string {
	for k, v := range replacement {
		pattern := ":" + k
		str = strings.Replace(str, pattern, v, -1)
	}
	return str
}

func editorCmd(filename string) *exec.Cmd {
	editorPath := os.Getenv("EDITOR")
	if editorPath == "" {
		editorPath = DefaultEditor
	}
	editor := exec.Command(editorPath, filename)

	editor.Stdin = os.Stdin
	editor.Stdout = os.Stdout
	editor.Stderr = os.Stderr

	return editor
}

func getJSONFromFile(filename string) (*IssueBodyRequest, error) {
	r, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	var result IssueBodyRequest
	if err = json.NewDecoder(r).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func getIssueFromFile() (*IssueBodyRequest, error) {
	tmpfile, err := ioutil.TempFile("", "example.*.json")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer os.Remove(tmpfile.Name()) // clean up

	cmd := editorCmd(tmpfile.Name())

	err = cmd.Run()
	if err != nil {
		return nil, err
	}
	result, err := getJSONFromFile(tmpfile.Name())
	if err != nil {
		return nil, err
	}
	return result, err
}

func makeRequest(url, httpMethod, token string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(httpMethod, url, body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	if body != nil {
		req.Header.Add("Content-Type", "application/json")
	}
	req.Header.Add("Authorization", token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 300 {
		return nil, fmt.Errorf("request failed: %s", resp.Status)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("can't parse Body of response")
	}
	return b, nil
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

func CreateAnIssue(token, owner, repo string) {
	issue, err := getIssueFromFile()
	if err != nil {
		log.Fatal(err)
	}
	replacement := map[string]string{
		"owner": owner,
		"repo":  repo,
	}
	url := replaceUrlParts(CreateIssueURL, replacement)
	fmt.Println("url", url)
	serialized, _ := json.Marshal(issue)
	fmt.Println("string serialized", string(serialized))
	resp, err := makeRequest(url, http.MethodPost, token, bytes.NewReader(serialized))
	if err != nil {
		log.Fatal(err)
	}
	var result Issue
	if err = json.NewDecoder(bytes.NewReader(resp)).Decode(&result); err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.Number)
}

func ReadAnIssue(token, owner, repo, number string) {
	replacement := map[string]string{
		"owner":  owner,
		"repo":   repo,
		"number": number,
	}
	url := replaceUrlParts(ReadIssueURL, replacement)
	resp, err := makeRequest(url, http.MethodGet, token, nil)
	if err != nil {
		log.Fatal(err)
	}
	var result Issue
	if err = json.NewDecoder(bytes.NewReader(resp)).Decode(&result); err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

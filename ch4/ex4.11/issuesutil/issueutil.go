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

func createIssueFromFile() (*IssueBodyRequest, error) {
	tmpFile, err := ioutil.TempFile("", "example.*.json")
	defer os.Remove(tmpFile.Name()) // clean up
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	// Close the file
	if err = tmpFile.Close(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return getIssueFromFile(tmpFile)
}

func updateIssueFromFile(issue []byte) (*IssueBodyRequest, error) {
	tmpFile, err := ioutil.TempFile("", "example.*.json")
	defer os.Remove(tmpFile.Name()) // clean up
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	_, err = tmpFile.Write(issue)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	// Close the file
	if err = tmpFile.Close(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return getIssueFromFile(tmpFile)
}

func getIssueFromFile(tmpfile *os.File) (*IssueBodyRequest, error) {
	cmd := editorCmd(tmpfile.Name())

	err := cmd.Run()
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
	req.Header.Add("Authorization", fmt.Sprintf("token %s", token))
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

func CreateAnIssue(token, owner, repo string) ([]byte, error) {
	issue, err := createIssueFromFile()
	if err != nil {
		log.Fatal(err)
		return nil, err
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
		return nil, err
	}
	var result Issue
	if err = json.NewDecoder(bytes.NewReader(resp)).Decode(&result); err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.Number)
	return resp, nil
}

func ReadAnIssue(token, owner, repo, number string) ([]byte, error) {
	replacement := map[string]string{
		"owner":  owner,
		"repo":   repo,
		"number": number,
	}
	url := replaceUrlParts(ReadIssueURL, replacement)
	resp, err := makeRequest(url, http.MethodGet, token, nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var result Issue
	if err = json.NewDecoder(bytes.NewReader(resp)).Decode(&result); err != nil {
		log.Fatal(err)
		return nil, err
	}
	fmt.Println(result)
	return resp, nil
}

func UpdateAnIssue(token, owner, repo, number string) ([]byte, error) {
	replacement := map[string]string{
		"owner":  owner,
		"repo":   repo,
		"number": number,
	}
	rawIssue, err := ReadAnIssue(token, owner, repo, number)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	issue, err := updateIssueFromFile(rawIssue)
	if err != nil {
		log.Fatal(err)
	}
	url := replaceUrlParts(UpdateIssueURL, replacement)
	serialized, _ := json.Marshal(issue)
	resp, err := makeRequest(url, http.MethodPatch, token, bytes.NewReader(serialized))
	if err != nil {
		log.Fatal(err)
	}
	var updatedIssue Issue
	if err = json.NewDecoder(bytes.NewReader(resp)).Decode(&updatedIssue); err != nil {
		log.Fatal(err)
	}
	fmt.Println(updatedIssue)
	return resp, nil
}

func DeleteAnIssue(token, owner, repo, number string) ([]byte, error) {
	replacement := map[string]string{
		"owner":  owner,
		"repo":   repo,
		"number": number,
	}
	url := replaceUrlParts(UpdateIssueURL, replacement)
	body := IssueBodyRequest{State: "closed", Title: "Closed"}
	serialized, _ := json.Marshal(body)
	resp, err := makeRequest(url, http.MethodPatch, token, bytes.NewReader(serialized))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	fmt.Println("url", url, body)
	var result Issue
	if err = json.NewDecoder(bytes.NewReader(resp)).Decode(&result); err != nil {
		log.Fatal(err)
		return nil, err
	}
	fmt.Println(string(resp))
	return resp, nil
}

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func getInitialPks(pks []string) []string {
	args := []string{"list", "-f", "'{{.ImportPath}}'"}

	for _, v := range pks {
		args = append(args, v)
	}

	out, err := exec.Command("go", args...).Output()
	if err != nil {
		log.Fatal(err)
	}

	return strings.Fields(string(out))
}
func getAncestors(depends []string) []string {
	result := make([]string, 0)
	dependsMap := make(map[string]struct{}, len(depends))
	for _, v := range depends {
		dependsMap[v] = struct{}{}
	}

	args := []string{"list", `-f='{{.ImportPath}} {{join .Deps " "}}'`, "..."}
	out, err := exec.Command("go", args...).Output()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(bytes.NewReader(out))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		res := strings.Fields(scanner.Text())
		ances := res[0]

		for _, v := range res[1:] {
			if _, ok := dependsMap[v]; !ok {
				continue
			}
			result = append(result, ances)
			break

		}

	}

	return result
}

func main() {
	pks := os.Args[1:]
	res := getAncestors(getInitialPks(pks))
	fmt.Println("Res", res)
}

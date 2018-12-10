package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/chedom/go_prog_lang/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	var lessThenMonth, lessThenYear, pastThenYear []*github.Issue

	if err != nil {
		log.Fatal(err)
	}
	now := time.Now()
	for _, v := range result.Items {
		switch h := now.Sub(v.CreatedAt).Hours(); {
		case h < 24*31:
			lessThenMonth = append(lessThenMonth, v)
		case h < 24*31:
			lessThenYear = append(lessThenYear, v)
		default:
			pastThenYear = append(pastThenYear, v)
		}
	}

	fmt.Println("Less then month")
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}

	fmt.Println("Less then year")
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}

	fmt.Println("Past then year")
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}

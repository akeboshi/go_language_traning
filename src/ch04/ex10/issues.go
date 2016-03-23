package main

import (
	"ch04/ex10/github"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("After 1 year:")
	printBeforeTime(result.Items, getDateOfLastYear())
	fmt.Println("before 1 year:")
	printAfterTime(result.Items, getDateOfLastYear())
	fmt.Println("before 1 month:")
	printAfterTime(result.Items, getDateOfLastMonth())
}

func printAfterTime(items []*github.Issue, t time.Time) {
	for _, item := range items {
		if item.CreatedAt.After(t) {
			fmt.Printf("#%-5d %9.9s %.55s %s\n",
				item.Number, item.User.Login,
				item.Title, timeToDateString(item.CreatedAt))
		}
	}
}

func printBeforeTime(items []*github.Issue, t time.Time) {
	for _, item := range items {
		if item.CreatedAt.Before(t) {
			fmt.Printf("#%-5d %9.9s %.55s %s\n",
				item.Number, item.User.Login,
				item.Title, timeToDateString(item.CreatedAt))
		}
	}
}

func timeToDateString(t time.Time) string {
	return fmt.Sprintf("%04d-%02d-%02d", t.Year(), t.Month(), t.Day())
}

func getDateOfLastMonth() time.Time {
	t := time.Now()
	return time.Date(t.Year(), t.Month()-1, t.Day(), 0, 0, 0, 0, t.Location())
}

func getDateOfLastYear() time.Time {
	t := time.Now()
	return time.Date(t.Year()-1, t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

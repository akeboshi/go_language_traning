package main

import (
	"bufio"
	"ch04/ex11/github"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	issueFunc := map[string]func(){"c": create, "u": update, "r": read, "d": delete}
	issueFunc[selectFeature()]()
}

func create() {
	var issue github.IssueRequest
	issue.Title = editLine("Title")
	issue.Body = editMessage("vim", "Body")
	if issue.Title == "" || issue.Body == "" {
		println("TitleとBodyを入力しておいてください")
		os.Exit(1)
	}
	result := github.CreateIssue(issue)
	fmt.Println(result.HTMLURL)
}

func update() {
	var issue github.IssueRequest
	num, err := strconv.Atoi(editLine("Number of Issue"))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	issue.Title = editLine("Title")
	issue.Body = editMessage("vim", "Body")
	if issue.Title == "" || issue.Body == "" {
		println("TitleとBodyを入力しておいてください")
		os.Exit(1)
	}
	issue.State = editLine("State(open or closed)")
	result := github.UpdateIssue(num, issue)
	fmt.Println(result.HTMLURL)
}

func read() {
	println("not supported yet")
}

func delete() {
	println("not supported yet")
}

func printIssue(issue *github.IssuesSearchResult) {

}

func selectFeature() string {
	var f [4]*bool
	flagName := [4]string{"c", "r", "u", "d"}
	f[0] = flag.Bool(flagName[0], false, "create issue")
	f[1] = flag.Bool(flagName[1], false, "read issue")
	f[2] = flag.Bool(flagName[2], false, "update issue")
	f[3] = flag.Bool(flagName[3], false, "close issue")
	flag.Parse()

	flagNum := 0
	var selected string
	for i, v := range f {
		if *v {
			flagNum++
			selected = flagName[i]
		}
	}
	if flagNum != 1 {
		flag.Usage()
		fmt.Fprintln(os.Stderr, "オプションは一つだけ選んでほしい")
		os.Exit(1)
	}
	return selected
}

func editLine(message string) string {
	fmt.Printf("%s を入力してください\n", message)
	scanner := bufio.NewScanner(os.Stdin)
	ok := scanner.Scan()
	if !ok {
		println("sine")
	}
	return scanner.Text()
}

func editMessage(editor string, message string) string {
	fmt.Printf("%s を入力してください\n", message)
	messageFile := ".ISSUES_MESSAGE"
	cmd := exec.Command(editor, messageFile)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		fmt.Printf("can't exec %s\n", editor)
		os.Exit(1)
	}

	fileBody, err := ioutil.ReadFile(messageFile)
	if err != nil {
		fmt.Println(err)
		fmt.Println("can't read your message.")
		os.Exit(1)
	}

	err = os.Remove(messageFile)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("cant remove buffer file.\n plz remove %s.\n", messageFile)
		os.Exit(1)
	}

	return string(fileBody)
}

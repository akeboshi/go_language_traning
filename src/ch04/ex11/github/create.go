//Copyright (c) 2016 by akeboshi. All Rights Reserved.
package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func CreateIssue(issue IssueRequest) Issue {
	input, err := json.Marshal(issue)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	req, err := http.NewRequest("POST", IssueURL, bytes.NewBuffer(input))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	req.SetBasicAuth("akeboshi", "")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if resp.StatusCode != http.StatusCreated {
		fmt.Printf("%s", body)
		os.Exit(1)
	}

	output := Issue{}
	err = json.Unmarshal(body, &output)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return output
}

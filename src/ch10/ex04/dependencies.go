// Copyright (c) 2016 by akeboshi. All Rights Reserved.

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
)

type List struct {
	Dir         string
	ImportPath  string
	Name        string
	Target      string
	Root        string
	GoFiles     []string
	Imports     []string
	Deps        []string
	TestGoFiles []string
	TestImports []string
}

func goList(packageNames []string) ([]List, error) {
	args := append([]string{"list", "-e", "-json"}, packageNames...)
	out, err := exec.Command("go", args...).Output()
	if err != nil {
		return nil, err
	}

	dec := json.NewDecoder(bytes.NewReader(out))
	lists := []List{}
	for {
		var list List
		err = dec.Decode(&list)
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		lists = append(lists, list)
	}
	return lists, nil
}

func main() {
	if len(os.Args) < 1 {
		fmt.Println("you need to set package name in args.")
		os.Exit(1)
	}

	needPackages, err := goList(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	workspaceLists, err := goList([]string{"..."})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	result := []List{}

	for _, wList := range workspaceLists {
		appendFlag := true

		for _, p := range needPackages {
			found := false
			for _, dep := range wList.Deps {

				if p.ImportPath == dep {
					found = true
				}
			}
			if !found {
				appendFlag = false
				break
			}
		}
		if appendFlag {
			result = append(result, wList)
		}
	}

	for _, r := range result {
		fmt.Println(r.Dir)
	}
}

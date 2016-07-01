//Copyright (c) 2016 by akeboshi. All Rights Reserved.
package xkcd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const xkcdURL = "https://xkcd.com/%d/info.0.json"
const saveDir = "/tmp/golang/"
const comicFile = saveDir + "%d"
const indexFile = saveDir + "index.json"

var searcher []XKCDIndex

func FetchId(num int) XKCD {
	if len(searcher) == 0 {
		searcher = readIndex()
	}
	if exist(num) {
		return readComicId(num)
	} else {
		return fetchFromInternetID(num)
	}
}

func CacheComics(begin, end int) {
	// TODO: goroutine
	for i := begin; i < end; i++ {
		FetchId(i)
	}
}

func exist(num int) bool {
	_, err := os.Stat(fmt.Sprintf(comicFile, num))
	return err == nil
}

func fetchFromInternetID(num int) XKCD {
	url := fmt.Sprintf(xkcdURL, num)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("%s", body)
		return XKCD{}
	}

	output := XKCD{}
	err = json.Unmarshal(body, &output)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	saveComic(output)
	updateIndex(output)
	return output
}

func SearchFromIndex(str string) []XKCD {
	if len(searcher) == 0 {
		searcher = readIndex()
	}
	// TODO: go routine
	var result []XKCD
	for _, index := range searcher {
		if strings.Contains(index.Search, str) {
			result = append(result, readComicId(index.Id))
		}
	}
	return result
}

func saveComic(xkcd XKCD) bool {
	os.Mkdir(saveDir, os.ModePerm)
	content, err := json.Marshal(xkcd)
	if err != nil {
		println(err.Error())
		return false
	}
	file := fmt.Sprintf(comicFile, xkcd.Num)
	err = ioutil.WriteFile(file, content, os.ModePerm)
	if err != nil {
		println(err.Error())
		return false
	}
	return true
}

func readComicId(id int) XKCD {
	str, err := ioutil.ReadFile(fmt.Sprintf(comicFile, id))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	output := XKCD{}
	err = json.Unmarshal(str, &output)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return output
}

func updateIndex(index XKCD) bool {
	searcher = append(searcher, XKCD2Index(index))
	os.Mkdir(saveDir, os.ModePerm)
	_, err := os.Stat(indexFile)
	if err == nil {
		err := os.Remove(indexFile)
		if err != nil {
			println(err.Error())
			return false
		}
	}
	content, err := json.Marshal(searcher)
	if err != nil {
		println(err.Error())
		return false
	}
	err = ioutil.WriteFile(indexFile, content, os.ModePerm)
	if err != nil {
		println(err.Error())
		return false
	}
	return true
}

func readIndex() []XKCDIndex {
	str, err := ioutil.ReadFile(indexFile)
	if err != nil {
		return []XKCDIndex{}
	}
	output := []XKCDIndex{}
	err = json.Unmarshal(str, &output)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return output
}

func XKCD2Index(data XKCD) XKCDIndex {
	var index XKCDIndex
	index.Id = data.Num
	index.Search = data.News + " " + data.Year + " " +
		data.News + " " + data.SafeTitle + " " +
		data.Title + " " + data.Transcript
	return index
}

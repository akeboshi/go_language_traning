package omdb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const searchURL = "http://www.omdbapi.com/?t=%s&y=&plot=short&r=json"

func SearchMovie(str string) Movie {
	var output Movie
	resp, err := http.Get(fmt.Sprintf(searchURL, str))
	if err != nil {
		println(err.Error())
		return output
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		println(err.Error())
		return output
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("%s\n", body)
		return output
	}

	err = json.Unmarshal(body, &output)
	if err != nil {
		println(err.Error())
		return output
	}
	return output
}

func WritePoster(url string, filename string) bool {
	resp, err := http.Get(url)
	if err != nil {
		println(err.Error())
		return false
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		println(err.Error())
		return false
	}

	if resp.StatusCode != http.StatusOK {
		println(body)
		return false
	}

	err = ioutil.WriteFile("/tmp/"+filename, body, os.ModePerm)
	if err != nil {
		println(err.Error())
		return false
	}
	return true
}

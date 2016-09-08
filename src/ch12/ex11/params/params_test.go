// Copyright (c) 2016 by akeboshi. All Rights Reserved.
package params

import (
	"net/http"
	"net/url"
	"testing"
)

func TestPack(t *testing.T) {
	data := struct {
		Foo string `http:"bar"`
		Etc string
	}{"hoge", "hogehoge"}

	actual := Pack(data)
	if actual != "bar=hoge&etc=hogehoge" {
		t.Errorf("want bar=hoge&etc=hogehoge, actual %v ", actual)
	}
}

func TestUnpack(t *testing.T) {
	var data struct {
		Tell string `http:"tell" check:"tell"`
		Etc  string
	}
	testData, _ := url.ParseQuery("tell=abc&etc=bar")
	req := http.Request{}
	req.Form = testData
	err := Unpack(&req, &data)
	if err == nil {
		t.Errorf("cant valid")
	}
	testData, _ = url.ParseQuery("tell=123-456&etc=bar")
	req = http.Request{}
	req.Form = testData
	err = Unpack(&req, &data)
	if err != nil || data.Tell != "123-456" || data.Etc != "bar" {
		t.Errorf("want= %v, actual= %v", "123-456", data.Tell)
		t.Errorf("want= %v, actual= %v", "bar", data.Etc)
	}
}

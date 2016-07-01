//Copyright (c) 2016 by akeboshi. All Rights Reserved.
package xkcd

type XKCD struct {
	Month      string
	Num        int
	Link       string
	Year       string
	News       string
	SafeTitle  string `json:"safe_title"`
	Transcript string
	Alt        string
	Img        string
	Title      string
}

type XKCDIndex struct {
	Id     int
	Search string
}

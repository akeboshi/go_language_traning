package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []xml.StartElement
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			localList := getLocalFromStartElement(stack)
			if containsAttr(stack, os.Args[1:]) {
				fmt.Printf("%s: %s\n", strings.Join(localList, " "), tok)
			}
		}
	}
}

func getLocalFromStartElement(x []xml.StartElement) (ret []string) {
	for _, xx := range x {
		ret = append(ret, xx.Name.Local)
	}
	return ret
}

func containsAttr(x []xml.StartElement, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		foo := fmt.Sprintf("%s", x[0].Name.Local)
		if foo == y[0] {
			y = y[1:]
		} else {
			for _, attr := range x[0].Attr {
				bar := fmt.Sprintf("%s=%s", attr.Name.Local, attr.Value)
				if bar == y[0] {
					y = y[1:]
					break
				}
			}
		}
		x = x[1:]
	}
	return false
}

func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}

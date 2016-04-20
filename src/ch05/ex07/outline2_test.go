package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"

	"golang.org/x/net/html"
)

func ExampleForEachNode() {
	resp, err := ioutil.ReadFile("outline2_test.html")
	if err != nil {
		fmt.Errorf("File read error outline2_test.html")
	}
	in := bufio.NewReader(bytes.NewReader(resp))

	doc, err := html.Parse(in)

	forEachNode(doc, startElement, endElement)

	// Output:
	// <html>
	//   <head>
	//     <title>
	//       title_body
	//     </title>
	//   </head>
	//   <body>
	//     <!-- comment -->
	//     <div id='test'>
	//       div_body
	//     </div>
	//     <div>
	//       <img src='http://foo.com'/>
	//     </div>
	//   </body>
	// </html>
}

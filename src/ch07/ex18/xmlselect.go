package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

type Node interface{} // CharData あるいは *Element

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var root Node
	var stack []*Element
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
			element := Element{
				tok.Name, tok.Attr, nil,
			}
			if root == nil {
				root = &element
			}
			stack = append(stack, &element)
		case xml.EndElement:
			if len(stack) > 1 { // rootではやらない
				child := stack[len(stack)-1]
				parent := stack[len(stack)-2]
				parent.Children = append(parent.Children, child)
			}
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if len(stack) > 0 { // rootではやらない
				parent := stack[len(stack)-1]
				parent.Children = append(parent.Children, CharData(tok))
			}
		}
	}
	printTree(root, 0)
}

func printTree(n Node, width int) {
	switch n := n.(type) {
	case *Element:
		printSpace(width)
		fmt.Printf("<%s", n.Type.Local)
		printAttr(n.Attr)
		fmt.Printf(">\n")
		for _, child := range n.Children {
			printTree(child, width+1)
		}
		printSpace(width)
		fmt.Printf("</%s>\n", n.Type.Local)
	case CharData:
		nn := strings.Split(string(n), "\n")
		for _, nnn := range nn {
			printSpace(width)
			fmt.Println(nnn)
		}
	default:
		panic(fmt.Sprintf("invalid type:%q\n", n))
	}
}

func printAttr(attr []xml.Attr) {
	if len(attr) == 0 {
		return
	}
	for _, a := range attr {
		fmt.Printf(" %s=\"%s\"", a.Name.Local, a.Value)
	}
}

func printSpace(width int) {
	for i := 0; i < width; i++ {
		fmt.Print("  ")
	}
}

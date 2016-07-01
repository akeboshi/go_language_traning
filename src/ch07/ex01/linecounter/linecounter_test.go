//Copyright (c) 2016 by akeboshi. All Rights Reserved.
package linecounter

import (
	"fmt"
	"testing"
)

func TestLineCounter(t *testing.T) {
	var c LineCounter
	word := `hello
  world
  golang`
	c.Write([]byte(word))
	expected := 3
	if int(c) != expected {
		t.Errorf("%s is wrong, word:%s expected:%d actual:%d", "ByteCounter.Write", word, expected, c)
	}
	c = 0
	name := "Dolly"
	fmt.Fprintf(&c, "%s\n%s", word, name)
	word = fmt.Sprintf("%s\n%s", word, name)
	expected = 4
	if int(c) != expected {
		t.Errorf("%s is wrong, word:%s expected:%d actual:%d", "ByteCounter.Write", word, expected, c)
	}
}

//Copyright (c) 2016 by akeboshi. All Rights Reserved.
package wordcounter

import (
	"fmt"
	"testing"
)

func TestLineCounter(t *testing.T) {
	var c WordCounter
	word := `hello night
  world
  golang`
	c.Write([]byte(word))
	expected := 4
	if int(c) != expected {
		t.Errorf("%s is wrong, word:%s expected:%d actual:%d", "WordCounter.Write", word, expected, c)
	}
	c = 0
	name := "Dolly"
	fmt.Fprintf(&c, "%s\n%s", word, name)
	word = fmt.Sprintf("%s\n%s", word, name)
	expected = 5
	if int(c) != 5 {
		t.Errorf("%s is wrong, word:%s expected:%d actual:%d", "WordCounter.Write", word, expected, c)
	}
}

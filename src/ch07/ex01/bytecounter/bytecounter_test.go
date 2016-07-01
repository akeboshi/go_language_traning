//Copyright (c) 2016 by akeboshi. All Rights Reserved.
package bytecounter

import (
	"fmt"
	"testing"
)

func TestByteCounter(t *testing.T) {
	var c ByteCounter
	word := "hello"
	c.Write([]byte(word))
	expected := len(word)
	if int(c) != 5 {
		t.Errorf("%s is wrong, word:%s expected:%d actual:%d", "ByteCounter.Write", word, expected, c)
	}
	c = 0
	name := "Dolly"
	fmt.Fprintf(&c, "%s, %s", word, name)
	word = fmt.Sprintf("%s, %s", word, name)
	expected = (len(word))
	if int(c) != 12 {
		t.Errorf("%s is wrong, word:%s expected:%d actual:%d", "ByteCounter.Write", word, expected, c)
	}
}

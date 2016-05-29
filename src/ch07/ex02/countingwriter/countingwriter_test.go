package countingwriter

import (
	"bytes"
	"fmt"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	bytesBuf := &bytes.Buffer{}
	writer, count := CountingWriter(bytesBuf)
	word := "foo\n"
	fmt.Fprintf(writer, word)
	expectedCount := 4
	if *count != int64(expectedCount) {
		t.Errorf("count error actual:%d, expected:%d", count, expectedCount)
	}
	word = "bar\n"
	fmt.Fprintf(writer, "bar\n")
	expectedCount = 8
	if *count != int64(expectedCount) {
		t.Errorf("count error actual:%d, expected:%d", count, expectedCount)
	}
}

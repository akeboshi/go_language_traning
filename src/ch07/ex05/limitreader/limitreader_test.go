package limitreader

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestLimitReader(t *testing.T) {
	reader := LimitReader(strings.NewReader("foo bar"), 3)
	foo, _ := ioutil.ReadAll(reader)
	if string(foo) != "foo" {
		t.Errorf("LimitReaderError expected:%s actual:%s", "foo", string(foo))
	}
}

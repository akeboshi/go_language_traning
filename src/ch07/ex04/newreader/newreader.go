package newreader

import (
	"io"

	"golang.org/x/net/html"
)

type Reader struct {
	s        string
	i        int64 // current reading index
	prevRune int   // index of previous rune; or < 0
}

func (r *Reader) Read(b []byte) (n int, err error) {
	if r.i >= int64(len(r.s)) {
		return 0, io.EOF
	}
	r.prevRune = -1
	n = copy(b, r.s[r.i:])
	r.i += int64(n)
	return
}

func Parse(s string) (*html.Node, error) {
	return html.Parse(NewReader(s))
}

func NewReader(s string) io.Reader {
	return &Reader{s, 0, -1}
}

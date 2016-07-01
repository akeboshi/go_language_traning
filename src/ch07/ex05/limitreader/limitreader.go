//Copyright (c) 2016 by akeboshi. All Rights Reserved.
package limitreader

import "io"

type LReader struct {
	reader io.Reader
	num    int64
}

func (r *LReader) Read(p []byte) (n int, err error) {
	if r.num <= 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > r.num {
		p = p[0:r.num]
	}
	n, err = r.reader.Read(p)
	r.num -= int64(n)
	return
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &LReader{r, n}
}

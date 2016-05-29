package countingwriter

import "io"

type CWriter struct {
	src     io.Writer
	counter int64
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var dst *CWriter = new(CWriter)
	dst.src = w
	return dst, &dst.counter
}

func (c *CWriter) Write(p []byte) (int, error) {
	count, err := c.src.Write(p)
	c.counter += int64(count)
	return count, err
}

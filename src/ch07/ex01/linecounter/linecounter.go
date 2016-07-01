//Copyright (c) 2016 by akeboshi. All Rights Reserved.
package linecounter

import (
	"bufio"
	"bytes"
)

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanLines)
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		return -1, err
	}
	*c += LineCounter(count)
	return count, nil
}

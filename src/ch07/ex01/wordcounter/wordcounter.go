package wordcounter

import "bufio"

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	num := 0
	for i := 0; i < len(p); {
		len_word, _, err := bufio.ScanWords(p[i:], true)
		if err != nil {
			return -1, err
		}
		i += len_word
		num++
	}
	*c += WordCounter(num)
	return num, nil
}

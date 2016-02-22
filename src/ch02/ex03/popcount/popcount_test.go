package popcount

import (
	"ch02/ex03/popcount"
	"testing"
)

// 101010101010 2730
// 111111111111 4095
// 000000000000 0
// 000000000001 1
// 100000000000 2048

func TestPopcount(t *testing.T) {
	testData := map[uint64]int{2730: 6, 4095: 12, 0: 0, 1: 1, 2048: 1}
	for input, expected := range testData {
		popCountAssert(t, popcount.PopCount, input, expected)
		popCountAssert(t, popcount.PopCountLoopVer, input, expected)
		popCountAssert(t, popcount.PopCountCheck64, input, expected)
		popCountAssert(t, popcount.PopCountClearBit, input, expected)
	}
}

func popCountAssert(t *testing.T, popcount func(uint64) int, input uint64, expected int) {
	actual := popcount(input)
	if actual != expected {
		t.Errorf("popcount got %d\nwant %d\ninput is %d", actual, expected, input)
	}
}

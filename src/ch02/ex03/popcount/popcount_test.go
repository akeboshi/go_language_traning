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
		popCountAssert(t, "PopCount", popcount.PopCount, input, expected)
		popCountAssert(t, "PopCountLoopVer", popcount.PopCountLoopVer, input, expected)
		popCountAssert(t, "PopCountCheck64", popcount.PopCountCheck64, input, expected)
		popCountAssert(t, "PopCountClearBit", popcount.PopCountClearBit, input, expected)
		popCountAssert(t, "PopCountHackers", popcount.PopCountHackers, input, 3)
	}
}

func popCountAssert(t *testing.T, funcName string, popcount func(uint64) int, input uint64, expected int) {
	actual := popcount(input)
	if actual != expected {
		t.Errorf("%s got %d\nwant %d\ninput is %d", funcName, actual, expected, input)
	}
}

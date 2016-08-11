//Copyright (c) 2016 by akeboshi. All Rights Reserved.
package popcount

import (
	"testing"
)

func bench(n int, f func(uint64) int) {
	for i := 0; i < n*1000; i++ {
		f(uint64(i))
	}
}

func BenchmarkHyo10(b *testing.B)         { bench(10, PopCount) }
func BenchmarkHyo100(b *testing.B)        { bench(100, PopCount) }
func BenchmarkHyo1000(b *testing.B)       { bench(1000, PopCount) }
func BenchmarkHyo10000(b *testing.B)      { bench(10000, PopCount) }
func BenchmarkLoopVer10(b *testing.B)     { bench(10, PopCountLoopVer) }
func BenchmarkLoopVer100(b *testing.B)    { bench(100, PopCountLoopVer) }
func BenchmarkLoopVer1000(b *testing.B)   { bench(1000, PopCountLoopVer) }
func BenchmarkLoopVer10000(b *testing.B)  { bench(10000, PopCountLoopVer) }
func BenchmarkCheck6410(b *testing.B)     { bench(10, PopCountCheck64) }
func BenchmarkCheck64100(b *testing.B)    { bench(100, PopCountCheck64) }
func BenchmarkCheck641000(b *testing.B)   { bench(1000, PopCountCheck64) }
func BenchmarkCheck6410000(b *testing.B)  { bench(10000, PopCountCheck64) }
func BenchmarkClearBit10(b *testing.B)    { bench(10, PopCountClearBit) }
func BenchmarkClearBit100(b *testing.B)   { bench(100, PopCountClearBit) }
func BenchmarkClearBit1000(b *testing.B)  { bench(1000, PopCountClearBit) }
func BenchmarkClearBit10000(b *testing.B) { bench(10000, PopCountClearBit) }
func BenchmarkHackers10(b *testing.B)     { bench(10, PopCountHackers) }
func BenchmarkHackers100(b *testing.B)    { bench(100, PopCountHackers) }
func BenchmarkHackers1000(b *testing.B)   { bench(1000, PopCountHackers) }
func BenchmarkHackers10000(b *testing.B)  { bench(10000, PopCountHackers) }

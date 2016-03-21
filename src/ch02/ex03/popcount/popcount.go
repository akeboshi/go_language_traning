package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount は x のポピュレーションカウント(1が設定されているビット数)を返します
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// PopCountLoopVer は for文を用いたPopCountです
func PopCountLoopVer(x uint64) int {
	var sum byte
	for i := 0; i < 8; i++ {
		sum += pc[byte(x>>(uint(i)*8))]
	}
	return int(sum)
}

// PopCountCheck64 は for文で64回ビットチェックするPopCountです
func PopCountCheck64(x uint64) int {
	var sum int
	for i := 0; i < 64; i++ {
		if x&1 == 1 {
			sum++
		}
		x >>= 1
	}
	return sum
}

// PopCountClearBit は最下位ビットをクリアした回数を数えるPopCountです
func PopCountClearBit(x uint64) int {
	sum := 0
	for x > 0 {
		x = x & (x - 1)
		sum++
	}
	return sum
}

func PopCountHackers(bits uint64) int {
	bits = (bits & 0x55555555) + (bits >> 1 & 0x55555555)
	bits = (bits & 0x33333333) + (bits >> 2 & 0x33333333)
	bits = (bits & 0x0f0f0f0f) + (bits >> 4 & 0x0f0f0f0f)
	bits = (bits & 0x00ff00ff) + (bits >> 8 & 0x00ff00ff)
	return int((bits & 0x0000ffff) + (bits >> 16 & 0x0000ffff))
}

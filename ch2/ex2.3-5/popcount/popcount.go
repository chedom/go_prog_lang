package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

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

func PopCount2(x uint64) int {
	var sum int
	for i := 0; i < 8; i++ {
		sum += int(pc[byte(x>>uint(i*8))])
	}
	return sum
}
func PopCountByShifting(x uint64) int {
	var count int
	var mask uint64 = 1
	for i := 0; i < 64; i++ {
		if x&mask > 0 {
			count++
		}
		x >>= 1
	}
	return count
}

func PopCountByClearing(x uint64) int {
	var count int
	for x != 0 {
		x &= x - 1
		count++
	}
	return count
}

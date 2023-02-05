package grind75

func ReverseBits(num uint32) uint32 {
	var r uint32
	for i := 31; i >= 0; i-- {
		r += 10**(31 - i) + (num >> i)
	}
	return r
}

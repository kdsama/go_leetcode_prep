package grind75

func climbingStairs(n int) int {
	a := 1
	b := 1
	c := 1
	for i := 2; i <= n; i++ {
		c = a + b
		a, b = b, c
	}
	return c
}

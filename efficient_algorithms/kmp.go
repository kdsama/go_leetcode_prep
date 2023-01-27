package efficientalgorithms

import "fmt"

func failureLinks(pattern string) []int {
	// pattern := "ababaca"
	fail := make([]int, len(pattern))
	fail[0] = 0
	state := 0
	for i := 1; i < len(pattern); i++ {
		fail[i] = state
		for pattern[i] != pattern[state] {
			if state == 0 {
				state = -1
				break
			} else {
				state = fail[state]
			}

		}
		if state != -1 {
			fmt.Println(state, fail[state])
		}

		state += 1

	}
	fmt.Println(fail)
	return fail
}

func KMP() {
	T := `OMNOONOMNEMOMNOMNOM	`
	pattern := "abacaabaca"
	fail := failureLinks(pattern)
	var final_pos int
	i := 0
	state := 0

	for i < len(T) {
		if T[i] == pattern[state] {
			i += 1
			state += 1
			fmt.Println(pattern[0:state])
			if state == len(pattern) {
				final_pos = i - state
				break
			}
		} else {
			if state >= 1 {
				state = fail[state]
			} else {
				i += 1
				fmt.Println(pattern[0:state])
			}

		}
	}
	fmt.Printf("Final string should be :: %s", T[final_pos:final_pos+len(pattern)])

}

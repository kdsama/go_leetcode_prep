package efficientalgorithms

import "fmt"

func finite_automata() {
	// finite automata
	// finite_automata()
	text := "aabacaababacaa"
	pattern := "ababaca"
	state := 0
	position := 0
	len_pattern := len(pattern)
	// var position int
	for i := 0; i < len(text); i++ {
		fmt.Printf("State is ::: %d\n", state)

		if state == len_pattern-1 {
			fmt.Print("Got it , now time to say goodbye")
			position = i
			break
		}
		if pattern[state] == text[i] {
			state += 1
		} else {
			if pattern[0] == text[i] {
				state = 1
			} else {
				state = 0
			}
		}
	}
	fmt.Printf("Lets print the pattern %s\n", text[position-len_pattern:position+1])
}

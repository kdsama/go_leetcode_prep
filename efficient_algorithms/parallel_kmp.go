package efficientalgorithms

import (
	"fmt"
	"log"
	"sync"
	"time"
)

var waitgroup = sync.WaitGroup{}

func ParallelKMP(block string, pattern string) {

	// we gonna evaluate string matching in parallel
	// we need to just print the index positions if we find a match

	fail := wG_failureLinks(pattern)

	len_pattern := len(pattern)
	len_block := len(block)
	start := time.Now()
	for b := 0; b < len_block/len_pattern; b++ {
		waitgroup.Add(1)
		// how to do the overlap here ?
		end := (b + 2) * len_pattern
		if end > len(block) {
			end = len(block)
		}

		go wG_KMP(block[b*len_pattern:end], fail, pattern)
	}

	waitgroup.Wait()
	elapsed := time.Since(start)
	log.Printf("Time taken %s", elapsed)
}

func wG_failureLinks(pattern string) []int {
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

func wG_KMP(T string, fail []int, pattern string) {

	defer waitgroup.Done()
	var final_pos int
	i := 0
	state := 0

	for i < len(T) {
		if T[i] == pattern[state] {
			i += 1
			state += 1
			// fmt.Println(pattern[0:state])
			if state == len(pattern) {
				final_pos = i - state
				break
			}
		} else {
			if state >= 1 {
				state = fail[state]
			} else {
				i += 1
				// fmt.Println(pattern[0:state])
			}

		}
	}

	if T[final_pos:final_pos+len(pattern)] == pattern {
		fmt.Printf("Found string %s ", T[final_pos:final_pos+len(pattern)])
	}

}

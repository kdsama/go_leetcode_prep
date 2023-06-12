package efficientalgorithms

import (
	"fmt"
	"log"
	"sync"
	"time"
)

var (
	wg         = sync.WaitGroup{}
	numThreads = 4
	fail       = []int{}
	pattern    string
)

func TP_ParallelKMP(block string, patt string) {

	start := time.Now()
	// using channels to do a KMP
	pattern = patt
	fail = tp_failureLinks()

	// a maximum of 1000 requests.
	bf_channels := make(chan string, 1000)

	len_pattern := len(pattern)
	len_block := len(block)

	for i := 0; i < numThreads; i++ {
		go tp_KMP(bf_channels)
	}
	wg.Add(numThreads)
	for b := 0; b < len_block/len_pattern; b++ {

		// how to do the overlap here ?
		end := (b + 2) * len_pattern
		if end > len(block) {
			end = len(block)
		}
		fmt.Println("Sending ", block[b*len_pattern:end])
		bf_channels <- block[b*len_pattern : end]

	}
	close(bf_channels)
	wg.Wait()

	elapsed := time.Since(start)
	log.Printf("Time elapsed %s", elapsed)
}

func tp_failureLinks() []int {
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

	return fail
}

func tp_KMP(bf_channels chan string) {

	defer wg.Done()
	for T := range bf_channels {
		var final_pos int
		// T := string(t)
		// fmt.Println("?????????????", T, t)
		if len(T) < len(pattern) {

			return
		}
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
		fmt.Println(T[final_pos:final_pos+len(pattern)], pattern)
		if T[final_pos:final_pos+len(pattern)] == pattern {
			fmt.Printf("Found string %s \n", T[final_pos:final_pos+len(pattern)])
		}
	}

}

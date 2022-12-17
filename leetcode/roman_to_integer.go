package leetcode

import (
	"fmt"
)

// var mapping []map[string]int{}
var mapping = map[string]int{"M": 1000, "D": 500, "C": 100, "L": 50, "X": 10, "V": 5, "I": 1}
var Romans = []string{"M", "D", "C", "L", "X", "V", "I"}

func Nine() {
	s := "MCMXCIV"
	// {"M": 1000}, {"D": 500}, {"C": 100}, {"L": 50}, {"X": 10}, {"V": 5}, {"I": 1}

	integerResult := Summation(&s, 0)
	fmt.Print(integerResult)
}
func Summation(s *string, pos int) int {

	sum := 0
	// var y string
	// first_one := mapping[string(y[len(y)-1])]
	// val := mapping[string(y[len(y)-1])]
	index := 0
	for index := 1; index < len(*s); index++ {
		first_one := mapping[string((*s)[index-1])]
		junk := 0
		for mapping[string((*s)[index-1])] < mapping[string((*s)[index])] {
			if first_one != mapping[string((*s)[index])] {
				junk -= mapping[string((*s)[index])]
			} else {
				junk += mapping[string((*s)[index])]

			}
			index += 1

		}

		fmt.Println("sum of section is ", first_one)
		sum += mapping[string((*s)[index])] - junk
		fmt.Println("SUM IS ", sum)
	}

	fmt.Println(index, sum)
	// sum += newFunction(y)
	return sum
}

// func newFunction(y string) int {

// 	val := mapping[string(y[len(y)-1])]
// 	first_one := mapping[string(y[len(y)-1])]
// 	for i := len(y) - 2; i > -1; i-- {
// 		if first_one != mapping[string(y[i])] {
// 			val -= mapping[string(y[i])]
// 		} else {
// 			val += mapping[string(y[i])]

// 		}
// 	}
// 	return val
// }

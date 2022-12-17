package leetcode

import "fmt"

// Given two binary strings a and b, return their sum as a binary string.
func Addbb() {
	fmt.Println(AddBinary("11", "1"))
}

// 110010
//    100
//  110110
// 11
//  1
//  0

func AddBinary(a string, b string) string {
	constants := map[string][]string{"0+0+0": []string{"0", "0"},
		"0+0+1": []string{"1", "0"},
		"1+0+0": []string{"1", "0"},
		"1+0+1": []string{"0", "1"},
		"0+1+0": []string{"1", "0"},
		"0+1+1": []string{"0", "1"},
		"1+1+0": []string{"0", "1"},
		"1+1+1": []string{"1", "1"},
	}
	minString := a
	maxString := b
	if len(a) > len(b) {
		minString = b
		maxString = a
	}
	c := "0"
	return_string := ""
	// len_minString = len(minString) - 1
	len_maxString := len(maxString) - 1
	for len_minString := len(minString) - 1; len_minString >= 0; len_minString-- {
		res := constants[string(minString[len_minString])+"+"+string(maxString[len_maxString])+"+"+c]
		// fmt.Println("ADDING ", string(minString[len_minString]), "and", string(maxString[len_maxString]), "with remainder ", c, "gives", res)
		return_string = res[0] + return_string
		c = res[1]
		len_maxString--
	}
	for j := len(maxString) - len(minString) - 1; j >= 0; j-- {

		res := constants[string(maxString[j])+"+0+"+c]
		return_string = res[0] + return_string
		c = res[1]
	}
	if c != "0" {
		return "1" + return_string
	}
	return return_string
}

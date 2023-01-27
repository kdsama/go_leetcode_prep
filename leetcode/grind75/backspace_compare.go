package grind75

import "fmt"

func BackspaceCompare(s string, t string) bool {
	symbol := "#"
	str_s := getExtractedStringInverted(s, symbol)
	str_j := getExtractedStringInverted(t, symbol)
	fmt.Println(str_j, str_s)
	return str_s == str_j

}

func getExtractedStringInverted(s string, symbol string) string {
	var res string
	var skipTimes int

	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) == symbol {
			skipTimes++
		} else if skipTimes > 0 {
			skipTimes--
		} else {
			res += string(s[i])
		}
	}

	return res
}

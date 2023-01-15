package grind75

// Given a string s which consists of lowercase or uppercase letters, return the length of the longest palindrome that can be built with those letters.

// Letters are case sensitive, for example, "Aa" is not considered a palindrome here.

func LongestPalindrome(s string) int {

	// occurance of alphabets as multiples of 2 can be considered
	// one of them can be 1 or 3 to be considered
	// not more than 1 , occuring 1 or 3 times have to be considered.
	// if length is 1 , it is a palindrome
	// upper and lower case are different
	// my mind always go to hash . get all occurances in a hash . Then count all the evne ones and the odd ones.
	if len(s) == 1 {
		return 1
	}
	x := make(map[byte]int)
	odd_found := false
	even_count := 0
	for i := range s {
		x[s[i]] += 1
	}
	if len(x) == 1 {
		return x[s[0]]
	}
	for i := range x {
		if x[i]%2 == 0 {
			even_count += x[i]
		} else {
			odd_found = true
			even_count += x[i] - 1

		}
	}
	if odd_found {
		even_count++
	}

	return even_count

}

// banana

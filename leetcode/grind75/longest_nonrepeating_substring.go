package grind75

func LengthOfLongestSubstring(s string) int {

	if len(s) < 2 {
		return len(s)
	}
	maxLen := 0
	// d v d f
	i := 0

	repeaterObj := map[byte]int{}
	j := 0
	for i < len(s) {
		if _, ok := repeaterObj[s[i]]; ok {

			for {
				delete(repeaterObj, s[j])
				j++
				if s[j-1] == s[i] {
					break
				}
			}

		}
		repeaterObj[s[i]] = 1
		if len(repeaterObj) > maxLen {
			maxLen = len(repeaterObj)
		}
		// The code is not able to pick largest at the end .
		i++

	}

	return maxLen
}

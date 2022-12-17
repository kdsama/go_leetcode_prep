package leetcode

func LongestCommonPrefix(strs []string) string {
	prefix := ""
	// flag := 0
	firstItem := strs[0]
	for i := 0; i < len(firstItem); i++ {

		for j := 0; j < len(strs); j++ {

			if i >= len(strs[j]) {
				return prefix
			} else {
				if strs[j][i] != firstItem[i] {
					return prefix
				}
			}

		}
		prefix += string(firstItem[i])
	}
	return prefix
}

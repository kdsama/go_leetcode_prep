package grind75

func longestCommonPrefix(strs []string) string {

	// flag := 0

	// longest common prefix
	// should be easy to read
	// If possible in < O(mn)
	prefix := strs[0]
	for i := 0; i < len(strs); i++ {
		prefix = findPrefix(strs[i], prefix)

	}
	return prefix
}

func findPrefix(a string, b string) string {
	// b is prefix
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] != b[j] {
			if i < j {
				return string(a[:i])
			}
			return string(b[:j])
		}
		i++
		j++
	}
	if i < j {
		return string(a[:i])
	}
	return string(b[:j])
}

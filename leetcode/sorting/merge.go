package sorting

func MergeSort(el []int) []int {

	// fmt.Println(el)
	if len(el) <= 1 {
		return el

	}
	mid := int(len(el) / 2)
	return merge(MergeSort(el[0:mid]), MergeSort(el[mid:]))
}
func merge(a []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}

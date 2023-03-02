package efficientalgorithms

import "fmt"

func Hoare(listing []int) int {

	i := 0
	counter := 0
	for i < len(listing)-2 {
		j := i + 1
		for j < len(listing)-1 {
			k := j + 1
			for k < len(listing) {
				if listing[i]+listing[j]+listing[k] == 5 {
					fmt.Println(listing[i], listing[j], listing[k])
					counter++
				}
				k++
			}
			j++
		}
		i++
	}
	return counter
}

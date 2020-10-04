package week209

import "sort"

func specialArray(nums []int) int {
	length := len(nums)
	var result = -1
	sort.Ints(nums)
	for i := 0; i <= 1000; i++ {
		var c int = -1
		for count, number := range nums {
			if number >= i {
				c = count
				break
			}
		}
		if c != -1 {
			d := length - c
			if d == i {
				result = i
				break
			}
		}
	}
	return result
}

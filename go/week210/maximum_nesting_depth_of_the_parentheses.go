package week210

func maxDepth(s string) int {
	var maxResult int
	var result int
	for _, item := range s {
		if item == '(' {
			result++
			if result > maxResult {
				maxResult = result
			}
		} else if item == ')' {
			result--
		}
	}
	return maxResult
}

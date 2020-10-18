package week211

func maxLengthBetweenEqualCharacters(s string) int {
	if len(s) == 0 {
		return 0
	}
	j := len(s) - 1
	maxValue := -1
	for i := 0; i <= j; i++ {
		for k := j; k > i; k-- {
			if s[i] == s[k] {
				tmp := k - i - 1
				if maxValue < tmp {
					maxValue = tmp
				}
			}
		}
	}
	return maxValue
}

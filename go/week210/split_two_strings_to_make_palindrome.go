package week210

func isParamdome(a string, left int, right int) bool {
	if len(a) == 0 || len(a) == 1 {
		return true
	}
	for left < right {
		if a[left] != a[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func checkPalindromeFormation(a string, b string) bool {
	lena := len(a)
	if (isParamdome(a, 0, lena-1)) || isParamdome(b, 0, lena-1) {
		return true
	}
	if a[0] == b[lena-1] {
		if check(a, b) {
			return true
		}
	}
	if b[0] == a[lena-1] {
		if check(b, a) {
			return true
		}
	}
	return false
}

func check(s1, s2 string) bool {
	left := 0
	right := len(s2) - 1
	for left < right {
		if s1[left] == s2[right] {
			left++
			right--
		} else {
			return (isParamdome(s1, left, right) || isParamdome(s2, left, right))
		}
	}
	return true
}

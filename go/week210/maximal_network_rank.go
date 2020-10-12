package week210

func Contains(array []int, elem int) bool {
	for _, item := range array {
		if item == elem {
			return true
		}
	}
	return false
}

func maximalNetworkRank(n int, roads [][]int) int {
	edges := make(map[int][]int)
	indeg := make([]int, n)

	for _, road := range roads {
		edges[road[0]] = append(edges[road[0]], road[1])
		edges[road[1]] = append(edges[road[1]], road[0])
		indeg[road[1]]++
		indeg[road[0]]++
	}

	var first = -1
	var second = -2
	for _, item := range indeg {
		if item > first {
			second = first
			first = item
		} else if item > second {
			second = item
		}
	}

	fisrList, secondList := make([]int, 0), make([]int, 0)

	for i := 0; i < n; i++ {
		if indeg[i] == first {
			fisrList = append(fisrList, i)
		} else if indeg[i] == second {
			secondList = append(secondList, i)
		}
	}

	if len(fisrList) == 1 {
		u := fisrList[0]
		for _, item := range secondList {
			if !Contains(edges[u], item) {
				return first + second
			}
		}
		return first + second - 1
	}
	m := len(roads)
	if (len(fisrList)*(len(fisrList)-1))/2 > m {
		return first * 2
	}
	for _, item := range fisrList {
		for _, j := range fisrList {
			if item != j && !Contains(edges[item], j) {
				return first * 2
			}
		}
	}
	return first*2 - 1
}

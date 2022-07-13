func minDif(a, b int) int {
	if b > a {
		return b - a
	}
	return a - b
}

func minimumAbsDifference(arr []int) [][]int {
	sort.SliceStable(arr, func(i, j int) bool { return (arr)[i] < (arr)[j] })
	result := [][]int{}
	guess := minDif(arr[0], arr[1])
	for i := 0; i < len(arr)-1; i++ {
		if guess > minDif(arr[i], arr[i+1]) {
			guess = minDif(arr[i], arr[i+1])
			result = [][]int{}
		}
		if guess == minDif(arr[i], arr[i+1]) {
			result = append(result, []int{arr[i], arr[i+1]})
		}
	}
	return result
}

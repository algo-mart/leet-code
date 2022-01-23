package amazon

func GetRow(rowIndex int) []int {
	result := []int{}
	for i := 0; i <= rowIndex; i++ {
		temp := make([]int, len(result))
		copy(temp, result)
		result = append(result, 1)
		for i := 1; i < len(result)-1; i++ {
			result[i] = temp[i-1] + temp[i]
		}
	}
	return result
}

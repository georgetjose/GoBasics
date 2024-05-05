package arrayslice

func SumAll(numberSlices ...[]int) []int {
	var sumAllSlice []int
	for _, numberSlice := range numberSlices {
		sumAllSlice = append(sumAllSlice, Sum(numberSlice))
	}
	return sumAllSlice
}

func Sum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}

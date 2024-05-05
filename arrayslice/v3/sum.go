package arrayslice

func SumAll(numbers1, numbers2 []int) []int {
	var sum1, sum2 int
	var sumAllValue []int
	for _, number1 := range numbers1 {
		sum1 += number1
	}
	for _, number2 := range numbers2 {
		sum2 += number2
	}
	sumAllValue = append(sumAllValue, sum1, sum2)
	return sumAllValue
}

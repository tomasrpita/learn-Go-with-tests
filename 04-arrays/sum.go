package arrays

func Sum(numbers []int) int {
	total := 0
	for _, i := range numbers {
		total += i
	}
	return total
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

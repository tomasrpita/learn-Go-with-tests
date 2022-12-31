package arrays

// Sum calculates the total from a slice of numbers.
func Sum(numbers []int) int {
	total := 0
	for _, i := range numbers {
		total += i
	}
	return total
}

// SumAll calculates the total from n slices of numbers.
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

// SumAllTails calculates the sums of all but first number given a colletion of slices.
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {

			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums

}

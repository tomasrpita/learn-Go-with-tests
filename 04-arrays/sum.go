package arrays

func Sum(numbers []int) int {
	total := 0
	for _, i := range numbers {
		total += i
	}
	return total
}

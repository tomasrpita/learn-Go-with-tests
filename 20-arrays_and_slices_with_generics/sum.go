package arrays_with_generics

// Sum calculates the total from a slice of numbers.
func Sum(numbers []int) int {
	// total := 0
	// for _, i := range numbers {
	// 	total += i
	// }
	// return total
	add := func(acc, x int) int { return acc + x }
	return Reduce(numbers, add, 0)
}

// SumAll calculates the total from n slices of numbers.
// func SumAll(numbersToSum ...[]int) []int {
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
	// sunAll := func(acc, x []int) []int {
	// 	return append(acc, Sum(...numbers))
	// }
}

// SumAllTails calculates the sums of all but first number given a colletion of slices.
func SumAllTails(numbers ...[]int) []int {
	// var sums []int
	// for _, numbers := range numbers {
	// 	if len(numbers) == 0 {
	// 		sums = append(sums, 0)
	// 	} else {

	// 		tail := numbers[1:]
	// 		sums = append(sums, Sum(tail))
	// 	}
	// }
	// return sums

	sumTail := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		} else {
			tail := x[1:]
			return append(acc, Sum(tail))

		}
	}
	return Reduce(numbers, sumTail, []int{})

}

func Reduce[A, B any](colletion []A, acumulator func(B, A) B, initialValue B) B {
	var result = initialValue
	for _, x := range colletion {
		result = acumulator(result, x)
	}
	return result
}

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func BalanceFor(transactions []Transaction, name string) float64 {
	// var balance float64
	// for _, t := range transactions {
	// 	if t.From == name {
	// 		balance -= t.Sum
	// 	}
	// 	if t.To == name {
	// 		balance += t.Sum
	// 	}
	// }
	// return balance
	adjustBalance := func(currentBalance float64, t Transaction) float64 {
		if t.From == name {
			return currentBalance - t.Sum
		}
		if t.To == name {
			return currentBalance + t.Sum
		}
		return currentBalance
	}
	return Reduce(transactions, adjustBalance, 0.0)
}

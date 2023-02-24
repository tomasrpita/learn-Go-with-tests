package arrays_with_generics

// Sum calculates the total from a slice of numbers.
func Sum(numbers []int) int {
	add := func(acc, x int) int { return acc + x }
	return Reduce(numbers, add, 0)
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
func SumAllTails(numbers ...[]int) []int {

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

func Find[A any](colletion []A, finder func(A) bool) (A, bool) {
	var noFound A
	for _, x := range colletion {
		if finder(x) {
			return x, true
		}
	}
	return noFound, false

}

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func NewTransaction(from, to Account, sum float64) Transaction {
	return Transaction{From: from.Name, To: to.Name, Sum: sum}
}

type Account struct {
	Name    string
	Balance float64
}

func NewBalanceFor(account Account, transactions []Transaction) Account {
	return Reduce(
		transactions,
		applyTransaction,
		account,
	)
}

func applyTransaction(a Account, transaction Transaction) Account {
	if transaction.From == a.Name {
		a.Balance -= transaction.Sum
	}
	if transaction.To == a.Name {
		a.Balance += transaction.Sum
	}
	return a
}

func BalanceFor(transactions []Transaction, name string) float64 {

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

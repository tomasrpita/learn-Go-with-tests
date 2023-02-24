package arrays_with_generics

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("Collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}
		got := Sum(numbers)
		want := 21

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestSumAllfTails(t *testing.T) {
	checkSum := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSum(t, got, want)

	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSum(t, got, want)

	})
}

func TestReduce(t *testing.T) {
	t.Run("multiplication of all element", func(t *testing.T) {
		multiply := func(x, y int) int {
			return x * y
		}

		AssertEqual(t, Reduce([]int{1, 2, 3}, multiply, 1), 6)
	})

	t.Run("concatenate strings", func(t *testing.T) {
		concatenate := func(x, y string) string {
			return x + y
		}

		AssertEqual(t, Reduce([]string{"a", "b", "c"}, concatenate, ""), "abc")
	})

}

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestBadBank(t *testing.T) {
	// transactions := []Transaction{
	// 	{
	// 		From: "Tom",
	// 		To:   "Nela",
	// 		Sum:  100,
	// 	},
	// 	{
	// 		From: "Charlie",
	// 		To:   "Tom",
	// 		Sum:  25,
	// 	},
	// }

	// AssertEqual(t, BalanceFor(transactions, "Nela"), 100)
	// AssertEqual(t, BalanceFor(transactions, "Tom"), -75)
	// AssertEqual(t, BalanceFor(transactions, "Charlie"), -25)

	var (
		nela    = Account{Name: "Nela", Balance: 100}
		tom     = Account{Name: "Tom", Balance: 75}
		charlie = Account{Name: "Charlie", Balance: 200}

		transactions = []Transaction{
			NewTransaction(tom, nela, 100),
			NewTransaction(charlie, tom, 25),
		}
	)

	newBalanceFor := func(account Account) float64 {
		return NewBalanceFor(account, transactions).Balance
	}

	AssertEqual(t, newBalanceFor(nela), 200)
	AssertEqual(t, newBalanceFor(tom), 0)
	AssertEqual(t, newBalanceFor(charlie), 175)
}

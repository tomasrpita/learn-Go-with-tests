package wallet

import (
	"errors"
	"fmt"
)

// ErrInsufficientFunds means a wallet does not have enough Bitcoin to perform a withdraw.
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficiente funds")

// Bitcoin represents a number of Bitcoin.
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet stores the number of Bitcoin someone owns.
type Wallet struct {
	balance Bitcoin
}

// Balance retunrs the number of Bitcoin a wallet has.
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// deposit will add some Bitcoin to a wallet.
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Withdraw substrats some Bitcoin from the wallet, returnig on error if it cannot be performed
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

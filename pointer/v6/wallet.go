package pointer

import (
	"errors"
	"fmt"
)

// Bitcoin is an OriginalType of int
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet is a struct
type Wallet struct {
	balance Bitcoin
}

// Deposit add money to account
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance return balance of this account
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

//Withdraw minus money to account
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return errors.New("oh no")
	}
	w.balance -= amount
	return nil
}

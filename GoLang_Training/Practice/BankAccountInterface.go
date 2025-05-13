package main

import "fmt"

type Account interface {
	Deposit(amount float64)
	Withdraw(amount float64) bool
	GetBalance() float64
}

type SavingsAccount struct {
	Owner   string
	Balance float64
}

func (s *SavingsAccount) Deposit(amount float64) {
	s.Balance += amount
	fmt.Printf("Deposited Amount: %.2f\n", amount)
}

func (s *SavingsAccount) Withdraw(amount float64) bool {
	if amount <= s.Balance {
		s.Balance -= amount
		fmt.Printf("Withdrwal Amount: %.2f\n", amount)
		return true
	} else {
		fmt.Printf("Amount not Withdrwan. Insufficient Balance:\n")
		return false
	}
}
func (s *SavingsAccount) GetBalance() float64 {
	return s.Balance
}

func operateAccount(acc Account) {
	acc.Deposit(1000)
	acc.Withdraw(500)
	acc.Withdraw(700)

	fmt.Println("Get the Final Balance: ", acc.GetBalance())
}

func main() {
	account := &SavingsAccount{
		Owner: "Rahul",
	}
	operateAccount(account)
}

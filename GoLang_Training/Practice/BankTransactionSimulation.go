package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type BankAccount struct {
	balance int
	mutex   sync.Mutex
}

func (b *BankAccount) Deposit(amount int) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.balance += amount
	fmt.Printf("Deposited: %d | Balance: %d\n", amount, b.balance)
}

func (b *BankAccount) Withdraw(amount int) bool {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	if b.balance >= amount {
		b.balance -= amount
		fmt.Printf("Withdrawn: %d | Balance: %d\n", amount, b.balance)
		return true
	}
	fmt.Printf("Withdrawal of %d failed: Insufficient Funds | Balance: %d\n", amount, b.balance)
	return false
}

func Worker(id int, transaction chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		amount := rand.Intn(200) - 100
		fmt.Printf("[Worker %d] Sending transaction: %d\n", id, amount)
		transaction <- amount
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
	}
}

func main() {
	account := &BankAccount{}
	transaction := make(chan int)
	var wg sync.WaitGroup

	go func() {
		for tx := range transaction {
			if tx >= 0 {
				account.Deposit(tx)
			} else {
				account.Withdraw(-tx)
			}
		}
	}()

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go Worker(i, transaction, &wg)
	}
	wg.Wait()
	close(transaction)

	fmt.Println("Final Balance: ", account.balance)
}

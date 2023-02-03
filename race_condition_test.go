package go_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	x := 0

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				x++
			}
		}()
	}

	time.Sleep(4 * time.Second)
	fmt.Println("Counter = ", x)
}

func TestMutexToFixRaceCondition(t *testing.T) {
	x := 0
	var mutex sync.Mutex
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				mutex.Lock()
				x++
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(4 * time.Second)
	fmt.Println("Counter = ", x)
}

/*
Kasus membuat struct dengan RWMutex ketika struct tersebut bakal diakses oleh
beberapa goroutine sekaligus
*/
type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	var balance int = account.Balance
	account.RWMutex.RUnlock()

	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				// fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Println("Total balance:", account.GetBalance())
}

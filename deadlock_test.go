package go_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type UserBalance struct {
	Mutex   sync.Mutex
	Name    string
	Balance int
}

func (userBalance *UserBalance) Lock() {
	userBalance.Mutex.Lock()
}

func (userBalance *UserBalance) Unlock() {
	userBalance.Mutex.Unlock()
}

func (userBalance *UserBalance) Change(amount int) {
	userBalance.Balance += amount
}

func Transfer(userBalance1 *UserBalance, userBalance2 *UserBalance, amount int) {
	userBalance1.Lock()
	fmt.Println("Lock user1", userBalance1.Name)
	userBalance1.Change(-amount)

	time.Sleep(1 * time.Second) // simulasi latensi

	userBalance2.Lock()
	fmt.Println("Lock user2", userBalance2.Name)
	userBalance2.Change(amount)

	time.Sleep(1 * time.Second)

	userBalance1.Unlock()
	userBalance2.Unlock()
}

func TestDeadLock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Danaerys",
		Balance: 1000,
	}

	user2 := UserBalance{
		Name:    "Rhaegar",
		Balance: 1000,
	}

	go Transfer(&user1, &user2, 100)
	go Transfer(&user2, &user1, 100)

	time.Sleep(3 * time.Second)

	fmt.Println("User 1", user1.Name, "balance", user1.Balance)
	fmt.Println("User 2", user2.Name, "balance", user2.Balance)

}

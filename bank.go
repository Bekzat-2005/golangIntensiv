package main

import (
	"errors"
	"fmt"
)

type Account struct {
	owner   string
	balance float64
}

// Создание нового счёта
func NewAccount(owner string, initial float64) (*Account, error) {
	if initial < 0 {
		return nil, errors.New("начальный баланс не может быть отрицательным")
	}
	return &Account{owner: owner, balance: initial}, nil
}

// Пополнение счёта
func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("сумма пополнения должна быть > 0")
	}
	a.balance += amount
	return nil
}

// Снятие средств
func (a *Account) Withdraw(amount float64) error {
	if amount > a.balance {
		return fmt.Errorf("недостаточно средств: %.2f", a.balance)
	}
	a.balance -= amount
	return nil
}

// Проверка баланса
func (a Account) Balance() float64 {
	return a.balance
}

// Вывод информации
func (a Account) Info() {
	fmt.Printf("Владелец: %s | Баланс: %.2f ₸\n", a.owner, a.balance)
}

func main() {
	account, err := NewAccount("Bekzat", 5000)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	account.Info()
	account.Deposit(1500)
	account.Withdraw(2000)
	account.Info()
}

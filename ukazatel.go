package main

import "fmt"

type Account1 interface {
	Deposit(amount float64)
	Balance() float64
}

type BankAccount struct {
	balance float64
}

func (b *BankAccount) Deposit(amount float64) {
	b.balance += amount
}

func (b BankAccount) Balance() float64 {
	return b.balance
}

func main() {
	var acc Account1
	a := BankAccount{balance: 1000}
	acc = &a // передаём указатель!
	acc.Deposit(500)
	fmt.Println(acc.Balance()) // 1500
}

// cherez ukazatel
//type Updater interface {
//	Update()
//}
//
//type User struct {
//	name string
//}
//
//func (u *User) Update() {
//	u.name = "Updated"
//}
//
//func main() {
//	var u Updater
//	user := User{name: "Bekzat"}
//	u = &user // нужно передать указатель!
//	u.Update()
//	fmt.Println(user.name) // Updated
//}

// 2
//type Counter struct {
//	value int
//}
//
//// Копия (изменения не сохранятся)
//func (c Counter) Add1() {
//	c.value++
//}
//
//// Указатель (изменения сохранятся)
//func (c *Counter) Add2() {
//	c.value++
//}
//
//func main() {
//	c := Counter{value: 10}
//	c.Add1()
//	fmt.Println(c.value) // 10 — не изменился
//	c.Add2()
//	fmt.Println(c.value) // 11 — изменился
//}

// 1
//func changeValue(num *int) {
//	*num = *num + 5
//}
//
//func main() {
//	x := 10
//	changeValue(&x)
//	fmt.Println(x) // 15
//}

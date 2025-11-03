package main

import "fmt"

// 1. Создаём собственный тип
type Age int

// 2. Вложенная структура (внутри Student будет Address)
type Address struct {
	city   string
	street string
}

// 3. Основная структура
type Student2 struct {
	name    string
	age     Age
	address Address
}

// 4. Метод для структуры Student
func (s Student2) Info() {
	fmt.Printf("Имя: %s\nВозраст: %d\nГород: %s\nУлица: %s\n",
		s.name, s.age, s.address.city, s.address.street)
}

func main() {
	// 5. Создаём объект и заполняем данные
	student := Student2{
		name: "Бекзат",
		age:  20,
		address: Address{
			city:   "Алматы",
			street: "Абылай хана 55",
		},
	}

	// 6. Вызываем метод
	student.Info()
}

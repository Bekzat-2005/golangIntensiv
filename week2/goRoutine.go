package main

import (
	"fmt"
	"time"
)

func main() {
	go printNumbers() // запуск функции в отдельной горутине
	go printLetters() // ещё одна горутина

	time.Sleep(2 * time.Second) // ждём, чтобы горутины успели завершиться
}

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Число:", i)
		time.Sleep(300 * time.Millisecond)
	}
}

func printLetters() {
	for _, c := range []rune{'A', 'B', 'C', 'D', 'E'} {
		fmt.Println("Буква:", string(c))
		time.Sleep(400 * time.Millisecond)
	}
}

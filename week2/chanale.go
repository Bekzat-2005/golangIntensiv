package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		time.Sleep(5 * time.Second)
		ch <- "Готово!" // отправляем сообщение в канал
	}()

	fmt.Println("Ждём сообщение...")
	msg := <-ch // получаем из канала
	fmt.Println("Получено:", msg)
}

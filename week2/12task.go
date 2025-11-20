package main

import (
	"fmt"
	"sync"
	"time"
)

func workerr(id int, tasks <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range tasks {
		fmt.Printf("👷 Worker %d начал задачу %d\n", id, task)
		time.Sleep(time.Millisecond * 500) // работа
		fmt.Printf("✅ Worker %d завершил задачу %d\n", id, task)
	}
}

func main() {
	tasks := make(chan int, 5) // буфер
	var wg sync.WaitGroup

	// 3 рабочих
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go workerr(i, tasks, &wg)
	}

	// отправляем 10 задач
	for i := 1; i <= 10; i++ {
		tasks <- i
	}

	close(tasks) // больше задач нет

	wg.Wait()
	fmt.Println("🎉 Все задачи обработаны!")
}

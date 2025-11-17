package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Засекаем общее время выполнения всей программы
	start := time.Now()

	// Список сайтов для проверки
	sites := []string{
		"https://google.com",
		"https://github.com",
		"https://golang.org",
		"https://example.com",
		"https://bekzat.com",
	}

	// Создаём канал для результатов
	ch := make(chan string)

	fmt.Println("🔍 Начало проверки сайтов:", start.Format("15:04:05"))

	// Запускаем горутину для каждого сайта
	for _, site := range sites {
		go checkSite(site, ch)
	}

	// Получаем результаты
	for i := 0; i < len(sites); i++ {
		fmt.Println(<-ch)
	}

	end := time.Now()
	fmt.Println("✅ Проверка завершена:", end.Format("15:04:05"))
	fmt.Println("⏱ Общее время выполнения:", time.Since(start))
}

// Проверка сайта
func checkSite(url string, ch chan string) {
	start := time.Now()
	_, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprintf("❌ %s недоступен (%v)", url, err)
		return
	}

	duration := time.Since(start)
	ch <- fmt.Sprintf("✅ %s доступен (время: %v)", url, duration)
}

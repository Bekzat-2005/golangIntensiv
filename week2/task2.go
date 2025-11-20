package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	sites := []string{
		"https://google.com",
		"https://github.com",
		"https://golang.org",
		"https://example.com",
	}

	var wg sync.WaitGroup

	fmt.Println("🔍 Начало проверки:", start.Format("15:04:05"))

	for _, site := range sites {
		wg.Add(1) // Увеличиваем счётчик

		go func(url string) {
			defer wg.Done() // Уменьшить счётчик по завершению
			checkSite(url)
		}(site)
	}

	wg.Wait() // Ждём всех горутин

	fmt.Println("⏱ Общее время:", time.Since(start))
}

func checkSite(url string) {
	start := time.Now()
	_, err := http.Get(url)

	if err != nil {
		fmt.Printf("❌ %s недоступен (%v)\n", url, err)
		return
	}

	fmt.Printf("✅ %s доступен (время: %v)\n", url, time.Since(start))
}

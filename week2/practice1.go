package main

import (
	"fmt"
	"time"
)

func downloadFile(name string, ch chan string) {
	fmt.Println("Начинаю загрузку:", name)
	time.Sleep(1 * time.Second)
	ch <- name + " загружен ✅"
}

func main() {
	ch := make(chan string)

	files := []string{"photo.jpg", "video.mp4", "doc.pdf"}

	for _, file := range files {
		go downloadFile(file, ch)
	}

	for range files {
		fmt.Println(<-ch)
	}
}

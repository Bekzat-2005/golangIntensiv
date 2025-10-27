package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	var name string
	fmt.Println("Name please:")
	fmt.Scan(&name)
	fmt.Println("Welcome ", name)

}

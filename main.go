package main

import (
	"fmt"
)

func main() {

	// 1
	// fmt.Println("Hello World")

	// var name string
	// fmt.Println("Name please:")
	// fmt.Scan(&name)
	// fmt.Println("Welcome ", name)

	// 2
	// var age int = 20
	// fmt.Print(age)
	// var name = "Bekzat"
	// fmt.Print(name)
	// city := "Almaty"
	// fmt.Print(city)
	// const pi = 3.14
	// greet("Bekzat")

	// result := sum(8, 10)
	// fmt.Println(result)

	// resultt, err := divide(10.9, 0)
	// if err != nil {
	// 	fmt.Print("Error", err)
	// } else {
	// 	fmt.Print(resultt)
	// }

	// var a int;
	// var b int;

	// fmt.Scan(&a, &b)
	// fmt.Println(sum(a, b))
	// fmt.Println(raz(a, b))
	// fmt.Println(kob(a, b))
	// fmt.Println(bol(a, b))

	// 3
	// nums := []int{10, 20, 30}
	// nums = append(nums, 100)

	// for index, num := range nums {
	// 	fmt.Println(index, num)
	// }

	var n int
	fmt.Println("san: ")
	fmt.Scan(&n)

	nums := make([]int, 64)

	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	fmt.Println(nums)

	var sum int
	for _, v := range nums {
		sum += v
	}

	average := sum / n
	fmt.Println(average)

}

// func sum (a int,b int)  float64{
// 	return a + b, a - b
// }

// 2
// func divide(a, b float64) (float64, error) {
// 	if b == 0 {
// 		return 0, fmt.Errorf("Nol ne dlitcya")
// 	}
// 	return a / b, nil
// }
// func sum(a int, b int) int {
// 	return a * b
// }
// func greet(name string) {
// 	fmt.Println(name)
// }

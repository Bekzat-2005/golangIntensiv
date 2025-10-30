package main

import "fmt"

type Rectangle struct {
	width, height float64
}

// Метод для площади
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Метод для периметра
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

// Метод для масштабирования
func (r *Rectangle) Scale(factor float64) {
	r.width *= factor
	r.height *= factor
}

func main() {
	var w, h float64
	fmt.Print("Введите ширину и высоту: ")
	fmt.Scan(&w, &h)

	r := Rectangle{width: w, height: h}

	fmt.Println("Площадь:", r.Area())
	fmt.Println("Периметр:", r.Perimeter())

	r.Scale(2)
	fmt.Println("\nПосле увеличения в 2 раза:")
	fmt.Println("Новая площадь:", r.Area())
	fmt.Println("Новый периметр:", r.Perimeter())
}

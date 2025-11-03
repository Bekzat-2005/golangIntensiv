package main

import (
	"fmt"
	_ "fmt"
)

type Student struct {
	name  string
	grade float64
}

func (s Student) Info() {
	fmt.Printf("name: %s | grade: %.1f\n", s.name, s.grade)
}
func addStudent(students []Student, name string, grade float64) []Student {
	newStudent := Student{name: name, grade: grade}
	students = append(students, newStudent)
	return students
}
func averageStudents(students []Student) float64 {
	if len(students) == 0 {
		return 0
	}
	var sum float64
	for _, i := range students {
		sum += i.grade
	}
	return sum / float64(len(students))
}
func findStudent(students []Student, name string) *Student {
	for i := range students {
		if students[i].name == name {
			return &students[i]
		}
	}
	return nil
}

func main() {

	var students []Student

	students = addStudent(students, "Bekzat", 95.2)
	students = addStudent(students, "Aruzhan", 89.4)
	students = addStudent(students, "Dias", 72.8)
	students = addStudent(students, "Asel", 81.6)

	for _, s := range students {
		s.Info()
	}
	fmt.Println(averageStudents(students))
	name := "Aruzhan"
	fmt.Println(findStudent(students, name))

}

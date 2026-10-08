package main

import "fmt"

type Student struct {
	Name        string
	IsMagistr   bool
	InQueue     bool
	QueueNumber int
}

// quickSortStudents - быстрая сортировка студентов по номеру в очереди
func quickSortStudents(students []Student) []Student {
	if len(students) < 2 {
		return students
	}

	midIndex := len(students) / 2
	pivot := students[midIndex]

	var less []Student
	var equal []Student
	var greater []Student

	for _, s := range students {
		if s.QueueNumber < pivot.QueueNumber {
			less = append(less, s)
		} else if s.QueueNumber > pivot.QueueNumber {
			greater = append(greater, s)
		} else {
			equal = append(equal, s)
		}
	}

	result := append(quickSortStudents(less), equal...)
	result = append(result, quickSortStudents(greater)...)

	return result
}

func main() {
	groupData := []Student{
		{Name: "Егорова Мария", IsMagistr: true, InQueue: true, QueueNumber: 150},
		{Name: "Фролова Кристина", IsMagistr: true, InQueue: false, QueueNumber: 0},
		{Name: "Костюченко Тимофей", IsMagistr: false, InQueue: true, QueueNumber: 342},
		{Name: "Портнова Ксения", IsMagistr: true, InQueue: true, QueueNumber: 612},
		{Name: "Кenжаев Рахим", IsMagistr: true, InQueue: true, QueueNumber: 89},
		{Name: "Даньшин Семён", IsMagistr: true, InQueue: true, QueueNumber: 540},
		{Name: "Шихайло Сергей", IsMagistr: false, InQueue: false, QueueNumber: 0},
		{Name: "Бабушкин Александр", IsMagistr: true, InQueue: true, QueueNumber: 221},
		{Name: "Якунин Андрей", IsMagistr: true, InQueue: true, QueueNumber: 431},
		{Name: "Казанцев Александр", IsMagistr: true, InQueue: true, QueueNumber: 641},
		{Name: "Юркин Александр", IsMagistr: true, InQueue: true, QueueNumber: 112},
		{Name: "Панкратова Анна", IsMagistr: false, InQueue: true, QueueNumber: 299},
		{Name: "Колесникова Лариса", IsMagistr: true, InQueue: false, QueueNumber: 0},
	}

	fmt.Println("Исходный список:")
	for _, s := range groupData {
		fmt.Printf("Студент: %-20s | Очередь: %d\n", s.Name, s.QueueNumber)
	}
	fmt.Println()

	sortedData := quickSortStudents(groupData)
	fmt.Println("Отсортированный список:")
	for _, s := range sortedData {
		fmt.Printf("Студент: %-20s | Очередь: %d\n", s.Name, s.QueueNumber)
	}
}

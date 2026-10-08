package main

import "fmt"

type Student struct {
	Name        string
	IsMagistr   bool
	InQueue     bool
	QueueNumber int
}

// quickSortStudents - сортировка студентов по номеру в очереди
func quickSortStudents(students []Student) []Student {
	if len(students) < 2 {
		return students
	}

	pivot := students[0]

	var less []Student
	var greater []Student

	for _, s := range students[1:] {
		if s.QueueNumber <= pivot.QueueNumber {
			less = append(less, s)
		} else {
			greater = append(greater, s)
		}
	}

	result := append(quickSortStudents(less), pivot)
	return append(result, quickSortStudents(greater)...)
}

// findSettledStudents - поиск студентов в утвержденных приказах ИТМО
func findSettledStudents(students []Student, activeOrders []int) []string {
	var settled []string

	for _, s := range students {
		low := 0
		high := len(activeOrders) - 1
		found := false
		for low <= high {
			mid := low + (high-low)/2

			if activeOrders[mid] == s.QueueNumber {
				found = true
				break
			}

			if activeOrders[mid] < s.QueueNumber {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}

		if found && s.QueueNumber != 0 {
			settled = append(settled, s.Name)
		}
	}

	return settled
}

func main() {
	groupData := []Student{
		{Name: "Егорова Мария", IsMagistr: true, InQueue: true, QueueNumber: 150},
		{Name: "Фролова Кристина", IsMagistr: true, InQueue: false, QueueNumber: 0},
		{Name: "Костюченко Тимофей", IsMagistr: false, InQueue: true, QueueNumber: 342},
		{Name: "Портнова Ксения", IsMagistr: true, InQueue: true, QueueNumber: 612},
		{Name: "Кенжаев Рахим", IsMagistr: true, InQueue: true, QueueNumber: 89},
		{Name: "Даньшин Семён", IsMagistr: true, InQueue: true, QueueNumber: 540},
		{Name: "Шихайло Сергей", IsMagistr: false, InQueue: false, QueueNumber: 0},
		{Name: "Бабушкин Александр", IsMagistr: true, InQueue: true, QueueNumber: 221},
		{Name: "Якунин Андрей", IsMagistr: true, InQueue: true, QueueNumber: 431},
		{Name: "Казанцев Александр", IsMagistr: true, InQueue: true, QueueNumber: 641},
		{Name: "Юркин Александр", IsMagistr: true, InQueue: true, QueueNumber: 112},
		{Name: "Панкратова Анна", IsMagistr: false, InQueue: true, QueueNumber: 299},
		{Name: "Колесникова Лариса", IsMagistr: true, InQueue: false, QueueNumber: 0},
	}

	sorted := quickSortStudents(groupData)
	fmt.Println("Студенты, отсортированные по номеру в очереди:")
	for _, s := range sorted {
		if s.InQueue {
			fmt.Printf("- %s: №%d\n", s.Name, s.QueueNumber)
		}
	}
	fmt.Println()

	ordersInViazma := []int{89, 112, 221, 299, 641}

	luckyStudents := findSettledStudents(groupData, ordersInViazma)

	fmt.Println("Студенты, чьи номера найдены в приказах на заселение в Вязьму:")
	for _, name := range luckyStudents {
		fmt.Println("-", name)
	}
}

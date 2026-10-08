package main

import "fmt"

type Student struct {
	Name        string
	IsMagistr   bool
	InQueue     bool
	QueueNumber int
}

// combSortStudents - сортировка расчёской массива студентов по номеру в очереди
func combSortStudents(students []Student) {
	n := len(students)
	gap := n
	// оптимальный коэффицент уменьшения шага
	shrink := 1.3
	// флаг - были ли перестановки на текущей итерации
	swapped := true

	for gap > 1 || swapped {
		gap = int(float64(gap) / shrink)
		if gap < 1 {
			gap = 1
		}

		swapped = false

		for i := 0; i < n-gap; i++ {
			if students[i].QueueNumber > students[i+gap].QueueNumber {
				students[i], students[i+gap] = students[i+gap], students[i]
				swapped = true
			}
		}
	}
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

	combSortStudents(groupData)

	fmt.Println("Отсортированный список:")
	for _, s := range groupData {
		fmt.Printf("Студент: %-20s | Очередь: %d\n", s.Name, s.QueueNumber)
	}
}

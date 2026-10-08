package main

import "fmt"

type Student struct {
	Name        string
	IsMagistr   bool
	InQueue     bool
	QueueNumber int
}

// heapify — перемещает элемент с маленьким приоритетом вниз, если его потомки больше
func heapify(students []Student, root, size int) {
	for {
		left := (root << 1) + 1
		right := left + 1
		largest := root // временно считаю корень самым большим элементом

		if left < size && students[left].QueueNumber > students[largest].QueueNumber {
			largest = left
		}

		if right < size && students[right].QueueNumber > students[largest].QueueNumber {
			largest = right
		}

		if largest == root {
			break
		}

		students[root], students[largest] = students[largest], students[root]

		root = largest
	}
}

// heapSortStudents — приамидальная сортировка
func heapSortStudents(students []Student) {
	n := len(students)

	for i := (n >> 1) - 1; i >= 0; i-- {
		heapify(students, i, n)
	}

	for i := n - 1; i > 0; i-- {
		students[0], students[i] = students[i], students[0]

		heapify(students, 0, i)
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

	heapSortStudents(groupData)

	fmt.Println("Отсортированный список:")
	for _, s := range groupData {
		fmt.Printf("Студент: %-20s | Очередь: %d\n", s.Name, s.QueueNumber)
	}
}

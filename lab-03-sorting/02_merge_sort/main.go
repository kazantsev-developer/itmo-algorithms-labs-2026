package main

import "fmt"

type Student struct {
	Name        string
	IsMagistr   bool
	InQueue     bool
	QueueNumber int
}

// merge объединяет два предварительно отсортированных слайса (left и right) в один общий упорядоченный
func merge(left, right []Student) []Student {
	result := make([]Student, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i].QueueNumber <= right[j].QueueNumber {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	for i < len(left) {
		result = append(result, left[i])
		i++
	}

	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	return result
}

// mergeSortStudents рекурсивно разбиваем массив пополам, пока длина подмассивов не станет меньше 2
func mergeSortStudents(students []Student) []Student {
	if len(students) < 2 {
		return students
	}

	mid := len(students) / 2

	left := mergeSortStudents(students[:mid])
	right := mergeSortStudents(students[mid:])

	return merge(left, right)
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

	sortedData := mergeSortStudents(groupData)

	fmt.Println("Отсортированный список:")
	for _, s := range sortedData {
		fmt.Printf("Студент: %-20s | Очередь: %d\n", s.Name, s.QueueNumber)
	}
}

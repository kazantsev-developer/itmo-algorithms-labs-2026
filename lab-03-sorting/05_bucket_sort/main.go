package main

import "fmt"

type Student struct {
	Name        string
	IsMagistr   bool
	InQueue     bool
	QueueNumber int
}

// bucketSortStudents - блочная сортировка студентов по номерам в очереди
func bucketSortStudents(students []Student) []Student {
	if len(students) == 0 {
		return students
	}

	// корзины под диапазоны номеров очередей
	var bucket1 []Student // категория до 200 включительно
	var bucket2 []Student // от 201 до 500 включительно
	var bucket3 []Student // от 501+

	for _, s := range students {
		if s.QueueNumber <= 200 {
			bucket1 = append(bucket1, s)
		} else if s.QueueNumber <= 500 {
			bucket2 = append(bucket2, s)
		} else {
			bucket3 = append(bucket3, s)
		}
	}

	insertionSort(bucket1)
	insertionSort(bucket2)
	insertionSort(bucket3)

	result := make([]Student, 0, len(students))
	result = append(result, bucket1...)
	result = append(result, bucket2...)
	result = append(result, bucket3...)

	return result
}

func insertionSort(arr []Student) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j].QueueNumber > key.QueueNumber {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
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

	sortedData := bucketSortStudents(groupData)

	fmt.Println("Отсортированный список:")
	for _, s := range sortedData {
		fmt.Printf("Студент: %-20s | Очередь: %d\n", s.Name, s.QueueNumber)
	}
}

package main

import (
	"fmt"
	"strings"
)

type Student struct {
	Name        string
	IsMagistr   bool
	InQueue     bool // true - ждёт общагу, false - уже заселился
	QueueNumber int  // номер в очереди, для заселенных 0
}

// getQueueStats - сбор аналитики по очередникам
func getQueueStats(students []Student) (int, int, int) {
	magistrsInQueue := 0
	totalInQueue := 0
	maxQueueNumber := 0

	for _, s := range students {
		if s.IsMagistr && s.InQueue {
			magistrsInQueue++
		}
	}

	for _, s := range students {
		if s.InQueue {
			totalInQueue++
		}
	}

	for _, s := range students {
		if s.InQueue && s.QueueNumber > maxQueueNumber {
			maxQueueNumber = s.QueueNumber
		}
	}

	return magistrsInQueue, totalInQueue, maxQueueNumber
}

// cleanStudentNames — очистка текстовых данных
func cleanStudentNames(names []string) []string {
	for i := range names {
		names[i] = strings.TrimSpace(names[i])
	}

	for i := range names {
		names[i] = strings.ToLower(names[i])
	}

	var validNames []string

	for i := range names {
		if names[i] != "" {
			validNames = append(validNames, names[i])
		}
	}

	return validNames
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
		{Name: "Казанцев Александр", IsMagistr: true, InQueue: true, QueueNumber: 641}, // Это ты, твоя реальная позиция
		{Name: "Юркин Александр", IsMagistr: true, InQueue: true, QueueNumber: 112},
		{Name: "Панкратова Анна", IsMagistr: false, InQueue: true, QueueNumber: 299},
		{Name: "Колесникова Лариса", IsMagistr: true, InQueue: false, QueueNumber: 0},
	}

	magInQueue, totalInQueue, maxNum := getQueueStats(groupData)

	fmt.Printf("Магистров в очереди: %d\n", magInQueue)
	fmt.Printf("Всего студентов в очереди: %d\n", totalInQueue)
	fmt.Printf("Максимальный номер в очереди потока: %d\n\n", maxNum)

	rawDorms := []string{"  Студгородок Вязьма ", "Общежитие №2 Ленсовета  ", "", "Альпы "}

	cleanedDorms := cleanStudentNames(rawDorms)

	fmt.Println("Результат работы очистки названий:")
	for _, name := range cleanedDorms {
		fmt.Println("-", name)
	}
}

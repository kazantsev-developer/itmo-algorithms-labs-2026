package main

import "fmt"

var studentDB = map[int]string{
	368088: "Даньшин Семён",
	506428: "Егорова Мария",
	563346: "Казанцев Александр",
	354511: "Кенжаев Рахим",
	564670: "Колесникова Лариса",
	563558: "Костюченко Тимофей",
	563102: "Панкратова Анна",
	367496: "Портнова Ксения",
	368978: "Фролова Кристина",
	182717: "Шихайло Сергей",
	367662: "Юркин Александр",
	369105: "Якунин Андрей",
}

// findStudentID - поиск ID студента по трем базам общежитий
func findStudentID(viazma, lensoveta, msg []int, targetID int) string {
	low, high := 0, len(viazma)-1
	for low <= high {
		mid := low + (high-low)/2
		if viazma[mid] == targetID {
			return fmt.Sprintf("%s в списках Вязьмы", studentDB[targetID])
		}
		if viazma[mid] < targetID {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	low, high = 0, len(lensoveta)-1
	for low <= high {
		mid := low + (high-low)/2
		if lensoveta[mid] == targetID {
			return fmt.Sprintf("%s в списках на Ленсовета", studentDB[targetID])
		}
		if lensoveta[mid] < targetID {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	low, high = 0, len(msg)-1
	for low <= high {
		mid := low + (high-low)/2
		if msg[mid] == targetID {
			return fmt.Sprintf("%s в списках Межвузовского студгородка", studentDB[targetID])
		}
		if msg[mid] < targetID {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return "Студент нигде не найден!"
}

// findFreeDate - поиск свободной даты для заселения
func findFreeDate(schedule1, schedule2, schedule3 []int, targetDate int) string {
	low, high := 0, len(schedule1)-1
	for low <= high {
		mid := low + (high-low)/2
		if schedule1[mid] == targetDate {
			return "свободный слот найден в первом графике"
		}
		if schedule1[mid] < targetDate {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	low, high = 0, len(schedule2)-1
	for low <= high {
		mid := low + (high-low)/2
		if schedule2[mid] == targetDate {
			return "свободный слот найден во втором графике"
		}
		if schedule2[mid] < targetDate {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	low, high = 0, len(schedule3)-1
	for low <= high {
		mid := low + (high-low)/2
		if schedule3[mid] == targetDate {
			return "свободный слот найден в третьем графике"
		}
		if schedule3[mid] < targetDate {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return "все слоты уже заняты!"
}

func main() {
	// Массивы отсортированы для бинарного поиска
	viazmaIDs := []int{367496, 368978, 563346, 563558}
	lensovetaIDs := []int{182717, 354511, 367662, 368088}
	msgIDs := []int{369105, 506428, 563102, 564670}

	myID := 563346
	resultDorm := findStudentID(viazmaIDs, lensovetaIDs, msgIDs, myID)
	fmt.Printf("Поиск студента с ID %d: %s\n", myID, resultDorm)

	mariaID := 506428
	resultGirl := findStudentID(viazmaIDs, lensovetaIDs, msgIDs, mariaID)
	fmt.Printf("Поиск студента с ID %d: %s\n\n", mariaID, resultGirl)

	dates1 := []int{1, 3, 5, 7}
	dates2 := []int{10, 12, 14, 15}
	dates3 := []int{20, 22, 25, 28}

	myDay := 27
	mySlot := findFreeDate(dates1, dates2, dates3, myDay)
	fmt.Printf("Поиск свободного слота для Казанцева Александра на %d октября: %s\n", myDay, mySlot)

	girlDay := 14
	girlSlot := findFreeDate(dates1, dates2, dates3, girlDay)
	fmt.Printf("Поиск свободного слота для Егоровой Марии на %d октября: %s\n", girlDay, girlSlot)
}

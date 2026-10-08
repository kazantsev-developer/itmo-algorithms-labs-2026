package main

import "fmt"

type Student struct {
	Name        string
	IsMale      bool // true - парень, false - девушка
	IsMagistr   bool
	InQueue     bool
	QueueNumber int
}

type RoomBlock struct {
	BlockNumber int
	Rooms       [][]bool // true - кровать свободна
}

// findRoomTriplets - подбор уникальных троек студентов для трехместной комнаты
// в комнату могут заселяться студенты только одного пола!
func findRoomTriplets(students []Student, maxCombineQueue int) [][]string {
	var triplets [][]string
	n := len(students)

	for i := range n {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if students[i].InQueue && students[j].InQueue && students[k].InQueue {
					allMales := students[i].IsMale && students[j].IsMale && students[k].IsMale
					allFemales := !students[i].IsMale && !students[j].IsMale && !students[k].IsMale

					if allMales || allFemales {
						sumQueue := students[i].QueueNumber + students[j].QueueNumber + students[k].QueueNumber
						if sumQueue <= maxCombineQueue {
							triplets = append(triplets, []string{students[i].Name, students[j].Name, students[k].Name})
						}
					}
				}
			}
		}
	}

	return triplets
}

// findFreeBeds - поиск свободных коек
func findFreeBeds(blocks []RoomBlock) []string {
	var freeBedsReport []string

	for _, block := range blocks {
		for roomIdx, room := range block.Rooms {
			for bedIdx, isFree := range room {
				if isFree {
					report := fmt.Sprintf("Блок %d, Комната %d, Кровать %d", block.BlockNumber, roomIdx+1, bedIdx+1)
					freeBedsReport = append(freeBedsReport, report)
				}
			}
		}
	}

	return freeBedsReport
}

func main() {
	groupData := []Student{
		{Name: "Егорова Мария", IsMale: false, IsMagistr: true, InQueue: true, QueueNumber: 150},
		{Name: "Фролова Кристина", IsMale: false, IsMagistr: true, InQueue: false, QueueNumber: 0},
		{Name: "Костюченко Тимофей", IsMale: true, IsMagistr: false, InQueue: true, QueueNumber: 342},
		{Name: "Портнова Ксения", IsMale: false, IsMagistr: true, InQueue: true, QueueNumber: 612},
		{Name: "Кенжаев Рахим", IsMale: true, IsMagistr: true, InQueue: true, QueueNumber: 89},
		{Name: "Даньшин Семён", IsMale: true, IsMagistr: true, InQueue: true, QueueNumber: 540},
		{Name: "Шихайло Сергей", IsMale: true, IsMagistr: false, InQueue: false, QueueNumber: 0},
		{Name: "Бабушкин Александр", IsMale: true, IsMagistr: true, InQueue: true, QueueNumber: 221},
		{Name: "Якунин Андрей", IsMale: true, IsMagistr: true, InQueue: true, QueueNumber: 431},
		{Name: "Казанцев Александр", IsMale: true, IsMagistr: true, InQueue: true, QueueNumber: 641},
		{Name: "Юркин Александр", IsMale: true, IsMagistr: true, InQueue: true, QueueNumber: 112},
		{Name: "Панкратова Анна", IsMale: false, IsMagistr: false, InQueue: true, QueueNumber: 299},
		{Name: "Колесникова Лариса", IsMale: false, IsMagistr: true, InQueue: false, QueueNumber: 0},
	}

	fmt.Println("Подобраны тройки студентов для Вязьмы:")
	triplets := findRoomTriplets(groupData, 500)
	for _, t := range triplets {
		fmt.Printf("- %s, %s, %s\n", t[0], t[1], t[2])
	}
	fmt.Println()

	viazmaBlocks := []RoomBlock{
		{
			BlockNumber: 4,
			Rooms: [][]bool{
				{false, false, true},
				{false, false, false},
			},
		},
		{
			BlockNumber: 5,
			Rooms: [][]bool{
				{true, false, true},
			},
		},
	}

	fmt.Println("Свободные места обнаружены:")
	freeBeds := findFreeBeds(viazmaBlocks)
	for _, bed := range freeBeds {
		fmt.Println("-", bed)
	}
}

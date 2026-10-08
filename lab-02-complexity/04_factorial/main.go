package main

import "fmt"

// findRoomPaths - находит все варианты порядка обхода комнат
func findRoomPaths(rooms []string, start int, result *[][]string) {
	if start == len(rooms) {
		temp := make([]string, len(rooms))
		copy(temp, rooms)
		*result = append(*result, temp)
		return
	}

	for i := start; i < len(rooms); i++ {
		rooms[start], rooms[i] = rooms[i], rooms[start]
		findRoomPaths(rooms, start+1, result)
		rooms[start], rooms[i] = rooms[i], rooms[start]
	}
}

// findShortestRoute - поиск кратчайшего пути
func findShortestRoute(
	locations []string,
	startIdx int,
	currentRoute []string,
	visited []bool,
	matrix [][]int,
	currentDist int,
	bestDist *int,
	finalRoute *[]string,
) {
	if len(currentRoute) == len(locations) {
		finalDist := currentDist + matrix[startIdx][0]
		if finalDist < *bestDist {
			*bestDist = finalDist
			*finalRoute = append([]string(nil), currentRoute...)
			*finalRoute = append(*finalRoute, locations[0])
		}
		return
	}

	for i := range locations {
		if !visited[i] {
			visited[i] = true
			findShortestRoute(locations, i, append(currentRoute, locations[i]), visited, matrix, currentDist+matrix[startIdx][i], bestDist, finalRoute)
			visited[i] = false
		}
	}
}

func main() {
	roomsList := []string{"Комната 201", "Комната 202", "Комната 203"}
	var options [][]string
	findRoomPaths(roomsList, 0, &options)

	fmt.Printf("Всего вариантов обхода для %d комнат: %d\n", len(roomsList), len(options))
	fmt.Println("Варианты маршрутов на Ленсовета:")
	for _, p := range options {
		fmt.Printf("- %s -> %s -> %s\n", p[0], p[1], p[2])
	}
	fmt.Println()

	itmoLocations := []string{"Общага Вавиловых", "Корпус Кронверкский", "Корпус Ломоносова", "Корпус Биржевая"}

	distanceMatrix := [][]int{
		{0, 45, 50, 55},
		{45, 0, 25, 20},
		{50, 25, 0, 30},
		{55, 20, 30, 0},
	}

	visited := make([]bool, len(itmoLocations))
	visited[0] = true

	minMinutes := 999999
	var pathResult []string

	findShortestRoute(itmoLocations, 0, []string{itmoLocations[0]}, visited, distanceMatrix, 0, &minMinutes, &pathResult)

	fmt.Printf("Самый короткий путь через метро: %d минут\n", minMinutes)
	fmt.Println("Порядок движения:")
	for i, loc := range pathResult {
		if i > 0 {
			fmt.Print(" -> ")
		}
		fmt.Print(loc)
	}
	fmt.Println()
}

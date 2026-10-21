package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/kazantsev-developer/itmo-algorithms-labs-2026/lab-06-structures/internal/algo"
)

func main() {
	fmt.Println("Задание 2. Поиск пути между объектами ИТМО")

	graph := algo.Graph{
		"Кронверкский":     {"Горьковская", "Петроградская"},
		"Ломоносова":       {"Достоевская", "Владимирская"},
		"Биржа":            {"Спортивная", "Василеостровская"},
		"Чайковского":      {"Чернышевская"},
		"Гривцова":         {"Сенная", "Садовая"},
		"Вязьма":           {"Петроградская"},
		"Альпы":            {"Проспект Славы"},
		"Белка":            {"Ладожская"},
		"Ленсовета":        {"Московская"},
		"Серебристый":      {"Пионерская"},
		"Пионерская":       {"Серебристый", "Петроградская"},
		"Петроградская":    {"Вязьма", "Кронверкский", "Пионерская", "Горьковская"},
		"Горьковская":      {"Кронверкский", "Петроградская", "Невский"},
		"Невский":          {"Горьковская", "Сенная", "Владимирская"},
		"Сенная":           {"Гривцова", "Невский", "Московская", "Садовая", "Спасская"},
		"Московская":       {"Ленсовета", "Сенная"},
		"Спасская":         {"Сенная", "Садовая", "Достоевская"},
		"Достоевская":      {"Ломоносова", "Владимирская", "Спасская", "Ладожская"},
		"Ладожская":        {"Белка", "Достоевская"},
		"Спортивная":       {"Биржа", "Садовая"},
		"Василеостровская": {"Биржа"},
		"Садовая":          {"Гривцова", "Сенная", "Спасская", "Спортивная", "Проспект Славы"},
		"Проспект Славы":   {"Альпы", "Садовая"},
		"Площадь Ленина":   {"Чернышевская"},
		"Чернышевская":     {"Чайковского", "Площадь Ленина", "Владимирская"},
		"Владимирская":     {"Ломоносова", "Достоевская", "Чернышевская", "Невский"},
	}

	campuses := []string{"Кронверкский", "Ломоносова", "Биржа", "Чайковского", "Гривцова"}
	dorms := []string{"Вязьма", "Альпы", "Белка", "Ленсовета", "Серебристый"}

	var stations []string
	for k := range graph {
		isCampus := false
		for _, c := range campuses {
			if k == c {
				isCampus = true
				break
			}
		}
		isDorm := false
		for _, d := range dorms {
			if k == d {
				isDorm = true
				break
			}
		}
		if !isCampus && !isDorm {
			stations = append(stations, k)
		}
	}
	sort.Strings(stations)

	fmt.Println("\nДоступные локации:")
	fmt.Printf("Корпуса ИТМО: %s\n", strings.Join(campuses, ", "))
	fmt.Printf("\nОбщежития: %s\n", strings.Join(dorms, ", "))
	fmt.Printf("\nСтанции метро: %s\n", strings.Join(stations, ", "))

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("\nВведите начальную локацию: ")
	start, _ := reader.ReadString('\n')
	start = strings.TrimSpace(start)

	fmt.Print("Введите конечную локацию: ")
	target, _ := reader.ReadString('\n')
	target = strings.TrimSpace(target)

	if _, ok := graph[start]; !ok {
		fmt.Printf("\nОшибка: '%s' не найдена. Проверьте список доступных локаций.\n", start)
		return
	}
	if _, ok := graph[target]; !ok {
		fmt.Printf("\nОшибка: '%s' не найдена. Проверьте список доступных локаций.\n", target)
		return
	}

	fmt.Printf("\nИщем путь от '%s' до '%s'\n", start, target)

	pathBFS, visitedBFS := algo.BFS(graph, start, target)
	fmt.Println("\nBFS (поиск в ширину):")
	if pathBFS != nil {
		fmt.Printf("  Путь: %s\n", strings.Join(pathBFS, " -> "))
	} else {
		fmt.Println("  Путь не найден")
	}
	fmt.Printf("  Посещено вершин: %d\n", visitedBFS)

	pathDFS, visitedDFS := algo.DFS(graph, start, target)
	fmt.Println("\nDFS (поиск в глубину):")
	if pathDFS != nil {
		fmt.Printf("  Путь: %s\n", strings.Join(pathDFS, " -> "))
	} else {
		fmt.Println("  Путь не найден")
	}
	fmt.Printf("  Посещено вершин: %d\n", visitedDFS)
}

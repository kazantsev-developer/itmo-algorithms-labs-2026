package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/kazantsev-developer/itmo-algorithms-labs-2026/lab-05-dynamics/internal/algo"
)

func main() {
	fmt.Println("Задача 1: Эвакуация ценных вещей")

	in := bufio.NewReader(os.Stdin)

	n := readInt(in, "Введите количество предметов (N): ", func(v int) bool { return v > 0 })
	maxW := readInt(in, "Введите грузоподъемность (W): ", func(v int) bool { return v >= 0 })
	maxT := readInt(in, "Введите лимит времени (T): ", func(v int) bool { return v >= 0 })

	items := make([]algo.Item, n)
	fmt.Printf("Введите %d предметов (вес, ценность, время):\n", n)
	for i := 0; i < n; i++ {
		fmt.Printf("  предмет %d:\n", i+1)
		w := readInt(in, "    вес: ", func(v int) bool { return v > 0 })
		v := readInt(in, "    ценность: ", func(v int) bool { return v >= 0 })
		t := readInt(in, "    время: ", func(v int) bool { return v > 0 })
		items[i] = algo.Item{Weight: w, Value: v, Time: t}
	}

	result := algo.SolveKnapsack(items, maxW, maxT)

	fmt.Printf("\nМаксимальная ценность: $%d\n", result.MaxValue)
}

func readInt(in *bufio.Reader, prompt string, ok func(int) bool) int {
	for {
		fmt.Print(prompt)
		line, err := in.ReadString('\n')
		if err != nil {
			log.Fatalf("Ошибка чтения ввода: %v", err)
		}
		line = strings.TrimSpace(line)
		v, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("  ожидалось целое число, попробуйте снова")
			continue
		}
		if !ok(v) {
			fmt.Println("  недопустимое значение, попробуйте снова")
			continue
		}
		return v
	}
}

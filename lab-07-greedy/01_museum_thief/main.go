package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/kazantsev-developer/itmo-algorithms-labs-2026/lab-07-greedy/internal/algo"
)

func main() {
	fmt.Println("Задача 1: Вор в музее")

	in := bufio.NewReader(os.Stdin)

	n := readInt(in, "Введите количество предметов (N): ", func(v int) bool { return v > 0 })
	m := readInt(in, "Введите количество заходов (M): ", func(v int) bool { return v > 0 })
	k := readInt(in, "Введите грузоподъемность за один заход (K): ", func(v int) bool { return v > 0 })

	items := make([]algo.Item, n)
	fmt.Printf("Введите %d предметов (вес, цена):\n", n)
	for i := 0; i < n; i++ {
		w := readInt(in, "  вес: ", func(v int) bool { return v > 0 })
		v := readInt(in, "  цена: ", func(v int) bool { return v >= 0 })
		items[i] = algo.Item{Weight: w, Value: v}
	}

	res := algo.SolveMuseumThief(items, m*k)
	fmt.Printf("\nМаксимальная сумма украденного: $%d\n", res)
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

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
	fmt.Println("Задача 2: Наибольшая непрерывная возрастающая последовательность")

	in := bufio.NewReader(os.Stdin)

	n := readInt(in, "Введите количество элементов (n): ", func(v int) bool { return v > 0 })

	arr := make([]int, n)
	fmt.Printf("Введите %d элементов (от -100 до 100):\n", n)
	for i := 0; i < n; i++ {
		arr[i] = readInt(in, fmt.Sprintf("  элемент %d: ", i+1), func(v int) bool { return v >= -100 && v <= 100 })
	}

	res := algo.SolveLongestIncreasing(arr)
	fmt.Printf("\nДлина наибольшей непрерывной возрастающей последовательности: %d\n", res)
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

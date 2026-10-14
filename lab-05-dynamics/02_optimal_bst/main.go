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
	fmt.Println("Задача 2: Оптимальное бинарное дерево поиска")

	in := bufio.NewReader(os.Stdin)

	n := readInt(in, "Введите количество ключей (N): ", func(v int) bool { return v > 0 })

	p := make([]int, n)
	fmt.Printf("Введите %d частот обращения (целые неотрицательные числа):\n", n)
	for i := 0; i < n; i++ {
		p[i] = readInt(in, fmt.Sprintf("  частота ключа %d: ", i+1), func(v int) bool { return v >= 0 })
	}

	res := algo.SolveOptimalBST(p)

	fmt.Printf("\nМинимальная стоимость поиска: %d\n", res.Cost)
	fmt.Println("Структура дерева (префиксная запись, точка - пустое поддерево):")
	fmt.Println("  " + renderTree(res.Root, 1, n))

	depths := make([]int, n)
	collectDepths(res.Root, 1, n, 0, depths)

	fmt.Println("\nВклад каждого ключа:")
	total := 0
	for i, d := range depths {
		contrib := (d + 1) * p[i]
		total += contrib
		fmt.Printf("  ключ %d: глубина %d, (глубина+1)*частота = %d * %d = %d\n",
			i+1, d, d+1, p[i], contrib)
	}
	fmt.Printf("  итого: %d\n", total)
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

func renderTree(root [][]int, i, j int) string {
	if i > j {
		return "."
	}
	k := root[i][j]
	return fmt.Sprintf("(%s k%d %s)", renderTree(root, i, k-1), k, renderTree(root, k+1, j))
}

func collectDepths(root [][]int, i, j, depth int, out []int) {
	if i > j {
		return
	}
	k := root[i][j]
	out[k-1] = depth
	collectDepths(root, i, k-1, depth+1, out)
	collectDepths(root, k+1, j, depth+1, out)
}

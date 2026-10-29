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
	fmt.Println("Задача 4: Резка бруска с учётом стоимости разрезов")

	in := bufio.NewReader(os.Stdin)

	L := readInt(in, "Введите длину бруска (L): ", func(v int) bool { return v > 0 })
	n := readInt(in, "Введите количество заказов (N): ", func(v int) bool { return v >= 0 })
	C := readInt(in, "Введите стоимость одного разреза (C): ", func(v int) bool { return v >= 0 })

	orders := make([]algo.Order, n)
	fmt.Printf("Введите %d заказов (длина, цена):\n", n)
	for i := 0; i < n; i++ {
		fmt.Printf("  заказ %d:\n", i+1)
		d := readInt(in, "    длина: ", func(v int) bool { return v > 0 })
		p := readInt(in, "    цена: ", func(v int) bool { return v >= 0 })
		orders[i] = algo.Order{Length: d, Price: p}
	}

	res := algo.SolveRodCutting(orders, L, C)

	fmt.Printf("\nМаксимальная прибыль: %d\n", res.MaxProfit)
	if len(res.Chosen) == 0 {
		fmt.Println("Ни один заказ невыгодно выполнять.")
		return
	}
	fmt.Printf("Выбрано заказов: %d\n", len(res.Chosen))
	for i, o := range res.Chosen {
		fmt.Printf("  %d) длина %d, цена %d, чистая прибыль %d\n",
			i+1, o.Length, o.Price, o.Price-C)
	}
	fmt.Printf("Суммарная длина выбранных: %d из %d\n", res.TotalLen, L)
	cuts := len(res.Chosen)
	fmt.Printf("Стоимость разрезов: %d × %d = %d\n", cuts, C, cuts*C)
	fmt.Printf("Проверка: сумма цен %d − разрезы %d = %d\n",
		res.MaxProfit+cuts*C, cuts*C, res.MaxProfit)
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

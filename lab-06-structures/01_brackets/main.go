package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/kazantsev-developer/itmo-algorithms-labs-2026/lab-06-structures/internal/algo"
)

func main() {
	fmt.Println("Задание 1. Проверка скобочной структуры")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите скобочную последовательность: ")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка чтения ввода:", err)
		return
	}

	input = strings.TrimSpace(input)
	isValid, errorPos := algo.CheckBrackets(input)

	if isValid {
		fmt.Println("Структура правильная.")
	} else {
		fmt.Printf("Структура неправильная. Первая ошибка в символе под номером: %d\n", errorPos)
	}
}

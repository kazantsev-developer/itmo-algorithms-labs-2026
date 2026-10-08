package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

const (
	maxValue = 100
	minValue = -100
)

type run struct {
	start int
	end   int
}

func (r run) length() int {
	return r.end - r.start + 1
}

func readArraySize(reader *bufio.Reader) int {
	for {
		fmt.Print("Введите размер массива n (от 2 до 1000): ")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка ввода, попробуйте снова.")
			continue
		}

		n, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			fmt.Println("Это не целое число, попробуйте снова.")
			continue
		}
		if n < 2 || n > 1000 {
			fmt.Println("Число должно быть в диапазоне от 2 до 1000.")
			continue
		}

		return n
	}
}

func generateRandomSlice(size int) []int {
	nums := make([]int, size)
	for i := range nums {
		nums[i] = rand.IntN(maxValue-minValue+1) + minValue
	}
	return nums
}

func longestAscendingRun(nums []int) run {
	if len(nums) == 0 {
		return run{start: -1, end: -1}
	}

	best := run{start: 0, end: 0}
	current := run{start: 0, end: 0}

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			current.end = i
		} else {
			if current.length() > best.length() {
				best = current
			}
			current = run{start: i, end: i}
		}
	}

	if current.length() > best.length() {
		best = current
	}

	return best
}

func printSlice(nums []int) {
	for i, v := range nums {
		fmt.Printf("%5d", v)
		if (i+1)%10 == 0 {
			fmt.Println()
		}
	}
	if len(nums)%10 != 0 {
		fmt.Println()
	}
}

func printRun(nums []int, r run) {
	for i := r.start; i <= r.end; i++ {
		fmt.Printf("%d", nums[i])
		if i < r.end {
			fmt.Print(" < ")
		}
	}
	fmt.Println()
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Поиск наибольшей непрерывной возрастающей последовательности")
	fmt.Println()

	n := readArraySize(reader)
	nums := generateRandomSlice(n)

	fmt.Println()
	fmt.Println("Исходный массив:")
	printSlice(nums)
	fmt.Println()

	best := longestAscendingRun(nums)

	if best.length() <= 1 {
		fmt.Println("Возрастающих последовательностей длиной больше 1 не найдено.")
		return
	}

	fmt.Printf("Длина наибольшей возрастающей серии: %d\n", best.length())
	fmt.Printf("Индексы: [%d..%d]\n", best.start, best.end)
	fmt.Print("Сама последовательность: ")
	printRun(nums, best)
}

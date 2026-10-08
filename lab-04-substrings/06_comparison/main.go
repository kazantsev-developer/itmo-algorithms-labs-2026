package main

import (
	"fmt"
	"os"
	"slices"
	"text/tabwriter"
	"time"

	"github.com/kazantsev-developer/itmo-algorithms-labs-2026/lab-04-substrings/internal/algo"
)

const primeCount = 500

type benchmarkResult struct {
	name      string
	elapsed   time.Duration
	bestNum   string
	bestCount int
}

func benchmarkAlgorithm(name string, search func(string, string) int, text string) benchmarkResult {
	start := time.Now()
	results := algo.CountTwoDigitFrequencies(text, search)
	elapsed := time.Since(start)

	slices.SortFunc(results, func(a, b algo.Freq) int {
		return b.Count - a.Count
	})

	return benchmarkResult{
		name:      name,
		elapsed:   elapsed,
		bestNum:   results[0].Number,
		bestCount: results[0].Count,
	}
}

func main() {
	fmt.Println("Сравнение алгоритмов поиска подстрок")
	fmt.Println()

	primes := algo.FirstNPrimes(primeCount)
	text := algo.BuildText(primes)

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintf(w, "Сгенерировано простых чисел:\t%d\n", primeCount)
	fmt.Fprintf(w, "Длина склеенной строки:\t%d символов\n", len(text))
	fmt.Fprintf(w, "Двузначных шаблонов:\t90\n")
	w.Flush()

	fmt.Println()

	algorithms := []struct {
		name string
		fn   func(string, string) int
	}{
		{"Наивный", algo.NaiveSearch},
		{"Рабина-Карпа", algo.RabinKarp},
		{"Бойера-Мура", algo.BoyerMoore},
		{"Кнута-Морриса-Пратта", algo.KMP},
	}

	results := make([]benchmarkResult, 0, len(algorithms))
	for _, a := range algorithms {
		results = append(results, benchmarkAlgorithm(a.name, a.fn, text))
	}

	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(tw, "Алгоритм\tВремя\tМаксимум")
	for _, r := range results {
		fmt.Fprintf(tw, "%s\t%v\t%s - %d\n", r.name, r.elapsed, r.bestNum, r.bestCount)
	}
	tw.Flush()

	fmt.Println()

	complexity := []struct {
		name    string
		average string
		worst   string
		memory  string
		feature string
	}{
		{"Наивный", "O(n·m)", "O(n·m)", "O(1)", "простой, но медленный на длинных паттернах"},
		{"Рабина-Карпа", "O(n+m)", "O(n·m)", "O(1)", "требует проверки коллизий хэшей"},
		{"Бойера-Мура", "O(n/m)", "O(n·m)", "O(σ)", "самый быстрый на практике, сложнее в реализации"},
		{"Кнута-Морриса-Пратта", "O(n+m)", "O(n+m)", "O(m)", "гарантированная линейность на любых данных"},
	}

	cw := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(cw, "Алгоритм\tСреднее\tХудшее\tПамять\tОсобенность")
	for _, c := range complexity {
		fmt.Fprintf(cw, "%s\t%s\t%s\t%s\t%s\n",
			c.name, c.average, c.worst, c.memory, c.feature)
	}
	cw.Flush()

	fmt.Println()
	fmt.Println("где n - длина текста, m - длина паттерна, σ - размер алфавита")
}

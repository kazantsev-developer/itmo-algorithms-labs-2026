// Package report - отчёты о поиске подстрок
package report

import (
	"fmt"
	"os"
	"slices"
	"text/tabwriter"

	"github.com/kazantsev-developer/itmo-algorithms-labs-2026/lab-04-substrings/internal/algo"
)

const (
	primeCount  = 500
	previewSize = 100
	topSize     = 10
)

func Search(title string, search func(string, string) int) {
	fmt.Println(title)
	fmt.Println()

	primes := algo.FirstNPrimes(primeCount)
	text := algo.BuildText(primes)

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintf(w, "Сгенерировано простых чисел:\t%d\n", primeCount)
	fmt.Fprintf(w, "Первое простое число:\t%d\n", primes[0])
	fmt.Fprintf(w, "Последнее простое число:\t%d\n", primes[len(primes)-1])
	fmt.Fprintf(w, "Длина склеенной строки:\t%d символов\n", len(text))
	w.Flush()

	fmt.Println()

	preview := text
	if len(preview) > previewSize {
		preview = preview[:previewSize]
	}
	fmt.Println("Первые 100 символов строки:")
	fmt.Println(preview)
	fmt.Println()

	results := algo.CountTwoDigitFrequencies(text, search)

	slices.SortFunc(results, func(a, b algo.Freq) int {
		return b.Count - a.Count
	})

	maxCount := results[0].Count
	fmt.Printf("Максимум вхождений: %d\n", maxCount)
	fmt.Println("Двузначные числа с максимальной частотой:")
	for _, r := range results {
		if r.Count != maxCount {
			break
		}
		fmt.Printf("  %s - %d\n", r.Number, r.Count)
	}
	fmt.Println()

	fmt.Println("Топ-10 двузначных чисел:")
	for i := 0; i < topSize && i < len(results); i++ {
		fmt.Printf("  %2d. %s - %d\n", i+1, results[i].Number, results[i].Count)
	}
}

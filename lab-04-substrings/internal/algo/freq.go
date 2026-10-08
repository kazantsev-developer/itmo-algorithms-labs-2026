package algo

import "strconv"

type Freq struct {
	Number string
	Count  int
}

func CountTwoDigitFrequencies(text string, search func(string, string) int) []Freq {
	results := make([]Freq, 0, 90)
	for num := 10; num <= 99; num++ {
		pattern := strconv.Itoa(num)
		results = append(results, Freq{pattern, search(text, pattern)})
	}
	return results
}

package main

import (
	"github.com/kazantsev-developer/itmo-algorithms-labs-2026/lab-04-substrings/internal/algo"
	"github.com/kazantsev-developer/itmo-algorithms-labs-2026/lab-04-substrings/internal/report"
)

func main() {
	report.Search("Наивный поиск подстрок", algo.NaiveSearch)
}

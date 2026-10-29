// Package algo содержит реализации жадных алгоритмов
package algo

import "sort"

type Item struct {
	Weight int
	Value  int
	Time   int
}

type Result struct {
	MaxValue    int
	ChosenItems []Item
}

type Order struct {
	Length int
	Price  int
}

type RodResult struct {
	MaxProfit int
	Chosen    []Order
	TotalLen  int
}

func SolveMuseumThief(items []Item, maxWeight int) int {
	sort.Slice(items, func(i, j int) bool {
		return float64(items[i].Value)/float64(items[i].Weight) > float64(items[j].Value)/float64(items[j].Weight)
	})
	totalValue := 0
	currentWeight := 0
	for _, item := range items {
		if currentWeight+item.Weight <= maxWeight {
			currentWeight += item.Weight
			totalValue += item.Value
		}
	}
	return totalValue
}

func SolveLongestIncreasing(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	maxLen := 1
	currentLen := 1
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			currentLen++
		} else {
			if currentLen > maxLen {
				maxLen = currentLen
			}
			currentLen = 1
		}
	}
	if currentLen > maxLen {
		maxLen = currentLen
	}
	return maxLen
}

func SolveKnapsack(items []Item, maxW int, maxT int) Result {
	sortedItems := make([]Item, len(items))
	copy(sortedItems, items)

	sort.Slice(sortedItems, func(i, j int) bool {
		return sortedItems[i].Value > sortedItems[j].Value
	})

	totalValue := 0
	currentWeight := 0
	currentTime := 0
	var chosen []Item

	for _, item := range sortedItems {
		if currentWeight+item.Weight <= maxW && currentTime+item.Time <= maxT {
			currentWeight += item.Weight
			currentTime += item.Time
			totalValue += item.Value
			chosen = append(chosen, item)
		}
	}

	return Result{
		MaxValue:    totalValue,
		ChosenItems: chosen,
	}
}

func SolveRodCutting(orders []Order, L, C int) RodResult {
	sortedOrders := make([]Order, len(orders))
	copy(sortedOrders, orders)

	sort.Slice(sortedOrders, func(i, j int) bool {
		return (sortedOrders[i].Price - C) > (sortedOrders[j].Price - C)
	})

	totalProfit := 0
	totalLen := 0
	var chosen []Order

	for _, o := range sortedOrders {
		if o.Price-C > 0 && totalLen+o.Length <= L {
			totalLen += o.Length
			totalProfit += (o.Price - C)
			chosen = append(chosen, o)
		}
	}

	return RodResult{
		MaxProfit: totalProfit,
		Chosen:    chosen,
		TotalLen:  totalLen,
	}
}

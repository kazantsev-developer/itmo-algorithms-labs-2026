// Package algo реализует структуры LIFO, FIFO, BFS и DFS
package algo

type bracketItem struct {
	char rune
	pos  int
}

func CheckBrackets(s string) (bool, int) {
	stack := make([]bracketItem, 0)
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for i, char := range s {
		pos := i + 1

		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, bracketItem{char, pos})
			continue
		}

		if expectedOpen, ok := pairs[char]; ok {
			if len(stack) == 0 {
				return false, pos
			}

			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if top.char != expectedOpen {
				return false, pos
			}
		}
	}

	if len(stack) > 0 {
		return false, stack[0].pos
	}

	return true, 0
}

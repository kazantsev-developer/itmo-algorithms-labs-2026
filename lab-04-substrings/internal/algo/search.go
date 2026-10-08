package algo

const (
	hashBase = 31
	hashMod  = 1_000_000_007
)

func NaiveSearch(text, pattern string) int {
	n, m := len(text), len(pattern)
	if m == 0 || m > n {
		return 0
	}

	count := 0
	for i := 0; i+m <= n; i++ {
		j := 0
		for j < m && text[i+j] == pattern[j] {
			j++
		}
		if j == m {
			count++
		}
	}
	return count
}

func RabinKarp(text, pattern string) int {
	n, m := len(text), len(pattern)
	if m == 0 || m > n {
		return 0
	}

	highestPow := 1
	for i := 0; i < m-1; i++ {
		highestPow = (highestPow * hashBase) % hashMod
	}

	var patternHash, windowHash int
	for i := 0; i < m; i++ {
		patternHash = (patternHash*hashBase + int(pattern[i])) % hashMod
		windowHash = (windowHash*hashBase + int(text[i])) % hashMod
	}

	count := 0
	for i := 0; i+m <= n; i++ {
		if patternHash == windowHash {
			match := true
			for j := 0; j < m; j++ {
				if text[i+j] != pattern[j] {
					match = false
					break
				}
			}
			if match {
				count++
			}
		}
		if i+m < n {
			windowHash = (windowHash - int(text[i])*highestPow) % hashMod
			if windowHash < 0 {
				windowHash += hashMod
			}
			windowHash = (windowHash*hashBase + int(text[i+m])) % hashMod
		}
	}
	return count
}

func BoyerMoore(text, pattern string) int {
	n, m := len(text), len(pattern)
	if m == 0 || m > n {
		return 0
	}

	badChar := make(map[byte]int, m)
	for i := 0; i < m; i++ {
		badChar[pattern[i]] = i
	}

	count := 0
	shift := 0
	for shift <= n-m {
		j := m - 1
		for j >= 0 && pattern[j] == text[shift+j] {
			j--
		}

		if j < 0 {
			count++
			shift++
		} else {
			last, ok := badChar[text[shift+j]]
			if !ok {
				last = -1
			}
			skip := j - last
			if skip < 1 {
				skip = 1
			}
			shift += skip
		}
	}
	return count
}

func KMP(text, pattern string) int {
	n, m := len(text), len(pattern)
	if m == 0 || m > n {
		return 0
	}

	pi := buildPrefixFunction(pattern)

	count := 0
	j := 0
	for i := 0; i < n; i++ {
		for j > 0 && text[i] != pattern[j] {
			j = pi[j-1]
		}
		if text[i] == pattern[j] {
			j++
		}
		if j == m {
			count++
			j = pi[j-1]
		}
	}
	return count
}

func buildPrefixFunction(pattern string) []int {
	m := len(pattern)
	pi := make([]int, m)

	for i := 1; i < m; i++ {
		j := pi[i-1]
		for j > 0 && pattern[i] != pattern[j] {
			j = pi[j-1]
		}
		if pattern[i] == pattern[j] {
			j++
		}
		pi[i] = j
	}
	return pi
}

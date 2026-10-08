package algo

func FirstNPrimes(n int) []int {
	primes := make([]int, 0, n)
	for candidate := 2; len(primes) < n; candidate++ {
		if isPrime(candidate, primes) {
			primes = append(primes, candidate)
		}
	}
	return primes
}

func isPrime(candidate int, knownPrimes []int) bool {
	for _, p := range knownPrimes {
		if p*p > candidate {
			break
		}
		if candidate%p == 0 {
			return false
		}
	}
	return true
}

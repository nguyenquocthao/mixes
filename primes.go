package util

var notprimes = make([]bool, 300001)
var primes = []int{}

func init() {
	for i := 2; i < len(notprimes); i++ {
		if notprimes[i] {
			continue
		}
		primes = append(primes, i)
		for j := i * i; j < len(notprimes); j += i {
			notprimes[j] = true
		}
	}
}

var factors = make([][]int, 1000001)

func init() {
	for i := 2; i < len(factors); i++ {
		if len(factors[i]) > 0 {
			continue
		}
		for j := i; j < len(factors); j += i {
			factors[j] = append(factors[j], i)
		}
	}
}

func factorize(v int) []int {
	lo, hi := []int{}, []int{}
	for i := 1; i*i <= v; i++ {
		if i*i == v {
			lo = append(lo, i)
			break
		} else if v%i == 0 {
			lo = append(lo, i)
			hi = append(hi, v/i)
		}
	}
	for j := len(hi) - 1; j >= 0; j-- {
		lo = append(lo, hi[j])
	}
	return lo
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * (b / gcd(a, b))
}

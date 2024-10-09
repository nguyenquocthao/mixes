package util

func FuzzyEquals(c1, c2 byte) bool {
	return c1 == '-' || c2 == '-' || c1 == c2
}

const Q = 5

var LF int

func FuzzyMatch(t, s string) bool {
	n := len(t)
	if len(s) != n {
		return false
	}
	if n <= 2*Q {
		for i := 0; i < n; i++ {
			if !FuzzyEquals(t[i], s[i]) {
				return false
			}
		}
		return true
	}

	for i := 0; i < Q; i++ {
		if !FuzzyEquals(t[i], s[i]) {
			return false
		}
	}
	for i := n - Q; i < n; i++ {
		if !FuzzyEquals(t[i], s[i]) {
			return false
		}
	}

	l := clamp(LF, 0, n-1)
	r := l + 1

	for l >= Q || r < n-Q {
		if l >= Q {
			if !FuzzyEquals(t[l], s[l]) {
				LF = l
				return false
			}
			l--
		}
		if r < n-Q {
			if !FuzzyEquals(t[r], s[r]) {
				LF = r
				return false
			}
			r++
		}
	}
	return true
}

func FuzzyFirstIndexOf(t, s string) int {
	n := len(t)
	m := len(s)

	for i := 0; i <= m-n; i++ {
		if FuzzyMatch(t, s[i:i+n]) {
			return i
		}
	}
	return -1
}

func clamp(v, min, max int) int {
	if v < min {
		return min
	} else if v > max {
		return max
	}
	return v
}

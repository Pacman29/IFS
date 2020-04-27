package main

func NaiveStringMatcher(t, p string) (out []int) {
	n, m := len(t), len(p)

	for i := 0; i <= n-m; i++ {
		if p == t[i: i + m] {
			out = append(out, i)
		}
	}

	return out
}

func uniqueLetters(s string) []uint8 {
	tmp := make([]uint8, 0, len(s)/2)
	for j := range s {
		flag := false
		for i := 0; i < len(tmp) && !flag; i++ {
			flag = tmp[i] == s[j]
		}
		if !flag {
			tmp = append(tmp, s[j])
		}
	}
	return tmp
}

func pow(a, b int) int {
	p := 1
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}

func RabinKarpMatcher(t, p string) (out []int) {
	n, m := len(t), len(p)
	d := uint64(len(uniqueLetters(t)))
	h := uint64(pow(int(d), m-1))

	var _p uint64 = 0
	var _t uint64 = 0

	var tmp = uint64(0)
	for i := 0; i < m; i++ {
		tmp = uint64(p[i])
		_p, _t = d*_p + uint64(p[i]), d*_t + uint64(t[i])
	}
	tmp += 1

	for i := 0; i <= n-m; i++ {
		if _p == _t && p == t[i: i + m] {
			out = append(out, i)
		}
		if i < n-m {
			_t = d * (_t - uint64(t[i]) * h) + uint64(t[i + m])
		}
	}
	return out
}

func getNextState(p string, state int, x uint8) int {
	m := len(p)

	if state < m && x == p[state] {
		return state+1
	}

	for nextState := state; nextState > 0; nextState-- {
		if p[nextState-1] == x {
			i := 0
			for ; i < nextState-1 && p[i] == p[state-nextState+1+i]; i++ {}
			if i == nextState-1 {
				return nextState
			}
		}
	}

	return 0
}

func initFA(p string, letters []uint8) []map[uint8]int {
	m := len(p)
	fa := make([]map[uint8]int, m+1)
	for i := range fa {
		fa[i] = make(map[uint8]int, len(letters))
	}

	for state := 0; state <= m; state++ {
		for x := 0; x < len(letters); x++ {
			letter := letters[x]
			fa[state][letter] = getNextState(p, state, letter)
		}
	}

	return fa
}

func FAMatcher(t, p string) (out []int) {
	n, m := len(t), len(p)

	fa := initFA(p, uniqueLetters(t))

	for state, i := 0, 0 ; i < n; i++ {
		state = fa[state][t[i]]
		if state == m {
			out = append(out, i - m + 1)
		}
	}

	return out
}

func computePrefixFunction(p string) []int {
	m := len(p)
	pi := make([]int, m)
	k := -1
	pi[0] = -1
	for q := 1; q < m; q++ {
		for k > -1 && p[k] != p[q] {
			k = pi[k]
		}
		if p[k+1] == p[q] {
			k++
		}
		pi[q] = k
	}
	return pi
}

func KnuthMorrisPratMatcher(t, p string) (out []int) {
	n, m := len(t), len(p)
	pi := computePrefixFunction(p)
	q := -1

	for i := 0; i < n; i++ {
		for q > -1 && p[q+1] != t[i] {
			q = pi[q]
		}

		if p[q+1] == t[i] {
			q++
		}

		if q == m-1 {
			out = append(out, i-m+1)
			q = pi[q]
		}
	}
	return out
}

func initOccurrence(occ map[uint8]int, p string) {
	m := len(p)
	for j:= 0; j<m; j++ {
		occ[p[j]] = j
	}
}

func preprocessStrongSuffix(bpos, shift []int, p string) {
	m := len(p)
	i, j := m, m+1
	bpos[i] = j

	for ; i > 0; {
		for ; j <= m && p[i-1] != p[j-1]; {
			if shift[j] == 0 {
				shift[j] = j-i
			}

			j = bpos[j]
		}
		i--
		j--
		bpos[i] = j
	}
}

func preprocessBorder(bpos, shift []int, p string) {
	j, m := bpos[0], len(p)
	for i := 0; i <= m; i++ {
		if shift[i] == 0 {
			shift[i] = j
		}
		if i == j {
			j = bpos[j]
		}
	}
}

func BoyerMooreMatcher(t, p string) (out []int) {
	n, m := len(t), len(p)
	s, j := 0, 0

	letters := uniqueLetters(t)

	occ, bpos, shift := make(map[uint8]int, len(letters)), make([]int, m+1), make([]int, m+1)

	for _, letter := range letters {
		occ[letter] = -1
	}

	initOccurrence(occ, p)
	preprocessStrongSuffix(bpos, shift, p)
	preprocessBorder(bpos, shift, p)
	
	for ; s <= n-m; {
		j = m-1
		for ; j >= 0 && p[j] == t[s+j]; j-- {}

		if j < 0 {
			out = append(out, s)
			s += shift[0]
		} else {
			s += max(shift[j+1], j - occ[t[s+j]])
		}
	}
	return out
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
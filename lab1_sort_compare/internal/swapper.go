package internal

type Swapper struct {
	Count int
}

func NewSwapper() *Swapper {
	return &Swapper{Count: 0}
}

func (s *Swapper) Swap(a, b *int) {
	*a, *b = *b, *a
	s.Count++
}

func (s *Swapper) SwapArr(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
	s.Count++
}



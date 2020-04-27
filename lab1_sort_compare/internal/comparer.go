package internal

type Comparator func(i, j int) bool

type Comparer struct {
	Count int
	Comparator Comparator
}

func NewComparer(comparator Comparator) *Comparer {
	return &Comparer{Count: 0, Comparator: comparator}
}

func (c *Comparer) Compare(i, j int) bool {
	c.Count++
	return c.Comparator(i, j)
}

func (c *Comparer) Equal(i, j int) bool {
	c.Count++
	return i == j
}
package utils

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h IntHeap) Peek() int { return h[0] }

func (h *IntHeap) Push(x any) {
	// Push uses pointer receivers because they modify the slice's length
	// not able to use primitives here because interface states any
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h // turns it into a value!

	x := old[h.Len()-1]
	*h = old[:h.Len()-1]

	return x
}

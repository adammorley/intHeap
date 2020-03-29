/*
   a simple integer heap
*/
package intHeap

// min heap of ints
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// pointer receivers because modifying slice
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() (r interface{}) {
	old := *h
	n := len(old)
	r = old[n-1]
	*h = old[0 : n-1]
	return
}

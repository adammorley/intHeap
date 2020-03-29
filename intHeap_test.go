package intHeap

import (
	"container/heap"
	"math/rand"
	"testing"
)

func TestIntHeap(test *testing.T) {
	h := &IntHeap{}
	heap.Init(h)
	m := make(map[int]bool)
	for i := 0; i < 20; i++ {
		j := rand.Intn(1000)
		for m[j] {
			j = rand.Intn(1000)
		}
		heap.Push(h, j)
	}
	p := heap.Pop(h).(int)
	var c int
	for h.Len() > 0 {
		c = heap.Pop(h).(int)
		if c < p {
			test.Error("popped a value that's less than previous value!")
		}
		p = c
	}
}

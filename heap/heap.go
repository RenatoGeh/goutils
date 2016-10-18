package heap

import (
	"errors"
)

// Heap is a binary priority queue.
type Heap struct {
	data []int
	size int
	cap  int
	crit int
}

// Left child.
func (h *Heap) Left(i int) (int, error) {
	v := 2*i + 1
	if v > h.cap || v >= h.size {
		return -1, errors.New("There is no left child.")
	}
	return v, nil
}

// Right child.
func (h *Heap) Right(i int) (int, error) {
	v := 2*i + 2
	if v > h.cap || v >= h.size {
		return -1, errors.New("There is no right child.")
	}
	return v, nil
}

// Parent node.
func (h *Heap) Parent(i int) (int, error) {
	v := (i - 1) / 2
	if v > h.cap || v < 1 {
		return -1, errors.New("There is no such parent.")
	}
	return v, nil
}

// Cmp compares two values.
func (h *Heap) Cmp(i, j int) bool {
	if i < 0 || j < 0 || i >= h.size || j >= h.size {
		return false
	} else if h.crit > 0 {
		return h.data[i] > h.data[j]
	}
	return h.data[i] < h.data[j]
}

// NewHeap creates new Heap with cap initial capacity and criteria.
// Criteria is non-positive if min-heap and positive if max-heap.
func NewHeap(cap, crit int) *Heap {
	return &Heap{data: make([]int, cap+1), size: 0, cap: cap, crit: crit}
}

// e is index
func (h *Heap) bubble(e int) {
	ipa, err := h.Parent(e)
	for h.Cmp(e, ipa) {
		if err != nil {
			return
		}
		h.data[ipa], h.data[e] = h.data[e], h.data[ipa]
		ipa, e = e, ipa
		ipa, err = h.Parent(e)
	}
}

// Insert adds element e into heap.
func (h *Heap) Insert(e int) {
	if h.size >= h.cap {
		h.data = append(h.data, e)
		h.cap++
	} else {
		h.data[h.size] = e
	}
	h.size++
	h.bubble(h.size - 1)
}

func (h *Heap) minmax(r, l int) (int, int) {
	rval, lval := h.data[r], h.data[l]
	if rval > lval {
		return r, l
	}
	return l, r
}

// e is index
func (h *Heap) sink(e int) {
	left, lerr := h.Left(e)
	right, rerr := h.Right(e)
	lval, rval := h.Cmp(e, left), h.Cmp(e, right)
	for (left < h.size && right < h.size) && (lval || rval) {
		if lerr != nil || rerr != nil {
			return
		}
		min, max := h.minmax(left, right)
		if h.crit > 0 {
			h.data[max], h.data[e] = h.data[e], h.data[max]
			e = max
		} else {
			h.data[min], h.data[e] = h.data[e], h.data[min]
			e = min
		}
		left, lerr = h.Left(e)
		right, rerr = h.Right(e)
		lval, rval = h.Cmp(e, left), h.Cmp(e, right)
	}
}

// Extract returns the min or max.
func (h *Heap) Extract() int {
	r := h.data[1]
	h.size--
	h.data[1] = h.data[h.size]
	h.sink(h.size)
	return r
}

// Delete removes element of value e.
func (h *Heap) Delete(e int) {
	for i := 0; i < h.size; i++ {
		if h.data[i] == e {
			h.size--
			h.data[i] = h.data[h.size]
			h.sink(i)
			return
		}
	}
}

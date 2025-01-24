package toolz

import (
	"sort"
	"time"
)

type LessFunc[T any] func(p1, p2 T) bool

// MultiSorter implements the Sort interface, sorting the changes within.
type MultiSorter[T any] struct {
	elements []T
	less     []LessFunc[T]
}

// Sort sorts the argument slice according to the less functions passed to OrderedBy.
func (ms *MultiSorter[T]) Sort(elements []T) {
	ms.elements = elements
	sort.Sort(ms)
}

// OrderedBy returns a Sorter that sorts using the less functions, in order.
// Call its Sort method to sort the data.
func OrderedBy[T any](less ...LessFunc[T]) *MultiSorter[T] {
	return &MultiSorter[T]{
		less: less,
	}
}

// Len is part of sort.Interface.
func (ms *MultiSorter[T]) Len() int {
	return len(ms.elements)
}

// Swap is part of sort.Interface.
func (ms *MultiSorter[T]) Swap(i, j int) {
	ms.elements[i], ms.elements[j] = ms.elements[j], ms.elements[i]
}

// Less is part of sort.Interface. It is implemented by looping along the
// less functions until it finds a comparison that discriminates between
// the two items (one is less than the other). Note that it can call the
// less functions twice per call. We could change the functions to return
// -1, 0, 1 and reduce the number of calls for greater efficiency: an
// exercise for the reader.
func (ms *MultiSorter[T]) Less(i, j int) bool {
	p, q := ms.elements[i], ms.elements[j]
	// Try all but the last comparison.
	var k int
	for k = 0; k < len(ms.less)-1; k++ {
		less := ms.less[k]
		switch {
		case less(p, q):
			// p < q, so we have a decision.
			return true
		case less(q, p):
			// p > q, so we have a decision.
			return false
		}
		// p == q; try the next comparison.
	}
	// All comparisons to here said "equal", so just return whatever
	// the final comparison reports.
	return ms.less[k](p, q)
}

type SimpLessConstraint interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | string
}

func SimpLess[T SimpLessConstraint](p1, p2 T, asc bool) bool {
	if asc {
		return p1 < p2
	} else {
		return p1 > p2
	}
}

func TimeLess(p1, p2 time.Time, asc bool) bool {
	if asc {
		return p1.Before(p2)
	} else {
		return p2.Before(p1)
	}
}

package linear

// PriorityQueue priority queue
type PriorityQueue struct {
	elements []interface{}
}

// NewPriorityQueue returns new PriorityQueue
func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{}
}

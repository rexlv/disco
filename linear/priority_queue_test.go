package linear

import (
	"fmt"
	"testing"
)

func Test_PriorityQueue(t *testing.T) {
	q := NewPriorityQueue()
	fmt.Println("queue: ", q.elements)
}

package linear

import (
	"fmt"
	"testing"
)

func Test_Queue(t *testing.T) {
	q := NewQueue()
	q.Put("hello", 0, 12)
	fmt.Println("queue: ", q.elements)
	fmt.Println("queue: ", q.Pop(), q.Pop())
}

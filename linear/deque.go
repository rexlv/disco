package linear

// 基于栈的双向队列, 避免产生大量的指针，同时尽量降低栈拷贝
// Bucket Item元素也应该尽量是简单类型，否则仍然会产生大量的指针

// Bucket bucket
type Bucket struct {
	items []interface{}
	cap   int
}

type link struct {
	curr uintptr
	prev uintptr
	next uintptr
}

type Deque struct {
	buckets []*Bucket
	head    *link
}

// NewDeque returns new deque
func NewDeque() *Deque {
	q := &Deque{}
	q.buckets = make([]*Bucket, 0)
	for i := 0; i < 3; i++ {
		q.buckets = append(q.buckets, &Bucket{
			items: make([]interface{}, 64),
			cap:   64,
		})
	}

	return q
}

func (q *Deque) LPush(items ...interface{}) {

}

func (q *Deque) RPush(items ...interface{}) {

}

func (q *Deque) LPop() interface{} {
	return nil
}

func (q *Deque) RPop() interface{} {
	return nil
}

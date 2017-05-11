package linear

type Node struct {
	element interface{}
	prev    *Node
	next    *Node
}

type LinkedList struct {
	size int
	root *Node
}

func NewLinkedList() *LinkedList {
	l := &LinkedList{
		size: 0,
		root: &Node{element: nil},
	}
	l.root.next = l.root
	l.root.prev = l.root
	return l
}

func (l *LinkedList) LPush(elems ...interface{}) {

}
func (l *LinkedList) RPush(elems ...interface{}) {

}

func (l *LinkedList) LPop() interface{} {
	return nil
}
func (l *LinkedList) RPop() interface{} {
	return nil
}

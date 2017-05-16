package linear

import (
	"fmt"
	"strings"
)

// Queue list
type Queue struct {
	elements []interface{}
	size     int
}

const (
	initSize = 32
)

// NewQueue returns new Queue
func NewQueue() *Queue {
	return &Queue{
		size:     0,
		elements: make([]interface{}, initSize),
	}
}

// Put add elems
func (l *Queue) Put(elems ...interface{}) {
	l.adjust(len(elems))

	for _, elem := range elems {
		l.elements[l.size] = elem
		l.size++
	}
}

func (l *Queue) Pop() (ret interface{}) {
	if l.size <= 0 {
		return nil
	}

	ret = l.elements[l.size-1]
	l.elements = l.elements[:l.size-1]
	return
}

// Get returns indexed element
func (l *Queue) Get(index int) (interface{}, bool) {
	if index >= l.size {
		return nil, false
	}
	return l.elements[index], true
}

// Empty returns wether list is empty
func (l *Queue) Empty() bool {
	return 0 == l.size
}

// Size returns list's size
func (l *Queue) Size() int {
	return l.size
}

// Clear clear all elements
func (l *Queue) Clear() {
	l.size = 0
	l.elements = []interface{}{}
}

func (l *Queue) String() string {
	vals := make([]string, l.size)
	for i, val := range l.elements[:l.size] {
		vals[i] = fmt.Sprintf("%#v", val)
	}
	return strings.Join(vals, ",")
}

// Values returns all elements
func (l *Queue) Values() (vals []interface{}) {
	vals = make([]interface{}, l.size, l.size)
	copy(vals, l.elements[:l.size])
	return vals
}

func (l *Queue) resize(n int) {
	news := make([]interface{}, n, n)
	copy(news, l.elements)
	l.elements = news
}

func (l *Queue) adjust(n int) {
	c := cap(l.elements)

	for l.size+n >= c {
		c = c * 2
	}

	l.resize(c)
}

/************************************************************************************************************/

// Iterator iterator
func (l *Queue) Iterator() func() (int, interface{}) {
	var cursor int
	return func() (index int, val interface{}) {
		if cursor >= l.size {
			return 0, nil
		}

		val = l.elements[cursor]
		index = cursor
		cursor = index + 1
		return index, val
	}
}

// Each impletements Enum interface
func (l *Queue) Each(f func(index int, val interface{})) {
	iter := l.Iterator()
	for i, v := iter(); v != nil; {
		f(i, v)
	}
}

// Map impletements Enum interface
func (l *Queue) Map(f func(index int, val interface{}) interface{}) (m []interface{}) {
	iter := l.Iterator()
	m = make([]interface{}, l.size, cap(l.elements))
	for i, v := iter(); v != nil; {
		m[i] = f(i, v)
	}

	return
}

// Any impletements Enum interface
func (l *Queue) Any(f func(index int, val interface{}) bool) bool {
	iter := l.Iterator()
	for i, v := iter(); v != nil; {
		if f(i, v) {
			return true
		}
	}
	return false
}

// All impletements Enum interface
func (l *Queue) All(f func(index int, val interface{}) bool) bool {
	iter := l.Iterator()
	for i, v := iter(); v != nil; {
		if !f(i, v) {
			return false
		}
	}
	return true
}

// Filter impletement Enum interface
func (l *Queue) Filter(f func(index int, val interface{}) bool) (m []interface{}) {
	iter := l.Iterator()
	m = make([]interface{}, 0, cap(l.elements))
	for i, v := iter(); v != nil; {
		if f(i, v) {
			m = append(m, v)
		}
	}

	return
}

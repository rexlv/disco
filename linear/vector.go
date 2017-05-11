package linear

import (
	"fmt"
	"strings"
)

// Vector list
type Vector struct {
	elements []interface{}
	size     int
}

const (
	initSize = 32
)

// NewVector returns new Vector
func NewVector() *Vector {
	return &Vector{
		size:     0,
		elements: make([]interface{}, initSize),
	}
}

// Add add elems
func (l *Vector) Add(elems ...interface{}) {
	l.adjust(len(elems))

	for _, elem := range elems {
		l.elements[l.size] = elem
		l.size++
	}
}

// Get returns indexed element
func (l *Vector) Get(index int) (interface{}, bool) {
	if index >= l.size {
		return nil, false
	}
	return l.elements[index], true
}

// Empty returns wether list is empty
func (l *Vector) Empty() bool {
	return 0 == l.size
}

// Size returns list's size
func (l *Vector) Size() int {
	return l.size
}

// Clear clear all elements
func (l *Vector) Clear() {
	l.size = 0
	l.elements = []interface{}{}
}

func (l *Vector) String() string {
	vals := make([]string, l.size)
	for i, val := range l.elements[:l.size] {
		vals[i] = fmt.Sprintf("%#v", val)
	}
	return strings.Join(vals, ",")
}

// Values returns all elements
func (l *Vector) Values() (vals []interface{}) {
	vals = make([]interface{}, l.size, l.size)
	copy(vals, l.elements[:l.size])
	return vals
}

func (l *Vector) resize(n int) {
	news := make([]interface{}, n, n)
	copy(news, l.elements)
	l.elements = news
}

func (l *Vector) adjust(n int) {
	c := cap(l.elements)

	for l.size+n >= c {
		c = c * 2
	}

	l.resize(c)
}

// Iterator iterator
func (l *Vector) Iterator() func() (int, interface{}) {
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
func (l *Vector) Each(f func(index int, val interface{})) {
	iter := l.Iterator()
	for i, v := iter(); v != nil; {
		f(i, v)
	}
}

// Map impletements Enum interface
func (l *Vector) Map(f func(index int, val interface{}) interface{}) (m []interface{}) {
	iter := l.Iterator()
	m = make([]interface{}, l.size, cap(l.elements))
	for i, v := iter(); v != nil; {
		m[i] = f(i, v)
	}

	return
}

// Any impletements Enum interface
func (l *Vector) Any(f func(index int, val interface{}) bool) bool {
	iter := l.Iterator()
	for i, v := iter(); v != nil; {
		if f(i, v) {
			return true
		}
	}
	return false
}

// All impletements Enum interface
func (l *Vector) All(f func(index int, val interface{}) bool) bool {
	iter := l.Iterator()
	for i, v := iter(); v != nil; {
		if !f(i, v) {
			return false
		}
	}
	return true
}

// Filter impletement Enum interface
func (l *Vector) Filter(f func(index int, val interface{}) bool) (m []interface{}) {
	iter := l.Iterator()
	m = make([]interface{}, 0, cap(l.elements))
	for i, v := iter(); v != nil; {
		if f(i, v) {
			m = append(m, v)
		}
	}

	return
}

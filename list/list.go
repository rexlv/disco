package list

// Iterator interface
type Iterator interface {
	Next() interface{}
}

// List interface
type List interface {
	Has(...interface{}) bool
	Get(int) (interface{}, bool)
	Remove(int) bool
	Add(...interface{})
	Insert(int, ...interface{})

	Empty() bool
	Purge() bool
	Size() int
	Values() []interface{}
	Iter() Iterator
}

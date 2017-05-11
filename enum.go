package disco

// Enum interface
type Enum interface {
	Each(func(int, interface{}))
	All(func(int, interface{}) bool) bool
	Any(func(int, interface{}) bool) bool
	Filter(func(int, interface{}) bool) []interface{}
	Map(func(int, interface{}) interface{}) []interface{}
}

package set

// std::bitset
// http://www.cplusplus.com/reference/bitset/bitset/

const bpw uint = 8 << (^uint(0)>>8&1 + ^uint(0)>>16&1 + ^uint(0)>>32&1)

// BitSet bitset
type BitSet struct {
	size  uint
	bytes []uint
}

// New creates a new BitSet with a hint that length bits will be required
func New(length uint) (bset *BitSet) {
	if length < 0 {
		return nil
	}
	s := &BitSet{
		size:  length,
		bytes: []uint{0},
	}

	for length/bpw+1 > uint(len(s.bytes)) {
		s.bytes = append(s.bytes, 0)
	}

	return s
}

// Size returns bitset's size
func (s *BitSet) Size() uint {
	return s.size
}

// Test if bit i seted
func (s *BitSet) Test(i uint) bool {
	l := uint(len(s.bytes))
	if i/bpw+1 > l {
		return false
	}
	return (s.bytes[i/bpw] & (1 << uint(i%bpw))) != 0
}

// Any returns if any bit seted, false otherwise
func (s *BitSet) Any() bool {
	for _, field := range s.bytes {
		if field != 0 {
			return true
		}
	}
	return false
}

// None returns if no bit seted, false otherwise
func (s *BitSet) None() (none bool) {
	none = true
	for _, field := range s.bytes {
		none = none && (field == 0)
	}

	return
}

// Count count bits seted
func (s *BitSet) Count() (c uint) {
	// for i:=uint(0); i<s.size; i++ {
	// 	if (s.bytes[i/bpw] | 1<<uint(i/bpw)) == 1 {
	// 		c ++
	// 	}
	// }
	for _, field := range s.bytes {
		for ; field != 0; c++ {
			field &= field - 1
		}
	}

	return
}

// All returns if all bit seted, false otherwise
func (s *BitSet) All() (all bool) {
	return s.Count() == s.size
}

// Set set bit
func (s *BitSet) Set(is ...uint) *BitSet {
	for _, i := range is {
		s.bytes[i/bpw] |= 1 << uint(i%bpw)
	}

	return s
}

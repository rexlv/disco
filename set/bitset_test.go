package set

import (
	"fmt"
	"testing"
)

func Test_HellO(t *testing.T) {
	bs := New(69)
	bs.Set(11, 23, 67)
	for _, i := range []uint{11, 22, 23, 55, 56, 78, 67} {
		fmt.Println("hello: ", bs.Test(i))
	}

	fmt.Println("count: ", bs.Count())
}

func Test_Count(t *testing.T) {
	bs := &BitSet{
		size:  uint(65),
		bytes: []uint{1, 11111, 1},
	}
	fmt.Println("count: ", bs.Count())
}
